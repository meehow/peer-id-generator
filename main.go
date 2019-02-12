package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/libp2p/go-libp2p-crypto"
	"github.com/libp2p/go-libp2p-peer"
)

var alphabet = regexp.MustCompile("^[123456789abcdefghijklmnopqrstuvwxyz]+$")

func main() {
	if len(os.Args) != 2 {
		fmt.Printf(`
This tool generates IPFS public and private keypair until it finds public key
which contains required substring. Keys are stored in local directory. If you
like one of them, you can move it to ~/.ipfs/keystore/ to use it with IPFS.

Usage:
	%s {part}
		For fast results suggested length of public key part is 4-5 characters
`, os.Args[0])
		os.Exit(1)
	}
	part := strings.ToLower(os.Args[1])
	if !alphabet.MatchString(part) {
		fmt.Println("{part} must match the alphabet:", alphabet.String())
		os.Exit(2)
	}
	for {
		err := generateKey(part)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func generateKey(part string) error {
	for {
		privateKey, publicKey, err := crypto.GenerateEd25519Key(rand.Reader)
		if err != nil {
			return err
		}
		peerID, err := peer.IDFromPublicKey(publicKey)
		if err != nil {
			return err
		}
		prettyID := peerID.Pretty()
		lowerID := strings.ToLower(prettyID)
		idx := strings.Index(lowerID, part)
		if idx == -1 {
			continue
		}
		privateKeyBytes, err := privateKey.Bytes()
		if err != nil {
			return err
		}
		fmt.Printf("%s\u001b[32m%s\u001b[0m%s\n", prettyID[:idx], prettyID[idx:len(part)+idx], prettyID[len(part)+idx:])
		return ioutil.WriteFile(prettyID, privateKeyBytes, 0600)
	}
}
