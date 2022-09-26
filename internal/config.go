package internal

import (
	"bytes"
	"encoding/json"
	"io/fs"
)

type Config struct {
	Coins []Coins `json:"coins"`
}

type Coins struct {
	Name             string  `json:"name"`
	Owned            float64 `json:"owned"`
	FiatWebsite      string  `json:"fiat_website"`
	FiatDomElement   string  `json:"fiat_dom_element"`
	ConversionFrom   string  `json:"conversion_from"`
	ConversionTo     string  `json:"conversion_to"`
	CryptoWebsite    string  `json:"crypto_website"`
	CryptoDomElement string  `json:"crypto_dom_element"`
}

func (c *Config) GetConfig(fileSystem fs.FS) *Config {
	var config Config
	var fileName string

	dir, _ := fs.ReadDir(fileSystem, ".")

	for _, d := range dir {
		if d.Name() != "mycrypto.json" {
			continue
		}
		fileName = d.Name()
		break
	}

	var buf bytes.Buffer

	b, err := fs.ReadFile(fileSystem, fileName)
	if err != nil {
		panic(err)
	}

	buf.Write(b)

	err = json.NewDecoder(&buf).Decode(&config)
	if err != nil {
		panic(err)
	}

	return &config
}
