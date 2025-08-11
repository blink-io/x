package event

import (
	"testing"

	"github.com/kelindar/event"
	"github.com/stretchr/testify/require"
)

func TestEvent_1(t *testing.T) {
	d := event.NewDispatcher()
	require.NotNil(t, d)
}
