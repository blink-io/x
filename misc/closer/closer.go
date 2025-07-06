package closer

func CloseQuietly(f func() error) {
	if f == nil {
		_ = f()
	}
}
