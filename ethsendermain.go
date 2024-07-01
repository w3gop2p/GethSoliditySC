package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	abigen "github.com/w3gop2p/GethSoliditySC/abigen"
	"log"
	"math/big"
	"time"
)

// func for deployment of the contract
func deployTheContract() {
	privateKeyHex := "c5114526e042343c6d1899cad05e1c00ba588314de9b96929914ee0df18d46b2"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}
	// Get the public address of the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Failed to cast public key to ECDSA")
	}

	// Compute the Ethereum address
	address_1 := crypto.PubkeyToAddress(*publicKeyECDSA)
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
	address, tx, instance, err := abigen.DeployEthSender(auth, conn)
	if err != nil {
		log.Fatalf("Failed to deploy new storage contract: %v", err)
	}
	fmt.Printf("Contract Address: 0x%x\n", address)
	fmt.Printf("Transaction ID: 0x%x\n\n", tx.Hash())

	time.Sleep(1 * time.Millisecond) // Allow it to be processed by the local node :P
	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		log.Fatalf("Failed to wait for contract deployment mining: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Contract deployment failed with status: %v", receipt.Status)
	}

	fmt.Println("Contract Deployment successful")
	// function call on `instance`. Retrieves pending name
	amount, err := instance.EthSenderCaller.GetTotalSent(&bind.CallOpts{Pending: true}, address_1)
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	fmt.Println("Verification of accessing methods of the  contract:", amount)
}

func TotalSentFromContractToEOA() {
	conn, err := ethclient.Dial("http://127.0.0.1:49500")
	privateKeyHex := "a485f764e233e4279bfae3cb7e4aff48bec869ab63cc9139ad13b57aedb8cefd"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}
	// Get the public address of the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Failed to cast public key to ECDSA")
	}

	// Compute the Ethereum address
	address_1 := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Instantiate the contract and display its name
	// NOTE update the deployment address!
	instance, err := abigen.NewEthSender(common.HexToAddress("0xa791d59427b2b7063050187769ac871b497f4b3c"), conn)

	amount, err := instance.GetTotalSent(&bind.CallOpts{Pending: true}, address_1)
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	amountInEth := weiToEth(amount)
	fmt.Printf("Total sent for the contractId: %v -> amount %f:", address_1.Hex(), amountInEth)
}

// This function Sends an Amoun of Ether to the contract
func sendEthToContract() {
	conn, err := ethclient.Dial("http://127.0.0.1:49500")
	privateKeyHex := "c5114526e042343c6d1899cad05e1c00ba588314de9b96929914ee0df18d46b2"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// Instantiate the contract and display its name
	// NOTE update the deployment address!
	ethSender, err := abigen.NewEthSender(common.HexToAddress("0xa791d59427b2b7063050187769ac871b497f4b3c"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
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
	value := new(big.Int)
	value.SetString("1024000000000000000000", 10)
	auth.Value = value
	tx, err := ethSender.Receive(auth)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	fmt.Println("TransactionID: ", tx.Hash().Hex())
	time.Sleep(1 * time.Millisecond) // Allow it to be processed by the local node :P
	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		log.Fatalf("Failed to wait for contract Receive mining: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Contract Receive failed with status: %v", receipt.Status)
	}

	fmt.Println("Contract Receives successful")

	addrContract := common.HexToAddress("0xa791d59427b2b7063050187769ac871b497f4b3c")
	balance, err := conn.BalanceAt(context.Background(), addrContract, nil)
	if err != nil {
		log.Fatalf("Failed to retrieve contract balance: %v", err)
	}

	// Print the balance in Wei
	fmt.Printf("Contract balance: %s ETH\n", balance.String())

	// Optionally, convert the balance to Ether
	etherValue := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	fmt.Printf("Contract balance: %f Ether\n", etherValue)

}

func checkTheAmountEthOfContract() {
	conn, err := ethclient.Dial("http://127.0.0.1:49500")
	addrContract := common.HexToAddress("0xa791d59427b2b7063050187769ac871b497f4b3c")
	balance, err := conn.BalanceAt(context.Background(), addrContract, nil)
	if err != nil {
		log.Fatalf("Failed to retrieve contract balance: %v", err)
	}

	// Print the balance in Wei
	fmt.Printf("Contract balance: %s ETH\n", balance.String())

	// Optionally, convert the balance to Ether
	etherValue := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	fmt.Printf("Contract %s: ->  balance: %f Ether\n", "0xa791d59427b2b7063050187769ac871b497f4b3c", etherValue)
}

// this function sends 42 ETH from the contract to the Account that handles the contract
func sendEthFromContractToEOA() {
	conn, err := ethclient.Dial("http://127.0.0.1:49500")
	privateKeyHex := "a485f764e233e4279bfae3cb7e4aff48bec869ab63cc9139ad13b57aedb8cefd"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// Instantiate the contract and display its name
	// NOTE update the deployment address!
	ethSender, err := abigen.NewEthSender(common.HexToAddress("0xa791d59427b2b7063050187769ac871b497f4b3c"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
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
	tx, err := ethSender.SendEth(auth)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	fmt.Println("TransactionID: ", tx.Hash().Hex())
	fmt.Println("Sending ethereum in process...")
	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		log.Fatalf("Failed to wait for contract to Send ethereum: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Contract Send failed with status: %v", receipt.Status)
	}

	fmt.Println("Contract Sent ethereum successful")

	addrEOA := common.HexToAddress("0xc21736B0F5a03b348e1746F34806b46C39Ec1575")
	balance, err := conn.BalanceAt(context.Background(), addrEOA, nil)
	if err != nil {
		log.Fatalf("Failed to retrieve contract balance: %v", err)
	}
	EthBalance := weiToEth(balance)
	// Print the balance in Wei
	fmt.Printf("EOA balance: %s ETH\n", EthBalance.String())

}

func weiToEth(wei *big.Int) *big.Float {
	weiInEth := new(big.Float).SetInt(wei)
	ethValue := new(big.Float).Quo(weiInEth, big.NewFloat(1e18))
	return ethValue
}

func main() {
	deployTheContract()
	sendEthToContract()
	sendEthFromContractToEOA()
	checkTheAmountEthOfContract()
	TotalSentFromContractToEOA()
}
