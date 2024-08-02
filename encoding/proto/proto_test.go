package proto

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"

	"github.com/blink-io/x/i18n/grpc"
	"github.com/blink-io/x/internal/testdata/pb"
	"github.com/segmentio/encoding/proto"
	"github.com/stretchr/testify/require"
)

func TestProto_1(t *testing.T) {
	cpb := &pb.TestingResponse{
		Code:    200,
		Message: "ok",
		Data: &pb.TestingResponse_Data{
			Action: "testing",
		},
	}
	data, err := proto.Marshal(cpb)
	require.NoError(t, err)
	require.NotNil(t, data)

	//fCode := eproto.FieldNumber(100).Int32(200)
	//fMessage := eproto.FieldNumber(200).String("ok")
	//fDataPayload := eproto.FieldNumber(1).String("testing")
	//fData := eproto.FieldNumber(1).Value(fDataPayload)
	//eproto.MessageRewriter{}
	data2 := testData1()

	fmt.Println("d1: ", base64.StdEncoding.EncodeToString(data))
	fmt.Println("d2: ", base64.StdEncoding.EncodeToString(data2))
}

func TestProto_2(t *testing.T) {
	data := testData2()
	var mm = new(pb.TestingResponse)
	err := proto.Unmarshal(data, mm)
	require.NoError(t, err)
}

func TestProto_3(t *testing.T) {
	pyd1, err1 := os.ReadFile("./testdata/zh-Hans.json")
	require.NoError(t, err1)

	pyd2, err2 := os.ReadFile("./testdata/en-US.json")
	require.NoError(t, err2)

	pyd3 := make([]byte, 0)

	e1 := createEntry("zh-Hans.json", "zh-Hans", true, pyd1)
	e2 := createEntry("en-US.json", "en-US", true, pyd2)
	e3 := createEntry("en-IN.json", "en-IN", false, pyd3)

	//entries := combineEntries(1, e1, e2)
	//require.NotNil(t, entries)

	var m proto.RawMessage
	m = proto.AppendVarlen(m, 1, e1)
	m = proto.AppendVarlen(m, 1, e2)
	m = proto.AppendVarlen(m, 1, e3)
	m = proto.AppendVarint(m, 2, 1701148888)

	var res = &grpc.ListLanguagesResponse{}
	require.NotNil(t, res)
	errx := proto.Unmarshal(m, res)
	require.NoError(t, errx)
}

func testData1() []byte {
	md := proto.AppendVarlen(nil, 1, []byte("testing"))
	m := proto.AppendVarlen(nil, 1, md)
	m = proto.AppendVarint(m, 100, 200)
	m = proto.AppendVarlen(m, 200, []byte("ok"))
	return m
}

func testData2() []byte {
	md := proto.AppendVarlen(nil, 1, []byte("testing为是一个测试瓦兹"))

	var m proto.RawMessage
	m = proto.AppendVarint(m, 100, 200)
	m = proto.AppendVarlen(m, 200, []byte("ok"))
	m = proto.AppendVarlen(m, 1, md)
	return m
}

func createEntry(path, language string, valid bool, payload []byte) []byte {
	ed := proto.AppendVarlen(nil, 1, []byte(path))
	ed = proto.AppendVarlen(ed, 2, []byte(language))
	ed = proto.AppendVarint(ed, 3, boolToUint64(valid))
	ed = proto.AppendVarlen(ed, 20, payload)

	kvd := proto.AppendVarlen(nil, 1, []byte(language))
	kvd = proto.AppendVarlen(kvd, 2, ed)
	return kvd
}

func boolToUint64(v bool) uint64 {
	if v {
		return 1
	} else {
		return 0
	}
}
