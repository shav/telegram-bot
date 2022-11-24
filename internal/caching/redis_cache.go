package caching

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// Номер БД по-умолчанию.
const defaultDB = 0

var notSupportEmptyMapError = errors.New("empty maps are not supported by Redis cache")

// RedisCache реализует универсальный механизм кеширования данных на основе Redis-а.
type RedisCache struct {
	// Название кеша.
	name string
	// Клиент для подключения к сервису кеширования Redis.
	client *redis.Client
}

// NewRedisCache создаёт новый кеш данных в Redis-е.
func NewRedisCache(ctx context.Context, name string, connectionString string, checkConnection bool) (*RedisCache, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("New RedisCache: cache name is not assigned")
	}
	connectionString = strings.TrimSpace(connectionString)
	if connectionString == "" {
		return nil, errors.New("New RedisCache: connection string for Redis is not assigned")
	}

	client := redis.NewClient(&redis.Options{
		Addr:            connectionString,
		MaxRetries:      2,
		MaxRetryBackoff: 100 * time.Millisecond,
		ReadTimeout:     100 * time.Millisecond,
		WriteTimeout:    100 * time.Millisecond,
		DialTimeout:     100 * time.Millisecond,
		// TODO: Пока для простоты храним все данные в одной БД, разделяя данные из разных кешей через префикс с названием кеша.
		// Если будут проблемы с производительностью на больших объёмах данных, то можно будет хранить данные из разных кешей в разных БД
		// (но это предполагает разработку механизма соответствия кешей номерам БД, синхронизированного на разных инстансах чат-бота).
		// Также можно будет шардировать данные по разным инстансам Redis-кластера.
		DB: defaultDB,
	})

	if checkConnection {
		err := client.Ping(ctx).Err()
		if err != nil {
			return nil, err
		}
	}

	return &RedisCache{
		name:   name,
		client: client,
	}, nil
}

// SetMap устанавливает для ключа в кеше key непустое значение value типа map со временем окончания жизни записи expireAt.
func (c *RedisCache) SetMap(ctx context.Context, key string, value map[string]string, expireAt time.Time) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "RedisCache.SetMap")
	defer span.Finish()

	if len(value) == 0 {
		return notSupportEmptyMapError
	}

	cacheKey := c.getCacheKey(key)
	err := c.client.HSet(ctx, cacheKey, value).Err()
	if err != nil {
		tracing.SetError(span)
		return err
	}

	err = c.client.ExpireAt(ctx, cacheKey, expireAt).Err()
	if err != nil {
		tracing.SetError(span)
	}
	return err
}

// GetMap получает из кеша по ключу key значение типа map.
func (c *RedisCache) GetMap(ctx context.Context, key string) (value map[string]string, exists bool, err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "RedisCache.GetMap")
	defer span.Finish()

	value, err = c.client.HGetAll(ctx, c.getCacheKey(key)).Result()
	if err == redis.Nil {
		return nil, false, nil
	}
	if err != nil {
		tracing.SetError(span)
		return nil, false, err
	}

	// Если карты с указанным ключом нет в кеше, то Redis возвращает пустую карту.
	// И это не может произойти в случае, если по заданному лючу в кеше хранится пустая карта, т.к. хранение пустых карт в Redis не поддерживается
	// (см. https://stackoverflow.com/questions/71479411/inserting-an-empty-map-into-redis-using-hset-fails-in-golang)
	if len(value) == 0 {
		return nil, false, nil
	}
	return value, true, nil
}

// TODO: По необходимости добавить Get- и Set-методы для других типов данных.

// Delete удаляет из кеша записи с ключами keys.
func (c *RedisCache) Delete(ctx context.Context, keys ...string) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "RedisCache.Delete")
	defer span.Finish()

	if len(keys) == 0 {
		return nil
	}

	pipeline := c.client.TxPipeline()
	for _, key := range keys {
		pipeline.Del(ctx, c.getCacheKey(key))
	}
	_, err := pipeline.Exec(ctx)

	if err != nil {
		tracing.SetError(span)
	}
	return err
}

// Clear очищает все кеши в сервисе кеширования.
func (c *RedisCache) Clear(ctx context.Context) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "RedisCache.Clear")
	defer span.Finish()

	err := c.client.FlushDB(ctx).Err()

	if err != nil {
		tracing.SetError(span)
	}
	return err
}

// Close закрывает соединение с сервисом кеширования.
func (c *RedisCache) Close() error {
	return c.client.Close()
}

// getCacheKey возвращает полный ключ записи в кеше, с учётом названия кеша.
func (c *RedisCache) getCacheKey(key string) string {
	namespace := fmt.Sprintf("%s:", c.name)
	if strings.HasPrefix(key, namespace) {
		return key
	}
	return namespace + key
}
