package cryptoformulas

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFiatConvertsCurrency(t *testing.T) {
	var h http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<body>
					<div class="YMlKec fxKbKc">0.9985</div>
				</body>
			</html>
		`))
	}

	s := httptest.NewServer(h)

	fiat := Fiat{
		From:       "USD",
		To:         "GBP",
		URL:        s.URL,
		DOMElement: ".YMlKec",
	}

	want, _ := fiat.ConvertFiat()
	got := 0.9985

	if got != want {
		t.Errorf("got:%.2f, want: %.2f", got, want)
	}
}

func TestFiatCannotConvert(t *testing.T) {
	var h http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<body>
					<div class="">0.9985</div>
				</body>
			</html>
		`))
	}

	s := httptest.NewServer(h)

	fiat := Fiat{
		From: "USD",
		To:   "GBP",
		URL:  s.URL,
	}

	_, err := fiat.ConvertFiat()

	want := true
	got := errors.Is(err, ErrorFiat)

	if want != got {
		t.Errorf("want: %t, got: %t", want, got)
	}
}
