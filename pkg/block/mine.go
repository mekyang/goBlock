package block

import (
	"strconv"
	"strings"
)

func mineBlock(pre, time, data []byte) ([]byte, int) {

	var nonce int = 0
	var hash []byte

	for true {
		nonce_ := []byte(strconv.FormatInt(int64(nonce), 10))
		hash := calculateHash(pre, time, data, nonce_)
		if string(hash[:difficult]) == strings.Repeat("0", difficult) {
			break
		}
		nonce++
	}
	return hash, nonce
}
