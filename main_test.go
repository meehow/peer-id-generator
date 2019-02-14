package main

import (
	"crypto/rand"
	"testing"

	crypto "github.com/libp2p/go-libp2p-crypto"
)

func BenchmarkRSA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		crypto.GenerateRSAKeyPair(2048, rand.Reader)
	}

}

func BenchmarkEd25519(b *testing.B) {
	for i := 0; i < b.N; i++ {
		crypto.GenerateEd25519Key(rand.Reader)
	}

}
