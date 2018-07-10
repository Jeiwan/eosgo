package eosgo

import "encoding/json"

// FullTransaction ...
type FullTransaction struct {
	ID       json.RawMessage `json:"id"`
	BlockNum uint64          `json:"block_num"`
	Status   string          `json:"status"`
	// Trx Trx
	Traces []Trace `json:"traces"`
}

// UnmarshalJSON ...
func (tx *FullTransaction) UnmarshalJSON(data []byte) error {
	type mirror FullTransaction
	var check mirror

	if err := json.Unmarshal(data, &check); err != nil {
		return err
	}

	*tx = FullTransaction(check)
	if len(tx.ID) > 0 {
		tx.ID = tx.ID[1 : len(tx.ID)-1]
	}

	return nil
}
