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

// func  for deployment of the contract
func deployTokenERC20ToBlockChain() {
	privateKeyHex := "c5114526e042343c6d1899cad05e1c00ba588314de9b96929914ee0df18d46b2"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}
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
	value := new(big.Int)
	value.SetString("5000", 10)
	// Deploy the contract passing the newly created `auth` and `conn` vars
	address, tx, _, err := abigen.DeployTokenERC20(auth, conn, value)
	if err != nil {
		log.Fatalf("Failed to deploy new storage contract: %v", err)
	}
	fmt.Println("Contract is in deploing state")
	fmt.Printf("Contract Address: 0x%x\n", address)
	fmt.Printf("Transaction ID: 0x%x\n\n", tx.Hash())
	// Optional: wait for the deployment transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		log.Fatalf("Failed to wait for the deployment transaction to be mined: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Deployment transaction failed with status: %v", receipt.Status)
	}

	fmt.Println("Contract successfully deployed!")
}

// EOA_1(*222) Transfers Tokens to EOA_2(*052)
func transfer_from_TokenErc20_to_EOA(accstring string) {
	// This function Sends an Amoun of Ether to the contract
	conn, err := ethclient.Dial("http://127.0.0.1:49500")
	privateKeyHex := "c5114526e042343c6d1899cad05e1c00ba588314de9b96929914ee0df18d46b2"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}
	// Instantiate the contract and display its name
	// NOTE update the deployment address!
	tokenERC20, err := abigen.NewTokenERC20(common.HexToAddress("0x08cf4ecfb623f8c4ede4d01ade77d517b64871b0"), conn)
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
	transferAmount := new(big.Int)
	transferAmount.SetString("1000000000000000000", 10)
	addrEOA := common.HexToAddress(accstring)
	tx, err := tokenERC20.Transfer(auth, addrEOA, transferAmount)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	fmt.Println("Transfering token...")
	// functionality to wait until the transaction is written in blockchain
	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		log.Fatalf("Failed to wait for transaction mining: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Transaction failed with status: %v", receipt.Status)
	}
	fmt.Println("Transfered successfully!")
	bl, _ := tokenERC20.BalanceOf(nil, addrEOA)
	fmt.Printf("New Balance for: %v -> %d ", addrEOA.Hex(), bl)
	fmt.Println()

	fmt.Printf("Transered amount of tokens to the EOA %s in wei is: %d", addrEOA, transferAmount)
	fmt.Println()
	etherValue := new(big.Float).Quo(new(big.Float).SetInt(transferAmount), big.NewFloat(1e18))
	fmt.Printf("Transered amount of tokens to the EOA %s in ETH is: %v", addrEOA, etherValue)
	fmt.Println()
	fmt.Printf("TransactionID: %s", tx.Hash().Hex())

}
func getTokensAmountforEOAs(accstring string) (*big.Int, error) {
	// This function Sends an Amoun of Ether to the contract
	conn, err := ethclient.Dial("http://127.0.0.1:49500")
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	acc := common.HexToAddress(accstring)

	// Instantiate the contract and display its name
	// NOTE update the deployment address!
	tokenERC20, err := abigen.NewTokenERC20(common.HexToAddress("0x08cf4ecfb623f8c4ede4d01ade77d517b64871b0"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
	}
	// Retrieve the current chain ID

	totalTokens, err := tokenERC20.BalanceOf(nil, acc)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Print the balance in Wei
	fmt.Printf("Contract Tokens: %s Tokens\n", totalTokens)

	// Optionally, convert the balance to Ether
	etherValue := new(big.Float).Quo(new(big.Float).SetInt(totalTokens), big.NewFloat(1e18))
	fmt.Printf("Contract tokens balance: %f Tokens\n", etherValue)

	return totalTokens, nil
}

func getTotalSupply_of_tokenERC20() {
	// This function Sends an Amoun of Ether to the contract
	conn, err := ethclient.Dial("http://127.0.0.1:49500")

	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// Instantiate the contract and display its name
	// NOTE update the deployment address!
	tokenERC20, err := abigen.NewTokenERC20(common.HexToAddress("0x08cf4ecfb623f8c4ede4d01ade77d517b64871b0"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
	}
	// Retrieve the current chain ID

	totalSupply, err := tokenERC20.TokenERC20Caller.TotalSupply(nil)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	fmt.Println("Total Supply of tokens is in 'wei': ", totalSupply)
	// Optionally, convert the balance to Ether
	etherValue := new(big.Float).Quo(new(big.Float).SetInt(totalSupply), big.NewFloat(1e18))
	fmt.Printf("Total Supply of tokens is in 'eth': %f", etherValue)
}

func EOA123Approach() {
	conn, err := ethclient.Dial("http://127.0.0.1:49500")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Private key of EOA1 (the owner of the tokens)
	privateKeyHex := "c5114526e042343c6d1899cad05e1c00ba588314de9b96929914ee0df18d46b2"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	chainID, err := conn.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to retrieve chain ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Instantiate the token contract
	tokenAddress := common.HexToAddress("0x08cf4ecfb623f8c4ede4d01ade77d517b64871b0")
	tokenERC20, err := abigen.NewTokenERC20(tokenAddress, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate token contract: %v", err)
	}

	// EOA1 approves EOA3 to spend tokens on its behalf
	spender := common.HexToAddress("0xac1121D80D44D9A17C88B2E5C34bE3359e33456A")
	approveAmount := new(big.Int)
	approveAmount.SetString("9000000000000000000", 10) // 1 token (assuming 18 decimals)
	tx, err := tokenERC20.Approve(auth, spender, approveAmount)
	if err != nil {
		log.Fatalf("Failed to approve spender: %v", err)
	}
	fmt.Printf("Approval transaction submitted: %s\n", tx.Hash().Hex())

	// Wait for the approval transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		log.Fatalf("Failed to wait for approval transaction mining: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Approval transaction failed with status: %v", receipt.Status)
	}
	fmt.Println("Amount from EOA1 to EOA3 approved to be spent")
	// EOA3 transfers tokens to EOA2
	privateKeyHexEOA3 := "49d44b1e652028ac7e3385c5ead8a7caea0c830c6f067d06d8e3acbf9164a946"
	privateKeyEOA3, err := crypto.HexToECDSA(privateKeyHexEOA3)
	if err != nil {
		log.Fatalf("Failed to load private key of EOA3: %v", err)
	}

	authEOA3, err := bind.NewKeyedTransactorWithChainID(privateKeyEOA3, chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor for EOA3: %v", err)
	}

	sender := common.HexToAddress("0xD8F3183DEF51A987222D845be228e0Bbb932C222")
	recipient := common.HexToAddress("0xac1121D80D44D9A17C88B2E5C34bE3359e33456A")
	transferAmount := new(big.Int)
	transferAmount.SetString("5000000000000000000", 10) // 0.5 tokens (assuming 18 decimals)

	tx, err = tokenERC20.TransferFrom(authEOA3, sender, recipient, transferAmount)
	if err != nil {
		log.Fatalf("Failed to transfer tokens: %v", err)
	}
	fmt.Printf("Transfer transaction submitted: %s\n", tx.Hash().Hex())

	// Wait for the transfer transaction to be mined
	receipt, err = bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		log.Fatalf("Failed to wait for transfer transaction mining: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Transfer transaction failed with status: %v", receipt.Status)
	}

	fmt.Println("Transfer successful")
	spenderAddr := common.HexToAddress("0xac1121D80D44D9A17C88B2E5C34bE3359e33456A")
	tokenBalanceRemained, _ := tokenERC20.Allowance(nil, sender, spenderAddr)
	fmt.Println("6A a ramas cu: ", tokenBalanceRemained)
}

func main() {
	deployTokenERC20ToBlockChain()
	getTotalSupply_of_tokenERC20()
	transfer_from_TokenErc20_to_EOA("0xc21736B0F5a03b348e1746F34806b46C39Ec1575")
	getTokensAmountforEOAs("0xac1121D80D44D9A17C88B2E5C34bE3359e33456A")
	EOA123Approach()
}

//0xD8F3183DEF51A987222D845be228e0Bbb932C222  EOA1
//0xc21736B0F5a03b348e1746F34806b46C39Ec1575  EOA2
//0xac1121D80D44D9A17C88B2E5C34bE3359e33456A  EOA3
