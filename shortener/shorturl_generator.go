package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/itchyny/base58-go"
)

func GenerateShortLink(initialLink string, userId string) string {
	urlHashBytes := sha2560f(initialLink + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}

func sha2560f(input string )[]byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		panic(fmt.Sprintln(err.Error()))
	}
	return string(encoded)
}