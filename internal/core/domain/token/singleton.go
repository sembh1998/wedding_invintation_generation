package token

import (
	"crypto/ed25519"
	"encoding/hex"
	"log"
	"sync"
)

type Tokenizer struct {
	Maker Maker
}

var tokenizer *Tokenizer
var once sync.Once

func NewTokenizer() *Tokenizer {
	if tokenizer == nil {
		once.Do(
			func() {
				publicKey, privateKey, err := ed25519.GenerateKey(nil)
				if err != nil {
					log.Fatalf("Error generating key pair: %v", err)
				}
				maker, err := NewPasetoMaker(hex.EncodeToString(privateKey), hex.EncodeToString(publicKey))
				if err != nil {
					log.Fatalf("Error creating Paseto maker: %v", err)
				}
				tokenizer = &Tokenizer{
					Maker: maker,
				}
			},
		)
	}
	return tokenizer
}
