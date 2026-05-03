package indexer

import (
	"crypto/sha1"
	"encoding/hex"
)

func GenerateID(content string) string {
	h := sha1.New()
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil))
}
