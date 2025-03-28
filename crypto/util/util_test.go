package util

import (
	"fmt"
	"github.com/go-crypt/crypt"
	"github.com/stretchr/testify/require"
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

	hashstr := "$argon2id$v=19$m=15360,t=2,p=1$dS6iORUFCpnJR3gewouTuiikildY6g9POB4Z6vXVCcN98hImCYg7EoFWpG7Hs1PjKZVxPw9Vm8qRCJoCmKFYkw$OYqDMR8p/TOVDvFgKFGpxVqSkcHZpqx01bg0qD/nH28\n"
	decoder, err := crypt.NewDefaultDecoder()
	require.NoError(t, err)

	decoded, err := decoder.Decode(hashstr)
	require.NoError(t, err)
	flag := decoded.Match("123456")
	require.True(t, flag)

}

func TestBCrypt1(t *testing.T) {
	hashstr := "$2b$12$kqB8qjhuxpEUYZYjKz/RHe9VgoZmg/3vIAjHXrQ0ruaq.AMBrqN1u"

	decoder, err := crypt.NewDefaultDecoder()
	require.NoError(t, err)

	d, err := decoder.Decode(hashstr)
	require.NoError(t, err)

	flag := d.Match("123456")
	require.True(t, flag)
}
