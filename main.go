package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		fmt.Println(err)
	}
	networkid, err := client.NetworkID(context.Background())
	fmt.Printf("%v", networkid.Text(10))
	if err != nil {
		fmt.Println(err)
	}
	newBigInt := big.NewInt(int64(1))
	fmt.Println(newBigInt.Text(10))
	block, err := client.BlockByNumber(context.Background(), newBigInt)
	if err != nil {
		fmt.Println(err)
	}
	jsonblock, err := json.Marshal(block)
	if err != nil {
		fmt.Println(err)
	}
	myaddress := common.BytesToAddress([]byte("0xca32f313b2b51ff902bb984c4cfff32bad9c2a46"))
	fmt.Println(string(jsonblock))
	balance, _ := client.BalanceAt(context.Background(), myaddress, nil)
	fmt.Printf("%v", balance.Text(10))
	header, err := client.HeaderByNumber(context.Background(), newBigInt)
	if err != nil {
		fmt.Println(err)
	}
	jsonheader, err := json.Marshal(header)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonheader))

}
