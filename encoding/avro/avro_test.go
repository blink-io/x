package avro

import (
	"testing"

	"github.com/linkedin/goavro/v2"
	"github.com/stretchr/testify/require"
)

func TestAvro_Codec_1(t *testing.T) {
	codec, err := goavro.NewCodec(`
    {
      "type": "record",
      "name": "LongList",
      "fields" : [
        {"name": "next", "type": ["null", "LongList"], "default": null}
      ]
    }`)
	require.NoError(t, err)
	require.NotNil(t, codec)
}
