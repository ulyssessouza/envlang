package compat_env

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestEnvVariablePrecedence(t *testing.T) {
	testcases := []struct {
		name     string
		dotEnv   string
		osEnv    []string
		expected map[string]string
	}{
		{
			"no value set in environment",
			"FOO=foo\nBAR=${FOO}",
			nil,
			map[string]string{
				"FOO": "foo",
				"BAR": "foo",
			},
		},
		{
			"conflict with value set in environment",
			"FOO=foo\nBAR=${FOO}",
			[]string{"FOO=zot"},
			map[string]string{
				"FOO": "zot",
				"BAR": "zot",
			},
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			wd := t.TempDir()
			err := os.WriteFile(filepath.Join(wd, ".env"), []byte(test.dotEnv), 0o700)
			assert.NilError(t, err)

			envMap, err := GetEnvFromFile(env2Map(test.osEnv), []string{filepath.Join(wd, ".env")})
			assert.NilError(t, err)

			assert.DeepEqual(t, test.expected, envMap)
		})
	}
}

func env2Map(env []string) map[string]string {
	m := make(map[string]string)
	for _, e := range env {
		k, v, b := strings.Cut(e, "=")
		if !b {
			v = ""
		}
		m[k] = v
	}
	return m
}
