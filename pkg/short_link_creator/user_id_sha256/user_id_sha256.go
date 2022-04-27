package useridsha256

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/itchyny/base58-go"
)
type ShortLinkCreator struct {}

func (slc *ShortLinkCreator) GenerateShortLink(original_link string, user_id string) string {
	sha2560f := func (input string )[]byte {
		algorithm := sha256.New()
		algorithm.Write([]byte(input))
		return algorithm.Sum(nil)
	}
	base58encode:= func(bytes []byte) string {
		encoding := base58.BitcoinEncoding
		encoded, err := encoding.Encode(bytes)
		if err != nil {
			panic(fmt.Sprintln(err.Error()))
		}
		return string(encoded)
	}
	urlHashBytes := sha2560f(original_link + user_id)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58encode([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}

