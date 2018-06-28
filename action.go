package eosgo

import "encoding/json"

// Action ...
type Action struct {
	Account       string                 `json:"account"`
	Name          string                 `json:"name"`
	Authorization []Authorization        `json:"authorization"`
	Data          map[string]interface{} `json:"data"`
	HexData       json.RawMessage        `json:"hex_data"`
}

// UnmarshalJSON ...
func (a *Action) UnmarshalJSON(data []byte) error {
	type mirror Action
	var check mirror

	if err := json.Unmarshal(data, &check); err != nil {
		typeErr, ok := err.(*json.UnmarshalTypeError)
		if ok && typeErr.Field == "data" {
			dummy := struct {
				HexData json.RawMessage `json:"data"`
			}{}
			if err = json.Unmarshal(data, &dummy); err != nil {
				return err
			}

			check.Data = map[string]interface{}{}
			check.HexData = dummy.HexData
		} else {
			return err
		}
	}

	*a = Action(check)
	if len(a.HexData) > 0 {
		a.HexData = []byte(a.HexData)[1 : len(a.HexData)-1]
	}

	return nil
}
