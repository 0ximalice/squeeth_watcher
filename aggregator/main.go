package main

import (
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/imalic3/iam-opyn/watcher/abi/opyn"
	"github.com/imalic3/iam-opyn/watcher/abi/uniswapv3"
	"github.com/imalic3/iam-opyn/watcher/abi/uniswapv3oracle"
	"github.com/imalic3/iam-opyn/watcher/pkg/rpcpool"
	"github.com/imalic3/iam-opyn/watcher/pkg/uniswap_helper"
	"github.com/imalic3/iam-opyn/watcher/pkg/utils/persist"
	"github.com/shopspring/decimal"
	"golang.org/x/exp/slog"
)

var (
	INDEX_SCALE = big.NewInt(1e4)
	ONE         = big.NewInt(1e18)
	ONE_DEC     = decimal.NewFromInt(1e18)
	HUNDRED_DEC = decimal.NewFromInt(100)

	// TODO: use env var
	LOG_FILE = "/app/data/ETH².json"
)

type Premium struct {
	At                  int64   `json:"at"`
	Rate                float64 `json:"rate"`
	NormalizationFactor int64   `json:"normalization_factor"`
	SpotPrice           float64 `json:"spot_price"`
	IndexPrice          float64 `json:"index_price"`
	MarkPriceEth        float64 `json:"mark_price_eth"`
	MarkPrice           float64 `json:"mark_price"`
}

var wait = func() {
	time.Sleep(5 * time.Second)
}

func main() {

	logger := slog.Default()
	rpcPoolList := rpcpool.NewRPCPoolList([]string{
		"https://eth.llamarpc.com",
		"https://eth.rpc.blxrbdn.com",
		"https://eth-mainnet.public.blastapi.io",
		"https://rpc.ankr.com/eth",
		"https://rpc.mevblocker.io",
		"https://eth.drpc.org",
		"https://endpoints.omniatech.io/v1/eth/mainnet/public",
		"https://eth-pokt.nodies.app",
		"https://rpc.payload.de",
		"https://rpc.notadegen.com/eth",
		"https://ethereum.publicnode.com",
	})

	weth := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	usdc := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	ethPoolAddress := common.HexToAddress("0x8ad599c3A0ff1De082011EFDDc58f1908eb6e6D8")
	osqthPoolAddress := common.HexToAddress("0x82c427AdFDf2d245Ec51D8046b41c4ee87F0d29C")
	powerTokenCtrlAddress := common.HexToAddress("0x64187ae08781B09368e6253F9E94951243A493D5")
	oracleAddress := common.HexToAddress("0x65D66c76447ccB45dAf1e8044e918fA786A483A1")

	premiums := []Premium{}
	persist.Load(LOG_FILE, &premiums)

	for {
		rpc := rpcPoolList.Iter()
		logger.Info("dialing...", "rpc", rpc)
		client, err := ethclient.Dial(rpc)
		if err != nil {
			log.Println(err)
			continue
		}

		oracle, err := uniswapv3oracle.NewUniswapv3oracle(oracleAddress, client)
		if err != nil {
			log.Println(err)
			continue
		}

		//
		// 1. Query normalization factor from power perp token controller (oSQTH minter)
		//
		powerTokenCtrl, err := opyn.NewOpyn(powerTokenCtrlAddress, client)
		if err != nil {
			log.Println(err)
			continue
		}
		normalizationFactor, err := powerTokenCtrl.NormalizationFactor(&bind.CallOpts{})
		if err != nil {
			log.Println(err)
			continue
		}
		logger.Info("fetch wPowerPerp state", "normalization_factor", normalizationFactor.String())

		wait()

		//
		// 2. Query ETH twap price using Uniswap v3 oracle
		//
		spotTwapPrice, err := oracle.GetTwap(&bind.CallOpts{}, ethPoolAddress, weth, usdc, 1, false)
		if err != nil {
			log.Println(err)
			continue
		}
		logger.Info("fetch ETH twap price oracle", "twap_price", spotTwapPrice.String())

		wait()

		//
		// Query spot state (USDC/ETH at 0.3%)
		//
		ethPool, err := uniswapv3.NewUniswapv3(ethPoolAddress, client)
		if err != nil {
			log.Println(err)
			continue
		}
		ethPoolSlot0, err := ethPool.Slot0(&bind.CallOpts{})
		if err != nil {
			log.Println(err)
			continue
		}
		underlyingPrice := uniswap_helper.SqrtPriceX96ToPrice(ethPoolSlot0.SqrtPriceX96, false)
		underlyingPrice = underlyingPrice.Mul(decimal.NewFromBigInt(big.NewInt(1), 12)) // 1e18 / 1e6
		logger.Info("fetch ETH price", "pool", ethPoolAddress, "price", underlyingPrice.String())

		wait()

		//
		// Query wPowerPerp state (WETH/oSQTH at 0.3%)
		//
		osqthPool, err := uniswapv3.NewUniswapv3(osqthPoolAddress, client)
		if err != nil {
			log.Println(err)
			continue
		}
		osqthPoolSlot0, err := osqthPool.Slot0(&bind.CallOpts{})
		if err != nil {
			log.Println(err)
			continue
		}
		squaredPrice := uniswap_helper.SqrtPriceX96ToPrice(osqthPoolSlot0.SqrtPriceX96, false)
		squaredPriceUsd := squaredPrice.Mul(underlyingPrice)
		logger.Info("fetch oSQTH price", "pool", osqthPoolAddress, "price_eth", squaredPrice.String(), "price_usd", squaredPriceUsd.String())

		wait()

		// Calculate index price of ETH² using spot twap price
		unscaledIndex := big.NewInt(0)
		unscaledIndex.Mul(spotTwapPrice, spotTwapPrice)
		unscaledIndex.Div(unscaledIndex, ONE)
		indexPrice := unscaledIndex.Div(unscaledIndex, ONE)

		// Calculate mark price (ETH² market price)
		markPrice := squaredPriceUsd.Mul(decimal.NewFromInt(1e4)).Mul(ONE_DEC).Div(decimal.NewFromBigInt(normalizationFactor, 0))

		// Calculate premium
		premium := markPrice.Mul(HUNDRED_DEC).Div(decimal.NewFromBigInt(indexPrice, 0)).Sub(HUNDRED_DEC)
		logger.Info("finalize", "index_price", indexPrice.String(), "mark_price", markPrice.String(), "premium", premium.String())

		// Save to file
		premiums = append(premiums, Premium{
			At:                  time.Now().Unix(),
			Rate:                premium.InexactFloat64(),
			NormalizationFactor: normalizationFactor.Int64(),
			SpotPrice:           underlyingPrice.InexactFloat64(),
			IndexPrice:          decimal.NewFromBigInt(indexPrice, 0).InexactFloat64(),
			MarkPriceEth:        squaredPrice.InexactFloat64(),
			MarkPrice:           markPrice.InexactFloat64(),
		})
		persist.Save(LOG_FILE, premiums)

		time.Sleep(30 * time.Second)
	}

}
