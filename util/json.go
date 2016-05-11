package util

import (
	"encoding/json"
)

func ParseJSON(raw []byte, data interface{}) error {
	err := json.Unmarshal(raw, data)
	switch e := err.(type) {
	case nil:
		return nil
	case *json.SyntaxError:
		return Error("JSON syntax error at offset %d: %v", e.Offset, e)
	default:
		return Error("Unrecognised JSON unmarshaling error: %v", e)
	}
}

func ParseJSONList(raw []byte) ([]json.RawMessage, error) {
	var list []json.RawMessage
	if err := ParseJSON(raw, &list); err != nil {
		return []json.RawMessage{}, err
	}
	return list, nil
}
