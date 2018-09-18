package eosgo_test

import (
	"testing"

	"github.com/Jeiwan/eosgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVerifySignature(t *testing.T) {
	data := []byte("hello world")
	pubKey := "EOS5Y3fq2HX74VyC1YegjZ6NZojCzGHRZXPcsbVtNbMxHtTUH1jQq"
	signature := "SIG_K1_KYtkgdsvygPiSSUPmSU8bhhWCJLZXiDpEeAUSfGt2HgNCmZmYJWf9fhTgZ4EiSSp1vRGAgNBBXq6mw1q6BjcnxWw5uw8Vr"

	result, err := eosgo.VerifySignature(
		data,
		pubKey,
		signature,
	)

	require.Nil(t, err)
	assert.True(t, result)
}
