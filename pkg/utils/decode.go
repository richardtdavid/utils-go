package jsondecode

import (
	"encoding/json"
	"log"
)

func DecodeBody(body []byte) (map[string]interface{}, error) {
	var jsonMap map[string]interface{}
	err := json.Unmarshal(body, &jsonMap)
	if err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("Syntax error at byte offset: %d", e.Offset)
		}
		return nil, err
	}
	return jsonMap, nil
}
