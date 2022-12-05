package main

import (
	utils "code/main/utils"
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"

	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	store "code/src"
)

func main() {
	client, err := ethclient.Dial(utils.GoerliWSS)
	if err != nil {
		log.Fatal("Dial err:", err)
	}

	// 1. 准备合约地址
	contractAddr := "0xe4a220e0bd37673a90e2114abc98e4a22445c32e"
	address := common.HexToAddress(contractAddr)

	// 2. 构造过滤查询条件
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
		FromBlock: new(big.Int).SetUint64(8059333),
		ToBlock:   new(big.Int).SetUint64(8059380),
	}

	// 3. 所有匹配的事件日志将存储在logs中
	// sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal("FilterLogs err:", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println("find new event:")
		fmt.Println("\tblock hash:", vLog.BlockHash.Hex())
		fmt.Println("\tblock num :", vLog.BlockNumber)
		fmt.Println("\ttx   hash :", vLog.TxHash.Hex())

		// 解析Event中的Data
		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}

		// 注意这里，与原文不一样，原文无法编译
		// err := contractAbi.Unpack(&event, "ItemSet", vLog.Data)
		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\tData:")
		fmt.Printf("\t\tkey  :%x\n", event.Key)
		fmt.Printf("\t\tvalue:%x\n", event.Value)

		// 解析Event中的Topic
		var topics [4]string
		fmt.Println("\tTopic:")
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
			fmt.Printf("\t\ttopic[%d]: %s\n", i, topics[i])
		}
	}

	eventSignature := []byte("ItemSet(address,bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println("topic hash:", hash.Hex())
}
