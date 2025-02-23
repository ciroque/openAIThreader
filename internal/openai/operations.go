package openai

import "encoding/json"

func UnmarshalResponse(data []byte) (*Response, error) {
	var r Response
	err := json.Unmarshal(data, &r)
	return &r, err
}
