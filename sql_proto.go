package ccvt

import (
	"database/sql"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"strings"
	"time"
)

func FromSQLFloatToProto(value sql.NullFloat64) *wrapperspb.FloatValue {
	if !value.Valid {
		return nil
	}

	return &wrapperspb.FloatValue{Value: float32(value.Float64)}
}

func FromSQLFloatToProtoDouble(value sql.NullFloat64) *wrapperspb.DoubleValue {
	if !value.Valid {
		return nil
	}

	return &wrapperspb.DoubleValue{Value: value.Float64}
}

func FromSQLStringToProto(value sql.NullString) *wrapperspb.StringValue {
	if !value.Valid {
		return nil
	}

	return &wrapperspb.StringValue{Value: value.String}
}

func FromSQLToLowerStringToProto(value sql.NullString) *wrapperspb.StringValue {
	if !value.Valid {
		return nil
	}
	result := strings.ToLower(value.String)

	return &wrapperspb.StringValue{Value: result}
}

func FromSQLIntToProto(value sql.NullInt32) *wrapperspb.Int32Value {
	if !value.Valid {
		return nil
	}

	return &wrapperspb.Int32Value{Value: value.Int32}
}

func FromSQLInt64ToProto(value sql.NullInt64) *wrapperspb.Int64Value {
	if !value.Valid {
		return nil
	}

	return &wrapperspb.Int64Value{Value: value.Int64}
}

func FromSQLIntToProtoUInt(value sql.NullInt32) *wrapperspb.UInt32Value {
	if !value.Valid {
		return nil
	}

	return &wrapperspb.UInt32Value{Value: uint32(value.Int32)}
}

func FromSQLBoolToProto(value sql.NullBool) *wrapperspb.BoolValue {
	if !value.Valid {
		return nil
	}

	return &wrapperspb.BoolValue{Value: value.Bool}
}

func FromSQLTimeToProto(value sql.NullTime) *wrapperspb.Int64Value {
	if !value.Valid {
		return nil
	}

	return &wrapperspb.Int64Value{Value: value.Time.Unix()}
}

func ToSQLStringFromProto(value *wrapperspb.StringValue) sql.NullString {
	if value == nil {
		return sql.NullString{}
	}

	return sql.NullString{String: value.Value, Valid: true}
}

func ToSQLFloatFromProto(value *wrapperspb.FloatValue) sql.NullFloat64 {
	if value == nil {
		return sql.NullFloat64{}
	}

	return sql.NullFloat64{Float64: float64(value.Value), Valid: true}
}

func ToSQLIntFromProto(value *wrapperspb.Int32Value) sql.NullInt32 {
	if value == nil {
		return sql.NullInt32{}
	}

	return sql.NullInt32{Int32: value.Value, Valid: true}
}

func ToSQLInt64FromProto(value *wrapperspb.Int64Value) sql.NullInt64 {
	if value == nil {
		return sql.NullInt64{}
	}

	return sql.NullInt64{Int64: value.Value, Valid: true}
}

func ToSQLBoolFromProto(value *wrapperspb.BoolValue) sql.NullBool {
	if value == nil {
		return sql.NullBool{}
	}

	return sql.NullBool{Bool: value.Value, Valid: true}
}

func ToSQLTimeFromProto(value *wrapperspb.Int64Value) sql.NullTime {
	if value == nil {
		return sql.NullTime{}
	}

	return sql.NullTime{Time: time.Unix(value.Value, 0), Valid: true}
}
