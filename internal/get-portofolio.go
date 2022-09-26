package internal

import (
	"log"

	"github.com/smlbeltran/crypto-portfolio-tracker/domain"
)

type Exchange interface {
	GetCurrentValue(url, domElement string) (float64, error)
	CurrentProfit(totalCoins, currentCryptoPrice float64) float64
}

func GetPortfolioData(cfg Config, ex Exchange) ([]*domain.PorfolioResult, error) {
	var results []*domain.PorfolioResult

	for _, v := range cfg.Coins {
		cryv, err := ex.GetCurrentValue(v.CryptoWebsite, v.CryptoDomElement)
		if err != nil {
			log.Fatal(err)
		}

		rewards := ex.CurrentProfit(v.Owned, cryv)

		pf := &domain.PorfolioResult{
			Name:   v.Name,
			Owned:  v.Owned,
			Fiat:   v.ConversionTo,
			Price:  cryv,
			Reward: rewards,
		}

		results = append(results, pf)

	}

	return results, nil
}
