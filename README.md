# Envlang

## Description

Envlang is a Go library that tries to treat environment variables as a proper language. By having a proper grammar, it can parse and evaluate expressions that are based on environment variables.
It's meant to be configurable through a simple API that lets you define your own Data Access Objects (DAOs) and use them to access the environment variables you need so they can serve as connectors to anything you need.

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
	"github.com/ulyssessouza/envlang/dao"
)

func main() {
	file, err := os.Open("input.env")
	if err != nil {
		panic(err)
	}
	d := dao.NewDefaultDao()
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

## License

This project is licensed under the [MIT](LICENSE) License - see the LICENSE.md file for details


_________________


Made with love ❤️. PRs are welcome 🚀
