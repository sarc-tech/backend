package utils

import (
	"context"
	"database/sql"
	"os"
	"time"

	"go.opentelemetry.io/otel/trace"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetTraceID(ctx context.Context) string {
    span := trace.SpanFromContext(ctx)
    sc := span.SpanContext()
    if sc.HasTraceID() {
        return sc.TraceID().String()
    }
    return ""
}

func ToNullStringFromString(s *string) sql.NullString {
    if s != nil && *s != "" {
        return sql.NullString{String: *s, Valid: true}
    }
    return sql.NullString{Valid: false}
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
