package eosgo

import "encoding/json"

// Trace ...
type Trace struct {
	Receipt       TraceReceipt    `json:"receipt"`
	Act           Action          `json:"act"`
	Elapsed       int64           `json:"elapsed"`
	CPUUsage      int64           `json:"cpu_usage"`
	Console       string          `json:"console"`
	TotalCPUUsage int64           `json:"total_cpu_usage"`
	TrxID         json.RawMessage `json:"trx_id"`
	InlineTraces  []Trace         `json:"inline_traces"`
}

// TraceReceipt ...
type TraceReceipt struct {
	Receiver       string          `json:"receiver"`
	ActDigest      json.RawMessage `json:"act_digest"`
	GlobalSequence uint64          `json:"global_sequence"`
	RecvSequence   uint64          `json:"recv_sequence"`
	AuthSequence   [][]interface{} `json:"auth_sequence"`
	CodeSequence   uint64          `json:"code_sequence"`
	AbiSequence    uint64          `json:"abi_sequence"`
}

// UnmarshalJSON ...
func (t *Trace) UnmarshalJSON(data []byte) error {
	type mirror Trace
	var check mirror

	if err := json.Unmarshal(data, &check); err != nil {
		return err
	}

	*t = Trace(check)
	if len(t.TrxID) > 0 {
		t.TrxID = t.TrxID[1 : len(t.TrxID)-1]
	}

	return nil
}
