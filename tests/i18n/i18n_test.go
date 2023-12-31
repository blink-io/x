package i18n

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/blink-io/x/i18n"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Loader_1(t *testing.T) {
	bundle := i18n.Default()
	apath, err := filepath.Abs(filepath.Join(".", "locales"))
	require.NoError(t, err)
	fmt.Println("abs path:  ", apath)

	langs := []string{"zh-Hans", "en-US"}
	require.NoError(t, i18n.NewDirLoader(apath).Load(bundle))

	for _, lang := range langs {
		tr := i18n.GetT(lang)

		msg := tr("hello", i18n.MD{"Accessor": "兜兜"}.O())
		fmt.Printf("msg: %s\n", msg)
	}
}

func TestParsePath_1(t *testing.T) {
	urlstr := "https://xxx.com/languages/zh-Hans.toml"
	langTag, format := parsePath(urlstr)
	assert.Equal(t, "zh-Hans", langTag)
	assert.Equal(t, "toml", format)

	urlstr2 := "https://xxx.com/languages/en-US.yaml"
	langTag2, format2 := parsePath(urlstr2)
	assert.Equal(t, "en-US", langTag2)
	assert.Equal(t, "yaml", format2)
}

func TestBytesLoader_1(t *testing.T) {
	bundle := i18n.Default()
	apath, err := filepath.Abs(filepath.Join(".", "locales", "zh-Hans.toml"))
	require.NoError(t, err)

	bytes, err := os.ReadFile(apath)
	require.NoError(t, err)

	require.NoError(t,
		i18n.NewBytesLoader(apath, bytes).Load(bundle),
	)
	tr := i18n.GetT("zh-Hans")
	msg := tr("hello", i18n.MD{"Accessor": "兜兜", "PhoneCount": 18}.O())
	fmt.Println("translated:  ", msg)
}

func parsePath(path string) (langTag, format string) {
	formatStartIdx := -1
	for i := len(path) - 1; i >= 0; i-- {
		c := path[i]
		if os.IsPathSeparator(c) {
			if formatStartIdx != -1 {
				langTag = path[i+1 : formatStartIdx]
			}
			return
		}
		if path[i] == '.' {
			if formatStartIdx != -1 {
				langTag = path[i+1 : formatStartIdx]
				return
			}
			if formatStartIdx == -1 {
				format = path[i+1:]
				formatStartIdx = i
			}
		}
	}
	if formatStartIdx != -1 {
		langTag = path[:formatStartIdx]
	}
	return
}
