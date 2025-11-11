package store

import (
	"testing"

	"gotest.tools/v3/assert"
)

const (
	testValue = "test_value"
	mapValue  = "map_value"
)

func TestNewDefaultStore(t *testing.T) {
	d := NewDefaultStore()
	assert.Assert(t, d != nil)

	_, ok := d.Get("NONEXISTENT")
	assert.Assert(t, !ok)
}

func TestNewDefaultStoreFromMap(t *testing.T) {
	value := testValue
	env := map[string]*string{
		"KEY1": &value,
	}

	d := NewDefaultStoreFromMap(env)
	assert.Assert(t, d != nil)

	v, ok := d.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, testValue)
}

func TestNewDefaultStoreFromMap_NilMap(t *testing.T) {
	d := NewDefaultStoreFromMap(nil)
	assert.Assert(t, d != nil)

	_, ok := d.Get("ANY_KEY")
	assert.Assert(t, !ok)
}

func TestNewDefaultStoreFromEnv(t *testing.T) {
	env := []string{
		"KEY1=value1",
		"KEY2=value2",
	}

	d := NewDefaultStoreFromEnv(env)
	assert.Assert(t, d != nil)

	v, ok := d.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value1")

	v, ok = d.Get("KEY2")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value2")
}

func TestNewDefaultStoreFromEnv_WithLookupFn(t *testing.T) {
	env := []string{"KEY1=value1"}

	lookupFn := func(k string) (string, bool) {
		if k == "CUSTOM_KEY" {
			return "custom_value", true
		}
		return "", false
	}

	d := NewDefaultStoreFromEnv(env, WithLookupFn(lookupFn))

	v, ok := d.Get("CUSTOM_KEY")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "custom_value")
}

func TestDefaultStore_PutAndGet(t *testing.T) {
	d := NewDefaultStore()

	value := testValue
	d.Put("KEY1", &value)

	v, ok := d.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, testValue)
}

func TestDefaultStore_Put_NilValue(t *testing.T) {
	d := NewDefaultStore()

	d.Put("KEY1", nil)

	v, ok := d.Get("KEY1")
	assert.Assert(t, ok)
	assert.Assert(t, v == nil)
}

func TestDefaultStore_Get_NonExistent(t *testing.T) {
	d := NewDefaultStore()

	_, ok := d.Get("NONEXISTENT")
	assert.Assert(t, !ok)
}

func TestDefaultStore_Get_PriorityOrder(t *testing.T) {
	envValue := "env_value"
	mapVal := mapValue
	lookupValue := "lookup_value"

	env := map[string]*string{
		"KEY": &envValue,
	}

	d := NewDefaultStoreFromMap(env)

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

func TestDefaultStore_Remove(t *testing.T) {
	d := NewDefaultStore()

	value := testValue
	d.Put("KEY1", &value)

	v, ok := d.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, testValue)

	assert.Equal(t, d.Remove("KEY1"), true)

	_, ok = d.Get("KEY1")
	assert.Assert(t, !ok)
}

func TestDefaultStore_Remove_NonExistent(t *testing.T) {
	d := NewDefaultStore()

	assert.Equal(t, d.Remove("NONEXISTENT"), false)
}

func TestDefaultStore_ImportList(t *testing.T) {
	d := NewDefaultStore()

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

func TestDefaultStore_ImportList_WithEquals(t *testing.T) {
	d := NewDefaultStore()

	env := []string{
		"KEY1=value=with=equals",
	}

	d.ImportList(env)

	v, ok := d.Get("KEY1")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "value=with=equals")
}

func TestDefaultStore_ImportMap(t *testing.T) {
	d := NewDefaultStore()

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

func TestDefaultStore_ImportMap_EmptyMap(t *testing.T) {
	d := NewDefaultStore()

	d.ImportMap(map[string]string{})

	_, ok := d.Get("ANY_KEY")
	assert.Assert(t, !ok)
}

func TestDefaultStore_ExportMap(t *testing.T) {
	d := NewDefaultStore()

	value1 := "value1"
	value2 := "value2"
	d.Put("KEY1", &value1)
	d.Put("KEY2", &value2)

	exported := d.ExportMap()
	assert.Equal(t, len(exported), 2)
	assert.Equal(t, *exported["KEY1"], value1)
	assert.Equal(t, *exported["KEY2"], value2)
}

func TestDefaultStore_ExportMap_Empty(t *testing.T) {
	d := NewDefaultStore()

	exported := d.ExportMap()
	assert.Equal(t, len(exported), 0)
}

func TestDefaultStore_ExportMap_DoesNotIncludeEnv(t *testing.T) {
	envValue := "env_value"
	env := map[string]*string{
		"ENV_KEY": &envValue,
	}
	d := NewDefaultStoreFromMap(env)

	mapVal := mapValue
	d.Put("MAP_KEY", &mapVal)

	exported := d.ExportMap()
	assert.Equal(t, len(exported), 1)
	assert.Equal(t, *exported["MAP_KEY"], mapValue)
	_, exists := exported["ENV_KEY"]
	assert.Assert(t, !exists, "ExportMap should only export from m, not env")
}

func TestDefaultStore_WithLookupFn(t *testing.T) {
	lookupFn := func(k string) (string, bool) {
		if k == "LOOKUP_KEY" {
			return "lookup_value", true
		}
		return "", false
	}

	d := NewDefaultStore()
	WithLookupFn(lookupFn)(d)

	v, ok := d.Get("LOOKUP_KEY")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "lookup_value")
}

func TestDefaultStore_WithLookupFn_FallbackToMap(t *testing.T) {
	lookupFn := func(k string) (string, bool) {
		return "", false
	}

	d := NewDefaultStore()
	WithLookupFn(lookupFn)(d)

	mapVal := mapValue
	d.Put("KEY", &mapVal)

	v, ok := d.Get("KEY")
	assert.Assert(t, ok)
	assert.Equal(t, *v, mapValue)
}

func TestDefaultStore_Put_WithLookupFn(t *testing.T) {
	lookupFn := func(k string) (string, bool) {
		if k == "KEY" {
			return "lookup_override", true
		}
		return "", false
	}

	d := NewDefaultStore()
	WithLookupFn(lookupFn)(d)

	originalValue := "original_value"
	d.Put("KEY", &originalValue)

	v, ok := d.Get("KEY")
	assert.Assert(t, ok)
	assert.Equal(t, *v, "lookup_override", "Put should use lookupFn value when available")
}
