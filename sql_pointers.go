package ccvt

import (
	"database/sql"
	"strings"
	"time"
)

func FromSQLFloat(value sql.NullFloat64) *float64 {
	if !value.Valid {
		return nil
	}

	return &value.Float64
}

func FromSQLString(value sql.NullString) *string {
	if !value.Valid {
		return nil
	}

	return &value.String
}

func FromSQLToLowerString(value sql.NullString) *string {
	if !value.Valid {
		return nil
	}
	result := strings.ToLower(value.String)

	return &result
}

func FromSQLInt(value sql.NullInt32) *int32 {
	if !value.Valid {
		return nil
	}

	return &value.Int32
}

func FromSQLInt64(value sql.NullInt64) *int64 {
	if !value.Valid {
		return nil
	}

	return &value.Int64
}

func FromSQLBool(value sql.NullBool) *bool {
	if !value.Valid {
		return nil
	}

	return &value.Bool
}

func FromSQLTime(value sql.NullTime) *time.Time {
	if !value.Valid {
		return nil
	}

	return &value.Time
}

func ToSQLString(value *string) sql.NullString {
	if value == nil {
		return sql.NullString{}
	}

	return sql.NullString{String: *value, Valid: true}
}

func ToSQLFloat(value *float64) sql.NullFloat64 {
	if value == nil {
		return sql.NullFloat64{}
	}

	return sql.NullFloat64{Float64: *value, Valid: true}
}

func ToSQLInt(value *int32) sql.NullInt32 {
	if value == nil {
		return sql.NullInt32{}
	}

	return sql.NullInt32{Int32: *value, Valid: true}
}

func ToSQLInt64(value *int64) sql.NullInt64 {
	if value == nil {
		return sql.NullInt64{}
	}

	return sql.NullInt64{Int64: *value, Valid: true}
}

func ToSQLBool(value *bool) sql.NullBool {
	if value == nil {
		return sql.NullBool{}
	}

	return sql.NullBool{Bool: *value, Valid: true}
}

func ToSQLTime(value *time.Time) sql.NullTime {
	if value == nil {
		return sql.NullTime{}
	}

	return sql.NullTime{Time: *value, Valid: true}
}
