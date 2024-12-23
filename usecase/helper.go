package usecase

import (
	"context"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/cache/v8"
)

// ParseUnwantedError parse and delete unwanted error message
func ParseUnwantedError(err error) string {
	if err == nil {
		return ""
	}

	var s string
	estr := err.Error()

	switch true {
	case strings.Contains(estr, "context deadline exceeded"):
		s = "context deadline exceeded, connection timeout"

	case strings.Contains(estr, "Timeout exceeded"):
		s = "dial tcp error timeout"

	case strings.Contains(estr, "SQLSTATE 42P01"):
		s = "desired table does not exist error"

	case strings.Contains(estr, "SQLSTATE 42703"):
		s = "desired column does not exist error"

	case strings.Contains(estr, "SQLSTATE 23502"):
		s = "some column value cannot be null"

	case strings.Contains(estr, "SQLSTATE 22001"):
		s = "value too long"

	case strings.Contains(estr, "SQLSTATE 23503"):
		s = "violates foreign key constraint"

	default:
		s = estr
	}

	return s
}

// GetCache2 get value from cache2 adapter
// Cache value will be set to wanted argument
func GetCache2(
	instance *cache.Cache,
	key string,
	wanted interface{},
) error {
	key = "republish-seamless-wallet:cache" + ":" + key
	ctx := context.TODO()

	return instance.Get(ctx, key, wanted)
}

// SetCache2 set given value to cache2 adapter
func SetCache2(
	insatnce *cache.Cache,
	key string,
	val interface{},
	exp time.Duration,
) error {
	key = "republish-seamless-wallet:cache" + ":" + key
	ctx := context.TODO()

	return insatnce.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: val,
		TTL:   exp,
	})
}

// DeleteCache2 delete given value in cache2 adapter
func DeleteCache2(
	insatnce *cache.Cache,
	key string,
) error {
	key = "republish-seamless-wallet:cache" + ":" + key
	ctx := context.TODO()

	return insatnce.Delete(ctx, key)
}

// ValidateUUIDV4 validating uuid v4
func ValidateUUIDV4(uuid string) error {
	validator := validator.New()

	return validator.Var(uuid, "uuid4")
}
