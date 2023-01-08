package main

import "encoding/json"

type Asset struct {
	Aclass          string      `json:"aclass"`
	Altname         string      `json:"altname"`
	Decimals        int         `json:"decimals"`
	DisplayDecimals int         `json:"display_decimals"`
	CollateralValue json.Number `json:"collateral_value"`
	Status          string      `json:"status"`
}

type Pair struct {
	Altname             string          `json:"altname"`
	Wsname              string          `json:"wsname"`
	Aclass_base         string          `json:"aclass_base"`
	Base                string          `json:"base"`
	Aclass_quote        string          `json:"aclass_quote"`
	Quote               string          `json:"quote"`
	Lot                 string          `json:"lot"`
	Pair_decimals       int             `json:"pair_decimals"`
	Cost_decimals       int             `json:"cost_decimals"`
	Lot_decimals        int             `json:"lot_decimals"`
	Lot_multiplier      int             `json:"lot_multiplier"`
	Leverage_buy        []int           `json:"leverage_buy"`
	Leverage_sell       []int           `json:"leverage_sell"`
	Fees                [][]json.Number `json:"fees"`
	Fees_maker          [][]json.Number `json:"fees_maker"`
	Fee_volume_currency string          `json:"fee_volume_currency"`
	Margin_call         int             `json:"margin_call"`
	Margin_stop         int             `json:"margin_stop"`
	Ordermin            string          `json:"ordermin"`
	Costmin             string          `json:"costmin"`
	Tick_size           string          `json:"tick_size"`
	Status              string          `json:"status"`
}

type Assets struct {
	Result map[string]Asset `json:"result"`
}

type AssetPairs struct {
	Result map[string]Pair `json:"result"`
}

type Time struct {
	Result struct {
		Unixtime uint   `json:"unixtime"`
		Rfc1123  string `json:"rfc1123"`
	} `json:"result"`
}

type Status struct {
	Result struct {
		Status    string `json:"status"`
		Timestamp string `json:"timestamp"`
	} `json:"result"`
}
