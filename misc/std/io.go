package std

import "io"

func CloseQuietly(c io.Closer) {
	func() {
		_ = c.Close()
	}()
}
