package client

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func InitClient() *ethclient.Client {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	return client
}
