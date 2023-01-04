package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/smlbeltran/crypto-portfolio-tracker/domain"
)

type Currency string

const (
	Dollar Currency = "$"
	GBP Currency = "£"
)

func Render(w io.Writer, d []*domain.PorfolioResult) error {
	t := table.NewWriter()
	t.SetOutputMirror(w)
	t.AppendHeader(table.Row{"Cryptocurrency", "Quantity", "FIAT", "Price", "Reward"})

	var totalReward float64

	for _, coin := range d {
		t.AppendRow([]interface{}{coin.Name, coin.Owned, coin.Fiat, fmt.Sprintf("%s%.2f", Dollar, coin.Price), fmt.Sprintf("%s%.2f", Dollar, coin.Reward)})
		t.AppendSeparator()

		totalReward += coin.Reward
	}

	resp, err := http.Get("https://www.currency.me.uk/charts-fetch.php?c1=USD&c2=GBP&t=1")
	if err != nil {
		return fmt.Errorf("unable to retrieve currrency")
	}	
	
	defer resp.Body.Close()

	r, _ := io.ReadAll(resp.Body)

	var currencyVal []map[string]interface{}

	json.Unmarshal(r, &currencyVal)
	
	v, ok := currencyVal[0]["y"].(float64)
	if !ok {
		panic("not available")
	}

	t.AppendFooter(table.Row{"Total", fmt.Sprintf("%s%.2f", Dollar, totalReward), "", ""})
	t.AppendFooter(table.Row{"Conversion", fmt.Sprintf("%s%.2f", GBP, totalReward * v ), "", ""})

	t.Render()

	return nil
}
