package main

import (
	"embed"
	"fmt"
	"log"
	"math/big"
	"os"
	"text/template"
	"unicode"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/internal/engine"
	sourceethereum "github.com/rss3-network/node/internal/engine/source/ethereum"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	networkx "github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

//go:embed template/*.tmpl
var embedFS embed.FS

var command = &cobra.Command{
	Use:           "mock",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, _ []string) (err error) {
		var (
			task engine.Task
			tmpl = template.New("")

			source   = networkx.Source(lo.Must(cmd.PersistentFlags().GetString("source")))
			endpoint = lo.Must(cmd.PersistentFlags().GetString("endpoint"))
			activity = lo.Must(cmd.PersistentFlags().GetString("activity"))
		)

		switch source {
		case networkx.EthereumSource:
			tmpl.Funcs(template.FuncMap{
				"BytesToHex": hexutil.Encode,
				"BigIntToHex": func(value *big.Int) string {
					return (*hexutil.Big)(value).String()
				},
				"UppercaseFirst": func(network networkx.Network) string {
					a := []rune(network.String())
					a[0] = unicode.ToUpper(a[0])
					return string(a)
				},
			})

			ethereumClient, err := ethereum.Dial(cmd.Context(), endpoint)
			if err != nil {
				return fmt.Errorf("dial to endpoint: %w", err)
			}

			chainID, err := ethereumClient.ChainID(cmd.Context())
			if err != nil {
				return fmt.Errorf("get chain id: %w", err)
			}

			chain := networkx.EthereumChainID(chainID.Uint64())
			if !chain.IsAEthereumChainID() {
				return fmt.Errorf("unsupported chain id %d", chainID.Uint64())
			}

			network, err := networkx.NetworkString(chain.String())
			if err != nil {
				return fmt.Errorf("get network: %w", err)
			}

			transaction, err := ethereumClient.TransactionByHash(cmd.Context(), common.HexToHash(activity))
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

			task = &sourceethereum.Task{
				Network:     network,
				ChainID:     chainID.Uint64(),
				Header:      header,
				Transaction: transaction,
				Receipt:     receipt,
			}
		default:
			return fmt.Errorf("unsupport source %s", source)
		}

		if tmpl, err = tmpl.ParseFS(embedFS, "template/*.tmpl"); err != nil {
			return fmt.Errorf("parse templates: %w", err)
		}

		if err := tmpl.ExecuteTemplate(os.Stdout, string(source), task); err != nil {
			return fmt.Errorf("execute template: %w", err)
		}

		return nil
	},
}

func init() {
	command.PersistentFlags().String("source", string(networkx.EthereumSource), "")
	command.PersistentFlags().String("endpoint", endpoint.MustGet(networkx.Ethereum), "")
	command.PersistentFlags().String("activity", "0x1f0c0bd550c111d76d3dfca67616f6f7968d10c673de1ad391c5141fdb336b97", "")
}

func main() {
	if err := command.Execute(); err != nil {
		log.Fatalf("execute command: %s", err)
	}
}
