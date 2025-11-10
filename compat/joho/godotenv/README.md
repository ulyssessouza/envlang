# godotenv Compatibility Layer

This package provides 100% API compatibility with [github.com/joho/godotenv](https://github.com/joho/godotenv), allowing you to use our improved ANTLR4 grammar-based parser as a drop-in replacement.

## Why Use This?

- **Better Grammar**: Uses a proper ANTLR4 grammar for parsing instead of regex
- **Improved Error Messages**: Grammar-based parsing provides clearer error diagnostics
- **Same API**: 100% compatible with the original godotenv
- **Better Testing**: More comprehensive test coverage
- **Modern Features**: Support for escape sequences (`\n`, `\t`, etc.)

## Migration Options

### Option 1: Drop-in Replacement (No Code Changes)

Simply add a replace directive to your `go.mod`:

```go
replace github.com/joho/godotenv => github.com/ulyssessouza/envlang/compat/joho/godotenv v1.0.0
```

Then run:
```bash
go mod tidy
```

**That's it!** Your existing code continues to work without any changes:

```go
import "github.com/joho/godotenv"  // This now uses envlang!

func main() {
    godotenv.Load()  // Works exactly the same
}
```

### Option 2: Explicit Import (Recommended for New Projects)

Update your imports to be explicit:

```go
import "github.com/ulyssessouza/envlang/compat/joho/godotenv"
```

This makes it clear you're using the envlang implementation.

## Complete API Support

All functions from the original godotenv are implemented:

### Core Functions
- `Load(filenames ...string) error` - Load env files without overriding existing vars
- `Overload(filenames ...string) error` - Load env files and override existing vars
- `Read(filenames ...string) (map[string]string, error)` - Read into map without loading
- `Parse(r io.Reader) (map[string]string, error)` - Parse from io.Reader
- `Unmarshal(str string) (map[string]string, error)` - Parse from string
- `UnmarshalBytes(src []byte) (map[string]string, error)` - Parse from bytes

### Write Functions
- `Marshal(envMap map[string]string) (string, error)` - Serialize to dotenv format
- `Write(envMap map[string]string, filename string) error` - Write to file

### Advanced Functions
- `Exec(filenames []string, cmd string, cmdArgs []string, overload bool) error` - Load and execute command

### Bonus Features (Not in Original)
- `ParseWithLookup(r io.Reader, lookupFn dao.LookupFn) (map[string]string, error)`
- `UnmarshalWithLookup(src string, lookupFn dao.LookupFn) (map[string]string, error)`
- `UnmarshalBytesWithLookup(src []byte, lookupFn dao.LookupFn) (map[string]string, error)`
- `ReadWithLookup(_ dao.LookupFn, filenames ...string) (map[string]string, error)`
- `ReadFile(filename string, lookupFn dao.LookupFn) (map[string]string, error)`

The `*WithLookup` variants allow custom variable resolution functions for advanced use cases.

## Examples

### Basic Usage

```go
package main

import (
    "log"
    "os"
    "github.com/joho/godotenv"  // Works with replace directive!
)

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Use environment variables
    dbHost := os.Getenv("DB_HOST")
    log.Printf("DB Host: %s", dbHost)
}
```

### Load Multiple Files

```go
godotenv.Load(".env", ".env.local", ".env.production")
```

### Overload Existing Variables

```go
// Override environment variables that already exist
godotenv.Overload(".env.override")
```

### Read Without Loading

```go
envMap, err := godotenv.Read(".env")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("DATABASE_URL: %s\n", envMap["DATABASE_URL"])
```

### Parse from String

```go
envStr := `
KEY1=value1
KEY2="value with spaces"
KEY3='single quoted'
`

envMap, err := godotenv.Unmarshal(envStr)
```

### Write to File

```go
envMap := map[string]string{
    "API_KEY": "secret123",
    "DEBUG": "true",
}

err := godotenv.Write(envMap, ".env.production")
```

### Execute Command with Env

```go
// Load env and execute command
err := godotenv.Exec(
    []string{".env"},
    "node",
    []string{"server.js"},
    false, // don't overload
)
```

## Grammar Improvements

Our ANTLR4 grammar provides better handling of:

- **Escape sequences**: `\n`, `\t`, `\r`, `\\`, `\$`, `\"`
- **Comments**: Proper `#` comment handling in grammar
- **Multiline values**: Quoted strings can span multiple lines
- **Variable substitution**: `$VAR`, `${VAR}`, `${VAR:-default}`, `${VAR-default}`
- **Edge cases**: Better handling of spaces, quotes, and special characters

## Testing

The compatibility layer has comprehensive test coverage:

```bash
cd compat/joho/godotenv
go test -v
```

Coverage: **90.1%** with 96 passing tests

## Autoload Support

Just like the original, you can auto-load `.env` on import:

```go
import _ "github.com/ulyssessouza/envlang/compat/joho/godotenv/autoload"
```

## License

This compatibility layer maintains compatibility with the original godotenv license while being part of the envlang project.

## Contributing

Contributions are welcome! Please see the main [envlang repository](https://github.com/ulyssessouza/envlang) for contribution guidelines.
