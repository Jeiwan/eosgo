package types_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/Jeiwan/eosgo"
	"github.com/stretchr/testify/assert"
)

func TestBlockProducersUnmarshalling(t *testing.T) {
	data, _ := ioutil.ReadFile("./fixtures/block_producers.json")

	var bps eosgo.ProducersInfo
	err := json.Unmarshal(data, &bps)

	assert.Nil(t, err)

	assert.Equal(t, "8144169062527346688.00000000000000000", bps.TotalProducerVoteWeight.Text('f', 17))
	assert.Equal(t, "anotherbp", bps.More)
	assert.Len(t, bps.Rows, 2)

	bp := bps.Rows[0]
	assert.Equal(t, "zbeosbp11111", bp.Owner)
	assert.Equal(t, "225494609004205664.00000000000000000", bp.TotalVotes.Text('f', 17))
	assert.Equal(t, "EOS7rhgVPWWyfMqjSbNdndtCK8Gkza3xnDbUupsPLMZ6gjfQ4nX81", bp.ProducerKey)
	assert.Equal(t, 1, bp.IsActive)
	assert.Equal(t, "http://www.zbeos.com", bp.URL)
	assert.Equal(t, 5356, bp.UnpaidBlocks)
	assert.Equal(t, 1530201589500000, bp.LastClaimTime)
	assert.Equal(t, 0, bp.Location)

	bp = bps.Rows[1]
	assert.Equal(t, "worblieosbp1", bp.Owner)
	assert.Equal(t, "641963941752872.50000000000000000", bp.TotalVotes.Text('f', 17))
	assert.Equal(t, "EOS1111111111111111111111111111111114T1Anm", bp.ProducerKey)
	assert.Equal(t, 0, bp.IsActive)
	assert.Equal(t, "https://eos.worbli.io", bp.URL)
	assert.Equal(t, 0, bp.UnpaidBlocks)
	assert.Equal(t, 0, bp.LastClaimTime)
	assert.Equal(t, 0, bp.Location)
}
