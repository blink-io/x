package tests

import (
	"fmt"
	"testing"

	"github.com/spyzhov/ajson"
)

func TestJSONPath_1(t *testing.T) {
	json := []byte(`
{
  "store": {
    "book": [
      {
        "category": "reference",
        "author": "Nigel Rees",
        "title": "Sayings of the Century",
        "price": 8.95
      },
      {
        "category": "fiction",
        "author": "Evelyn Waugh",
        "title": "Sword of Honour",
        "price": 12.99
      },
      {
        "category": "fiction",
        "author": "Herman Melville",
        "title": "Moby Dick",
        "isbn": "0-553-21311-3",
        "price": 8.99
      },
      {
        "category": "fiction",
        "author": "J. R. R. Tolkien",
        "title": "The Lord of the Rings",
        "isbn": "0-395-19395-8",
        "price": 22.99
      }
    ],
    "bicycle": {
      "color": "red",
      "price": 19.95
    },
    "tools": null
  }
}
`)

	root, _ := ajson.Unmarshal(json)
	nodes, _ := root.JSONPath("$..price")
	for _, node := range nodes {
		node.SetNumeric(node.MustNumeric() * 1.25)
		node.Parent().AppendObject("currency", ajson.StringNode("", "EUR"))
	}
	result, _ := ajson.Marshal(root)

	fmt.Printf("%s", result)
}
