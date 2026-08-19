package id

import (
	"fmt"
	"testing"

	gofrsuuid "github.com/gofrs/uuid/v5"
)

func TestID_1(t *testing.T) {
	uuidv := UUID()
	uuidv4v := UUIDV4()
	shortiuuidv := ShortUUID()
	nanoidv := NanoID(16)
	ulidv := ULID()
	mid := MachineID()

	uuidv7str := gofrsuuid.Must(gofrsuuid.NewV7())
	fmt.Printf("uuid v7: %s\n", uuidv7str)

	fmt.Println("UUID(),        len: ", len(uuidv), " id: ", uuidv)
	fmt.Println("UUIDV4(),      len: ", len(uuidv4v), " id: ", uuidv4v)
	fmt.Println("NanoID(),      len: ", len(nanoidv), " id: ", nanoidv)
	fmt.Println("ShortUUID(),   len: ", len(shortiuuidv), " id: ", shortiuuidv)
	fmt.Println("ULID(),        len: ", len(ulidv), " id: ", ulidv)
	fmt.Println("MachineID(),   len: ", len(mid), " id: ", mid)
}
