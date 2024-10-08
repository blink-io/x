package pie

import (
	"github.com/elliotchance/pie/v2"
)

func Contains[T comparable](ss []T, lookingFor T) bool {
	return pie.Contains(ss, lookingFor)
}

func First[T any](ss []T) T {
	return pie.First(ss)
}

func Last[T any](ss []T) T {
	return pie.First(ss)
}
