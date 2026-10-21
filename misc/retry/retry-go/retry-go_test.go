package retry_go

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestDoRetriesUntilSuccess(t *testing.T) {
	calls := 0

	err := Do(func() error {
		calls++
		if calls < 3 {
			return errors.New("not yet")
		}

		return nil
	}, Attempts(5), Delay(time.Millisecond))

	if err != nil {
		t.Fatalf("Do() error = %v, want nil", err)
	}

	if calls != 3 {
		t.Errorf("Do() called func %d times, want 3", calls)
	}
}

func TestDoReturnsErrorAfterAttempts(t *testing.T) {
	calls := 0
	wantErr := errors.New("always fails")

	err := Do(func() error {
		calls++

		return wantErr
	}, Attempts(3), Delay(time.Millisecond))

	if err == nil {
		t.Fatal("Do() error = nil, want an error")
	}

	if calls != 3 {
		t.Errorf("Do() called func %d times, want 3", calls)
	}

	if !errors.Is(err, wantErr) {
		t.Errorf("Do() error = %v, want it to wrap %v", err, wantErr)
	}
}

func TestDoWithData(t *testing.T) {
	calls := 0

	got, err := DoWithData(func() (string, error) {
		calls++
		if calls < 2 {
			return "", errors.New("not yet")
		}

		return "ok", nil
	}, Attempts(5), Delay(time.Millisecond))

	if err != nil {
		t.Fatalf("DoWithData() error = %v, want nil", err)
	}

	if got != "ok" {
		t.Errorf("DoWithData() = %q, want %q", got, "ok")
	}
}

func TestUnrecoverableStopsRetrying(t *testing.T) {
	calls := 0

	err := Do(func() error {
		calls++

		return Unrecoverable(errors.New("fatal"))
	}, Attempts(5), Delay(time.Millisecond))

	if err == nil {
		t.Fatal("Do() error = nil, want an error")
	}

	if calls != 1 {
		t.Errorf("Do() called func %d times, want 1", calls)
	}
}

func TestNewRetrierIsReusable(t *testing.T) {
	retrier := New(Attempts(3), Delay(time.Millisecond))

	for i := range 3 {
		calls := 0

		err := retrier.Do(func() error {
			calls++

			return nil
		})

		if err != nil {
			t.Fatalf("run %d: Do() error = %v, want nil", i, err)
		}

		if calls != 1 {
			t.Errorf("run %d: Do() called func %d times, want 1", i, calls)
		}
	}
}

func TestContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := Do(func() error {
		return errors.New("always fails")
	}, Attempts(5), Delay(time.Millisecond), Context(ctx))

	if err == nil {
		t.Fatal("Do() error = nil, want an error")
	}
}

func BenchmarkDoFreshRetrier(b *testing.B) {
	for b.Loop() {
		_ = Do(func() error { return nil }, Attempts(3))
	}
}

func BenchmarkDoReusedRetrier(b *testing.B) {
	retrier := New(Attempts(3))

	for b.Loop() {
		_ = retrier.Do(func() error { return nil })
	}
}
