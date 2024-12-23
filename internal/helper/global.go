package helper

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"time"

	"github.com/go-redis/cache/v8"
)

func CheckBetValid(category string, wAmount, wlAmount float64) (betValid float64) {
	switch {
	case InArray(category, []string{"S", "L", "P"}):
		betValid = wAmount
	case InArray(category, []string{"C", "SB", "LG"}):
		if math.Abs(wlAmount) > wAmount {
			betValid = wAmount
		} else {
			betValid = math.Abs(wlAmount)
		}
	case wAmount == wlAmount || InArray(category, []string{"A", "T"}):
		betValid = 0
	}

	return
}

func InArray(v interface{}, in interface{}) (ok bool) {
	val := reflect.Indirect(reflect.ValueOf(in))
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if ok = v == val.Index(i).Interface(); ok {
				return
			}
		}
	}
	return
}

func GetCache(instance *cache.Cache, ck string, wanted interface{}) error {
	key := "wallet:cache-" + ck
	ctx := context.TODO()

	return instance.Get(ctx, key, wanted)
}

func SetCache(instance *cache.Cache, ck string, val interface{}, ttl ...time.Duration) error {
	fmt.Println(instance)
	key := "wallet:cache-" + ck
	ctx := context.TODO()

	cacheItem := &cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: val,
	}
	if len(ttl) > 0 {
		cacheItem.TTL = ttl[0]
	}

	return instance.Set(cacheItem)
}

// Cache delete
func DeleteCache(
	instance *cache.Cache,
	ck string,
) error {
	key := "wallet:cache-" + ck
	ctx := context.TODO()

	return instance.Delete(ctx, key)
}
