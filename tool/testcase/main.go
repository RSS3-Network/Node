package main

import (
	"embed"
	"fmt"
	"log"
	"math/big"
	"os"
	"text/template"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/endpoint"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

//go:embed template/*.tmpl
var embedFS embed.FS

var command = &cobra.Command{
	Use:           "mock",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var (
			task engine.Task
			tmpl = template.New("")

			network  = lo.Must(cmd.PersistentFlags().GetString("network"))
			endpoint = lo.Must(cmd.PersistentFlags().GetString("endpoint"))
			feed     = lo.Must(cmd.PersistentFlags().GetString("feed"))
		)

		switch network {
		case filter.NetworkEthereum.String():
			tmpl.Funcs(template.FuncMap{
				"BytesToHex": func(value []byte) string {
					return hexutil.Encode(value)
				},
				"BigIntToHex": func(value *big.Int) string {
					return (*hexutil.Big)(value).String()
				},
			})

			chainID, err := filter.EthereumChainIDString(network)
			if err != nil {
				return fmt.Errorf("unsupported ethereum chain: %w", err)
			}

			ethereumClient, err := ethereum.Dial(cmd.Context(), endpoint)
			if err != nil {
				return fmt.Errorf("dial to endpoint: %w", err)
			}

			transaction, err := ethereumClient.TransactionByHash(cmd.Context(), common.HexToHash(feed))
			if err != nil {
				return fmt.Errorf("get transacion by hash: %w", err)
			}

			receipt, err := ethereumClient.TransactionReceipt(cmd.Context(), transaction.Hash)
			if err != nil {
				return fmt.Errorf("get receipt by hash: %w", err)
			}

			header, err := ethereumClient.HeaderByHash(cmd.Context(), receipt.BlockHash)
			if err != nil {
				return fmt.Errorf("get block by hash: %w", err)
			}

			task = &source.Task{
				ChainID:     uint64(chainID),
				Header:      header,
				Transaction: transaction,
				Receipt:     receipt,
			}
		default:
			return fmt.Errorf("unsupport network %s", network)
		}

		if tmpl, err = tmpl.ParseFS(embedFS, "template/*.tmpl"); err != nil {
			return fmt.Errorf("parse templates: %w", err)
		}

		if err := tmpl.ExecuteTemplate(os.Stdout, network, task); err != nil {
			return fmt.Errorf("execute template: %w", err)
		}

		return nil
	},
}

func init() {
	command.PersistentFlags().String("network", filter.NetworkEthereum.String(), "")
	command.PersistentFlags().String("endpoint", endpoint.MustGet(filter.NetworkEthereum), "")
	command.PersistentFlags().String("feed", "0xf74008a8fde35012c5bc9c897c1d413fe0befbc9e6fc9b6d8bfab38b7dd3c6bd", "")
}

func main() {
	if err := command.Execute(); err != nil {
		log.Fatalf("execute command: %s", err)
	}
}
