//go:build !go1.10
// +build !go1.10

package hooks

import (
	"context"
	"database/sql/driver"
	"errors"
)

func isSessionResetter(conn driver.Conn) bool {
	return false
}

func (s *SessionResetter) ResetSession(ctx context.Context) error {
	return errors.New("SessionResetter not implemented")
}
