package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/vavilen84/geth_example/client"
	"log"
	"math/big"
	"os"
	"strings"
)

var AppCommands = []*cobra.Command{
	&cobra.Command{
		Use:  "print-balance",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cl := client.InitClient()
			printBalance(cl, args[0], args[0])
		},
	},
	&cobra.Command{
		Use: "print-alice-balance",
		Run: func(cmd *cobra.Command, args []string) {
			cl := client.InitClient()
			printBalance(cl, os.Getenv("ALICE_ADDRESS"), "ALICE")
		},
	},
	&cobra.Command{
		Use: "print-bob-balance",
		Run: func(cmd *cobra.Command, args []string) {
			cl := client.InitClient()
			printBalance(cl, os.Getenv("BOB_ADDRESS"), "BOB")
		},
	},
	&cobra.Command{
		Use: "send-from-alice-to-bob",
		Run: func(cmd *cobra.Command, args []string) {
			cl := client.InitClient()
			privateKey, err := crypto.HexToECDSA(
				strings.TrimPrefix(os.Getenv("ALICE_PRIVATE_KEY"), "0x"),
			)
			if err != nil {
				log.Fatal(err)
			}

			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Fatal("error casting public key to ECDSA")
			}

			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
			nonce, err := cl.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				log.Fatal(err)
			}

			value := big.NewInt(1000000000000000000) // in wei (1 eth)
			gasLimit := uint64(21000)                // in units
			gasPrice, err := cl.SuggestGasPrice(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			toAddress := common.HexToAddress(os.Getenv("BOB_ADDRESS"))
			var data []byte
			tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

			chainID, err := cl.NetworkID(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
			if err != nil {
				log.Fatal(err)
			}

			err = cl.SendTransaction(context.Background(), signedTx)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("tx sent: %s \r\n", signedTx.Hash().Hex())

			printBalance(cl, os.Getenv("BOB_ADDRESS"), "BOB")
			printBalance(cl, os.Getenv("ALICE_ADDRESS"), "ALICE")
		},
	},
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var rootCmd = &cobra.Command{}
	for _, command := range AppCommands {
		rootCmd.AddCommand(command)
	}
	rootCmd.Execute()
}

func printBalance(client *ethclient.Client, address, name string) {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name, convertWeiToETH(balance))
}

func convertWeiToETH(wei *big.Int) string {
	fl := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e18))
	return fl.String()
}
