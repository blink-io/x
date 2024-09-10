package misc

func ToOffset(page int, perPage int) int {
	if page > 1 {
		return (page - 1) * perPage
	} else {
		return 0
	}
}
