package i18n

import (
	"context"
	"fmt"
	"testing"

	"github.com/apache/thrift/lib/go/thrift"
	i18nthrift "github.com/blink-io/x/i18n/thrift"
	"github.com/stretchr/testify/require"
)

func TestThriftLoader_1(t *testing.T) {
	langs := []string{"zh-Hans"}
	ld, err := NewThriftLoader("localhost:19099", langs)
	require.NoError(t, err)

	require.NoError(t, ld.Load(bb))
}

func TestThriftLoader_HTTP_1(t *testing.T) {
	langs := []string{"zh-Hans"}
	ld, err := NewThriftLoader("http://localhost:19099", langs, WithUseHTTP())
	require.NoError(t, err)

	require.NoError(t, ld.Load(bb))
}

func TestThriftClient_1(t *testing.T) {
	addr := "localhost:19099"

	var cfg = &thrift.TConfiguration{
		SocketTimeout:  DefaultTimeout,
		ConnectTimeout: DefaultTimeout,
	}
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(cfg)

	transport, err := transportFactory.GetTransport(thrift.NewTSocketConf(addr, cfg))
	require.NoError(t, err)

	defer transport.Close()

	require.NoError(t, transport.Open())

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	cc := i18nthrift.NewI18NClient(thrift.NewTStandardClient(iprot, oprot))

	res, err := cc.ListLanguages(context.Background(), &i18nthrift.ListLanguagesRequest{
		Languages: []string{"zh-Hans"},
	})
	require.NoError(t, err)

	fmt.Println("res: ", res)
}

func TestThriftServer_1(t *testing.T) {
	zhHansJSON := `{"name":"广州", "language":"简体中文"}`
	enUSJSON := `{"name":"gz", "language":"American English"}`

	entries := map[string]*Entry{
		"zh-Hans": {
			Path:     "zh-Hans.json",
			Language: "zh-Hans",
			Valid:    true,
			Payload:  []byte(zhHansJSON),
		},
		"en-US": {
			Path:     "en-US.json",
			Language: "en-US",
			Valid:    true,
			Payload:  []byte(enUSJSON),
		},
		"en-UK": {
			Path:     "en-UK.json",
			Language: "en-UK",
			Valid:    false,
			Payload:  []byte(""),
		},
	}

	var ff = Entries(entries)

	addr := "localhost:19099"
	useHTTP := true

	p := i18nthrift.NewI18NProcessor(&ThriftHandler{h: ff})

	trans, err := thrift.NewTServerSocket(addr)
	require.NoError(t, err)

	var transportFactory thrift.TTransportFactory
	if useHTTP {
		transportFactory = thrift.NewTHttpClientTransportFactory("http://" + addr)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(nil)

	server := thrift.NewTSimpleServer4(p, trans, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", addr)

	serr := server.Serve()
	require.NoError(t, serr)
}

func TestThriftServer_2(t *testing.T) {
	zhHansJSON := `{"name":"广州", "language":"简体中文"}`
	enUSJSON := `{"name":"gz", "language":"American English"}`

	entries := map[string]*Entry{
		"zh-Hans": {
			Path:     "zh-Hans.json",
			Language: "zh-Hans",
			Valid:    true,
			Payload:  []byte(zhHansJSON),
		},
		"en-US": {
			Path:     "en-US.json",
			Language: "en-US",
			Valid:    true,
			Payload:  []byte(enUSJSON),
		},
		"en-UK": {
			Path:     "en-UK.json",
			Language: "en-UK",
			Valid:    false,
			Payload:  []byte(""),
		},
	}

	var ff = Entries(entries)

	addr := "localhost:19099"

	srv, err := NewTBinaryServer(addr, ff)
	require.NoError(t, err)

	fmt.Println("Starting the simple server... on ", addr)

	require.NoError(t, srv.Serve())
}
