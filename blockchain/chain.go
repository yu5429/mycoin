package blockchain

import (
	"fmt"
	"sync"

	"github.com/yu5429/mycoin/db"
	"github.com/yu5429/mycoin/utils"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) { //블록체인 전체 decode
	utils.FromBytes(b, data)
}

func (b *blockchain) persist() {
	db.SaveCheckpoint(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func (b *blockchain) Blocks() []*Block { //전체 블록체인 가져오기
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func Blockchain() *blockchain { //starting point
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			if db.Checkpoint() == nil { //db의 체크포인트 확인
				b.AddBlock("Genesis")
			} else {
				b.restore(db.Checkpoint())
			}

		})
	}
	fmt.Println(b.NewestHash)
	return b
}
