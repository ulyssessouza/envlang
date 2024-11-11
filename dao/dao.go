package dao

type LookupFn func(string) (string, bool)

type EnvLangDao interface {
	ImportList([]string)
	ImportMap(map[string]string)
	Get(string) (*string, bool)
	Put(string, *string)
	Remove(string)
	ExportMap() map[string]*string
}
