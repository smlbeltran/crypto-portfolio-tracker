package internal

import (
	"fmt"
	"io"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/smlbeltran/crypto-portfolio-tracker/domain"
)

type Currency string

const (
	Dollar Currency = "$"
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

	t.AppendFooter(table.Row{"Total", fmt.Sprintf("%s%.2f", Dollar, totalReward), "", ""})

	t.Render()

	return nil
}
