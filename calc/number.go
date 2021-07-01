package calc

import "encoding/binary"

func NumberToBytes(number uint64) []byte {
	amountBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountBytes, number)
	return amountBytes
}
