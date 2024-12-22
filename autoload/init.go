package autoload

import "github.com/ulyssessouza/envlang"

func init() {
	envlang.Load() //nolint:errcheck
}
