package utils

import (
	"database/sql"
	"os"
	"time"

	"github.com/sarc-tech/backend/internal/oas"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func ToNullStringFromOptString(s oas.OptString) sql.NullString {
	if s.IsSet() {
        return sql.NullString{String: s.Value, Valid: true}
    }
    return sql.NullString{Valid: false}
}

func ToNullStringFromString(s *string) sql.NullString {
    if s != nil && *s != "" {
        return sql.NullString{String: *s, Valid: true}
    }
    return sql.NullString{Valid: false}
}

func ToNullDateFromOptDate(t oas.OptDate) sql.NullTime {
	if t.IsSet() {
		return sql.NullTime{Time: t.Value, Valid: true}
    }
    return sql.NullTime{Valid: false}
}

func ToNullDateFromString(s string) sql.NullTime {
    if s == "" {
        return sql.NullTime{Valid: false}
    }

    t, err := time.Parse("2006-01-02", s)
    if err != nil {
        return sql.NullTime{}
    }

    return sql.NullTime{Time: t, Valid: true}
}
