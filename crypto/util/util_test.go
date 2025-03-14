package util

import (
	"fmt"
	"testing"

	"github.com/go-crypt/crypt/algorithm"
	"github.com/go-crypt/crypt/algorithm/argon2"
)

func TestArgon2d(t *testing.T) {
	var (
		hasher *argon2.Hasher
		err    error
		digest algorithm.Digest
	)

	if hasher, err = argon2.New(
		argon2.WithProfileRFC9106Recommended(),
	); err != nil {
		panic(err)
	}

	if digest, err = hasher.Hash("123456"); err != nil {
		panic(err)
	}

	fmt.Printf("Encoded Digest With Password 'example': %s\n", digest.Encode())
}
