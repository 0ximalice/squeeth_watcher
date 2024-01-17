import streamlit as st
import json
from vault import Vault
from calculator import calculate_short_pnl, calculate_long_pnl, Trading, Tick

vaults = {
    "ETH²": Vault("ETH²", 1e4),
}

st.markdown("# Squeeth: Position Simulator")

with st.sidebar:
    st.markdown("## Global Settings")

    AMOUNT_IN = st.text_input('Amount In (Ether)', value='2')
    CR_RATIO = st.text_input('Collateral Ratio (%)', value='200')

    with st.expander("Swap Parameters", expanded=True):
        SWAP_TIER = st.text_input('Pool Tier (%)', value='0.3')
        SLIPPAGE = st.text_input('Slippage (%)', value='0.25')
    
    with st.expander("Gas Parameters", expanded=True):
        GAS_TIER_STR = ["Slow", "Average", "Fast"]
        GAS_TIER_VALUES = [0.00000001, 0.00000003, 0.00000006]
        GAS_TIER = st.sidebar.radio('Gas Price', GAS_TIER_STR)

# Select action type
action_type = st.radio('Action', ['Long', 'Short'])

# Select vault type
vault_type = st.selectbox("Vault", list(vaults.keys()))
vault_info = vaults[vault_type]

# Load tick data
open_at_tick = None
close_at_tick = None
with open("/app/data/"+vault_type+".json", "r") as f:
    ticks = [Tick(tick) for tick in json.load(f)]

    # TODO: Sampling same rate from the data

    tick_keys = ["{:.2f}% ({})".format(tick.rate, tick.datetime()) for tick in ticks]
    st.markdown("Loaded {} ticks from {}.json".format(len(ticks), vault_type))

    col1, col2 = st.columns(2)
    with col1:
        open_at_tick = st.selectbox(
            "Select rate at tick do you want to open/close the position",
            tick_keys,
            placeholder="Select open tick",
        )
    with col2:
        close_at_tick = st.selectbox(" ",
            tick_keys,
            placeholder="Select close tick",
        )

# Calculate PnL
if open_at_tick is not None and close_at_tick is not None:

    open_tick = ticks[tick_keys.index(open_at_tick)]
    close_tick = ticks[tick_keys.index(close_at_tick)]

    if close_tick.at <= open_tick.at:
        st.error('Close tick must be after open tick', icon="🚨")
        st.stop()
        
    st.markdown("#### Simulation Result")
    st.markdown("Earn premiums for selling ETH collateralized squeeth")

    gas_price = GAS_TIER_VALUES[GAS_TIER_STR.index(GAS_TIER)]
    trading = Trading(AMOUNT_IN, CR_RATIO, gas_price, SWAP_TIER, SLIPPAGE)
    if action_type == "Short":
        trade_sections = calculate_short_pnl(open_tick, close_tick, vault_info.scale_factor, trading)
        st.json(trade_sections)
    elif action_type == "Long":
        trade_sections = calculate_long_pnl(open_tick, close_tick, trading)
        st.json(trade_sections)
        


