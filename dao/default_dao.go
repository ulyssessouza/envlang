package dao

import (
	"strings"
	"sync"
)

const (
	pairSeparator = "="
	pairNumber    = 2
)

var _ EnvLangDao = &DefaultDao{}

type DefaultDao struct {
	sync.RWMutex

	env map[string]*string
	m   map[string]*string

	lookupFn LookupFn
}

func NewDefaultDao() EnvLangDao {
	return &DefaultDao{
		m:   make(map[string]*string),
		env: make(map[string]*string),
	}
}

type EnvlangDaoOptionsFn func(d EnvLangDao)

func WithLookupFn(fn LookupFn) EnvlangDaoOptionsFn {
	return func(d EnvLangDao) {
		if d, ok := d.(*DefaultDao); ok {
			d.lookupFn = fn
		}
	}
}

func NewDefaultDaoFromEnv(env []string, opts ...EnvlangDaoOptionsFn) EnvLangDao {
	d := &DefaultDao{
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

func NewDefaultDaoFromMap(env map[string]*string) EnvLangDao {
	if env == nil {
		env = make(map[string]*string)
	}
	return &DefaultDao{
		m:   make(map[string]*string),
		env: env,
	}
}

func (d *DefaultDao) ImportList(env []string) {
	d.Lock()
	defer d.Unlock()

	for _, e := range env {
		splitEnv := strings.SplitN(e, pairSeparator, pairNumber)
		v := splitEnv[1]
		d.m[splitEnv[0]] = &v
	}
}

func (d *DefaultDao) ImportMap(m map[string]string) {
	d.Lock()
	defer d.Unlock()

	for k, v := range m {
		d.m[k] = &v
	}
}

func (d *DefaultDao) ExportMap() map[string]*string {
	return d.m
}

func (d *DefaultDao) Get(k string) (*string, bool) {
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

func (d *DefaultDao) Put(k string, v *string) {
	d.Lock()
	defer d.Unlock()

	if d.lookupFn != nil {
		if lookupValue, ok := d.lookupFn(k); ok {
			v = &lookupValue
		}
	}

	d.m[k] = v
}

func (d *DefaultDao) Remove(k string) {
	d.Lock()
	defer d.Unlock()

	delete(d.m, k)
}
