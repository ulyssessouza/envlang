package store

type LookupFn func(string) (string, bool)

type Store interface {
	ImportList([]string)
	ImportMap(map[string]string)
	Get(string) (*string, bool)
	Put(string, *string)
	Remove(string) bool
	ExportMap() map[string]*string
}
