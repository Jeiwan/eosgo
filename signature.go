package eosgo

import (
	"crypto/sha256"
	"fmt"
	"strings"

	b58 "github.com/Jeiwan/eos-b58"
	secp256k1 "github.com/ethereum/go-ethereum/crypto/secp256k1"
)

// VerifySignature verifies that a signature is correct
func VerifySignature(data []byte, pubKey, signature string) (bool, error) {
	pubKeyBytes, err := parsePubKey(pubKey)
	if err != nil {
		return false, fmt.Errorf("parsePubKey: %s", err)
	}

	sigBytes, err := parseSignature(signature)
	if err != nil {
		return false, fmt.Errorf("parseSignature: %s", err)
	}

	hashData := sha256.Sum256(data)

	verified := secp256k1.VerifySignature(pubKeyBytes[:], hashData[:], sigBytes[1:])

	return verified, nil
}

// TODO: support PUB, K1, R1, prefixes maybe
func parsePubKey(pubKey string) ([]byte, error) {
	if !strings.HasPrefix(pubKey, "EOS") {
		return nil, fmt.Errorf("Wrong public key format")
	}

	encoded := pubKey[3:]
	decoded, err := b58.CheckDecode(encoded)
	if err != nil {
		return nil, err
	}

	return decoded, nil
}

func parseSignature(signature string) ([]byte, error) {
	if !strings.HasPrefix(signature, "SIG_K1_") {
		return nil, fmt.Errorf("Wrong signature type")
	}

	encoded := signature[7:]

	decoded, err := b58.CheckDecodeWithType(encoded, "K1")
	if err != nil {
		return nil, err
	}

	return decoded, nil
}
