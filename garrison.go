//go:generate abigen --sol Spawn.sol --pkg main --out Spawn.go
package main

import (
	"fmt"
	"log"
	"strings"

	"time"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

const key = `{"address":"b5c423d9b40cb35ae2d70dc119a2db4d36a99ed7","crypto":{"cipher":"aes-128-ctr","ciphertext":"26876906681a87b4b90281d3ba85f6ed06cbbd037b9fdb7a95cafa2af95507d4","cipherparams":{"iv":"56ef1a75c04a9e7e1c54844ffe393bb3"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"d1c866674cff14a416acd2754aaea95188f0492b754e27daba1e3d0ba666c002"},"mac":"1d34f643b9a4a0b0ae44c4b5ac0f3d1152aa9d72f2c930ebe3fdb80df7b8af6a"},"id":"9b483bab-8006-4845-9248-48dbb734e068","version":3}`

func main() {

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial("/home/dave/testGeth/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	fmt.Println("success: ", conn)
	auth, err := bind.NewTransactor(strings.NewReader(key), "a")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	fmt.Println("auth succesful: ", auth)

	address, tx, token, err := DeploySpawn(auth, conn)
	if err != nil {
		log.Fatalf("Failed to deploy new contract %v", err)
	}
	fmt.Printf("Contract pending deploy 0x%x\n\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	fmt.Print("token: ", token)
	time.Sleep(4000 * time.Millisecond)

	session := &SpawnSession{
		Contract: token,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: big.NewInt(3141592),
		},
	}

	// sub, err := token.SubAddress(&bind.CallOpts{})
	sub, err := session.SubAddress()
	if err != nil {
		log.Fatalf("Failed to retrieve subaddress: %v", err)
	}

	fmt.Println("subaddress:", sub)

	//	trans, err := token.CreateTransfer(&bind.TransactOpts{})
	trans, err := session.CreateTransfer()
	if err != nil {
		log.Fatalf("failed to make sub transfer %v", err)
	}

	time.Sleep(4000 * time.Millisecond)
	fmt.Println("trans:", trans)
	fmt.Println("sub", sub)

	sub, err = session.SubAddress()
	fmt.Println("final Sub:", sub)

}
