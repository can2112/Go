package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func eth_main() {
	conn := connectToRPC(POLYGON_RPC_URL)
	gasPrice := fetchGasPrice(conn)
	fmt.Println(gasPrice.String())
	approve(conn, SENDER, IINCH, USDC, big.NewInt(1000000))
}

func connectToRPC(url string) *ethclient.Client {
	conn, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}
	return conn
}

func approve(conn *ethclient.Client, from string, spender string, tokenAddress string, amount *big.Int) {
	privateKey, _ := crypto.HexToECDSA(PRIVATE_KEY)
	address := common.HexToAddress(tokenAddress)
	tokenErc, err := NewERC20(address, conn)
	if err != nil {
		panic(err)
	}
	gasPrice := fetchGasPrice(conn)
	chainId, err := conn.ChainID(context.TODO())
	if err != nil {
		panic(err)
	}
	options := &bind.TransactOpts{
		From:     common.HexToAddress(from),
		GasPrice: gasPrice,
		GasLimit: 70000,
		Signer: func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(t, types.NewEIP155Signer(chainId), privateKey)
		},
	}

	txn, err := tokenErc.Approve(options, common.HexToAddress(spender), amount)
	if err != nil {
		panic(err)
	}
	fmt.Println(txn.Hash().Hex())
}

func fetchGasPrice(conn *ethclient.Client) *big.Int {
	gasPrice, err := conn.SuggestGasPrice(context.TODO())
	if err != nil {
		panic(err)
	}
	return gasPrice
}
