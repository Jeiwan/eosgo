package eosgo

import (
	"encoding/json"
	"math/big"
	"strconv"
)

// ProducersInfo ...
type ProducersInfo struct {
	Rows                    []ProducerInfo `json:"rows"`
	TotalProducerVoteWeight *big.Float     `json:"total_producer_vote_weight"`
	More                    string         `json:"more"`
}

// ProducerInfo ...
type ProducerInfo struct {
	Owner         string     `json:"owner"`
	TotalVotes    *big.Float `json:"total_votes"`
	ProducerKey   string     `json:"producer_key"`
	IsActive      int        `json:"is_active"`
	URL           string     `json:"url"`
	UnpaidBlocks  int        `json:"unpaid_blocks"`
	LastClaimTime int        `json:"last_claim_time"`
	Location      int        `json:"location"`
}

// UnmarshalJSON ...
func (pi *ProducerInfo) UnmarshalJSON(data []byte) error {
	type mirror ProducerInfo
	var check mirror

	if err := json.Unmarshal(data, &check); err != nil {
		typeErr, ok := err.(*json.UnmarshalTypeError)
		if ok && typeErr.Field == "last_claim_time" {
			dummy := struct {
				LastClaimTime string `json:"last_claim_time"`
			}{}
			if err = json.Unmarshal(data, &dummy); err != nil {
				return err
			}

			i, err := strconv.Atoi(dummy.LastClaimTime)
			if err != nil {
				return err
			}
			check.LastClaimTime = i
		} else {
			return err
		}
	}

	*pi = ProducerInfo(check)

	return nil
}
