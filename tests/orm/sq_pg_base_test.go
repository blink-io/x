package orm

import (
	"testing"

	"github.com/bokwoon95/sq"
	"github.com/stretchr/testify/require"
)

func TestSq_Fields_1(t *testing.T) {
	fields := sq.Fields{}

	require.NotNil(t, fields)
}
