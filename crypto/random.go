package crypto

import (
	"crypto/rand"
	"fmt"
)

// use crypto/rand Reader to generate keys
func Key() string {
	buf := make([]byte, 16)

	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", buf)
}