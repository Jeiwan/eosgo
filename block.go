package eosgo

import "encoding/json"

// Block ...
type Block struct {
	// BlockExtensions   []BlockExtension `jons:"block_extensions"`
	// HeaderExtensions []HeaderExtension `json:"header_extension"`
	// NewProducers []Producer `json:"new_producers"`
	ActionMroot       json.RawMessage     `json:"action_mroot"`
	BlockNum          int                 `json:"block_num"`
	Confirmed         int                 `json:"confirmed"`
	ID                json.RawMessage     `json:"id"`
	Previous          json.RawMessage     `json:"previous"`
	Producer          string              `json:"producer"`
	ProducerSignature string              `json:"producer_signature"`
	RefBlockPrefix    int                 `json:"ref_block_prefix"`
	ScheduleVersion   int                 `json:"schedule_version"`
	Timestamp         Time                `json:"timestamp"`
	TransactionMroot  json.RawMessage     `json:"transaction_mroot"`
	Transactions      []TransactionHeader `json:"transactions"`
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
