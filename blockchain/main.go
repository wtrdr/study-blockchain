package main

import "fmt"
import "time"
import "encoding/json"
import "./util"

func main() {
	blockchain := Blockchain{}
	bytes, _ := json.Marshal(&blockchain)
	fmt.Printf(string(bytes))
	fmt.Printf("hello, world\n")
}

type Blockchain struct {
	chain                []Block
	current_transactions []Transaction
}

type Block struct {
	Index         int           `json:"id"`
	Timestamp     float64       `json:"timestamp"`
	Transactions  []Transaction `json:"transactions"`
	Proof         string        `json:"proof"`
	Previous_hash string        `json:"previous_hash"`
}

type Transaction struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount    int    `json:"amount"`
}

func (b *Blockchain) NewBlock(proof string, previous_hash *string) Block {
	timestamp := float64(time.Now().UnixNano()/10 ^ 9)

	var hash string = *previous_hash
	if previous_hash == nil {
		hash = util.Hash()
	}

	block := Block{len(b.chain) + 1, timestamp, b.current_transactions, proof, hash}

	b.current_transactions = []Transaction{}
	b.chain = append(b.chain, block)
	return block
}

func (b *Blockchain) NewTransaction(sender string, recipient string, amount int) int {
	b.current_transactions = append(b.current_transactions, Transaction{sender, recipient, amount})
	return b.LastBlock().Index + 1
}

func (b *Blockchain) LastBlock() Block {
	return b.chain[len(b.chain)-1]
}

func GenesisBlock() *Block {
	return nil
}

//block = {
//	'index': 1,
//	'timestamp': 1506057125.900785,
//	'transactions': [
//	{
//		'sender': "8527147fe1f5426f9dd545de4b27ee00",
//		'recipient': "a77f5cdfa2934df3954a5c7c7da5df1f",
//		'amount': 5,
//	}
//	],
//	'proof': 324984774000,
//	'previous_hash': "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
//}
