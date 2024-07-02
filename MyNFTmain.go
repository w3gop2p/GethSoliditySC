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

// Function to deploy the MyNFT contract
func DeployMyNFTContract() {
	privateKeyHex := "c5114526e042343c6d1899cad05e1c00ba588314de9b96929914ee0df18d46b2"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// Create an RPC-based connection to the Ethereum client
	conn, err := ethclient.Dial("http://127.0.0.1:49500")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Retrieve the current chain ID
	chainID, err := conn.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to retrieve chain ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	baseURI := "http://127.0.0.1:8080/ipfs/QmZSwfkZB2dgGJgh7DNCGSH3S5ejyCEFgHmn2EyGsMYMmy"

	// Deploy the MyNFT contract
	address, tx, _, err := abigen.DeployMyNFT(auth, conn, baseURI)
	if err != nil {
		log.Fatalf("Failed to deploy MyNFT contract: %v", err)
	}

	fmt.Printf("Contract is in a state of deploying! Address: %s, Transaction: %s\n", address.Hex(), tx.Hash().Hex())

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

func mintNFTs() {
	// Connect to the Ethereum client
	conn, err := ethclient.Dial("http://127.0.0.1:49500")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Alice's credentials
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

	// Contract address
	tokenAddress := common.HexToAddress("0xaefB7846330307a20CFe2E88FCAF1e06ef10eE34")
	nft, err := abigen.NewMyNFT(tokenAddress, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate token contract: %v", err)
	}

	// Alice's address
	aliceAddress := common.HexToAddress("0xD8F3183DEF51A987222D845be228e0Bbb932C222")
	// Bob's address
	bobAddress := common.HexToAddress("0xc21736B0F5a03b348e1746F34806b46C39Ec1575")
	// Charlie's address
	charlieAddress := common.HexToAddress("0xac1121D80D44D9A17C88B2E5C34bE3359e33456A")

	// Mint NFT to Alice's address
	tx, err := nft.Mint(auth, aliceAddress)
	if err != nil {
		log.Fatalf("Failed to mint NFT to Alice: %v", err)
	}
	fmt.Printf("Minting NFT to Alice. Transaction: %s\n", tx.Hash().Hex())

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		log.Fatalf("Failed to wait for minting transaction: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Minting transaction failed: %v", receipt.Status)
	}
	fmt.Println("Minted NFT to Alice successfully.")

	// Mint NFT to Bob's address
	tx, err = nft.Mint(auth, bobAddress)
	if err != nil {
		log.Fatalf("Failed to mint NFT to Bob: %v", err)
	}
	fmt.Printf("Minting NFT to Bob. Transaction: %s\n", tx.Hash().Hex())

	// Wait for the transaction to be mined
	receipt, err = bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		log.Fatalf("Failed to wait for minting transaction: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Minting transaction failed: %v", receipt.Status)
	}
	fmt.Println("Minted NFT to Bob successfully.")

	// Mint NFT to Charlie's address
	tx, err = nft.Mint(auth, charlieAddress)
	if err != nil {
		log.Fatalf("Failed to mint NFT to Charlie: %v", err)
	}
	fmt.Printf("Minting NFT to Charlie. Transaction: %s\n", tx.Hash().Hex())

	// Wait for the transaction to be mined
	receipt, err = bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		log.Fatalf("Failed to wait for minting transaction: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Minting transaction failed: %v", receipt.Status)
	}
	fmt.Println("Minted NFT to Charlie successfully.")

}

func checkBalancesAndOwnership() {
	// Connect to the Ethereum client
	conn, err := ethclient.Dial("http://127.0.0.1:49500")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Contract address
	tokenAddress := common.HexToAddress("0xaefB7846330307a20CFe2E88FCAF1e06ef10eE34")
	nft, err := abigen.NewMyNFT(tokenAddress, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate token contract: %v", err)
	}

	// Alice's address
	aliceAddress := common.HexToAddress("0xD8F3183DEF51A987222D845be228e0Bbb932C222")
	// Bob's address
	bobAddress := common.HexToAddress("0xc21736B0F5a03b348e1746F34806b46C39Ec1575")
	// Charlie's address
	charlieAddress := common.HexToAddress("0xac1121D80D44D9A17C88B2E5C34bE3359e33456A")

	// Check Alice's Balance
	balance, err := nft.BalanceOf(nil, aliceAddress)
	if err != nil {
		log.Fatalf("Failed to get NFT balance of Alice: %v", err)
	}
	fmt.Printf("Balance NFT Alice : %s\n", balance)

	// Check Bob's Balance
	balance, err = nft.BalanceOf(nil, bobAddress)
	if err != nil {
		log.Fatalf("Failed to get NFT balance of Bob: %v", err)
	}
	fmt.Printf("Balance NFT Bob: %s\n", balance)

	// Check Charlie's Balance
	balance, err = nft.BalanceOf(nil, charlieAddress)
	if err != nil {
		log.Fatalf("Failed to get NFT balance of Charlie: %v", err)
	}
	fmt.Printf("Balance NFT Charlie: %s\n", balance)
}

func checkOwnership(tokenId *big.Int) {
	// Connect to the Ethereum client
	conn, err := ethclient.Dial("http://127.0.0.1:49500")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Contract address
	tokenAddress := common.HexToAddress("0xaefB7846330307a20CFe2E88FCAF1e06ef10eE34")
	nft, err := abigen.NewMyNFT(tokenAddress, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate token contract: %v", err)
	}
	owner, err := nft.OwnerOf(nil, tokenId)
	fmt.Printf("The owneer of the NFT with ID: %d  is: %s", tokenId, owner.Hex())
	fmt.Println()
}

func transferringNFT(from common.Address, to common.Address) {
	// Connect to the Ethereum client
	conn, err := ethclient.Dial("http://127.0.0.1:49500")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Alice's credentials
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

	// Contract address
	tokenAddress := common.HexToAddress("0xaefB7846330307a20CFe2E88FCAF1e06ef10eE34")
	nft, err := abigen.NewMyNFT(tokenAddress, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate token contract: %v", err)
	}

	// Transfer 1 Token from Alice to Bob
	var id int64 = 1
	tx, err := nft.TransferFrom(auth, from, to, new(big.Int).SetInt64(id))
	if err != nil {
		log.Fatalf("Failed to transfer NFT to Alice: %s: %v", auth.From, err)
	}
	fmt.Printf("Transfering NFT from %s to %s. Transaction: %s\n", from.Hex(), to.Hex(), tx.Hash().Hex())
	fmt.Println()
	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		log.Fatalf("Failed to wait for minting transaction: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Minting transaction failed: %v", receipt.Status)
	}
	fmt.Println("Transfered NFT to *575 successfully.")
}

func approvingAndTransferNFT(to common.Address, tokenId int64) {
	// Connect to the Ethereum client
	conn, err := ethclient.Dial("http://127.0.0.1:49500")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Bob's credentials
	privateKeyHex := "a485f764e233e4279bfae3cb7e4aff48bec869ab63cc9139ad13b57aedb8cefd" //*of *575 bob's account
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

	// Contract address
	tokenAddress := common.HexToAddress("0xaefB7846330307a20CFe2E88FCAF1e06ef10eE34")
	nft, err := abigen.NewMyNFT(tokenAddress, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate token contract: %v", err)
	}

	// Mint NFT to Alice's address
	tx, err := nft.Approve(auth, to, new(big.Int).SetInt64(tokenId))
	if err != nil {
		log.Fatalf("Failed to approve  NFT from *Bob *231 to Alice *222: %v", err)
	}
	fmt.Printf("Approving NFT with tokenId %d from %s  to %s  . Transaction: %s\n", tokenId, auth.From.Hex(), to, tx.Hash().Hex())

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		log.Fatalf("Failed to wait for approving transaction: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Approving transaction failed: %v", receipt.Status)
	}
	fmt.Printf("Aproving  NFT from bob %s to Alice %s successfully.", auth.From.Hex(), to)
	fmt.Println()
	addr, err := nft.GetApproved(nil, new(big.Int).SetInt64(tokenId))
	fmt.Printf("Approved NFT from  %s to %s successfully.", auth.From.Hex(), addr.Hex())
	fmt.Println()
	// transfering
}

func gettingTokenMetada(tokenId int64) {
	// Connect to the Ethereum client
	conn, err := ethclient.Dial("http://127.0.0.1:49500")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Contract address
	tokenAddress := common.HexToAddress("0xaefB7846330307a20CFe2E88FCAF1e06ef10eE34")
	nft, err := abigen.NewMyNFT(tokenAddress, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate token contract: %v", err)
	}

	// Mint NFT to Alice's address
	metadata, err := nft.TokenURI(nil, new(big.Int).SetInt64(tokenId))
	if err != nil {
		log.Fatalf("Failed to get metadata: %v", err)
	}
	fmt.Printf("TokenId: %d ->Metadata: %s\n", tokenId, metadata)
	fmt.Println()
}

func main() {
	// Alice's address
	aliceAddress := common.HexToAddress("0xD8F3183DEF51A987222D845be228e0Bbb932C222")
	// Bob's address
	bobAddress := common.HexToAddress("0xc21736B0F5a03b348e1746F34806b46C39Ec1575")
	// Charlie's address
	//charlieAddress := common.HexToAddress("0xac1121D80D44D9A17C88B2E5C34bE3359e33456A")

	DeployMyNFTContract()
	mintNFTs()
	checkBalancesAndOwnership()
	transferringNFT(bobAddress, aliceAddress)
	checkOwnership(new(big.Int).SetInt64(3))
	approvingAndTransferNFT(aliceAddress, int64(1))
	gettingTokenMetada(1)
}
