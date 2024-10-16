package backoff

import (
	"testing"

	"github.com/cenkalti/backoff/v4"
)

func TestBackOff_1(t *testing.T) {
	// An operation that may fail.
	operation := func() error {
		t.Log("Run operation")
		return nil // or an error
	}

	err := backoff.Retry(operation, backoff.NewExponentialBackOff())
	if err != nil {
		// Handle error.
		return
	}

	// Operation is successful.
}
