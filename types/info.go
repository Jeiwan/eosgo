package types

// Info ...
type Info struct {
	BlockCPULimit            int         `json:"block_cpu_limit"`
	BlockNetLimit            int         `json:"block_net_limit"`
	ChainID                  string      `json:"chain_id"`
	HeadBlockID              string      `json:"head_block_id"`
	HeadBlockNum             int         `json:"head_block_num"`
	HeadBlockProducer        string      `json:"head_block_producer"`
	HeadBlockTime            Time        `json:"head_block_time"`
	LastIrreversibleBlockID  string      `json:"last_irreversible_block_id"`
	LastIrreversibleBlockNum int         `json:"last_irreversible_block_num"`
	ParticipationRate        string      `json:"participation_rate"`
	RecentSlots              string      `json:"recent_slots"`
	ServerVersion            string      `json:"server_version"`
	VirtualBlockCPULimit     interface{} `json:"virtual_block_cpu_limit"`
	VirtualBlockNetLimit     interface{} `json:"virtual_block_net_limit"`
}
