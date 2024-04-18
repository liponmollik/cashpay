package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

// Blockchain represents the blockchain
type Blockchain struct {
	Chain []Block
}

// NewBlock creates a new block
func NewBlock(index int, data, prevHash string) *Block {
	timestamp := time.Now().String()
	hash := calculateHash(index, timestamp, data, prevHash)
	return &Block{Index: index, Timestamp: timestamp, Data: data, PrevHash: prevHash, Hash: hash}
}

// calculateHash calculates the hash of a block
func calculateHash(index int, timestamp, data, prevHash string) string {
	hash := sha256.New()
	hash.Write([]byte(string(index) + timestamp + data + prevHash))
	return hex.EncodeToString(hash.Sum(nil))
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := NewBlock(prevBlock.Index+1, data, prevBlock.Hash)
	bc.Chain = append(bc.Chain, *newBlock)
}

// IsValid checks if the blockchain is valid
func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Chain); i++ {
		currentBlock := bc.Chain[i]
		prevBlock := bc.Chain[i-1]

		if currentBlock.Index != prevBlock.Index+1 {
			return false
		}
		if currentBlock.PrevHash != prevBlock.Hash {
			return false
		}
		if currentBlock.Hash != calculateHash(currentBlock.Index, currentBlock.Timestamp, currentBlock.Data, currentBlock.PrevHash) {
			return false
		}
	}
	return true
}

// NewBlockchain creates a new blockchain with the genesis block
func NewBlockchain() *Blockchain {
	genesisBlock := Block{Index: 0, Timestamp: time.Now().String(), Data: "Genesis Block", PrevHash: "", Hash: ""}
	return &Blockchain{Chain: []Block{genesisBlock}}
}
