package main

import (
	"context"
	"github.com/NethermindEth/juno/configs"
	"github.com/NethermindEth/juno/internal/log"
	"github.com/NethermindEth/juno/pkg/db"
	"github.com/tarrencev/go-starknet/provider"
)

var logger = log.GetLogger()

func main() {
	logger.Info("Starting Juno, StarkNet Go Client")
	baseURL := configs.MainnetGateway
	prv := provider.NewProvider(baseURL)
	// opt := provider.BlockOptions{}
	ctx := context.Background()
	block, err := prv.Block(ctx, nil)
	if err != nil {
		logger.With("With Error", err).Error("Failed to retrieve block")
	}
	logger.With("blockHash", block.BlockHash).Debug("Block Hash retrieved from provider, ")

	database := db.NewKeyValueDatabase("example", 0)

	val, err := database.Get([]byte("blockHash"))
	if err != nil {
		logger.With("Error", err).Error("Error getting values from")
		return
	}
	logger.With("blockHash", string(val)).Info("Got latest BlockHash Value from DB")

	if block.BlockHash == string(val) {
		logger.Info("Still the same blockHash")
		return
	}
	logger.Info("Storing the new blockHash")

	err = database.Put([]byte("blockHash"), []byte(block.BlockHash))
	if err != nil {
		logger.With("Error", err).Error("Error putting in database")
		return
	}
	n, err := database.NumberOfItems()
	if err != nil {
		logger.With("Error", err).Error("Error getting the number of items in database")
		return
	}
	logger.With("Number of Items on DB", n).Info("Got the number of items on DB")
}
