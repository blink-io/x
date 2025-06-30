package typetostring

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestGeneric_1(t *testing.T) {
	s := GetType[sql.Null[string]]()
	fmt.Println(s)
}
