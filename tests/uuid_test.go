package tests

import (
	"fmt"
	"testing"

	fuuid "github.com/gofrs/uuid/v5"
	guuid "github.com/google/uuid"
	puuid "github.com/pborman/uuid"
)

type (
	GUUID = guuid.UUID
	FUUID = fuuid.UUID
	PUUID = puuid.UUID
)

func TestUUID_1(t *testing.T) {
	fu, _ := fuuid.NewV4()
	gu := guuid.NewString()
	pu, _ := fuuid.NewV4()
	fmt.Println(fu)
	fmt.Println(gu)
	fmt.Println(pu)
}
