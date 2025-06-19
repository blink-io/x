package orm

import (
	"fmt"
	"testing"

	"github.com/blink-io/sq"
	"github.com/stretchr/testify/require"
)

func TestSq_Fields_1(t *testing.T) {
	ts := sq.NewTableStruct("public", "test", "test")

	tf := sq.NewTimeField("created_at", ts)

	tf2 := tf.As("cat")

	aa := tf2.GetAlias()

	fmt.Println(aa)

	require.NotNil(t, tf2)
}

func TestSq_Fields_2(t *testing.T) {

}
