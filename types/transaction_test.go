package types_test

import (
	"encoding/json"
	"testing"

	"github.com/Jeiwan/eosgo"
	"github.com/stretchr/testify/assert"
)

func TestTransactionHeaderUnmarshalling(t *testing.T) {
	data := []byte(`
		{
			"status": "confirmed",
			"cpu_usage_us": 123,
			"net_usage_words": 456,
			"trx": "deadbeef"
		}
	`)

	var h eosgo.TransactionHeader
	err := json.Unmarshal(data, &h)

	assert.Nil(t, err)
	assert.Equal(t, "deadbeef", string(h.Trx.ID))
	assert.Nil(t, h.Trx.Signatures)
	assert.Empty(t, h.Trx.PackedTrx)

	data = []byte(`
		{
			"status": "confirmed",
			"cpu_usage_us": 123,
			"net_usage_words": 456,
			"trx": {
				"id": "1a2b3c",
				"signatures": ["SIG"],
				"packed_trx": "deadbeef"
			}
		}
	`)

	err = json.Unmarshal(data, &h)

	assert.Nil(t, err)
	assert.Equal(t, "1a2b3c", string(h.Trx.ID))
	assert.Equal(t, []string{"SIG"}, h.Trx.Signatures)
	assert.Equal(t, "deadbeef", string(h.Trx.PackedTrx))
}
