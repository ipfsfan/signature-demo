package main

import (
	"encoding/hex"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/lib/sigs"
	_ "github.com/filecoin-project/lotus/lib/sigs/secp"
)

func newSecp256k1Address() {
	privateKey, err := sigs.Generate(crypto.SigTypeSecp256k1)
	if err != nil {
		return
	}

	publicKey, err := sigs.ToPublic(crypto.SigTypeSecp256k1, privateKey)
	if err != nil {
		return
	}
	address, err := address.NewSecp256k1Address(publicKey)
	if err != nil {
		return
	}
	fmt.Printf("privateKey: %s\n", hex.EncodeToString(privateKey))
	fmt.Printf("publicKey: %s\n", hex.EncodeToString(publicKey))
	fmt.Printf("address: %s\n", address)
}

func testSignature() {
	//PrivateKey: 50afa18afc77814b633d94e2d62d0f62801d62ccb4405a794405ef35851f1c02
	//publicKey: 044dce231f161fccc3087a8476687c7ec6b60f1fa7515a45301f432de476db7f746943f2579f4c125066fb6cea6b0194f4d19c5a44279de65d316d8c2d2d1073da
	//Address: w14f24d34rwxjtqr7fh6omhfdhl7jfniio3cghuuq

	msg := "hello"
	privateKey, _ := hex.DecodeString("50afa18afc77814b633d94e2d62d0f62801d62ccb4405a794405ef35851f1c02")
	sign, err := sigs.Sign(crypto.SigTypeSecp256k1, privateKey, []byte(msg))
	if err != nil {
		return
	}

	addr, err := address.NewFromString("w14f24d34rwxjtqr7fh6omhfdhl7jfniio3cghuuq")
	if err != nil {
		return
	}

	err = sigs.Verify(sign, addr, []byte(msg))
	if err != nil{
		return
	}

	fmt.Printf("Verify signature  ok")
}

func main() {
	newSecp256k1Address()

	testSignature()
}
