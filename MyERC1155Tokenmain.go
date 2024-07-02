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

// *** --- ***
func DeployMyERC1155Tokens() {
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

	// Deploy the MyNFT contract
	address, tx, _, err := abigen.DeployMyERC1155Token(auth, conn, auth.From)
	if err != nil {
		log.Fatalf("Failed to deploy MyERC1155Token contract: %v", err)
	}

	fmt.Printf("Contract is the state of deploying! Address: %s, Transaction: %s\n", address.Hex(), tx.Hash().Hex())

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

func mintAllTokens() {
	// Replace with your private key hex
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

	// Replace with your deployed contract address and recipient address
	contractAddress := common.HexToAddress("0x8be3be168D5DCADfDEE88Fc37d305D70DE5Fe3c1")
	recipient := common.HexToAddress("0xc21736B0F5a03b348e1746F34806b46C39Ec1575")

	// Call the mintAllTokens function with client, auth, contractAddress, recipient
	err = mintAllTokensFunc(conn, auth, contractAddress, recipient)
	if err != nil {
		log.Fatalf("Failed to mint tokens: %v", err)
	}
}

func mintAllTokensFunc(client *ethclient.Client, auth *bind.TransactOpts, contractAddress common.Address, recipient common.Address) error {
	// Connect to the contract instance
	instance, err := abigen.NewMyERC1155Token(contractAddress, client)
	if err != nil {
		return err
	}
	var Gold, _ = instance.GOLD(nil)

	// Mint GOLD tokens
	tx, err := instance.Mint(auth, recipient, Gold, big.NewInt(1000), nil)
	if err != nil {
		return err
	}
	fmt.Printf("Minted Gold Token, Transaction: %s\n", tx.Hash().Hex())

	// Mint SILVER tokens
	var Silver, _ = instance.SILVER(nil)
	tx, err = instance.Mint(auth, recipient, Silver, big.NewInt(1000), nil)
	if err != nil {
		return err
	}
	fmt.Printf("Minted Silver Token, Transaction: %s\n", tx.Hash().Hex())

	// Mint SWORD tokens
	var Sword, _ = instance.SWORD(nil)
	tx, err = instance.Mint(auth, recipient, Sword, big.NewInt(10), nil)
	if err != nil {
		return err
	}
	fmt.Printf("Minted Gold Token, Transaction: %s\n", tx.Hash().Hex())
	// Mint SHIELD tokens
	var Shield, _ = instance.SHIELD(nil)
	tx, err = instance.Mint(auth, recipient, Shield, big.NewInt(10), nil)
	if err != nil {
		return err
	}
	fmt.Printf("Minted Shield Token, Transaction: %s\n", tx.Hash().Hex())

	// Optional: wait for the deployment transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("Failed to wait for the deployment transaction to be mined: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Deployment transaction failed with status: %v", receipt.Status)
	}

	fmt.Println("Tokens minted successfully!")

	log.Println("Tokens minted successfully!")
	return nil
}

func checkBalance() {
	// Create an RPC-based connection to the Ethereum client
	conn, err := ethclient.Dial("http://127.0.0.1:49500")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Replace with your deployed contract address and recipient address
	contractAddress := common.HexToAddress("0x8be3be168D5DCADfDEE88Fc37d305D70DE5Fe3c1")
	recipient := common.HexToAddress("0xac1121D80D44D9A17C88B2E5C34bE3359e33456A")
	//	recipient1 := common.HexToAddress("0xD8F3183DEF51A987222D845be228e0Bbb932C222")
	// 0xac1121D80D44D9A17C88B2E5C34bE3359e33456A
	//0xc21736B0F5a03b348e1746F34806b46C39Ec1575
	instance, err := abigen.NewMyERC1155Token(contractAddress, conn)
	if err != nil {
		fmt.Println(err)
	}
	// Define the token IDs you want to check
	tokenIDs := []*big.Int{
		big.NewInt(0), // GOLD
		big.NewInt(1), // SILVER
		big.NewInt(2), // SWORD
		big.NewInt(3), // SHIELD
	}

	// Check balances for each token ID
	for _, tokenID := range tokenIDs {
		balance, err := instance.BalanceOf(&bind.CallOpts{}, recipient, tokenID)
		if err != nil {
			log.Println(err)
		}

		// Print the balance
		fmt.Printf("Token ID %d: Balance %s\n", tokenID, balance.String())
	}
}

func mintBatchfunc() {
	privateKeyHex := "c5114526e042343c6d1899cad05e1c00ba588314de9b96929914ee0df18d46b2"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	conn, err := ethclient.Dial("http://127.0.0.1:49500")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	chainID, err := conn.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to retrieve chain ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	contractAddress := common.HexToAddress("0x8be3be168D5DCADfDEE88Fc37d305D70DE5Fe3c1") // Replace with your deployed contract address
	recipient := common.HexToAddress("0xc21736B0F5a03b348e1746F34806b46C39Ec1575")

	ids := []*big.Int{
		big.NewInt(0), // GOLD
		big.NewInt(1), // SILVER
		big.NewInt(2), // SWORD
		big.NewInt(3), // SHIELD
	}

	amounts := []*big.Int{
		big.NewInt(100), // Amount of GOLD
		big.NewInt(200), // Amount of SILVER
		big.NewInt(5),   // Amount of SWORD
		big.NewInt(5),   // Amount of SHIELD
	}

	instance, err := abigen.NewMyERC1155Token(contractAddress, conn)
	if err != nil {
		log.Fatalf("Failed to load contract instance: %v", err)
	}

	tx, err := instance.MintBatch(auth, recipient, ids, amounts, []byte{})
	if err != nil {
		log.Fatalf("Failed to mint tokens: %v", err)
	}

	fmt.Printf("MintBatch transaction sent: %s\n", tx.Hash().Hex())
	// Optional: wait for the deployment transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		log.Fatalf("Failed to wait for the deployment transaction to be mined: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Deployment transaction failed with status: %v", receipt.Status)
	}

	fmt.Println("MintBatch transaction completed!")

	log.Println("MintBatch transaction completed!")
}

func transferToken() {
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
	contractAddress := common.HexToAddress("0x8be3be168D5DCADfDEE88Fc37d305D70DE5Fe3c1")
	recipient := common.HexToAddress("0xc21736B0F5a03b348e1746F34806b46C39Ec1575")
	instance, err := abigen.NewMyERC1155Token(contractAddress, conn)
	if err != nil {
		fmt.Println(err)
	}
	tokenID := big.NewInt(1)
	amount := big.NewInt(123)
	tx, err := instance.SafeTransferFrom(auth, auth.From, recipient, tokenID, amount, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("SafeTransferFrom transaction sent: %s\n", tx.Hash().Hex())

	// Optional: wait for the deployment transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		log.Fatalf("Failed to wait for the deployment transaction to be mined: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Deployment transaction failed with status: %v", receipt.Status)
	}

	fmt.Println("SafeTransferFrom transaction sent successfully!")

	log.Println("SafeTransferFrom transaction sent successfully!")
}

func transferBatchToken() {
	privateKeyHex := "49d44b1e652028ac7e3385c5ead8a7caea0c830c6f067d06d8e3acbf9164a946"
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
	contractAddress := common.HexToAddress("0x8be3be168D5DCADfDEE88Fc37d305D70DE5Fe3c1")
	recipient := common.HexToAddress("0xac1121D80D44D9A17C88B2E5C34bE3359e33456A")
	//recipient1 := common.HexToAddress("0x7DE6c417Fe1937AC25f76752D6e1388A12A13036")
	instance, err := abigen.NewMyERC1155Token(contractAddress, conn)
	if err != nil {
		fmt.Println(err)
	}
	ids := []*big.Int{
		big.NewInt(0), // GOLD
		big.NewInt(1), // SILVER
		big.NewInt(2), // SWORD
		big.NewInt(3), // SHIELD
	}

	amounts := []*big.Int{
		big.NewInt(800),  // Amount of GOLD
		big.NewInt(1000), // Amount of SILVER
		big.NewInt(5),    // Amount of SWORD
		big.NewInt(8),    // Amount of SHIELD
	}
	fromAccount := "0xc21736B0F5a03b348e1746F34806b46C39Ec1575"
	fromAccountAddress := common.HexToAddress(fromAccount)
	tx, err := instance.SafeBatchTransferFrom(auth, fromAccountAddress, recipient, ids, amounts, []byte{})
	if err != nil {
		fmt.Println("The error is for SafeBatchTransferFrom method: ", err)
	}
	fmt.Printf("SafeBatchTransferFrom transaction sent: %s\n", tx.Hash().Hex())
	// Optional: wait for the deployment transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		log.Fatalf("Failed to wait for the deployment transaction to be mined: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Deployment transaction failed with status: %v", receipt.Status)
	}

	fmt.Println("SafeBatchTransferFrom transaction sent successfully!")

	log.Println("SafeBatchTransferFrom transaction sent successfully!")

}

func burnToken() {
	privateKeyHex := "a485f764e233e4279bfae3cb7e4aff48bec869ab63cc9139ad13b57aedb8cefd"
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

	contractAddress := common.HexToAddress("0x8be3be168D5DCADfDEE88Fc37d305D70DE5Fe3c1")
	instance, err := abigen.NewMyERC1155Token(contractAddress, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a MyERC1155Token contract: %v", err)
	}

	playerAddress := common.HexToAddress("0xc21736B0F5a03b348e1746F34806b46C39Ec1575")
	tokenID := big.NewInt(2) // SWORD token ID
	amount := big.NewInt(3)  // Amount to burn

	tx, err := instance.Burn(auth, playerAddress, tokenID, amount)
	if err != nil {
		log.Fatalf("Failed to burn token: %v", err)
	}

	fmt.Printf("Token burn transaction sent: %s\n", tx.Hash().Hex())

	// Optional: wait for the burn transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		log.Fatalf("Failed to wait for the burn transaction to be mined: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Burn transaction failed with status: %v", receipt.Status)
	}

	fmt.Println("Token burned successfully!")
}

func burnBatchTokens() {
	privateKeyHex := "a485f764e233e4279bfae3cb7e4aff48bec869ab63cc9139ad13b57aedb8cefd"
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

	contractAddress := common.HexToAddress("0x8be3be168D5DCADfDEE88Fc37d305D70DE5Fe3c1")
	instance, err := abigen.NewMyERC1155Token(contractAddress, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a MyERC1155Token contract: %v", err)
	}

	playerAddress := common.HexToAddress("0xc21736B0F5a03b348e1746F34806b46C39Ec1575")

	ids := []*big.Int{
		big.NewInt(0), // GOLD
		big.NewInt(1), // SILVER
		big.NewInt(2), // SWORD
		big.NewInt(3), // SHIELd
	}

	amounts := []*big.Int{
		big.NewInt(300), // Amount of GOLD
		big.NewInt(500), // Amount of SILVER
		big.NewInt(1),
		big.NewInt(8),
	}

	tx, err := instance.BurnBatch(auth, playerAddress, ids, amounts)
	if err != nil {
		log.Fatalf("Failed to burn tokens: %v", err)
	}

	fmt.Printf("Token burn transaction sent: %s\n", tx.Hash().Hex())

	// Optional: wait for the burn transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		log.Fatalf("Failed to wait for the burn transaction to be mined: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Burn transaction failed with status: %v", receipt.Status)
	}

	fmt.Println("Tokens burned successfully!")
}

func approveTokenSpending() {
	// Private key of the account that owns the tokens
	ownerPrivateKeyHex := "a485f764e233e4279bfae3cb7e4aff48bec869ab63cc9139ad13b57aedb8cefd"
	ownerPrivateKey, err := crypto.HexToECDSA(ownerPrivateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// Address of the account that will be approved to spend tokens
	spenderAddress := common.HexToAddress("0xac1121D80D44D9A17C88B2E5C34bE3359e33456A")

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

	// Create a transactor with the owner's private key and chain ID
	auth, err := bind.NewKeyedTransactorWithChainID(ownerPrivateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Address of the ERC1155 token contract
	contractAddress := common.HexToAddress("0x8be3be168D5DCADfDEE88Fc37d305D70DE5Fe3c1")
	instance, err := abigen.NewMyERC1155Token(contractAddress, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a MyERC1155Token contract: %v", err)
	}

	// Call the setApprovalForAll function of the ERC1155 contract
	tx, err := instance.SetApprovalForAll(auth, spenderAddress, true)
	if err != nil {
		log.Fatalf("Failed to approve spending: %v", err)
	}

	fmt.Printf("Approval transaction sent: %s\n", tx.Hash().Hex())

	// Optional: wait for the transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		log.Fatalf("Failed to wait for the approval transaction to be mined: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Approval transaction failed with status: %v", receipt.Status)
	}

	fmt.Println("Token approval for spending completed successfully!")
}

func main() {
	DeployMyERC1155Tokens()
	mintAllTokens()
	checkBalance()
	mintBatchfunc()
	transferToken()
	transferBatchToken()
	burnToken()
	burnBatchTokens()
	approveTokenSpending()
}
