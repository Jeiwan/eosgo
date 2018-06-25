package types

// Block ...
type Block struct {
	// BlockExtensions   []BlockExtension `jons:"block_extensions"`
	// HeaderExtensions []HeaderExtension `json:"header_extension"`
	// NewProducers []Producer `json:"new_producers"`
	ActionMroot       string        `json:"action_mroot"`
	BlockNum          int           `json:"block_num"`
	Confirmed         int           `json:"confirmed"`
	ID                string        `json:"id"`
	Previous          string        `json:"previous"`
	Producer          string        `json:"producer"`
	ProducerSignature string        `json:"producer_signature"`
	RefBlockPrefix    int           `json:"ref_block_prefix"`
	ScheduleVersion   int           `json:"schedule_version"`
	Timestmap         Time          `json:"timestamp"`
	TransactionMroot  string        `json:"transaction_mroot"`
	Transactions      []Transaction `json:"transactions"`
}
