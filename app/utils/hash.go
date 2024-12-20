package utils

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string, ctx context.Context) (string, error) {
	// _, span := otel.StartNewTrace(ctx, "HashPassword")
	// defer otel.EndSpan(span)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	// _, span := otel.StartNewTrace(ctx, "CheckPasswordHash")
	// defer otel.EndSpan(span)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		// otel.RecordError(span, err)
		return false
	}
	return err == nil
}
