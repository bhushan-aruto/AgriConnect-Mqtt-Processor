package entity

import "encoding/json"

type MoistureSettings struct {
	Max uint32 `json:"max"`
	Min uint32 `json:"min"`
}

func NewMoistureSettings(max, min uint32) []byte {
	jsonBytes, _ := json.Marshal(
		MoistureSettings{
			Max: max,
			Min: min,
		},
	)
	return jsonBytes
}
