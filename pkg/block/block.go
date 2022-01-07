package block

import (
	"fmt"
	"strconv"
	"time"
)

var idPoint uint64
var difficult int = 3

const format string = Format

type Block struct {
	blockID    uint64
	blockHash  []byte
	previHash  []byte
	timeStamp  int64
	timeFormat string
	blockData  []byte
	nonce      int
}

func GainBlock(id uint64, hash, pre, data []byte, nonce int) *Block {

	t := time.Now()
	year, month, day := t.Date()
	var timestamp int64 = t.Unix()
	var timeformat string = fmt.Sprintf("%d-%d-%d", year, month, day)

	blockThis := &Block{
		blockID:    id,
		blockHash:  hash,
		previHash:  pre,
		timeStamp:  timestamp,
		timeFormat: timeformat,
		blockData:  data,
		nonce:      nonce,
	}

	return blockThis
}

func (b *Block) PrintBlockInfo() {

	bm := b.FormatBlockInfo()

	fmt.Printf(format, "block id", bm["ID"], bm)

}

func (b *Block) FormatBlockInfo() map[string]string {

	blockInfo := make(map[string]string)

	blockInfo["ID"] = strconv.FormatUint(b.blockID, 10)
	blockInfo["Hash"] = string(b.blockHash)
	blockInfo["pHash"] = string(b.previHash)
	blockInfo["time"] = b.timeFormat
	blockInfo["data"] = string(b.blockData)
	blockInfo["nonce"] = strconv.FormatInt(int64(b.nonce), 10)

	return blockInfo
}
