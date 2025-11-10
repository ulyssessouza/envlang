package godotenv

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ulyssessouza/envlang"
	"github.com/ulyssessouza/envlang/dao"
)

const defaultFileMode = 0o600

//nolint:gocognit
func GetEnvFromFile(currentEnv map[string]string, filenames []string) (map[string]string, error) {
	envMap := make(map[string]string)

	for _, dotEnvFile := range filenames {
		abs, err := filepath.Abs(dotEnvFile)
		if err != nil {
			return envMap, err
		}
		dotEnvFile = abs

		s, err := os.Stat(dotEnvFile)
		if os.IsNotExist(err) {
			return envMap, fmt.Errorf("couldn't find env file: %s", dotEnvFile)
		}
		if err != nil {
			return envMap, err
		}

		if s.IsDir() {
			if len(filenames) == 0 {
				return envMap, nil
			}
			return envMap, fmt.Errorf("%s is a directory", dotEnvFile)
		}

		b, err := os.ReadFile(dotEnvFile)
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("couldn't read env file: %s", dotEnvFile)
		}
		if err != nil {
			return envMap, err
		}

		env, err := ParseWithLookup(bytes.NewReader(b), func(k string) (string, bool) {
			v, ok := currentEnv[k]
			if ok {
				return v, true
			}
			v, ok = envMap[k]
			return v, ok
		})
		if err != nil {
			return envMap, fmt.Errorf("failed to read %s: %w", dotEnvFile, err)
		}
		for k, v := range env {
			envMap[k] = v
		}
	}

	return envMap, nil
}

// Parse reads an env file from io.Reader, returning a map of keys and values.
func Parse(r io.Reader) (map[string]string, error) {
	return ParseWithLookup(r, nil)
}

// ParseWithLookup reads an env file from io.Reader, returning a map of keys and values.
func ParseWithLookup(r io.Reader, lookupFn dao.LookupFn) (map[string]string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	// seek past the UTF-8 BOM if it exists (particularly on Windows, some
	// editors tend to add it, and it'll cause parsing to fail)
	utf8BOM := []byte("\uFEFF")
	data = bytes.TrimPrefix(data, utf8BOM)

	return UnmarshalBytesWithLookup(data, lookupFn)
}

// Unmarshal parses env file from string, returning a map of keys and values.
func Unmarshal(str string) (map[string]string, error) {
	return UnmarshalWithLookup(str, nil)
}

// UnmarshalBytes parses env file from byte slice of chars, returning a map of keys and values.
func UnmarshalBytes(src []byte) (map[string]string, error) {
	return UnmarshalBytesWithLookup(src, nil)
}

// UnmarshalBytesWithLookup parses env file from byte slice of chars, returning a map of keys and values.
func UnmarshalBytesWithLookup(src []byte, lookupFn dao.LookupFn) (map[string]string, error) {
	return UnmarshalWithLookup(string(src), lookupFn)
}

// UnmarshalWithLookup parses env file from string, returning a map of keys and values.
func UnmarshalWithLookup(src string, lookupFn dao.LookupFn) (map[string]string, error) {
	m := make(map[string]string)
	defaultDAO := dao.NewDefaultDaoFromEnv(os.Environ(), dao.WithLookupFn(lookupFn))
	srcMap := envlang.GetVariablesFromInputStream(defaultDAO, bytes.NewBufferString(src))

	for k, v := range srcMap {
		value := ""
		if v != nil {
			value = *v
		}
		m[k] = value
	}
	return m, nil
}

// Load reads env file(s) and loads them into ENV for this process.
// It will NOT override an env variable that already exists.
func Load(filenames ...string) error {
	filenames = filenamesOrDefault(filenames)
	for _, filename := range filenames {
		err := loadFile(filename, false)
		if err != nil {
			return err
		}
	}
	return nil
}

// Overload reads env file(s) and loads them into ENV for this process.
// It WILL override an env variable that already exists.
func Overload(filenames ...string) error {
	filenames = filenamesOrDefault(filenames)
	for _, filename := range filenames {
		err := loadFile(filename, true)
		if err != nil {
			return err
		}
	}
	return nil
}

func filenamesOrDefault(filenames []string) []string {
	if len(filenames) == 0 {
		return []string{".env"}
	}
	return filenames
}

func loadFile(filename string, overload bool) error {
	defaultDAO := dao.NewDefaultDaoFromEnv(os.Environ())
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	m := envlang.GetVariablesFromInputStream(defaultDAO, f)
	for k, v := range m {
		if !overload {
			if _, ok := os.LookupEnv(k); ok {
				continue
			}
		}
		if v == nil {
			v = new(string)
		}
		os.Setenv(k, *v)
	}
	return nil
}

func Read(filenames ...string) (map[string]string, error) {
	return ReadWithLookup(nil, filenames...)
}

func ReadWithLookup(_ dao.LookupFn, filenames ...string) (map[string]string, error) {
	filenames = filenamesOrDefault(filenames)
	retMap := make(map[string]string)
	for _, filename := range filenames {
		m, err := lookupFile(filename)
		if err != nil {
			return nil, err
		}
		for k, v := range m {
			retMap[k] = v
		}
	}
	return retMap, nil
}

func lookupFile(filename string) (map[string]string, error) {
	defaultDAO := dao.NewDefaultDaoFromEnv(os.Environ())
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	m := envlang.GetVariablesFromInputStream(defaultDAO, f)
	retMap := make(map[string]string)
	for k, v := range m {
		if v == nil {
			v = new(string)
		}
		retMap[k] = *v
	}
	return retMap, nil
}

func ReadFile(filename string, lookupFn dao.LookupFn) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ParseWithLookup(file, lookupFn)
}

// Marshal outputs the given environment as a dotenv-formatted environment file.
// Each line is in the format: KEY="value" where value is escaped.
func Marshal(envMap map[string]string) (string, error) {
	lines := make([]string, 0, len(envMap))
	for k, v := range envMap {
		lines = append(lines, fmt.Sprintf("%s=%s", k, doubleQuoteEscape(v)))
	}
	return fmt.Sprintf("%s\n", string(bytes.Join(convertStringsToBytes(lines), []byte("\n")))), nil
}

func convertStringsToBytes(strs []string) [][]byte {
	result := make([][]byte, len(strs))
	for i, s := range strs {
		result[i] = []byte(s)
	}
	return result
}

func doubleQuoteEscape(s string) string {
	s = fmt.Sprintf(`"%s"`, s)
	s = fmt.Sprintf("%q", s[1:len(s)-1])
	return s
}

// Write serializes the given environment and writes it to a file.
func Write(envMap map[string]string, filename string) error {
	content, err := Marshal(envMap)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, []byte(content), defaultFileMode)
}

// Exec loads env vars from .env files and then executes the given cmd with those environment variables.
func Exec(filenames []string, cmd string, cmdArgs []string, overload bool) error {
	if overload {
		if err := Overload(filenames...); err != nil {
			return err
		}
	} else {
		if err := Load(filenames...); err != nil {
			return err
		}
	}

	// Execute the command with the current environment
	command := exec.Command(cmd, cmdArgs...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}
