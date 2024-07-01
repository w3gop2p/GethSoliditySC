package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"time"
)

func main() {
	//	seedBlockChain()
	info()
	//	info1()
}

func seedBlockChain() {
	client, err := ethclient.Dial("http://127.0.0.1:49500")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()
	// This is one of the default account created with geth by kurtosis for private ethereum blockchain
	defaultAccount := common.HexToAddress("0xD9211042f35968820A3407ac3d80C725f8F75c14")
	defaultPrivateKey := "a492823c3e193d6c595f37a18e3c06650cf4c74558cc818b16130b293716106f" // Replace with your private key

	// Parse the private key
	privateKey, err := crypto.HexToECDSA(defaultPrivateKey)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// Create new accounts

	var newAccounts []common.Address
	for i := 0; i < 1; i++ {
		newAccount := createNewAccount()
		newAccounts = append(newAccounts, newAccount)
	}

	// Send 23 ETH to each new account
	amount := new(big.Int)
	amount.SetString("23000000000000000000", 10) // 23 ETH in Wei
	for _, addr := range newAccounts {
		sendEth(client, defaultAccount, privateKey, addr, amount)
	}

	time.Sleep(5 * time.Second)
	// Print balances of all accounts
	printBalances(client, append([]common.Address{defaultAccount}, newAccounts...))

	// Print All Accounts with their private key and ETH Amount
	printInfoAccounts(client)
}

func info() {
	client, err := ethclient.Dial("http://127.0.0.1:49500")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()
	printInfoAccounts(client)
}

func printInfoAccounts(client *ethclient.Client) {
	// Connect to the RPC client to get the list of accounts
	rpcClient, err := rpc.DialContext(context.Background(), "http://127.0.0.1:49500")
	if err != nil {
		log.Fatalf("Failed to connect to the RPC client: %v", err)
	}

	// Retrieve the list of accounts from the RPC client
	var accountAddresses []string
	err = rpcClient.CallContext(context.Background(), &accountAddresses, "eth_accounts")
	if err != nil {
		log.Fatalf("Failed to retrieve accounts: %v", err)
	}

	// Convert string addresses to common.Address
	var accounts []common.Address
	for _, addr := range accountAddresses {
		accounts = append(accounts, common.HexToAddress(addr))
	}

	// Add the manually specified accounts to the list
	manualAccounts := []common.Address{
		common.HexToAddress("0xD9211042f35968820A3407ac3d80C725f8F75c14"),
		common.HexToAddress("0x39d39B884a19EE904C3558c4CF517573D21d5b46"),
		common.HexToAddress("0x95d7e6AF5fB568236D847DC5a008c4645ff4c736"),

		common.HexToAddress("0xB89c54F06Efd20c55156B0a5e2553a46B8019396"),
		common.HexToAddress("0xac440891426A1fd2970Cd7185CA046940A6C288C"),
		common.HexToAddress("0xc21736B0F5a03b348e1746F34806b46C39Ec1575"),
		common.HexToAddress("0x1634041B2fBa94CE73De13043865C9151205fC9E"),
		common.HexToAddress("0xac1121D80D44D9A17C88B2E5C34bE3359e33456A"),
	}
	accounts = append(accounts, manualAccounts...)
	balances, err := getAccountBalances(client, accounts)
	if err != nil {
		log.Fatalf("Failed to retrieve account balances: %v", err)
	}

	// Print the account balances
	for account, balance := range balances {
		fmt.Printf("Account: %s, Balance: %s Wei\n", account.Hex(), balance.String())
	}
}
func getAccountBalances(client *ethclient.Client, accounts []common.Address) (map[common.Address]*big.Int, error) {
	balances := make(map[common.Address]*big.Int)
	ctx := context.Background()

	for _, account := range accounts {
		balance, err := client.BalanceAt(ctx, account, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get balance for account %s: %v", account.Hex(), err)
		}
		balances[account] = balance
	}
	return balances, nil
}

func createNewAccount() common.Address {
	// Ensure the keystore directory exists
	keystoreDir := filepath.Join("C:", "tmp", "keystore")
	if _, err := os.Stat(keystoreDir); os.IsNotExist(err) {
		if err := os.MkdirAll(keystoreDir, os.ModePerm); err != nil {
			log.Fatalf("Failed to create keystore directory: %v", err)
		}
	}

	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount("password")
	if err != nil {
		log.Fatalf("Failed to create new account: %v", err)
	}
	fmt.Printf("Created new account: %s\n", account.Address.Hex())

	// Get the private key of the new account
	keyJSON, err := os.ReadFile(account.URL.Path)
	if err != nil {
		log.Fatalf("Failed to read key file: %v", err)
	}

	key, err := keystore.DecryptKey(keyJSON, "password")
	if err != nil {
		log.Fatalf("Failed to decrypt key: %v", err)
	}

	privateKeyHex := fmt.Sprintf("%x", crypto.FromECDSA(key.PrivateKey))
	fmt.Printf("Private key of account %s: %s\n", account.Address.Hex(), privateKeyHex)
	return account.Address
}

func sendEth(client *ethclient.Client, from common.Address, fromPrivKey *ecdsa.PrivateKey, to common.Address, amount *big.Int) {
	nonce, err := client.PendingNonceAt(context.Background(), from)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get gas price: %v", err)
	}

	tx := types.NewTransaction(nonce, to, amount, uint64(21000), gasPrice, nil)

	// Use the correct chain ID for Ganache
	//chainID := big.NewInt(1337)
	chainID, err := client.NetworkID(context.Background())
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), fromPrivKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	fmt.Printf("Sending %s Wei from %s to %s\n", amount.String(), from.Hex(), to.Hex())

}

func printBalances(client *ethclient.Client, addresses []common.Address) {
	for _, addr := range addresses {
		balance, err := client.BalanceAt(context.Background(), addr, nil)
		if err != nil {
			log.Fatalf("Failed to get balance: %v", err)
		}
		fmt.Printf("Balance of %s: %s Wei\n", addr.Hex(), balance.String())
	}
}
