package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	client2 "github.com/vavilen84/geth_example/client"
	"log"
)

var AppCommands = []*cobra.Command{
	&cobra.Command{
		Use:  "print-balance",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			client := client2.InitClient()
			account := common.HexToAddress(args[0])
			balance, err := client.BalanceAt(context.Background(), account, nil)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(balance)
		},
	},
}

func main() {
	var rootCmd = &cobra.Command{}
	for _, command := range AppCommands {
		rootCmd.AddCommand(command)
	}
	rootCmd.Execute()
}
