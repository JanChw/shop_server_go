package models

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Result struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (r Result) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, r)
}

func (r Result) MarshalBinary() ([]byte, error) {
	return json.Marshal(r)
}
