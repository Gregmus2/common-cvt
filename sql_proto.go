package convptr

import (
	"database/sql"
	"github.com/golang/protobuf/ptypes/wrappers"
	"strings"
	"time"
)

func FromSQLFloatToProto(value sql.NullFloat64) *wrappers.FloatValue {
	if !value.Valid {
		return nil
	}

	return &wrappers.FloatValue{Value: float32(value.Float64)}
}

func FromSQLFloatToProtoDouble(value sql.NullFloat64) *wrappers.DoubleValue {
	if !value.Valid {
		return nil
	}

	return &wrappers.DoubleValue{Value: value.Float64}
}

func FromSQLStringToProto(value sql.NullString) *wrappers.StringValue {
	if !value.Valid {
		return nil
	}

	return &wrappers.StringValue{Value: value.String}
}

func FromSQLToLowerStringToProto(value sql.NullString) *wrappers.StringValue {
	if !value.Valid {
		return nil
	}
	result := strings.ToLower(value.String)

	return &wrappers.StringValue{Value: result}
}

func FromSQLIntToProto(value sql.NullInt32) *wrappers.Int32Value {
	if !value.Valid {
		return nil
	}

	return &wrappers.Int32Value{Value: value.Int32}
}

func FromSQLInt64ToProto(value sql.NullInt64) *wrappers.Int64Value {
	if !value.Valid {
		return nil
	}

	return &wrappers.Int64Value{Value: value.Int64}
}

func FromSQLIntToProtoUInt(value sql.NullInt32) *wrappers.UInt32Value {
	if !value.Valid {
		return nil
	}

	return &wrappers.UInt32Value{Value: uint32(value.Int32)}
}

func FromSQLBoolToProto(value sql.NullBool) *wrappers.BoolValue {
	if !value.Valid {
		return nil
	}

	return &wrappers.BoolValue{Value: value.Bool}
}

func FromSQLTimeToProto(value sql.NullTime) *wrappers.Int64Value {
	if !value.Valid {
		return nil
	}

	return &wrappers.Int64Value{Value: value.Time.Unix()}
}

func ToSQLStringFromProto(value *wrappers.StringValue) sql.NullString {
	if value == nil {
		return sql.NullString{}
	}

	return sql.NullString{String: value.Value, Valid: true}
}

func ToSQLFloatFromProto(value *wrappers.FloatValue) sql.NullFloat64 {
	if value == nil {
		return sql.NullFloat64{}
	}

	return sql.NullFloat64{Float64: float64(value.Value), Valid: true}
}

func ToSQLIntFromProto(value *wrappers.Int32Value) sql.NullInt32 {
	if value == nil {
		return sql.NullInt32{}
	}

	return sql.NullInt32{Int32: value.Value, Valid: true}
}

func ToSQLInt64FromProto(value *wrappers.Int64Value) sql.NullInt64 {
	if value == nil {
		return sql.NullInt64{}
	}

	return sql.NullInt64{Int64: value.Value, Valid: true}
}

func ToSQLBoolFromProto(value *wrappers.BoolValue) sql.NullBool {
	if value == nil {
		return sql.NullBool{}
	}

	return sql.NullBool{Bool: value.Value, Valid: true}
}

func ToSQLTimeFromProto(value *wrappers.Int64Value) sql.NullTime {
	if value == nil {
		return sql.NullTime{}
	}

	return sql.NullTime{Time: time.Unix(value.Value, 0), Valid: true}
}
