package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"

	helloworld "ethereumdocker/helloworld"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
	fmt.Println(string(jsonblock))
	myaddress := common.HexToAddress("0xca32f313b2b51ff902bb984c4cfff32bad9c2a46")
	fmt.Println(myaddress)
	balance, _ := client.BalanceAt(context.Background(), myaddress, nil)
	fmt.Printf("\n balance is %v \n", balance.Text(10))
	header, err := client.HeaderByNumber(context.Background(), newBigInt)
	if err != nil {
		fmt.Println(err)
	}
	jsonheader, err := json.Marshal(header)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonheader))

	address := common.HexToAddress("0x81361134d039b7c2783f32fbc01864c39dd0befd")
	instance, err := helloworld.NewHelloworld(address, client)
	value, _ := instance.Get(nil)
	fmt.Println("===============================")
	fmt.Println(value)
	fmt.Println("===============================")

	privateKey, err := crypto.GenerateKey()
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	// in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	_, _ = instance.Set(auth, "Prueba 2 desde Go")

}
