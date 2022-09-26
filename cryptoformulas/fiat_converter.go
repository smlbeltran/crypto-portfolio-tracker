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

type Fiat struct {
	From       string
	To         string
	URL        string
	DOMElement string
}

func (f *Fiat) ConvertFiat() (float64, error) {
	var p float64

	resp, err := http.Get(fmt.Sprintf("%s/%s-%s", f.URL, f.From, f.To))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	docHTML, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	if err != nil {
		return p, fmt.Errorf("[Convert] unable to read document, %w", ErrorFiat)
	}

	p, err = strconv.ParseFloat(strings.TrimSpace(docHTML.Find(f.DOMElement).Text()), 64)
	if err != nil {
		return p, fmt.Errorf("[Convert] unable to find element, %w", ErrorFiat)
	}

	return p, nil
}
