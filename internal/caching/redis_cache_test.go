package caching_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/shav/telegram-bot/internal/caching"
	"github.com/shav/telegram-bot/internal/config"
	"github.com/shav/telegram-bot/internal/testing"
)

const networkError = "dial tcp: lookup unavailable-host"

var ctx = context.Background()

var never = time.Now().AddDate(9999, 0, 0)

var map1 = map[string]string{
	"key1": "value1",
}

var map2 = map[string]string{
	"key1": "value1",
	"key2": "value2",
}

var map3 = map[string]string{
	"key1": "value1",
	"key2": "value2",
	"key3": "value3",
}

var defaultMap map[string]string = nil

func configureCache(t *testing.T, cacheName string) *caching.RedisCache {
	config, err := config.NewEnvConfig(test_settings.ServiceName)
	assert.NoError(t, err)
	cacheConnString := config.CacheConnectionString()
	assert.NotEmpty(t, cacheConnString)
	cache, err := caching.NewRedisCache(context.Background(), cacheName, cacheConnString, true)
	assert.NoError(t, err)
	assert.NotNil(t, cache)
	err = cache.Clear(ctx)
	assert.NoError(t, err)
	return cache
}

func assertValueNotExist[TValue any](t *testing.T, actualValue TValue, exists bool, err error) {
	assert.NoError(t, err)
	assert.False(t, exists)
	var defaultValue TValue
	assert.Equal(t, defaultValue, actualValue)
}

func assertValueExist[TValue any](t *testing.T, expectedValue TValue, actualValue TValue, exists bool, err error) {
	assert.NoError(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedValue, actualValue)
}

func Test_OnCreateCache_ShouldReturnNetworkErrorAndNilCache_WhenCacheServiceIsNotAvailable_AndCheckConnectionOCreate(t *testing.T) {
	cache, err := caching.NewRedisCache(context.Background(), "UnavailableCache", "unavailable-host:666", true)
	assert.Nil(t, cache)
	assert.ErrorContains(t, err, networkError)

}

func Test_OnExecuteCommand_ShouldReturnNetworkError_WhenCacheServiceBecameNotAvailable(t *testing.T) {
	cache, err := caching.NewRedisCache(context.Background(), "UnavailableCache", "unavailable-host:666", false)
	assert.NoError(t, err)

	err = cache.SetMap(ctx, "key", map1, never)
	assert.ErrorContains(t, err, networkError)

	value, exists, err := cache.GetMap(ctx, "key")
	assert.ErrorContains(t, err, networkError)
	assert.False(t, exists)
	assert.Equal(t, defaultMap, value)

	err = cache.Delete(ctx, "key")
	assert.ErrorContains(t, err, networkError)
}

func Test_GetMap_ShouldReturnNotExistingFlag_WhenKeyNotExistInCache(t *testing.T) {
	cache := configureCache(t, "CacheWithoutValue")

	value, exists, err := cache.GetMap(ctx, "key")
	assertValueNotExist(t, value, exists, err)
}

func Test_GetMap_ShouldReturnValue_WhenKeyExistInCache(t *testing.T) {
	cache := configureCache(t, "CacheWithValue")

	err := cache.SetMap(ctx, "key", map1, never)
	assert.NoError(t, err)

	value, exists, err := cache.GetMap(ctx, "key")
	assertValueExist(t, map1, value, exists, err)
}

func Test_SetMap_ShouldOverrideValue_WhenKeyExistInCache(t *testing.T) {
	cache := configureCache(t, "CacheWithValue")

	err := cache.SetMap(ctx, "key", map1, never)
	assert.NoError(t, err)

	value, exists, err := cache.GetMap(ctx, "key")
	assertValueExist(t, map1, value, exists, err)

	err = cache.SetMap(ctx, "key", map2, never)
	assert.NoError(t, err)

	value, exists, err = cache.GetMap(ctx, "key")
	assertValueExist(t, map2, value, exists, err)
}

func Test_GetMap_ShouldReturnDifferentValues_WhenSameKeyExistInDifferentCaches(t *testing.T) {
	cache1 := configureCache(t, "CacheWithValue-1")
	cache2 := configureCache(t, "CacheWithValue-2")

	err1 := cache1.SetMap(ctx, "key", map1, never)
	assert.NoError(t, err1)

	err2 := cache2.SetMap(ctx, "key", map2, never)
	assert.NoError(t, err2)

	value1, exists1, err1 := cache1.GetMap(ctx, "key")
	assertValueExist(t, map1, value1, exists1, err1)

	value2, exists2, err2 := cache2.GetMap(ctx, "key")
	assertValueExist(t, map2, value2, exists2, err2)
}

func Test_GetMap_ShouldReturnDifferentValues_WhenDifferentKeysExistInDifferentCaches(t *testing.T) {
	cache1 := configureCache(t, "CacheWithValue-1")
	cache2 := configureCache(t, "CacheWithValue-2")

	err1 := cache1.SetMap(ctx, "key1", map1, never)
	assert.NoError(t, err1)

	err2 := cache2.SetMap(ctx, "key2", map2, never)
	assert.NoError(t, err2)

	value1, exists1, err1 := cache1.GetMap(ctx, "key1")
	assertValueExist(t, map1, value1, exists1, err1)

	value2, exists2, err2 := cache2.GetMap(ctx, "key2")
	assertValueExist(t, map2, value2, exists2, err2)
}

func Test_GetMap_ShouldReturnNotExistingFlag_WhenKeyNotExistInCache_AndSameKeyExistInOtherCache(t *testing.T) {
	cache1 := configureCache(t, "CacheWithoutValue")
	cache2 := configureCache(t, "CacheWithValue")

	err2 := cache2.SetMap(ctx, "key", map2, never)
	assert.NoError(t, err2)

	value1, exists1, err1 := cache1.GetMap(ctx, "key")
	assertValueNotExist(t, value1, exists1, err1)

	value2, exists2, err2 := cache2.GetMap(ctx, "key")
	assertValueExist(t, map2, value2, exists2, err2)
}

func Test_GetMap_ShouldReturnNotExistingFlag_WhenKeyInCacheWasExpired(t *testing.T) {
	cache := configureCache(t, "CacheWithExpiringValue")

	err := cache.SetMap(ctx, "key", map1, time.Now().Add(1000*time.Millisecond))
	assert.NoError(t, err)

	value, exists, err := cache.GetMap(ctx, "key")
	assertValueExist(t, map1, value, exists, err)

	time.Sleep(1500 * time.Millisecond)
	value, exists, err = cache.GetMap(ctx, "key")
	assertValueNotExist(t, value, exists, err)
}

func Test_Delete_ShouldDeleteNothing_WhenDeleteNoKeys(t *testing.T) {
	cache := configureCache(t, "CacheForDeleteNoKeys")

	err := cache.SetMap(ctx, "key1", map1, never)
	assert.NoError(t, err)
	err = cache.SetMap(ctx, "key2", map2, never)
	assert.NoError(t, err)

	err = cache.Delete(ctx)
	assert.NoError(t, err)

	value1, exists1, err1 := cache.GetMap(ctx, "key1")
	assertValueExist(t, map1, value1, exists1, err1)

	value2, exists2, err2 := cache.GetMap(ctx, "key2")
	assertValueExist(t, map2, value2, exists2, err2)
}

func Test_Delete_ShouldDeleteSingleValue_WhenDeleteSingleKey(t *testing.T) {
	cache := configureCache(t, "CacheForDeleteSingleKey")

	err := cache.SetMap(ctx, "key1", map1, never)
	assert.NoError(t, err)
	err = cache.SetMap(ctx, "key2", map2, never)
	assert.NoError(t, err)

	err = cache.Delete(ctx, "key1")
	assert.NoError(t, err)

	value1, exists1, err1 := cache.GetMap(ctx, "key1")
	assertValueNotExist(t, value1, exists1, err1)

	value2, exists2, err2 := cache.GetMap(ctx, "key2")
	assertValueExist(t, map2, value2, exists2, err2)
}

func Test_Delete_ShouldDeleteMultipleValues_WhenDeleteMultipleKeys(t *testing.T) {
	cache := configureCache(t, "CacheForDeleteMultipleKey")

	err := cache.SetMap(ctx, "key1", map1, never)
	assert.NoError(t, err)
	err = cache.SetMap(ctx, "key2", map2, never)
	assert.NoError(t, err)
	err = cache.SetMap(ctx, "key3", map3, never)
	assert.NoError(t, err)

	err = cache.Delete(ctx, "key1", "key2")
	assert.NoError(t, err)

	value1, exists1, err1 := cache.GetMap(ctx, "key1")
	assertValueNotExist(t, value1, exists1, err1)

	value2, exists2, err2 := cache.GetMap(ctx, "key2")
	assertValueNotExist(t, value2, exists2, err2)

	value3, exists3, err3 := cache.GetMap(ctx, "key3")
	assertValueExist(t, map3, value3, exists3, err3)
}
