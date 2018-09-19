package types_test

import (
	"encoding/json"
	"testing"

	types "github.com/Jeiwan/eosgo/types"
	"github.com/stretchr/testify/assert"
)

func TestActionUnmarshalling(t *testing.T) {
	data := []byte(`
		{
			"account": "eosio",
			"name": "test",
			"authorization": [{
				"actor": "eosio",
				"permission": "owner"
			}],
			"data": {
				"param": "value",
				"param2": "value2"
			},
			"hex_data": "deadbeef"
		}
	`)

	var a types.Action
	err := json.Unmarshal(data, &a)

	assert.Nil(t, err)
	assert.Equal(t, "eosio", a.Account)
	assert.Equal(t, "test", a.Name)
	assert.Len(t, a.Authorization, 1)
	assert.Equal(t, "eosio", a.Authorization[0].Actor)
	assert.Equal(t, "owner", a.Authorization[0].Permission)
	assert.Equal(t, "value", a.Data["param"])
	assert.Equal(t, "value2", a.Data["param2"])
	assert.Equal(t, "deadbeef", string(a.HexData))

	data = []byte(`
		{
			"account": "eosio",
			"name": "test",
			"authorization": [{
				"actor": "eosio",
				"permission": "owner"
			}],
			"data": "deadbeef",
			"hex_data": ""
		}
	`)

	a = types.Action{}
	err = json.Unmarshal(data, &a)

	assert.Nil(t, err)
	assert.Equal(t, "eosio", a.Account)
	assert.Equal(t, "test", a.Name)
	assert.Len(t, a.Authorization, 1)
	assert.Equal(t, "eosio", a.Authorization[0].Actor)
	assert.Equal(t, "owner", a.Authorization[0].Permission)
	assert.Empty(t, a.Data)
	assert.Equal(t, "deadbeef", string(*a.HexData))
}
