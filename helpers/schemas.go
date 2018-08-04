package helpers

import (
	"encoding/json"
)

// Ad represents an Localbitcoin's Ad buying BTC
type Ad struct {
	Data struct {
		Message   	string `json:"msg"`
		BankName  	string `json:"bank_name"`
		Price     	string `json:"temp_price"`
		MinAmount 	string `json:"min_amount"`
		MaxAmount 	string `json:"max_amount"`
		City 		string `json:"city"`
		Profile   struct {
			Username string `json:"username"`
		} `json:"profile"`
	} `json:"data"`
	Actions struct {
		URL string `json:"public_view"`
	} `json:"actions"`
}

// LBTCResponse represents a response from Localbitcoin's API
type LBTCResponse struct {
	Data struct {
		Ads []Ad `json:"ad_list"`
	} `json:"data"`
	Pagination struct {
		Next string `json:"next"`
	} `json:"pagination"`
}

type CountriesResponse struct {
	Data struct {
		Count int      `json:"cc_count"`
		List  []string `json:"cc_list"`
	} `json:"data"`
}

//TODO: automate later

type LBCurrency struct {
	Name    string `json:"name"`
	Altcoin bool   `json:"altcoin"`
}

type LBCurrencyWrapper struct {
	LBCurrency `json:”-”`
}

func (w *LBCurrencyWrapper) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &w.LBCurrency)
}

type CurrencyResponse struct {
	Data struct {
		Currencies []LBCurrency `json:"cur_list"`
	} `json:"data"`
}


type CurrencyResponse1 struct {
	Data struct {
		Currencies struct {
			items []LBCurrency
		} `json:"currencies"`
		CurrencyCount int `json:"currency_count"`
	} `json:"data"`
}
