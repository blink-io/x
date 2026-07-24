package id

import (
	"fmt"
	"testing"
)

func TestID_1(t *testing.T) {
	uuidv := UUID()
	uuidv4v := UUIDV4()
	shortiuuidv := ShortUUID()
	nanoidv := NanoID(16)
	ulidv := ULID()
	guidv := GUID()
	mid := MachineID()

	fmt.Println("UUID(),        len: ", len(uuidv), " id: ", uuidv)
	fmt.Println("UUIDV4(),      len: ", len(uuidv4v), " id: ", uuidv4v)
	fmt.Println("NanoID(),      len: ", len(nanoidv), " id: ", nanoidv)
	fmt.Println("ShortUUID(),   len: ", len(shortiuuidv), " id: ", shortiuuidv)
	fmt.Println("ULID(),        len: ", len(ulidv), " id: ", ulidv)
	fmt.Println("UserGUID(),        len: ", len(guidv), " id: ", guidv)
	fmt.Println("MachineID(),   len: ", len(mid), " id: ", mid)
}

func TestGUID_1(t *testing.T) {
	id := GUID()
	fmt.Println(id)
}
