package redis

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/ulyssessouza/envlang/store"
	"gotest.tools/v3/assert"
)

const (
	testValue = "test_value"
	mapValue  = "map_value"
)

func setupRedis(t *testing.T) (*redis.Client, func()) {
	t.Helper()

	mr := miniredis.RunT(t)

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return client, func() {
		client.Close()
		mr.Close()
	}
}

func TestNew(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	s := New(client)
	assert.Assert(t, s != nil)

	_, ok := s.Get("NONEXISTENT")
	assert.Assert(t, !ok)
}

func TestNew_WithOptions(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	ctx := context.Background()
	env := []string{"KEY1=value1"}

	s := New(client,
		WithPrefix("test:"),
		WithContext(ctx),
		WithInitialEnv(env),
	)
	assert.Assert(t, s != nil)

	v, ok := s.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value1")
}

func TestRedisStore_PutAndGet(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	s := New(client)

	value := testValue
	s.Put("KEY1", &value)

	v, ok := s.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, testValue)
}

func TestRedisStore_Put_NilValue(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	s := New(client)

	value := testValue
	s.Put("KEY1", &value)

	v, ok := s.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, testValue)

	s.Put("KEY1", nil)

	_, ok = s.Get("KEY1")
	assert.Assert(t, !ok, "nil value should delete the key")
}

func TestRedisStore_Get_NonExistent(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	s := New(client)

	_, ok := s.Get("NONEXISTENT")
	assert.Assert(t, !ok)
}

func TestRedisStore_Get_PriorityOrder(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	envValue := "env_value"
	lookupValue := "lookup_value"

	env := []string{"KEY=env_value"}
	s := New(client, WithInitialEnv(env))

	v, ok := s.Get("KEY")
	assert.Assert(t, ok)
	assert.Equal(t, *v, envValue, "should get from env when redis is empty")

	mapVal := mapValue
	s.Put("KEY", &mapVal)
	v, ok = s.Get("KEY")
	assert.Assert(t, ok)
	assert.Equal(t, *v, mapValue, "should get from redis when both exist")

	lookupFn := func(k string) (string, bool) {
		if k == "KEY" {
			return lookupValue, true
		}
		return "", false
	}
	WithLookupFn(lookupFn)(s.(*RedisStore))

	v, ok = s.Get("KEY")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "lookup_value", "should get from lookupFn when all exist")
}

func TestRedisStore_Remove(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	s := New(client)

	value := testValue
	s.Put("KEY1", &value)

	v, ok := s.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, testValue)

	assert.Equal(t, s.Remove("KEY1"), true)

	_, ok = s.Get("KEY1")
	assert.Assert(t, !ok)
}

func TestRedisStore_Remove_NonExistent(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	s := New(client)

	assert.Equal(t, s.Remove("NONEXISTENT"), false)
}

func TestRedisStore_ImportList(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	s := New(client)

	env := []string{
		"KEY1=value1",
		"KEY2=value2",
		"KEY3=value with spaces",
	}

	s.ImportList(env)

	v, ok := s.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value1")

	v, ok = s.Get("KEY2")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value2")

	v, ok = s.Get("KEY3")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value with spaces")
}

func TestRedisStore_ImportList_WithEquals(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	s := New(client)

	env := []string{
		"KEY1=value=with=equals",
	}

	s.ImportList(env)

	v, ok := s.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value=with=equals")
}

func TestRedisStore_ImportMap(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	s := New(client)

	m := map[string]string{
		"KEY1": "value1",
		"KEY2": "value2",
	}

	s.ImportMap(m)

	v, ok := s.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value1")

	v, ok = s.Get("KEY2")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value2")
}

func TestRedisStore_ImportMap_EmptyMap(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	s := New(client)

	s.ImportMap(map[string]string{})

	_, ok := s.Get("ANY_KEY")
	assert.Assert(t, !ok)
}

func TestRedisStore_ExportMap(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	s := New(client)

	value1 := "value1"
	value2 := "value2"
	s.Put("KEY1", &value1)
	s.Put("KEY2", &value2)

	exported := s.ExportMap()
	assert.Equal(t, len(exported), 2)
	assert.Equal(t, *exported["KEY1"], value1)
	assert.Equal(t, *exported["KEY2"], value2)
}

func TestRedisStore_ExportMap_Empty(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	s := New(client)

	exported := s.ExportMap()
	assert.Equal(t, len(exported), 0)
}

func TestRedisStore_ExportMap_DoesNotIncludeEnv(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	env := []string{"ENV_KEY=env_value"}
	s := New(client, WithInitialEnv(env))

	mapVal := mapValue
	s.Put("MAP_KEY", &mapVal)

	exported := s.ExportMap()
	assert.Equal(t, len(exported), 1)
	assert.Equal(t, *exported["MAP_KEY"], mapValue)
	_, exists := exported["ENV_KEY"]
	assert.Assert(t, !exists, "ExportMap should only export from redis, not env")
}

func TestRedisStore_WithLookupFn(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	lookupFn := func(k string) (string, bool) {
		if k == "LOOKUP_KEY" {
			return "lookup_value", true
		}
		return "", false
	}

	s := New(client)
	WithLookupFn(lookupFn)(s.(*RedisStore))

	v, ok := s.Get("LOOKUP_KEY")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "lookup_value")
}

func TestRedisStore_WithLookupFn_FallbackToRedis(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	lookupFn := func(k string) (string, bool) {
		return "", false
	}

	s := New(client)
	WithLookupFn(lookupFn)(s.(*RedisStore))

	mapVal := mapValue
	s.Put("KEY", &mapVal)

	v, ok := s.Get("KEY")
	assert.Assert(t, ok)
	assert.Equal(t, *v, mapValue)
}

func TestRedisStore_Put_WithLookupFn(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	lookupFn := func(k string) (string, bool) {
		if k == "KEY" {
			return "lookup_override", true
		}
		return "", false
	}

	s := New(client)
	WithLookupFn(lookupFn)(s.(*RedisStore))

	originalValue := "original_value"
	s.Put("KEY", &originalValue)

	v, ok := s.Get("KEY")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "lookup_override", "Put should use lookupFn value when available")
}

func TestRedisStore_WithPrefix(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	s := New(client, WithPrefix("custom:"))

	value := testValue
	s.Put("KEY1", &value)

	v, ok := s.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, testValue)

	rawValue, err := client.Get(context.Background(), "custom:KEY1").Result()
	assert.NilError(t, err)
	assert.Equal(t, rawValue, testValue)
}

func TestRedisStore_WithContext(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	ctx := context.Background()
	s := New(client, WithContext(ctx))

	value := testValue
	s.Put("KEY1", &value)

	v, ok := s.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, testValue)
}

func TestRedisStore_WithInitialEnv(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	env := []string{
		"KEY1=value1",
		"KEY2=value2",
	}

	s := New(client, WithInitialEnv(env))

	v, ok := s.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value1")

	v, ok = s.Get("KEY2")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value2")
}

func TestRedisStore_InterfaceCompliance(t *testing.T) {
	client, cleanup := setupRedis(t)
	defer cleanup()

	var _ store.Store = New(client)
}
