package store

import (
	"strings"
	"sync"
)

const (
	pairSeparator = "="
	pairNumber    = 2
)

var _ Store = &DefaultStore{}

type DefaultStore struct {
	sync.RWMutex

	env map[string]*string
	m   map[string]*string

	lookupFn LookupFn
}

func NewDefaultStore() Store {
	return &DefaultStore{
		m:   make(map[string]*string),
		env: make(map[string]*string),
	}
}

type StoreOptionsFn func(d Store)

func WithLookupFn(fn LookupFn) StoreOptionsFn {
	return func(d Store) {
		if d, ok := d.(*DefaultStore); ok {
			d.lookupFn = fn
		}
	}
}

func NewDefaultStoreFromEnv(env []string, opts ...StoreOptionsFn) Store {
	d := &DefaultStore{
		m:   make(map[string]*string),
		env: make(map[string]*string),
	}

	for _, e := range env {
		splitEnv := strings.SplitN(e, pairSeparator, pairNumber)
		v := splitEnv[1]
		d.env[splitEnv[0]] = &v
	}

	for _, opt := range opts {
		opt(d)
	}

	return d
}

func NewDefaultStoreFromMap(env map[string]*string) Store {
	if env == nil {
		env = make(map[string]*string)
	}
	return &DefaultStore{
		m:   make(map[string]*string),
		env: env,
	}
}

func (d *DefaultStore) ImportList(env []string) {
	d.Lock()
	defer d.Unlock()

	for _, e := range env {
		splitEnv := strings.SplitN(e, pairSeparator, pairNumber)
		v := splitEnv[1]
		d.m[splitEnv[0]] = &v
	}
}

func (d *DefaultStore) ImportMap(m map[string]string) {
	d.Lock()
	defer d.Unlock()

	for k, v := range m {
		d.m[k] = &v
	}
}

func (d *DefaultStore) ExportMap() map[string]*string {
	return d.m
}

func (d *DefaultStore) Get(k string) (*string, bool) {
	d.RLock()
	defer d.RUnlock()

	if d.lookupFn != nil {
		if v, ok := d.lookupFn(k); ok {
			return &v, true
		}
	}

	v, ok := d.m[k]
	if !ok {
		v, ok = d.env[k]
	}
	return v, ok
}

func (d *DefaultStore) Put(k string, v *string) {
	d.Lock()
	defer d.Unlock()

	if d.lookupFn != nil {
		if lookupValue, ok := d.lookupFn(k); ok {
			v = &lookupValue
		}
	}

	d.m[k] = v
}

func (d *DefaultStore) Remove(k string) bool {
	d.Lock()
	defer d.Unlock()

	_, ok := d.m[k]
	delete(d.m, k)
	return ok
}
