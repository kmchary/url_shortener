package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/itchyny/base58-go"
)

type UrlShortenerServiceInterface interface {
	ShortenURL(url, userId string) string
}

type UrlShortenerService struct {
}

var ctx = context.Background()

func (s *UrlShortenerService) ShortenURL(url string, userId string) string {
	hashAlgorithm := sha256.New()
	hashAlgorithm.Write([]byte(url + userId))
	hashBytes := hashAlgorithm.Sum(nil)
	generatedNumber := new(big.Int).SetBytes(hashBytes).Uint64()
	fmt.Println("generatedNumber: ", generatedNumber)
	encoder := base58.BitcoinEncoding
	encodedString, _ := encoder.Encode([]byte(fmt.Sprintf("%d", generatedNumber)))

	return string(encodedString)
}

func NewUrlShortenerService() UrlShortenerServiceInterface {
	return &UrlShortenerService{}
}
