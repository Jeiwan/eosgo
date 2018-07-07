package eosgo

import "encoding/json"

// Block ...
type Block struct {
	ActionMroot       json.RawMessage          `json:"action_mroot"`
	BlockExtensions   []map[string]interface{} `json:"block_extensions"`
	BlockNum          int                      `json:"block_num"`
	Confirmed         int                      `json:"confirmed"`
	HeaderExtensions  []map[string]interface{} `json:"header_extension"`
	ID                json.RawMessage          `json:"id"`
	NewProducers      []map[string]interface{} `json:"new_producers"`
	Previous          json.RawMessage          `json:"previous"`
	Producer          string                   `json:"producer"`
	ProducerSignature string                   `json:"producer_signature"`
	RefBlockPrefix    int                      `json:"ref_block_prefix"`
	ScheduleVersion   int                      `json:"schedule_version"`
	Timestamp         Time                     `json:"timestamp"`
	TransactionMroot  json.RawMessage          `json:"transaction_mroot"`
	Transactions      []TransactionHeader      `json:"transactions"`
}

// UnmarshalJSON ...
func (b *Block) UnmarshalJSON(data []byte) error {
	type mirror Block
	var check mirror

	if err := json.Unmarshal(data, &check); err != nil {
		return err
	}

	*b = Block(check)
	if len(b.ID) > 0 {
		b.ID = b.ID[1 : len(b.ID)-1]
	}
	if len(b.Previous) > 0 {
		b.Previous = b.Previous[1 : len(b.Previous)-1]
	}
	if len(b.ActionMroot) > 0 {
		b.ActionMroot = b.ActionMroot[1 : len(b.ActionMroot)-1]
	}
	if len(b.TransactionMroot) > 0 {
		b.TransactionMroot = b.TransactionMroot[1 : len(b.TransactionMroot)-1]
	}

	return nil
}
