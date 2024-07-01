package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	abigen "github.com/w3gop2p/GethSoliditySC/abigen"
	"log"
	"math/big"
)

func main() {
	//DeployStorage()
	StoreRetrieve()
}

// This main is deploying the contract
func DeployStorage() {
	privateKeyHex := "c5114526e042343c6d1899cad05e1c00ba588314de9b96929914ee0df18d46b2"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}
	// Get the public address of the private key
	//	publicKey := privateKey.Public()
	// Create an IPC-based RPC connection to a remote node and an authorized transactor
	conn, err := ethclient.Dial("http://127.0.0.1:49500")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Retrieve the current chain ID
	chainID, err := conn.ChainID(context.Background())
	if err != nil {
		log.Fatal("Failed to retrieve chain ID: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	// Deploy the contract passing the newly created `auth` and `conn` vars
	address, tx, instance, err := abigen.DeployStorage(auth, conn)
	if err != nil {
		log.Fatalf("Failed to deploy new storage contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	// Wait for the transfer transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		log.Fatalf("Failed to wait for contract deployment mining: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Contract deployment failed with status: %v", receipt.Status)
	}

	fmt.Println("Contract Deployment successful")
	// function call on `instance`. Retrieves pending name
	name, err := instance.StorageCaller.Retrieve(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	fmt.Println("Pending name:", name)
}

// This func is for actions after deployment of the contract, set the value, and read the value

func StoreRetrieve() {
	conn, err := ethclient.Dial("http://127.0.0.1:49500")
	privateKeyHex := "c5114526e042343c6d1899cad05e1c00ba588314de9b96929914ee0df18d46b2"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// Instantiate the contract and display its name
	store, err := abigen.NewStorage(common.HexToAddress("0x86a4f08103eade88ade0be164d7bda8fe166680d"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
	}
	chainID, err := conn.ChainID(context.Background())
	if err != nil {
		log.Fatal("Failed to retrieve chain ID: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	tx, err := store.Store(auth, big.NewInt(1023))
	if err != nil {
		log.Fatalf("Failed to store: %v, tx: %v", err, tx.Hash().Hex())
	}
	// Wait for the transfer transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		log.Fatalf("Failed to wait for store value on contract: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Store value on contract failed with status: %v", receipt.Status)
	}

	fmt.Println("Store value successful")

	value, err := store.Retrieve(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve value: %v", err)
	}
	fmt.Println("Value: ", value)
}
