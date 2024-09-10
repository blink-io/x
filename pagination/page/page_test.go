package page_test

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/blink-io/x/log/slog/handlers/color"
	"github.com/blink-io/x/log/slog/handlers/colorized"
	"github.com/blink-io/x/pagination/page"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _log = slog.New(color.New(os.Stderr, color.Options{Level: slog.LevelDebug}))
var _clog = slog.New(colorized.New(os.Stderr, &colorized.Options{Level: slog.LevelDebug}))

func toJSONStr[E any](v page.Result[E]) string {
	data, err := json.Marshal(&v)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func TestPageBased_Invalid_1(t *testing.T) {
	records := make([]string, 0)
	t.Run("page is 0", func(t *testing.T) {
		r := page.NewResult[string](0, 1, 0, records)

		_clog.Info(toJSONStr(r))

		assert.Equal(t, 1, r.Pagination.Page)
	})

	t.Run("per_page is 0", func(t *testing.T) {
		r := page.NewResult[string](10, 0, 0, records)

		_clog.Info(toJSONStr(r))

		assert.Equal(t, 10, r.Pagination.PerPage)
	})

	t.Run("page and per_page are 0", func(t *testing.T) {
		r := page.NewResult[string](0, 0, 0, records)

		_clog.Info(toJSONStr(r))

		assert.Equal(t, 1, r.Pagination.Page)
		assert.Equal(t, 10, r.Pagination.PerPage)
	})
}

func TestPageBased_NoTotals_1(t *testing.T) {
	emptyRecords := make([]string, 0)
	totalRecords := 0
	records := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
	}

	t.Run("NoTotal with non-empty records returns has_more is TRUE", func(t *testing.T) {
		p := page.NewResult[string](1, 5, totalRecords, records)

		_log.Info(toJSONStr(p))
		assert.True(t, p.Pagination.HasMore)
	})

	t.Run("NoTotal with empty records returns has_more is FALSE", func(t *testing.T) {
		p := page.NewResult[string](1, 5, totalRecords, emptyRecords)

		_log.Info(toJSONStr(p))
		assert.False(t, p.Pagination.HasMore)
	})
}

func TestPageBased_Totals_1(t *testing.T) {
	records := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
	}

	t.Run("Calculate totalPages 1", func(t *testing.T) {
		p := page.NewResult[string](1, 5, 122, records)

		_log.Info(toJSONStr(p))

		assert.Equal(t, 25, p.Pagination.TotalPages)
		assert.True(t, p.Pagination.HasMore)
	})

	t.Run("Calculate totalPages 2", func(t *testing.T) {
		p := page.NewResult[string](1, 5, 120, records)
		jstr2, err2 := json.Marshal(p)
		require.NoError(t, err2)

		fmt.Println(string(jstr2))

		assert.Equal(t, 24, p.Pagination.TotalPages)
		assert.True(t, p.Pagination.HasMore)
	})

	t.Run("Field has_more is true 1", func(t *testing.T) {
		p := page.NewResult[string](23, 5, 120, records)

		_log.Info(toJSONStr(p))

		assert.True(t, p.Pagination.HasMore)
	})

	t.Run("Field has_more is false 1", func(t *testing.T) {
		p := page.NewResult[string](24, 5, 120, records)

		_log.Info(toJSONStr(p))

		assert.False(t, p.Pagination.HasMore)
	})

	t.Run("Field has_more is false 2", func(t *testing.T) {
		p := page.NewResult[string](25, 5, 120, records)

		_log.Info(toJSONStr(p))

		assert.False(t, p.Pagination.HasMore)
	})
}
