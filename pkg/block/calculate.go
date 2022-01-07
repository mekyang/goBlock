package block

import (
	"bytes"
	"crypto/sha256"
)

func calculateHash(pre, times, data, nonce []byte) []byte {

	header := bytes.Join([][]byte{pre, times, data, nonce}, []byte{})
	hash := sha256.Sum256(header)
	return hash[:]
}

func VerificationHash(b *Block) bool {

	bmap := b.FormatBlockInfo() //map
	time := []byte(bmap["time"])
	nonce := []byte(bmap["nonce"])
	data := []byte(bmap["data"])

	if string(calculateHash(b.previHash, time, data, nonce)) == bmap["blockHash"] {
		return true
	} else {
		return false
	}
}
