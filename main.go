package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx := context.Background()

	client, _ := ethclient.Dial("https://rpc.xdaichain.com/")
	txHash := common.HexToHash("0xd5e104f67d69db3c18a8a43e378f77f5d279e53851ea2d5e19c577032b95079b")

	receipt, _ := client.TransactionReceipt(ctx, txHash)
	blockHeader, _ := client.HeaderByNumber(ctx, receipt.BlockNumber)
	blockByNumber, _ := client.BlockByNumber(ctx, receipt.BlockNumber)

	fmt.Printf("Transaction block number: %d\n", receipt.BlockNumber)
	fmt.Printf("Transaction block hash: %x\n", receipt.BlockHash)
	fmt.Println()

	fmt.Printf("Block number: %d\n", blockHeader.Number)
	fmt.Printf("Block hash: %x\n", blockHeader.Hash())
	fmt.Println()
	fmt.Printf("Block number (whole block): %d\n", blockByNumber.Number())
	fmt.Printf("Block hash (whole block): %x\n", blockByNumber.Hash())

	fmt.Println()
}
