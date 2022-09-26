package internal

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestGetConfigFromSpecificFile(t *testing.T) {
	res := fstest.MapFS{
		"mycrypto.json": &fstest.MapFile{
			Data: []byte(`{
"coins":[
	{
		"name":"XRP",
		"owned": 1000,
		"fiat_website":"somefiat.com",
		"fiat_dom_element": "classname",
		"conversion_from": "GBP",
		"conversion_to": "USD",
		"crypto_website":"example.com",
		"crypto_dom_element": "classname"
	},
	{
		"name":"BTC",
		"owned": 1000,
		"fiat_website":"somefiat.com",
		"fiat_dom_element": "classname",
		"conversion_from": "GBP",
		"conversion_to": "USD",
		"crypto_website":"example.com",
		"crypto_dom_element": "classname"
	}
	]
}`),
		},
	}

	cfg := Config{}
	got := *cfg.GetConfig(res)

	want := Config{
		[]Coins{
			{"XRP", 1000, "somefiat.com", "classname", "GBP", "USD", "example.com", "classname"},
			{"BTC", 1000, "somefiat.com", "classname", "GBP", "USD", "example.com", "classname"},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
