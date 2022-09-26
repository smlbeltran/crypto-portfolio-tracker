package internal

import (
	"reflect"
	"testing"

	"github.com/smlbeltran/crypto-portfolio-tracker/domain"
)

type InMemoryExchange struct{}

func (m *InMemoryExchange) GetCurrentValue(url, domElement string) (float64, error) { return 0.40, nil }

func (m *InMemoryExchange) CurrentProfit(totalCoins, currentCryptoPrice float64) float64 {
	return 1000.00
}

func TestMapFileToPorFolioStruct(t *testing.T) {
	t.Run("maps only a single crypto currency", func(t *testing.T) {
		portfolio := Config{
			[]Coins{
				{"XRP", 1000, "somefiat.com", "classname", "GBP", "USD", "example.com", "classname"},
			},
		}

		pf, _ := GetPortfolioData(portfolio, &InMemoryExchange{})

		got := *pf[0]
		want := domain.PorfolioResult{Name: "XRP", Owned: 1000, Fiat: "USD", Price: 0.40, Reward: 1000.00}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("maps a list of crypto currency", func(t *testing.T) {
		portfolio := Config{
			[]Coins{
				{"XRP", 1000, "somefiat.com", "classname", "GBP", "USD", "example.com", "classname"},
				{"BTC", 2000, "somefiat.com", "classname", "GBP", "USD", "example.com", "classname"},
			},
		}

		pf, _ := GetPortfolioData(portfolio, &InMemoryExchange{})

		got := pf
		want := []*domain.PorfolioResult{
			{Name: "XRP", Owned: 1000, Fiat: "USD", Price: 0.40, Reward: 1000.00},
			{Name: "BTC", Owned: 2000, Fiat: "USD", Price: 0.40, Reward: 1000.00},
		}

		got = []*domain.PorfolioResult{
			got[0],
			got[1],
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
