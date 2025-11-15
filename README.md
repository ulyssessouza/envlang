# Envlang

## Description

Envlang is a Go library that tries to treat environment variables as a proper language. By having a proper grammar, it can parse and evaluate expressions that are based on environment variables.
It's meant to be configurable through a simple API that lets you define your own Store and use them to access the environment variables you need so they can serve as connectors to anything you need.

## Getting Started

### Add to your project

```
go get -u github.com/ulyssessouza/envlang
```

### Usage

Given the following environment file `input.env`:
```env
FOO=foo
BAR="bar"
BAZ = 1
EMPTY=
UNSET

VAR1=$FOO
VAR2=${BAR}
VAR3=${EMPTY:-default_value_for_empty}
VAR4=${UNSET-default_value_for_unset}
```

Example code:

```go
package main

import (
	"fmt"
	"os"

	"github.com/ulyssessouza/envlang"
	"github.com/ulyssessouza/envlang/store"
)

func main() {
	file, err := os.Open("input.env")
	if err != nil {
		panic(err)
	}
	d := store.NewDefaultStore()
	vars := envlang.GetVariablesFromInputStream(d, file)
	fmt.Printf("FOO=%q\n", *vars["FOO"])
	fmt.Printf("BAR=%q\n", *vars["BAR"])
	fmt.Printf("BAZ=%q\n", *vars["BAZ"])
	fmt.Printf("EMPTY=%q\n", *vars["EMPTY"])
	fmt.Printf("UNSET=%q\n", vars["UNSET"])
	fmt.Printf("VAR1=%q\n", *vars["VAR1"])
	fmt.Printf("VAR2=%q\n", *vars["VAR2"])
	fmt.Printf("VAR3=%q\n", *vars["VAR3"])
	fmt.Printf("VAR4=%q\n", *vars["VAR4"])
}
```

Output:
```
FOO="foo"
BAR="bar"
BAZ="1"
EMPTY=""
UNSET=%!q(*string=<nil>)
VAR1="foo"
VAR2="bar"
VAR3="default_value_for_empty"
VAR4="default_value_for_unset"
```

Please note that it used `%q` to print empty strings, as `%s` would not print anything.

## Custom Stores

Envlang is designed to be extensible through a `store.Store` interface, allowing you to connect to any data source for your environment variables. This provides the flexibility to fetch variables from databases, cloud services, or any other configuration management system.

### Implementing a New Store

To create a custom store, you need to implement the `store.Store` interface:

```go
package store

type Store interface {
	ImportList([]string)
	ImportMap(map[string]string)
	Get(string) (*string, bool)
	Put(string, *string)
	Remove(string) bool
	ExportMap() map[string]*string
}
```

By implementing these methods, you can define how Envlang interacts with your chosen backend.

### Example: Redis Store

Envlang includes a ready-to-use Redis store as an example of a custom implementation. It allows you to use a Redis instance as a backend for storing and retrieving environment variables.

You can find the implementation in the `store/redis` directory. This serves as a practical guide for building your own custom stores.

## Shell Parameter Expansion

Envlang supports a comprehensive set of POSIX shell parameter expansion operators:

### Default Values
- `${VAR:-default}` - Use default value if variable is unset or empty
- `${VAR-default}` - Use default value if variable is unset (keeps empty values)

### Assignment
- `${VAR:=default}` - Assign and return default if variable is unset or empty
- `${VAR=default}` - Assign and return default if variable is unset

### Alternate Values
- `${VAR:+alternate}` - Use alternate value if variable is set and non-empty
- `${VAR+alternate}` - Use alternate value if variable is set (even if empty)

### Error Handling
- `${VAR:?message}` - Panic with error if variable is unset or empty
- `${VAR?message}` - Panic with error if variable is unset

### String Operations
- `${#VAR}` - Return the length of the variable's value

### Pattern Matching
- `${VAR#pattern}` - Remove shortest prefix match (supports glob patterns: `*`, `?`)
- `${VAR##pattern}` - Remove longest prefix match
- `${VAR%pattern}` - Remove shortest suffix match
- `${VAR%%pattern}` - Remove longest suffix match

### Example

```go
d := store.NewDefaultStoreFromMap(map[string]*string{
    "PATH":     strPtr("/usr/local/bin"),
    "FILENAME": strPtr("archive.tar.gz"),
    "EMPTY":    strPtr(""),
})

// Default values
vars := envlang.GetVariables(d, `DEFAULT="${UNSET:-fallback}"`)        // "fallback"
vars = envlang.GetVariables(d, `DEFAULT="${EMPTY-fallback}"`)          // ""

// Assignment
vars = envlang.GetVariables(d, `ASSIGNED="${NEW:=value}"`)             // "value", NEW is now set

// String length
vars = envlang.GetVariables(d, `LENGTH="${#PATH}"`)                    // "15"

// Pattern removal
vars = envlang.GetVariables(d, `BASENAME="${FILENAME##*/}"`)           // "archive.tar.gz"
vars = envlang.GetVariables(d, `NOEXT="${FILENAME%.*}"`)               // "archive.tar"
vars = envlang.GetVariables(d, `NAME="${FILENAME%%.*}"`)               // "archive"
```

## godotenv Compatibility Layer

Envlang includes a **100% API-compatible** drop-in replacement for [github.com/joho/godotenv](https://github.com/joho/godotenv), powered by our improved ANTLR4 grammar-based parser.

### Quick Start: Drop-in Replacement

Add this to your `go.mod`:

```go
replace github.com/joho/godotenv => github.com/ulyssessouza/envlang/compat/joho/godotenv v0.1.3
```

**That's it!** Your existing code using `github.com/joho/godotenv` will now use envlang's parser with **zero code changes**.

### Why Use the Compatibility Layer?

- ✅ **Better parsing**: ANTLR4 grammar instead of regex
- ✅ **Clearer errors**: Grammar-based error messages
- ✅ **More features**: Escape sequences (`\n`, `\t`, `\$`, etc.)
- ✅ **Same API**: All godotenv functions work identically

### Example

```go
import "github.com/joho/godotenv"  // Works with replace directive!

func main() {
    godotenv.Load()
    // Use your environment variables as usual
}
```

### Supported Functions

All godotenv functions are supported:
- `Load()`, `Overload()`, `Read()`, `Parse()`
- `Unmarshal()`, `UnmarshalBytes()`
- `Marshal()`, `Write()`
- `Exec()`
- Plus bonus `*WithLookup()` variants for advanced use cases

📖 **[Full documentation →](compat/joho/godotenv/README.md)**

## License

This project is licensed under the [MIT](LICENSE) License


_________________


Made with love ❤️. PRs are welcome 🚀
