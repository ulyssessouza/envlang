module github.com/ulyssessouza/envlang/store/redis

go 1.24.0

require (
	github.com/alicebob/miniredis/v2 v2.35.0
	github.com/redis/go-redis/v9 v9.16.0
	github.com/ulyssessouza/envlang v0.0.0
	gotest.tools/v3 v3.5.2
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/yuin/gopher-lua v1.1.1 // indirect
)

replace github.com/ulyssessouza/envlang => ../..
