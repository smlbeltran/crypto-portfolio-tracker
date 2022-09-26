package cmd

import (
	"os"
	"strings"

	"github.com/smlbeltran/crypto-portfolio-tracker/cryptoformulas"
	"github.com/smlbeltran/crypto-portfolio-tracker/internal"
	"github.com/spf13/cobra"
)

var DirPath string

var portfolioCmd = &cobra.Command{
	Use:   "portfolio",
	Short: "Display currenty portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		if len(strings.TrimSpace(DirPath)) == 0 {
			panic("no directory specified")
		}

		fs := os.DirFS(DirPath)

		cfg := internal.Config{}

		cr := cfg.GetConfig(fs)

		ex := &cryptoformulas.CryptoCurrency{}

		rs, err := internal.GetPortfolioData(*cr, ex)
		if err != nil {
			panic(err)
		}

		if internal.Render(os.Stdout, rs); err != nil {
			panic(err)
		}
	},
}

func init() {
	portfolioCmd.Flags().StringVarP(&DirPath, "dir-path", "d", "", "location for mycrypto.txt")
	rootCmd.AddCommand(portfolioCmd)
}
