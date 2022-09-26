package cryptoformulas

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type CryptoCurrency struct{}

func (cy *CryptoCurrency) GetCurrentValue(url, domElement string) (float64, error) {
	var v float64

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	docHTML, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	if err != nil {
		return v, fmt.Errorf("[GetCurrentValue] unable to read document, %w", ErrorFiat)
	}

	price := strings.ReplaceAll(strings.TrimSpace(docHTML.Find(domElement).Text()), "$", "")

	v, err = strconv.ParseFloat(price, 64)
	if err != nil {
		return v, fmt.Errorf("[GetCurrentValue] unable to find element, %w", ErrorFiat)
	}

	return v, nil
}

func (cy *CryptoCurrency) CurrentProfit(totalCoins, currentCryptoPrice float64) float64 {
	return totalCoins * currentCryptoPrice
}
