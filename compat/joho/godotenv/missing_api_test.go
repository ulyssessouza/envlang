package godotenv

import (
	"os"
	"path/filepath"
	"testing"

	"gotest.tools/v3/assert"
)

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
