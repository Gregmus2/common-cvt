package convptr

import (
	"strconv"
	"time"
)

func FromString(value string) *string {
	if value == "" {
		return nil
	}

	return &value
}

func FromInt(value int32) *int32 {
	return &value
}

func FromInt64(value int64) *int64 {
	return &value
}

func FromFloat32(value float32) *float32 {
	return &value
}

func FromFloat64(value float64) *float64 {
	return &value
}

func FromBool(value bool) *bool {
	return &value
}

func FromTime(value time.Time) *time.Time {
	return &value
}

func ToString(value *string) string {
	if value == nil {
		return ""
	}

	return *value
}

func ToInt(value *int32) int32 {
	if value == nil {
		return 0
	}

	return *value
}

func ToInt64(value *int64) int64 {
	if value == nil {
		return 0
	}

	return *value
}

func ToFloat32(value *float32) float32 {
	if value == nil {
		return 0.0
	}

	return *value
}

func ToFloat64(value *float64) float64 {
	if value == nil {
		return 0.0
	}

	return *value
}

func ToBool(value *bool) bool {
	if value == nil {
		return false
	}

	return *value
}

func ToTime(value *time.Time) time.Time {
	if value == nil {
		return time.Now()
	}

	return *value
}

func ToIntString(value *int32) string {
	if value == nil {
		return ""
	}

	return strconv.Itoa(int(*value))
}

func ToFloat32String(value *float32) string {
	if value == nil {
		return ""
	}

	return strconv.FormatFloat(float64(*value), 'G', -1, 32)
}

func ToFloat64String(value *float64) string {
	if value == nil {
		return ""
	}

	return strconv.FormatFloat(*value, 'G', -1, 64)
}

func ToBoolString(value *bool) string {
	if value == nil {
		return ""
	}

	return strconv.FormatBool(*value)
}
