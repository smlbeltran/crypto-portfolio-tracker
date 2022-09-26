package cryptoformulas

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCurrentValueForACryptoCurrency(t *testing.T) {
	var h http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
		<html>
			<body>
				<div>XRP</div>
				<div class="YMlKec">1.00</div>
			</body>
		</html>
		`))
	}

	server := httptest.NewServer(h)

	cy := CryptoCurrency{}

	got, _ := cy.GetCurrentValue(server.URL, ".YMlKec")
	want := 1.00

	if want != got {
		t.Errorf("want: %.2f, got: %.2f", want, got)
	}
}
