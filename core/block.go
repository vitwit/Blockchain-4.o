
package core

import (
	"github.com/vitwit/Blockchain-4.o/config"
	"crypto/sha256"
	"bytes"
)

//block function
type Block struct {
	    Timestamp string
        LastHash  string
        Hash      string
        Data      []string
        Nonce     int 
        Difficulty int
}

var  block Block

func (b *Block) Init(timestamp string,lastHash string,hash string,data string,nonce int,difficulty string) {

	b.Timestamp = timestamp
	b.LastHash = lastHash
	b.Hash = hash
	b.Data = data
	b.Nonce = nonce
	b.Difficulty = config.DIFFICULTY 
}

 func genesis() Block{
	

	block.timestamp = "Genesis Block"
	block.lastHash = "-----"
	block.hash = "S2H3I4V5-A"
	block.data = []string{ "0x01" }
	block.nonce = 0
	block.difficulty = config.DIFFICULTY
	
	return block
}
func mineBlock(lastBlock *Block,data string) Block{
	
    lastHash := lastBlock.hash
    //console.log("fedee",lastHash.toString());
    difficulty = lastBlock
	nonce=0;
	hash := lastBlock.

	if 

	if hash.substring(0,difficulty) != '0'.repeat(difficulty {
		nonce++;
        timestamp=Date.now();
        difficulty=Block.adjustDifficulty(lastBlock,timestamp);
        hash = Block.hash(timestamp,lastHash,data,nonce,difficulty)	

	})
       return block
}


func adjustDifficulty(block Block,data string) int {
	difficulty=lastBlock.timestamp+MINE_RATE > currentime ? difficulty+1 : difficulty-1;
    return difficulty;
}

func hash(timestamp,lastHash,data,nonce,difficulty) ( hash.Hash ){



	h := sha256.New()

	b := bytes.Buffer

	h
	
    return ChainUtil.hash(`${timestamp}${lastHash}${data}${nonce}${difficulty}`).toString();

}
static blockHash(block){
    const{timestamp,lastHash,data,nonce,difficulty} = block;
    return Block.hash(timestamp,lastHash,data,nonce,difficulty);
}
}
