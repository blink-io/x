package jsonx

import (
	"encoding/json"
	"testing"
)

func TestJSONMap_1(t *testing.T) {
	var jm = JSONMap{
		"name":  "Mary",
		"age":   18,
		"score": 88.1,
		"data": JSONMap{
			"gpa": 3.21,
		},
	}

	jsonstr, _ := json.Marshal(jm)
	t.Log(string(jsonstr))
}
