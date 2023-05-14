package main

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

type BlockData struct {
	timeStamp time.Time
	txCount   int
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ethereumNodeURL := os.Getenv("ETHEREUM_NODE_URL")
	if ethereumNodeURL == "" {
		log.Fatal("ETHEREUM_NODE_URL not set in .env file")
	}

	client, err := ethclient.Dial(ethereumNodeURL)
	if err != nil {
		log.Fatal(err)
	}

	var (
		mu        sync.Mutex
		blockData []BlockData
	)

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		for _ = range ticker.C {
			header, err := client.HeaderByNumber(context.Background(), nil)
			if err != nil {
				log.Fatal(err)
			}

			block, err := client.BlockByNumber(context.Background(), header.Number)
			if err != nil {
				log.Fatal(err)
			}

			mu.Lock()
			blockData = append(blockData, BlockData{
				timeStamp: time.Now(),
				txCount:   len(block.Transactions()),
			})
			mu.Unlock()
		}
	}()

	ticker := time.NewTicker(1 * time.Minute)
	for _ = range ticker.C {
		mu.Lock()
		var totalTx int
		cutoff := time.Now().Add(-1 * time.Minute)
		filteredBlockData := blockData[:0]
		for _, bd := range blockData {
			if bd.timeStamp.Before(cutoff) {
				continue
			}
			filteredBlockData = append(filteredBlockData, bd)
			totalTx += bd.txCount
		}
		blockData = filteredBlockData
		mu.Unlock()

		log.Printf("Number of transactions in the last minute: %d\n", totalTx)
	}
}
