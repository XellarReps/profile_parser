package parser

import "encoding/json"

type Arguments struct {
	Type string `json:"op_name"`
}

type ProfileInfo struct {
	Category string    `json:"cat"`
	Time     int       `json:"dur"`
	Name     string    `json:"name"`
	Args     Arguments `json:"args"`
}

func ParseJSON(buf []byte) ([]ProfileInfo, error) {
	var nodesInfo []ProfileInfo
	err := json.Unmarshal(buf, &nodesInfo)
	if err != nil {
		return nil, err
	}
	return nodesInfo, nil
}
