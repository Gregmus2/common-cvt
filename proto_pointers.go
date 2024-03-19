package convptr

import (
	"github.com/golang/protobuf/ptypes/wrappers"
)

func ToProtoStringWrapper(value *string) *wrappers.StringValue {
	if value != nil {
		return &wrappers.StringValue{Value: *value}
	}

	return nil
}

func ToProtoInt32Wrapper(value *int32) *wrappers.Int32Value {
	if value != nil {
		return &wrappers.Int32Value{Value: *value}
	}

	return nil
}

func ToProtoFloatWrapper(value *float32) *wrappers.FloatValue {
	if value != nil {
		return &wrappers.FloatValue{Value: *value}
	}

	return nil
}

func ToProtoFloat64Wrapper(value *float64) *wrappers.FloatValue {
	if value != nil {
		return &wrappers.FloatValue{Value: float32(*value)}
	}

	return nil
}

func ToProtoBoolWrapper(value *bool) *wrappers.BoolValue {
	if value != nil {
		return &wrappers.BoolValue{Value: *value}
	}

	return nil
}

func FromProtoStringWrapper(value *wrappers.StringValue) *string {
	if value != nil {
		return &value.Value
	}

	return nil
}

func FromProtoInt32Wrapper(value *wrappers.Int32Value) *int32 {
	if value != nil {
		return &value.Value
	}

	return nil
}

func FromProtoFloatWrapper(value *wrappers.FloatValue) *float32 {
	if value != nil {
		return &value.Value
	}

	return nil
}

func FromProtoFloat64Wrapper(value *wrappers.FloatValue) *float64 {
	if value != nil {
		val := float64(value.Value)
		return &val
	}

	return nil
}

func FromProtoBoolWrapper(value *wrappers.BoolValue) *bool {
	if value != nil {
		return &value.Value
	}

	return nil
}
