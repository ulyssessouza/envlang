package dao

import (
	"strings"
	"sync"
)

var _ EnvLangDao = &DefaultDao{}

type DefaultDao struct {
	sync.RWMutex

	m map[string]*string
}

func NewDefaultDao() EnvLangDao {
	return &DefaultDao{
		m: make(map[string]*string),
	}
}

func (d *DefaultDao) ImportList(env []string) {
	d.Lock()
	defer d.Unlock()

	for _, e := range env {
		splitEnv := strings.SplitN(e, "=", 2)
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

	v, ok := d.m[k]
	return v, ok
}

func (d *DefaultDao) Put(k string, v *string) {
	d.Lock()
	defer d.Unlock()

	d.m[k] = v
}

func (d *DefaultDao) Remove(k string) {
	d.Lock()
	defer d.Unlock()

	delete(d.m, k)
}
