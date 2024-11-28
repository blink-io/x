package tests

import (
	"testing"

	"github.com/amazon-ion/ion-go/ion"
	"github.com/stretchr/testify/require"
)

func TestIon_1(t *testing.T) {
	var d ion.Unmarshaler
	require.Nil(t, d)
}
