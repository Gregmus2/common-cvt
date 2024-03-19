package convptr

import (
	"google.golang.org/protobuf/types/known/wrapperspb"
	"time"
)

func FromFloatToProto(value float32) *wrapperspb.FloatValue {
	if value == 0.0 {
		return nil
	}

	return &wrapperspb.FloatValue{Value: value}
}

func FromFloat64ToProto(value float64) *wrapperspb.FloatValue {
	if value == 0.0 {
		return nil
	}

	return &wrapperspb.FloatValue{Value: float32(value)}
}

func FromStringToProto(value string) *wrapperspb.StringValue {
	if value == "" {
		return nil
	}

	return &wrapperspb.StringValue{Value: value}
}

func FromIntToProto(value int32) *wrapperspb.Int32Value {
	if value == 0 {
		return nil
	}

	return &wrapperspb.Int32Value{Value: value}
}

func FromInt64ToProto(value int64) *wrapperspb.Int64Value {
	if value == 0 {
		return nil
	}

	return &wrapperspb.Int64Value{Value: value}
}

func FromBoolToProto(value bool) *wrapperspb.BoolValue {
	if value == false {
		return nil
	}

	return &wrapperspb.BoolValue{Value: value}
}

func FromTimeToProto(value time.Time) *wrapperspb.Int64Value {
	if value.IsZero() {
		return nil
	}

	return &wrapperspb.Int64Value{Value: value.Unix()}
}

func ToStringFromProto(value *wrapperspb.StringValue) string {
	if value == nil {
		return ""
	}

	return value.Value
}

func ToFloatFromProto(value *wrapperspb.FloatValue) float64 {
	if value == nil {
		return 0.0
	}

	return float64(value.Value)
}

func ToIntFromProto(value *wrapperspb.Int32Value) int32 {
	if value == nil {
		return 0
	}

	return value.Value
}

func ToBoolFromProto(value *wrapperspb.BoolValue) bool {
	if value == nil {
		return false
	}

	return value.Value
}

func ToTimeFromProto(value *wrapperspb.Int64Value) time.Time {
	if value == nil {
		return time.Time{}
	}

	return time.Unix(value.Value, 0)
}

func FromFloat64ToDoubleProto(value float64) *wrapperspb.DoubleValue {
	if value == 0.0 {
		return nil
	}

	return &wrapperspb.DoubleValue{Value: value}
}
