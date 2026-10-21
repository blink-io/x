package std

import "io"

func CloseQuietly(c io.Closer) {
	if c == nil {
		return
	}
	_ = c.Close()
}
