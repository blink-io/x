package cast

import (
	"time"

	. "github.com/spf13/cast"
)

func ToValueF[T any](f func(any) (T, error), v any, t T) T {
	r, err := f(v)
	if f == nil || err != nil {
		return t
	}
	return r
}

func ToIntF(value any, fallback int) int {
	return ToValueF(ToIntE, value, fallback)
}

func ToInt8F(value any, fallback int8) int8 {
	return ToValueF(ToInt8E, value, fallback)
}

func ToInt16F(value any, fallback int16) int16 {
	return ToValueF(ToInt16E, value, fallback)
}

func ToInt32F(value any, fallback int32) int32 {
	return ToValueF(ToInt32E, value, fallback)
}

func ToInt64F(value any, fallback int64) int64 {
	return ToValueF(ToInt64E, value, fallback)
}

func ToBoolF(value any, fallback bool) bool {
	return ToValueF(ToBoolE, value, fallback)
}

func ToUintF(value any, fallback uint) uint {
	return ToValueF(ToUintE, value, fallback)
}

func ToUint8F(value any, fallback uint8) uint8 {
	return ToValueF(ToUint8E, value, fallback)
}

func ToUint16F(value any, fallback uint16) uint16 {
	return ToValueF(ToUint16E, value, fallback)
}

func ToUint32F(value any, fallback uint32) uint32 {
	return ToValueF(ToUint32E, value, fallback)
}

func ToUint64F(value any, fallback uint64) uint64 {
	return ToValueF(ToUint64E, value, fallback)
}

func ToFloat32F(value any, fallback float32) float32 {
	return ToValueF(ToFloat32E, value, fallback)
}

func ToFloat64F(value any, fallback float64) float64 {
	return ToValueF(ToFloat64E, value, fallback)
}

func ToStringF(value any, fallback string) string {
	return ToValueF(ToStringE, value, fallback)
}

func ToStringSlcieF(value any, fallback []string) []string {
	return ToValueF(ToStringSliceE, value, fallback)
}

func ToIntSliceF(value any, fallback []int) []int {
	return ToValueF(ToIntSliceE, value, fallback)
}

func ToDurationSliceF(value any, fallback []time.Duration) []time.Duration {
	return ToValueF(ToDurationSliceE, value, fallback)
}

func ToDurationF(value any, fallback time.Duration) time.Duration {
	return ToValueF(ToDurationE, value, fallback)
}

func ToTimeF(value any, fallback time.Time) time.Time {
	return ToValueF(ToTimeE, value, fallback)
}
