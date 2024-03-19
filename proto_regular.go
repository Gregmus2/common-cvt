package convptr

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"time"
)

func FromFloatToProto(value float32) *wrappers.FloatValue {
	if value == 0.0 {
		return nil
	}

	return &wrappers.FloatValue{Value: value}
}

func FromFloat64ToProto(value float64) *wrappers.FloatValue {
	if value == 0.0 {
		return nil
	}

	return &wrappers.FloatValue{Value: float32(value)}
}

func FromStringToProto(value string) *wrappers.StringValue {
	if value == "" {
		return nil
	}

	return &wrappers.StringValue{Value: value}
}

func FromIntToProto(value int32) *wrappers.Int32Value {
	if value == 0 {
		return nil
	}

	return &wrappers.Int32Value{Value: value}
}

func FromInt64ToProto(value int64) *wrappers.Int64Value {
	if value == 0 {
		return nil
	}

	return &wrappers.Int64Value{Value: value}
}

func FromBoolToProto(value bool) *wrappers.BoolValue {
	if value == false {
		return nil
	}

	return &wrappers.BoolValue{Value: value}
}

func FromTimeToProto(value time.Time) *wrappers.Int64Value {
	if value.IsZero() {
		return nil
	}

	return &wrappers.Int64Value{Value: value.Unix()}
}

func ToStringFromProto(value *wrappers.StringValue) string {
	if value == nil {
		return ""
	}

	return value.Value
}

func ToFloatFromProto(value *wrappers.FloatValue) float64 {
	if value == nil {
		return 0.0
	}

	return float64(value.Value)
}

func ToIntFromProto(value *wrappers.Int32Value) int32 {
	if value == nil {
		return 0
	}

	return value.Value
}

func ToBoolFromProto(value *wrappers.BoolValue) bool {
	if value == nil {
		return false
	}

	return value.Value
}

func ToTimeFromProto(value *wrappers.Int64Value) time.Time {
	if value == nil {
		return time.Time{}
	}

	return time.Unix(value.Value, 0)
}

func FromFloat64ToDoubleProto(value float64) *wrappers.DoubleValue {
	if value == 0.0 {
		return nil
	}

	return &wrappers.DoubleValue{Value: value}
}
