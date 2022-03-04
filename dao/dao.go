package dao

import (
	"strings"
	"sync"
)

type EnvLangDao interface {
	ImportList([]string)
	ImportMap(map[string]string)
	Get(string) (*string, bool)
	Put(string, *string)
	Remove(string)
	ExportMap() map[string]*string
}

var _ EnvLangDao = &DefaultEnvLangDao{}

type DefaultEnvLangDao struct {
	sync.RWMutex

	m map[string]*string
}

func NewDefaultEnvLangDao() EnvLangDao {
	return &DefaultEnvLangDao{
		m: make(map[string]*string),
	}
}

func (d *DefaultEnvLangDao) ImportList(env []string) {
	d.Lock()
	defer d.Unlock()

	for _, e := range env {
		splitEnv := strings.SplitN(e, "=", 2)
		v := splitEnv[1]
		d.m[splitEnv[0]] = &v
	}
}

func (d *DefaultEnvLangDao) ImportMap(m map[string]string) {
	d.Lock()
	defer d.Unlock()

	for k, v := range m {
		d.m[k] = &v
	}
}

func (d *DefaultEnvLangDao) ExportMap() map[string]*string {
	return d.m
}

func (d *DefaultEnvLangDao) Get(k string) (*string, bool) {
	d.RLock()
	defer d.RUnlock()

	v, ok := d.m[k]
	return v, ok
}

func (d *DefaultEnvLangDao) Put(k string, v *string) {
	d.Lock()
	defer d.Unlock()

	d.m[k] = v
}

func (d *DefaultEnvLangDao) Remove(k string) {
	d.Lock()
	defer d.Unlock()

	delete(d.m, k)
}
