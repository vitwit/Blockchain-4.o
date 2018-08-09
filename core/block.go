
package core

import (
	"time"
	"strings"
)

//Block is model for a single block in the blockchain
type Block struct {
	    Timestamp time.Time
        LastHash  string
        Hash      string
        Data      string
        Nonce     int32
        Difficulty int32
}

// Genesis will return the initial block for a blockchain,
// should only be called first time and once
 func Genesis() Block {
	b := Block{
		Timestamp: time.Now().UTC(),
		LastHash: "",
		Hash: "0x01",
		Data: "blockchain 4.0",
		Nonce: 0,
		Difficulty: Difficulty,
	}


	return b
}

// MineBlock will try to brute force nonce and if succeeded, generates a new block
func MineBlock(lastBlock Block, data string) Block {

		var nonce int32
		var t time.Time
		var hash string
	for {
		nonce += 1
		t = time.Now().UTC()
		Difficulty = adjustDifficulty(lastBlock)
		hash, _ = newHash(t, lastBlock.Hash, data, nonce, Difficulty)

		if strings.HasPrefix(hash, "0000") {
			break
		}
	}
	newBlock := Block{
		Timestamp: t,
		Hash: hash,
		Data: data,
		Nonce: nonce,
		Difficulty: Difficulty,
	}

	return newBlock
}

// adjustDifficulty changes difficulty to maintain a relative time for block mining
func adjustDifficulty(lastBlock Block) int32 {

	if lastBlock.Timestamp.Sub(time.Now().UTC()) < time.Second * 12 {
		return  Difficulty + 1
	}

	return Difficulty - 1
}

// newHash returns a new SHA256 hash, type-casted to a string
func newHash(timestamp time.Time, lastHash, data string, nonce, difficulty int32) ( string, error ){

	var h  bytes.Buffer

	h.Write([]byte(timestamp.String()))
	h.Write([]byte(lastHash))
	h.Write([]byte(data))
	h.Write([]byte(string(nonce)))
	h.Write([]byte(string(difficulty)))

	hash := sha256.Sum256(h.Bytes())

	return string(hash[:]), nil
}

// blockHash generates hash for a block
func  blockHash (block *Block) ( string, error ) {

	bHash, err := newHash(block.Timestamp, block.LastHash, block.Data, block.Nonce, block.Difficulty)
	utils.CheckError(err)
	fmt.Printf("%x", bHash)

	return bHash, nil
}