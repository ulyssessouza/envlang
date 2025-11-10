package dao

import (
	"sync"
	"testing"

	"gotest.tools/v3/assert"
)

const (
	testValue = "test_value"
	mapValue  = "map_value"
)

func TestNewDefaultDao(t *testing.T) {
	d := NewDefaultDao()
	assert.Assert(t, d != nil)

	_, ok := d.Get("NONEXISTENT")
	assert.Assert(t, !ok)
}

func TestNewDefaultDaoFromMap(t *testing.T) {
	value := testValue
	env := map[string]*string{
		"KEY1": &value,
	}

	d := NewDefaultDaoFromMap(env)
	assert.Assert(t, d != nil)

	v, ok := d.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, testValue)
}

func TestNewDefaultDaoFromMap_NilMap(t *testing.T) {
	d := NewDefaultDaoFromMap(nil)
	assert.Assert(t, d != nil)

	_, ok := d.Get("ANY_KEY")
	assert.Assert(t, !ok)
}

func TestNewDefaultDaoFromEnv(t *testing.T) {
	env := []string{
		"KEY1=value1",
		"KEY2=value2",
	}

	d := NewDefaultDaoFromEnv(env)
	assert.Assert(t, d != nil)

	v, ok := d.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value1")

	v, ok = d.Get("KEY2")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value2")
}

func TestNewDefaultDaoFromEnv_WithLookupFn(t *testing.T) {
	env := []string{"KEY1=value1"}

	lookupFn := func(k string) (string, bool) {
		if k == "CUSTOM_KEY" {
			return "custom_value", true
		}
		return "", false
	}

	d := NewDefaultDaoFromEnv(env, WithLookupFn(lookupFn))

	v, ok := d.Get("CUSTOM_KEY")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "custom_value")
}

func TestDefaultDao_PutAndGet(t *testing.T) {
	d := NewDefaultDao()

	value := testValue
	d.Put("KEY1", &value)

	v, ok := d.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, testValue)
}

func TestDefaultDao_Put_NilValue(t *testing.T) {
	d := NewDefaultDao()

	d.Put("KEY1", nil)

	v, ok := d.Get("KEY1")
	assert.Assert(t, ok)
	assert.Assert(t, v == nil)
}

func TestDefaultDao_Get_NonExistent(t *testing.T) {
	d := NewDefaultDao()

	_, ok := d.Get("NONEXISTENT")
	assert.Assert(t, !ok)
}

func TestDefaultDao_Get_PriorityOrder(t *testing.T) {
	envValue := "env_value"
	mapVal := mapValue
	lookupValue := "lookup_value"

	env := map[string]*string{
		"KEY": &envValue,
	}

	d := NewDefaultDaoFromMap(env)

	v, ok := d.Get("KEY")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "env_value", "should get from env when m is empty")

	d.Put("KEY", &mapVal)
	v, ok = d.Get("KEY")
	assert.Assert(t, ok)
	assert.Equal(t, *v, mapValue, "should get from m when both exist")

	lookupFn := func(k string) (string, bool) {
		if k == "KEY" {
			return lookupValue, true
		}
		return "", false
	}
	WithLookupFn(lookupFn)(d)

	v, ok = d.Get("KEY")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "lookup_value", "should get from lookupFn when all exist")
}

func TestDefaultDao_Remove(t *testing.T) {
	d := NewDefaultDao()

	value := testValue
	d.Put("KEY1", &value)

	v, ok := d.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, testValue)

	d.Remove("KEY1")

	_, ok = d.Get("KEY1")
	assert.Assert(t, !ok)
}

func TestDefaultDao_Remove_NonExistent(t *testing.T) {
	d := NewDefaultDao()

	d.Remove("NONEXISTENT")
}

func TestDefaultDao_ImportList(t *testing.T) {
	d := NewDefaultDao()

	env := []string{
		"KEY1=value1",
		"KEY2=value2",
		"KEY3=value with spaces",
	}

	d.ImportList(env)

	v, ok := d.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value1")

	v, ok = d.Get("KEY2")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value2")

	v, ok = d.Get("KEY3")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value with spaces")
}

func TestDefaultDao_ImportList_WithEquals(t *testing.T) {
	d := NewDefaultDao()

	env := []string{
		"KEY1=value=with=equals",
	}

	d.ImportList(env)

	v, ok := d.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value=with=equals")
}

func TestDefaultDao_ImportMap(t *testing.T) {
	d := NewDefaultDao()

	m := map[string]string{
		"KEY1": "value1",
		"KEY2": "value2",
	}

	d.ImportMap(m)

	v, ok := d.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value1")

	v, ok = d.Get("KEY2")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value2")
}

func TestDefaultDao_ImportMap_EmptyMap(t *testing.T) {
	d := NewDefaultDao()

	d.ImportMap(map[string]string{})

	_, ok := d.Get("ANY_KEY")
	assert.Assert(t, !ok)
}

func TestDefaultDao_ExportMap(t *testing.T) {
	d := NewDefaultDao()

	value1 := "value1"
	value2 := "value2"
	d.Put("KEY1", &value1)
	d.Put("KEY2", &value2)

	exported := d.ExportMap()
	assert.Equal(t, len(exported), 2)
	assert.Equal(t, *exported["KEY1"], value1)
	assert.Equal(t, *exported["KEY2"], value2)
}

func TestDefaultDao_ExportMap_Empty(t *testing.T) {
	d := NewDefaultDao()

	exported := d.ExportMap()
	assert.Equal(t, len(exported), 0)
}

func TestDefaultDao_ExportMap_DoesNotIncludeEnv(t *testing.T) {
	envValue := "env_value"
	env := map[string]*string{
		"ENV_KEY": &envValue,
	}
	d := NewDefaultDaoFromMap(env)

	mapVal := mapValue
	d.Put("MAP_KEY", &mapVal)

	exported := d.ExportMap()
	assert.Equal(t, len(exported), 1)
	assert.Equal(t, *exported["MAP_KEY"], mapValue)
	_, exists := exported["ENV_KEY"]
	assert.Assert(t, !exists, "ExportMap should only export from m, not env")
}

func TestDefaultDao_WithLookupFn(t *testing.T) {
	lookupFn := func(k string) (string, bool) {
		if k == "LOOKUP_KEY" {
			return "lookup_value", true
		}
		return "", false
	}

	d := NewDefaultDao()
	WithLookupFn(lookupFn)(d)

	v, ok := d.Get("LOOKUP_KEY")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "lookup_value")
}

func TestDefaultDao_WithLookupFn_FallbackToMap(t *testing.T) {
	lookupFn := func(k string) (string, bool) {
		return "", false
	}

	d := NewDefaultDao()
	WithLookupFn(lookupFn)(d)

	mapVal := mapValue
	d.Put("KEY", &mapVal)

	v, ok := d.Get("KEY")
	assert.Assert(t, ok)
	assert.Equal(t, *v, mapValue)
}

func TestDefaultDao_Put_WithLookupFn(t *testing.T) {
	lookupFn := func(k string) (string, bool) {
		if k == "KEY" {
			return "lookup_override", true
		}
		return "", false
	}

	d := NewDefaultDao()
	WithLookupFn(lookupFn)(d)

	originalValue := "original_value"
	d.Put("KEY", &originalValue)

	v, ok := d.Get("KEY")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "lookup_override", "Put should use lookupFn value when available")
}

func TestDefaultDao_ConcurrentAccess(t *testing.T) {
	d := NewDefaultDao()

	var wg sync.WaitGroup
	iterations := 100

	for i := 0; i < iterations; i++ {
		wg.Add(3)

		go func(_ int) {
			defer wg.Done()
			value := "value"
			d.Put("KEY", &value)
		}(i)

		go func(_ int) {
			defer wg.Done()
			d.Get("KEY")
		}(i)

		go func(_ int) {
			defer wg.Done()
			d.Remove("KEY")
		}(i)
	}

	wg.Wait()
}

func TestDefaultDao_ConcurrentImports(t *testing.T) {
	d := NewDefaultDao()

	var wg sync.WaitGroup
	iterations := 50

	for i := 0; i < iterations; i++ {
		wg.Add(2)

		go func(_ int) {
			defer wg.Done()
			d.ImportList([]string{"KEY=value"})
		}(i)

		go func(_ int) {
			defer wg.Done()
			d.ImportMap(map[string]string{"KEY2": "value2"})
		}(i)
	}

	wg.Wait()

	v, ok := d.Get("KEY")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value")

	v, ok = d.Get("KEY2")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value2")
}
