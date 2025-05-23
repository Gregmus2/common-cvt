package ccvt

import (
	"google.golang.org/protobuf/types/known/wrapperspb"
	"time"
)

func ToProtoStringWrapper(value *string) *wrapperspb.StringValue {
	if value != nil {
		return &wrapperspb.StringValue{Value: *value}
	}

	return nil
}

func ToProtoInt32Wrapper(value *int32) *wrapperspb.Int32Value {
	if value != nil {
		return &wrapperspb.Int32Value{Value: *value}
	}

	return nil
}

func ToProtoInt64Wrapper(value *int64) *wrapperspb.Int64Value {
	if value != nil {
		return &wrapperspb.Int64Value{Value: *value}
	}

	return nil
}

func ToProtoFloatWrapper(value *float32) *wrapperspb.FloatValue {
	if value != nil {
		return &wrapperspb.FloatValue{Value: *value}
	}

	return nil
}

func ToProtoFloat64Wrapper(value *float64) *wrapperspb.FloatValue {
	if value != nil {
		return &wrapperspb.FloatValue{Value: float32(*value)}
	}

	return nil
}

func ToProtoBoolWrapper(value *bool) *wrapperspb.BoolValue {
	if value != nil {
		return &wrapperspb.BoolValue{Value: *value}
	}

	return nil
}

func ToProtoTimeWrapper(value *time.Time) *wrapperspb.Int64Value {
	if value != nil {
		return &wrapperspb.Int64Value{Value: value.Unix()}
	}

	return nil
}

func FromProtoStringWrapper(value *wrapperspb.StringValue) *string {
	if value != nil {
		return &value.Value
	}

	return nil
}

func FromProtoInt32Wrapper(value *wrapperspb.Int32Value) *int32 {
	if value != nil {
		return &value.Value
	}

	return nil
}

func FromProtoInt64Wrapper(value *wrapperspb.Int64Value) *int64 {
	if value != nil {
		return &value.Value
	}

	return nil
}

func FromProtoFloatWrapper(value *wrapperspb.FloatValue) *float32 {
	if value != nil {
		return &value.Value
	}

	return nil
}

func FromProtoFloat64Wrapper(value *wrapperspb.FloatValue) *float64 {
	if value != nil {
		val := float64(value.Value)
		return &val
	}

	return nil
}

func FromProtoBoolWrapper(value *wrapperspb.BoolValue) *bool {
	if value != nil {
		return &value.Value
	}

	return nil
}

func FromProtoTimeWrapper(value *wrapperspb.Int64Value) *time.Time {
	if value != nil {
		t := time.Unix(value.Value, 0)
		return &t
	}

	return nil
}
