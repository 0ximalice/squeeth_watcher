import json
from datetime import datetime

ONE = 1e18

class Tick:
    
    AT = "at"
    RATE = "rate"
    NORMALIZATION_FACTOR = "normalization_factor"
    SPOT_PRICE = "spot_price"
    MARK_PRICE_ETH = "mark_price_eth"

    def __init__(self, json_obj):
        if json_obj is str:
            json_obj = json.loads(json_obj)

        self.at = Tick.have(json_obj, Tick.AT)
        self.rate = Tick.have(json_obj, Tick.RATE)
        self.normalization_factor = Tick.have(json_obj, Tick.NORMALIZATION_FACTOR)
        self.spot_price = Tick.have(json_obj, Tick.SPOT_PRICE)
        self.mark_price_eth =  Tick.have(json_obj, Tick.MARK_PRICE_ETH)

    def have(json, key):
        if key not in json:
            raise Exception(f"Tick must have {key}")
        return json[key]
    
    def datetime(self) -> str:
        return datetime.fromtimestamp(self.at)
    
class Trading:
     
    def __init__(self, amount_in: float, cr_percent: float, gas_price: float, pool_tier: float, slippage: float):
        self.amount_in = float(amount_in) * ONE
        self.cr_ratio = float(cr_percent) / 100
        self.gas_price = float(gas_price) 
        self.trading_fee_bps = float(pool_tier) / 100
        self.slippage_bps = float(slippage) / 100

    def fee(self, amount):
        return amount * self.trading_fee_bps
    
    def slippage(self, amount):
        return amount * self.slippage_bps
    
    def buy(self, amount, price):
        amount_out = amount / price
        slippage = self.slippage(amount_out)
        fee = self.fee(amount_out)
        return (amount_out - fee - slippage, fee, slippage)
    
    def sell(self, amount, price):
        amount_out = amount * price
        slippage = self.slippage(amount_out)
        fee = self.fee(amount_out)
        return (amount_out - fee - slippage, fee, slippage)
          

def calculate_short_pnl(open_on_tick: Tick, close_on_tick: Tick, scale_factor: float, trading: Trading):
    # The exact collateral formula is: 
    # CR = 10000 * (#collateral)/(#oSQTH * normFactor * ethPrice)
    mint_amount = scale_factor * trading.amount_in / (trading.cr_ratio * open_on_tick.normalization_factor * open_on_tick.spot_price)
    (initial_premium, short_swap_fee, short_slippage) = trading.sell(mint_amount, open_on_tick.mark_price_eth)
    (spend, close_swap_fee, close_slippage) = trading.sell(mint_amount, close_on_tick.mark_price_eth)

    gas_cost = trading.gas_price * (450000 * 3 + 350000)
    pnl = initial_premium - spend - gas_cost
    pnl_usd = pnl * close_on_tick.spot_price
    return {
        "Open Short": {
            "At (%)": open_on_tick.rate,
            "Collateral amount (Underlying)": trading.amount_in / ONE,
            "Minted (Perp)": mint_amount,
            "Trading Fee": short_swap_fee,
            "Slippage": short_slippage,
            "Initial Premium": initial_premium,
        },
        "Close": {
            "At (%)": close_on_tick.rate,
            "Trading Fee": close_swap_fee,
            "Slippage": close_slippage,
            "Close Spend (Underlying)": spend,
        },
        "Summary": {
            "Premium Spread (%)": open_on_tick.rate - close_on_tick.rate,
            "Estimated Gas Usage (ETH)": gas_cost,
            "PnL (ETH)": pnl,
            "PnL (USD)": pnl_usd,
            "PnL (%)": pnl * 100 / (trading.amount_in/ONE),
        },
    }

def calculate_long_pnl(open_on_tick: Tick, close_on_tick: Tick, trading: Trading):
    
    amount_in = trading.amount_in / ONE / 2
    (perp_amount, long_fee, long_slippage) = trading.buy(amount_in, open_on_tick.mark_price_eth)
    (usd_amount, hedge_fee, hedge_slippage) = trading.sell(amount_in, open_on_tick.spot_price)

    (closed_amount, close_swap_fee, close_slippage) = trading.sell(perp_amount, close_on_tick.mark_price_eth)
    (hedged_amount, close_hedge_fee, close_hedge_slippage) = trading.buy(usd_amount, close_on_tick.spot_price)

    gas_cost = trading.gas_price * (450000 * 2)
    pnl = (closed_amount + hedged_amount) - (trading.amount_in / ONE) - gas_cost
    pnl_usd = pnl * close_on_tick.spot_price
    return {
        "Open Long": {
            "At (%)": open_on_tick.rate,
            "Amount In (Underlying)": amount_in,
            "Trading Fee": long_fee,
            "Slippage": long_slippage,
            "Long Amount (Perp)": perp_amount,
        },
        "Open Hedging": {
            "Amount In (Underlying)": amount_in,
            "Spot Price": open_on_tick.spot_price,
            "Trading Fee": hedge_fee,
            "Slippage": hedge_slippage,
            "Hedged Amount (USD)": usd_amount,
        },
        "Close Long": {
            "At (%)": close_on_tick.rate,
            "Amount In (Perp)": perp_amount,
            "Trading Fee": close_swap_fee,
            "Slippage": close_slippage,
            "Close Long": closed_amount,
        },
        "Close Hedging": {
            "Amount In (USD)": usd_amount,
            "Spot Price": close_on_tick.spot_price,
            "Trading Fee": close_hedge_fee,
            "Slippage": close_hedge_slippage,
            "Close Long": hedged_amount,
        },
        "Summary": {
            "Premium Spread (%)": close_on_tick.rate - open_on_tick.rate,
            "Estimated Gas Usage (ETH)": gas_cost,
            "PnL (ETH)": pnl,
            "PnL (USD)": pnl_usd,
            "PnL (%)": pnl * 100 / (trading.amount_in/ONE),
        },
    }