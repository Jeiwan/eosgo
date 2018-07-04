package eosgo

import (
	"encoding/json"
)

// TransactionHeader ...
type TransactionHeader struct {
	Status        string `json:"status"`
	CPUUsageUs    int    `json:"cpu_usage_us"`
	NetUsageWords int    `json:"net_usage_words"`
	Trx           Trx    `json:"trx"`
}

// Trx ...
type Trx struct {
	Compression           string          `json:"compression"`
	ContextFreeData       []interface{}   `json:"context_free_data"`
	ID                    json.RawMessage `json:"id"`
	PackedContextFreeData json.RawMessage `json:"packed_context_free_data"`
	PackedTrx             json.RawMessage `json:"packed_trx"`
	Signatures            []string        `json:"signatures"`
	Transaction           Transaction     `json:"transaction"`
}

// UnmarshalJSON ...
func (t *Trx) UnmarshalJSON(data []byte) error {
	type mirror Trx
	var check mirror

	if err := json.Unmarshal(data, &check); err != nil {
		_, ok := err.(*json.UnmarshalTypeError)
		if ok {
			check.ID = json.RawMessage(data)
		} else {
			return err
		}
	}

	*t = Trx(check)
	if len(t.ID) > 0 {
		t.ID = t.ID[1 : len(t.ID)-1]
	}
	if len(t.PackedTrx) > 0 {
		t.PackedTrx = t.PackedTrx[1 : len(t.PackedTrx)-1]
	}
	if len(t.PackedContextFreeData) > 0 {
		t.PackedContextFreeData = t.PackedContextFreeData[1 : len(t.PackedContextFreeData)-1]
	}

	return nil
}

// Transaction ...
type Transaction struct {
	Actions            []Action `json:"actions"`
	ContextFreeActions []Action `json:"context_free_actions"`
	DelaySec           int      `json:"delay_sec"`
	Expiration         Time     `json:"expiration"`
	MaxCPUUsageMs      int      `json:"max_cpu_usage_ms"`
	MaxNetUsagWords    int      `json:"max_net_usage_words"`
	RefBlockNum        int      `json:"ref_block_num"`
	RefBlockPrefix     int      `json:"ref_block_prefix"`
}
