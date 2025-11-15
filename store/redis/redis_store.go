package redis

import (
	"context"
	"strings"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/ulyssessouza/envlang/store"
)

const (
	pairSeparator = "="
	pairNumber    = 2
)

type RedisStore struct {
	sync.RWMutex
	client   *redis.Client
	ctx      context.Context
	prefix   string
	env      map[string]*string
	lookupFn store.LookupFn
}

type Option func(*RedisStore)

func WithPrefix(prefix string) Option {
	return func(r *RedisStore) {
		r.prefix = prefix
	}
}

func WithContext(ctx context.Context) Option {
	return func(r *RedisStore) {
		r.ctx = ctx
	}
}

func WithLookupFn(fn store.LookupFn) Option {
	return func(r *RedisStore) {
		r.lookupFn = fn
	}
}

func WithInitialEnv(env []string) Option {
	return func(r *RedisStore) {
		r.env = make(map[string]*string)
		for _, e := range env {
			splitEnv := strings.SplitN(e, pairSeparator, pairNumber)
			if len(splitEnv) == pairNumber {
				v := splitEnv[1]
				r.env[splitEnv[0]] = &v
			}
		}
	}
}

func New(client *redis.Client, opts ...Option) store.Store {
	r := &RedisStore{
		client: client,
		ctx:    context.Background(),
		prefix: "envlang:",
		env:    make(map[string]*string),
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func (r *RedisStore) key(k string) string {
	return r.prefix + k
}

func (r *RedisStore) ImportList(env []string) {
	r.Lock()
	defer r.Unlock()

	for _, e := range env {
		splitEnv := strings.SplitN(e, pairSeparator, pairNumber)
		if len(splitEnv) == pairNumber {
			r.client.Set(r.ctx, r.key(splitEnv[0]), splitEnv[1], 0)
		}
	}
}

func (r *RedisStore) ImportMap(m map[string]string) {
	r.Lock()
	defer r.Unlock()

	for k, v := range m {
		r.client.Set(r.ctx, r.key(k), v, 0)
	}
}

func (r *RedisStore) Get(k string) (*string, bool) {
	r.RLock()
	defer r.RUnlock()

	if r.lookupFn != nil {
		if v, ok := r.lookupFn(k); ok {
			return &v, true
		}
	}

	val, err := r.client.Get(r.ctx, r.key(k)).Result()
	if err == nil {
		return &val, true
	}

	if v, ok := r.env[k]; ok {
		return v, true
	}

	return nil, false
}

func (r *RedisStore) Put(k string, v *string) {
	r.Lock()
	defer r.Unlock()

	if r.lookupFn != nil {
		if lookupValue, ok := r.lookupFn(k); ok {
			v = &lookupValue
		}
	}

	if v == nil {
		r.client.Del(r.ctx, r.key(k))
		return
	}

	r.client.Set(r.ctx, r.key(k), *v, 0)
}

func (r *RedisStore) Remove(k string) bool {
	r.Lock()
	defer r.Unlock()

	deleted, err := r.client.Del(r.ctx, r.key(k)).Result()
	return err == nil && deleted > 0
}

func (r *RedisStore) ExportMap() map[string]*string {
	r.RLock()
	defer r.RUnlock()

	result := make(map[string]*string)

	iter := r.client.Scan(r.ctx, 0, r.prefix+"*", 0).Iterator()
	for iter.Next(r.ctx) {
		key := iter.Val()
		unprefixedKey := strings.TrimPrefix(key, r.prefix)

		val, err := r.client.Get(r.ctx, key).Result()
		if err == nil {
			result[unprefixedKey] = &val
		}
	}

	return result
}

var _ store.Store = &RedisStore{}
