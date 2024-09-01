package dialect

import (
	"strings"
)

var (
	SQLite3 = sqlite3{}
)

func QuoteIdent(s, quote string) string {
	part := strings.SplitN(s, ".", 2)
	if len(part) == 2 {
		return QuoteIdent(part[0], quote) + "." + QuoteIdent(part[1], quote)
	}
	return quote + s + quote
}
