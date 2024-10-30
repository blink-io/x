package tests

import (
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/samber/oops"
	"github.com/stretchr/testify/require"
)

func TestOops_1(t *testing.T) {
	err1 := oops.
		In("repository").
		Tags("database", "sql").
		Recover(func() {
			panic("caramba!")
		})

	require.NoError(t, err1)

	req, _ := http.NewRequest("POST", "http://localhost:1337/foobar", strings.NewReader("hello world"))

	err3 := oops.
		Code("iam_authz_missing_permission").
		In("authz").
		Time(time.Now()).
		With("user_id", 1234).
		With("permission", "post.create").
		Hint("Runbook: https://doc.acme.org/doc/abcd.md").
		User("user-123", "firstname", "john", "lastname", "doe").
		Request(req, true).
		Errorf("permission denied")

	require.NoError(t, err3)
}
