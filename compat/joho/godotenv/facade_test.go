package godotenv

import (
	"bytes"
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
			err := os.WriteFile(filepath.Join(wd, ".env"), []byte(test.dotEnv), 0o600)
			assert.NilError(t, err)

			envMap, err := GetEnvFromFile(env2Map(test.osEnv), []string{filepath.Join(wd, ".env")})
			assert.NilError(t, err)

			assert.DeepEqual(t, test.expected, envMap)
		})
	}
}

func TestParse(t *testing.T) {
	input := "KEY1=value1\nKEY2=value2\n"
	reader := strings.NewReader(input)

	result, err := Parse(reader)
	assert.NilError(t, err)
	assert.Equal(t, len(result), 2)
	assert.Equal(t, result["KEY1"], "value1")
	assert.Equal(t, result["KEY2"], "value2")
}

func TestParse_EmptyInput(t *testing.T) {
	reader := strings.NewReader("")

	result, err := Parse(reader)
	assert.NilError(t, err)
	assert.Equal(t, len(result), 0)
}

func TestParse_WithQuotes(t *testing.T) {
	input := `KEY1="quoted value"
KEY2='single quoted'`
	reader := strings.NewReader(input)

	result, err := Parse(reader)
	assert.NilError(t, err)
	assert.Equal(t, result["KEY1"], "quoted value")
	assert.Equal(t, result["KEY2"], "single quoted")
}

func TestParseWithLookup(t *testing.T) {
	input := "KEY1=${EXTERNAL}\n"
	reader := strings.NewReader(input)

	lookupFn := func(k string) (string, bool) {
		if k == "EXTERNAL" {
			return "external_value", true
		}
		return "", false
	}

	result, err := ParseWithLookup(reader, lookupFn)
	assert.NilError(t, err)
	assert.Equal(t, result["KEY1"], "external_value")
}

func TestParseWithLookup_UTF8BOM(t *testing.T) {
	utf8BOM := []byte("\uFEFF")
	content := []byte("KEY=value")
	input := make([]byte, 0, len(utf8BOM)+len(content))
	input = append(input, utf8BOM...)
	input = append(input, content...)
	reader := bytes.NewReader(input)

	result, err := ParseWithLookup(reader, nil)
	assert.NilError(t, err)
	assert.Equal(t, result["KEY"], "value")
}

func TestUnmarshalBytesWithLookup(t *testing.T) {
	input := []byte("KEY1=value1\nKEY2=value2")

	result, err := UnmarshalBytesWithLookup(input, nil)
	assert.NilError(t, err)
	assert.Equal(t, len(result), 2)
	assert.Equal(t, result["KEY1"], "value1")
	assert.Equal(t, result["KEY2"], "value2")
}

func TestUnmarshalWithLookup(t *testing.T) {
	input := "KEY1=value1\nKEY2=value2"

	result, err := UnmarshalWithLookup(input, nil)
	assert.NilError(t, err)
	assert.Equal(t, len(result), 2)
	assert.Equal(t, result["KEY1"], "value1")
	assert.Equal(t, result["KEY2"], "value2")
}

func TestUnmarshalWithLookup_NilValues(t *testing.T) {
	input := "KEY1=\nKEY2=value2"

	result, err := UnmarshalWithLookup(input, nil)
	assert.NilError(t, err)
	assert.Equal(t, result["KEY1"], "")
	assert.Equal(t, result["KEY2"], "value2")
}

func TestLoad_SingleFile(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, "test.env")
	err := os.WriteFile(envFile, []byte("TEST_LOAD_KEY=test_value"), 0o600)
	assert.NilError(t, err)

	origDir, err := os.Getwd()
	assert.NilError(t, err)
	defer func() {
		_ = os.Chdir(origDir)
	}()

	err = os.Chdir(tmpDir)
	assert.NilError(t, err)

	os.Unsetenv("TEST_LOAD_KEY")

	err = Load("test.env")
	assert.NilError(t, err)

	value := os.Getenv("TEST_LOAD_KEY")
	assert.Equal(t, value, "test_value")

	os.Unsetenv("TEST_LOAD_KEY")
}

func TestLoad_MultipleFiles(t *testing.T) {
	tmpDir := t.TempDir()

	env1 := filepath.Join(tmpDir, "env1.env")
	err := os.WriteFile(env1, []byte("KEY1=value1"), 0o600)
	assert.NilError(t, err)

	env2 := filepath.Join(tmpDir, "env2.env")
	err = os.WriteFile(env2, []byte("KEY2=value2"), 0o600)
	assert.NilError(t, err)

	origDir, err := os.Getwd()
	assert.NilError(t, err)
	defer func() {
		_ = os.Chdir(origDir)
	}()

	err = os.Chdir(tmpDir)
	assert.NilError(t, err)

	os.Unsetenv("KEY1")
	os.Unsetenv("KEY2")

	err = Load("env1.env", "env2.env")
	assert.NilError(t, err)

	assert.Equal(t, os.Getenv("KEY1"), "value1")
	assert.Equal(t, os.Getenv("KEY2"), "value2")

	os.Unsetenv("KEY1")
	os.Unsetenv("KEY2")
}

func TestLoad_DefaultFilename(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, ".env")
	err := os.WriteFile(envFile, []byte("DEFAULT_KEY=default_value"), 0o600)
	assert.NilError(t, err)

	origDir, err := os.Getwd()
	assert.NilError(t, err)
	defer func() {
		_ = os.Chdir(origDir)
	}()

	err = os.Chdir(tmpDir)
	assert.NilError(t, err)

	os.Unsetenv("DEFAULT_KEY")

	err = Load()
	assert.NilError(t, err)

	value := os.Getenv("DEFAULT_KEY")
	assert.Equal(t, value, "default_value")

	os.Unsetenv("DEFAULT_KEY")
}

func TestLoad_FileNotFound(t *testing.T) {
	err := Load("nonexistent.env")
	assert.Assert(t, err != nil)
}

func TestLoad_PreservesExistingEnv(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, "test.env")
	err := os.WriteFile(envFile, []byte("PRESERVE_KEY=new_value"), 0o600)
	assert.NilError(t, err)

	origDir, err := os.Getwd()
	assert.NilError(t, err)
	defer func() {
		_ = os.Chdir(origDir)
	}()

	err = os.Chdir(tmpDir)
	assert.NilError(t, err)

	t.Setenv("PRESERVE_KEY", "original_value")

	err = Load("test.env")
	assert.NilError(t, err)

	value := os.Getenv("PRESERVE_KEY")
	assert.Equal(t, value, "original_value", "Load should not override existing env vars")
}

func TestRead_SingleFile(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, "test.env")
	err := os.WriteFile(envFile, []byte("READ_KEY1=read_value1\nREAD_KEY2=read_value2"), 0o600)
	assert.NilError(t, err)

	result, err := Read(envFile)
	assert.NilError(t, err)
	assert.Equal(t, len(result), 2)
	assert.Equal(t, result["READ_KEY1"], "read_value1")
	assert.Equal(t, result["READ_KEY2"], "read_value2")
}

func TestRead_MultipleFiles(t *testing.T) {
	tmpDir := t.TempDir()

	env1 := filepath.Join(tmpDir, "env1.env")
	err := os.WriteFile(env1, []byte("KEY1=value1"), 0o600)
	assert.NilError(t, err)

	env2 := filepath.Join(tmpDir, "env2.env")
	err = os.WriteFile(env2, []byte("KEY2=value2"), 0o600)
	assert.NilError(t, err)

	result, err := Read(env1, env2)
	assert.NilError(t, err)
	assert.Equal(t, len(result), 2)
	assert.Equal(t, result["KEY1"], "value1")
	assert.Equal(t, result["KEY2"], "value2")
}

func TestRead_DefaultFilename(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, ".env")
	err := os.WriteFile(envFile, []byte("DEFAULT_READ_KEY=default_read_value"), 0o600)
	assert.NilError(t, err)

	origDir, err := os.Getwd()
	assert.NilError(t, err)
	defer func() {
		_ = os.Chdir(origDir)
	}()

	err = os.Chdir(tmpDir)
	assert.NilError(t, err)

	result, err := Read()
	assert.NilError(t, err)
	assert.Equal(t, result["DEFAULT_READ_KEY"], "default_read_value")
}

func TestRead_FileNotFound(t *testing.T) {
	_, err := Read("nonexistent.env")
	assert.Assert(t, err != nil)
}

func TestReadWithLookup(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, "test.env")
	err := os.WriteFile(envFile, []byte("KEY=value"), 0o600)
	assert.NilError(t, err)

	lookupFn := func(k string) (string, bool) {
		return "", false
	}

	result, err := ReadWithLookup(lookupFn, envFile)
	assert.NilError(t, err)
	assert.Equal(t, result["KEY"], "value")
}

func TestReadFile(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, "test.env")
	err := os.WriteFile(envFile, []byte("READFILE_KEY=readfile_value"), 0o600)
	assert.NilError(t, err)

	result, err := ReadFile(envFile, nil)
	assert.NilError(t, err)
	assert.Equal(t, result["READFILE_KEY"], "readfile_value")
}

func TestReadFile_WithLookupFn(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, "test.env")
	err := os.WriteFile(envFile, []byte("KEY=${EXTERNAL}"), 0o600)
	assert.NilError(t, err)

	lookupFn := func(k string) (string, bool) {
		if k == "EXTERNAL" {
			return "external_value", true
		}
		return "", false
	}

	result, err := ReadFile(envFile, lookupFn)
	assert.NilError(t, err)
	assert.Equal(t, result["KEY"], "external_value")
}

func TestReadFile_FileNotFound(t *testing.T) {
	_, err := ReadFile("nonexistent.env", nil)
	assert.Assert(t, err != nil)
}

func TestGetEnvFromFile_FileNotFound(t *testing.T) {
	_, err := GetEnvFromFile(nil, []string{"nonexistent.env"})
	assert.Assert(t, err != nil)
	assert.Assert(t, strings.Contains(err.Error(), "couldn't find env file"))
}

func TestGetEnvFromFile_IsDirectory(t *testing.T) {
	tmpDir := t.TempDir()

	_, err := GetEnvFromFile(nil, []string{tmpDir})
	assert.Assert(t, err != nil)
	assert.Assert(t, strings.Contains(err.Error(), "is a directory"))
}

func TestGetEnvFromFile_MultipleFiles(t *testing.T) {
	tmpDir := t.TempDir()

	env1 := filepath.Join(tmpDir, "env1.env")
	err := os.WriteFile(env1, []byte("KEY1=value1"), 0o600)
	assert.NilError(t, err)

	env2 := filepath.Join(tmpDir, "env2.env")
	err = os.WriteFile(env2, []byte("KEY2=value2"), 0o600)
	assert.NilError(t, err)

	result, err := GetEnvFromFile(nil, []string{env1, env2})
	assert.NilError(t, err)
	assert.Equal(t, len(result), 2)
	assert.Equal(t, result["KEY1"], "value1")
	assert.Equal(t, result["KEY2"], "value2")
}

func TestGetEnvFromFile_EmptyCurrentEnv(t *testing.T) {
	tmpDir := t.TempDir()

	envFile := filepath.Join(tmpDir, "test.env")
	err := os.WriteFile(envFile, []byte("KEY1=value1\nKEY2=value2"), 0o600)
	assert.NilError(t, err)

	result, err := GetEnvFromFile(map[string]string{}, []string{envFile})
	assert.NilError(t, err)
	assert.Equal(t, len(result), 2)
	assert.Equal(t, result["KEY1"], "value1")
	assert.Equal(t, result["KEY2"], "value2")
}

func TestGetEnvFromFile_RelativePath(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, "test.env")
	err := os.WriteFile(envFile, []byte("KEY=value"), 0o600)
	assert.NilError(t, err)

	origDir, err := os.Getwd()
	assert.NilError(t, err)
	defer func() {
		_ = os.Chdir(origDir)
	}()

	err = os.Chdir(tmpDir)
	assert.NilError(t, err)

	result, err := GetEnvFromFile(nil, []string{"test.env"})
	assert.NilError(t, err)
	assert.Equal(t, result["KEY"], "value")
}

func TestFilenamesOrDefault_Empty(t *testing.T) {
	result := filenamesOrDefault([]string{})
	assert.Equal(t, len(result), 1)
	assert.Equal(t, result[0], ".env")
}

func TestFilenamesOrDefault_Provided(t *testing.T) {
	input := []string{"custom.env", "other.env"}
	result := filenamesOrDefault(input)
	assert.Equal(t, len(result), 2)
	assert.Equal(t, result[0], "custom.env")
	assert.Equal(t, result[1], "other.env")
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

func TestOverload(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, "test.env")
	err := os.WriteFile(envFile, []byte("OVERLOAD_KEY=new_value"), 0o600)
	assert.NilError(t, err)

	// Set existing env var
	t.Setenv("OVERLOAD_KEY", "original_value")

	origDir, err := os.Getwd()
	assert.NilError(t, err)
	defer func() {
		_ = os.Chdir(origDir)
	}()

	err = os.Chdir(tmpDir)
	assert.NilError(t, err)

	// Overload should override the existing value
	err = Overload("test.env")
	assert.NilError(t, err)

	value := os.Getenv("OVERLOAD_KEY")
	assert.Equal(t, value, "new_value")
}

func TestOverload_DefaultFilename(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, ".env")
	err := os.WriteFile(envFile, []byte("OVERLOAD_DEFAULT=value"), 0o600)
	assert.NilError(t, err)

	origDir, err := os.Getwd()
	assert.NilError(t, err)
	defer func() {
		_ = os.Chdir(origDir)
	}()

	err = os.Chdir(tmpDir)
	assert.NilError(t, err)

	err = Overload()
	assert.NilError(t, err)

	value := os.Getenv("OVERLOAD_DEFAULT")
	assert.Equal(t, value, "value")
}

func TestUnmarshal(t *testing.T) {
	input := "KEY1=value1\nKEY2=value2"
	result, err := Unmarshal(input)
	assert.NilError(t, err)
	assert.Equal(t, len(result), 2)
	assert.Equal(t, result["KEY1"], "value1")
	assert.Equal(t, result["KEY2"], "value2")
}

func TestUnmarshal_EmptyString(t *testing.T) {
	result, err := Unmarshal("")
	assert.NilError(t, err)
	assert.Equal(t, len(result), 0)
}

func TestUnmarshalBytes(t *testing.T) {
	input := []byte("KEY1=value1\nKEY2=value2")
	result, err := UnmarshalBytes(input)
	assert.NilError(t, err)
	assert.Equal(t, len(result), 2)
	assert.Equal(t, result["KEY1"], "value1")
	assert.Equal(t, result["KEY2"], "value2")
}

func TestUnmarshalBytes_EmptyBytes(t *testing.T) {
	result, err := UnmarshalBytes([]byte{})
	assert.NilError(t, err)
	assert.Equal(t, len(result), 0)
}

func TestMarshal(t *testing.T) {
	envMap := map[string]string{
		"KEY1": "value1",
		"KEY2": "value with spaces",
	}

	result, err := Marshal(envMap)
	assert.NilError(t, err)

	// Parse it back to verify
	parsed, err := Unmarshal(result)
	assert.NilError(t, err)
	assert.Equal(t, parsed["KEY1"], "value1")
	assert.Equal(t, parsed["KEY2"], "value with spaces")
}

func TestMarshal_EmptyMap(t *testing.T) {
	envMap := map[string]string{}
	result, err := Marshal(envMap)
	assert.NilError(t, err)
	assert.Equal(t, result, "\n")
}

func TestMarshal_SpecialCharacters(t *testing.T) {
	envMap := map[string]string{
		"KEY": "value\nwith\nnewlines",
	}

	result, err := Marshal(envMap)
	assert.NilError(t, err)

	// Parse it back to verify escaping works
	parsed, err := Unmarshal(result)
	assert.NilError(t, err)
	assert.Equal(t, parsed["KEY"], "value\nwith\nnewlines")
}

func TestWrite(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, "output.env")

	envMap := map[string]string{
		"WRITE_KEY1": "value1",
		"WRITE_KEY2": "value2",
	}

	err := Write(envMap, envFile)
	assert.NilError(t, err)

	// Verify file was created
	_, err = os.Stat(envFile)
	assert.NilError(t, err)

	// Read it back
	result, err := Read(envFile)
	assert.NilError(t, err)
	assert.Equal(t, result["WRITE_KEY1"], "value1")
	assert.Equal(t, result["WRITE_KEY2"], "value2")
}

func TestWrite_InvalidPath(t *testing.T) {
	envMap := map[string]string{"KEY": "value"}
	err := Write(envMap, "/nonexistent/path/file.env")
	assert.ErrorContains(t, err, "no such file or directory")
}

func TestExec(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, "exec.env")
	err := os.WriteFile(envFile, []byte("EXEC_TEST_KEY=exec_value"), 0o600)
	assert.NilError(t, err)

	// Test executing echo command
	err = Exec([]string{envFile}, "echo", []string{"test"}, false)
	assert.NilError(t, err)
}

func TestExec_Overload(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, "exec_overload.env")
	err := os.WriteFile(envFile, []byte("EXEC_OVERLOAD_KEY=new_value"), 0o600)
	assert.NilError(t, err)

	t.Setenv("EXEC_OVERLOAD_KEY", "original")

	// Test with overload=true
	err = Exec([]string{envFile}, "echo", []string{"test"}, true)
	assert.NilError(t, err)

	// The environment should have been overloaded
	value := os.Getenv("EXEC_OVERLOAD_KEY")
	assert.Equal(t, value, "new_value")
}

func TestExec_CommandNotFound(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, "exec.env")
	err := os.WriteFile(envFile, []byte("KEY=value"), 0o600)
	assert.NilError(t, err)

	err = Exec([]string{envFile}, "nonexistentcommand12345", []string{}, false)
	assert.ErrorContains(t, err, "")
}
