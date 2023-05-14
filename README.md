# Ethereum Transaction Monitor

## Overview

The Ethereum Transaction Monitor is a Go application that connects to an Ethereum node and periodically calculates the total number of transactions that have been processed within the last minute. 

## How it works

The application works in two main steps:

1. **Fetching blocks and recording transactions:** The application fetches the most recent block from the Ethereum node every 10 seconds. It then counts the number of transactions in that block and stores this data along with the current timestamp.

2. **Calculating transactions per minute:** Every minute, the application sums the number of transactions from all blocks received within the last minute and logs this total.

The program is designed to be run continuously as a background process, providing an ongoing calculation of transaction volume.

## Prerequisites

Before running this application, ensure that you have:

- The Go programming language [installed](https://golang.org/dl/).
- Go Ethereum and godotenv packages installed. You can install these packages with the following commands:

  ```bash
  go get github.com/ethereum/go-ethereum
  go get github.com/joho/godotenv

## Setup

1. **Environment Variables**: Create a `.env` file in the root directory of the project and add the following environment variable:

    ```env
    ETHEREUM_NODE_URL=http://localhost:8545
    ```

    Replace `http://localhost:8545` with the URL of your Ethereum node.

## Running the Application

1. Open your terminal.
2. Navigate to the root directory of the project.
3. Run `go run bc-number-transactions.go`.

The application will now start and monitor the Ethereum node. It logs the total number of transactions in the last minute every minute.
