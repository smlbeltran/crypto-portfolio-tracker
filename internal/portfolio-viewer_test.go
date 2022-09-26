package internal

import (
	"bytes"
	"testing"

	"github.com/smlbeltran/crypto-portfolio-tracker/domain"
	"github.com/stretchr/testify/assert"
)

func TestRenderCryptoPortfolio(t *testing.T) {
	t.Run("converts a set of crypto into a information row", func(t *testing.T) {
		crypto := []*domain.PorfolioResult{
			{
				Name:   "XRP",
				Owned:  100,
				Fiat:   "USD",
				Price:  0.41,
				Reward: 1000,
			},
			{
				Name:   "XLM",
				Owned:  100,
				Fiat:   "USD",
				Price:  0.41,
				Reward: 1000,
			},
		}

		var buf bytes.Buffer

		err := Render(&buf, crypto)
		if err != nil {
			t.Fatal(err)
		}
		got := buf.String()
		want := `+----------------+----------+------+-------+----------+
| CRYPTOCURRENCY | QUANTITY | FIAT | PRICE | REWARD   |
+----------------+----------+------+-------+----------+
| XRP            |      100 | USD  | $0.41 | $1000.00 |
+----------------+----------+------+-------+----------+
| XLM            |      100 | USD  | $0.41 | $1000.00 |
+----------------+----------+------+-------+----------+
| TOTAL          | $2000.00 |      |       |          |
+----------------+----------+------+-------+----------+
`

		assert.Equal(t, want, got)
	})
}
