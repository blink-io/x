package registrar

import (
	"testing"

	"github.com/go-co-op/gocron/v2"
	"github.com/stretchr/testify/require"
)

func TestIface(t *testing.T) {
	var rr ServiceRegistrar
	c, err := gocron.NewScheduler(nil)
	require.NoError(t, err)
	rr = c
	require.NotNil(t, rr)
}
