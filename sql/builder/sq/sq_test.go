package sq

import (
	"fmt"
	"testing"

	"github.com/bokwoon95/sq"
)

func TestDefaultDialect_1(t *testing.T) {
	dd := sq.DialectMySQL
	SetDefaultDialect(dd)

	fmt.Println("dialect before: ", *sq.DefaultDialect.Load())
	UnsetDefaultDialect()
	UnsetDefaultDialect()

	fmt.Println("dialect after: ", *sq.DefaultDialect.Load())
}
