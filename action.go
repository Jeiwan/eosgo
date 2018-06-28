package eosgo

import "encoding/json"

// Action ...
type Action struct {
	Account       string                 `json:"account"`
	Name          string                 `json:"name"`
	Authorization []Authorization        `json:"authorization"`
	Data          map[string]interface{} `json:"data"`
	HexData       string                 `json:"hex_data"`
}

// UnmarshalJSON ...
func (a *Action) UnmarshalJSON(data []byte) error {
	type mirror Action
	var check mirror

	if err := json.Unmarshal(data, &check); err != nil {
		typeErr, ok := err.(*json.UnmarshalTypeError)
		if ok && typeErr.Field == "data" {
			dummy := struct {
				HexData string `json:"data"`
			}{}
			if err = json.Unmarshal(data, &dummy); err != nil {
				return err
			}

			check.Data = map[string]interface{}{}
			check.HexData = dummy.HexData

			*a = Action(check)

			return nil
		}

		return err
	}

	*a = Action(check)

	return nil
}
