package envlang

import (
	"io"
	"os"
	"strings"

	antlr "github.com/antlr4-go/antlr/v4"

	"github.com/ulyssessouza/envlang/gen/fileparser"
	"github.com/ulyssessouza/envlang/handlers"
	"github.com/ulyssessouza/envlang/store"
)

const defaultEnvFile = ".env"

func GetVariablesFromInputStream(d store.Store, r io.Reader) map[string]*string {
	return getVariablesFromInputStream(d, antlr.NewIoStream(r))
}

func GetVariables(d store.Store, s string) map[string]*string {
	return getVariablesFromInputStream(d, antlr.NewInputStream(s))
}

func Load(paths ...string) error {
	if len(paths) == 0 {
		paths = []string{defaultEnvFile}
	}
	for _, p := range paths {
		if strings.TrimSpace(p) == "" {
			continue
		}
		f, err := os.Open(p)
		if err != nil {
			return err
		}
		m := GetVariablesFromInputStream(store.NewDefaultStoreFromEnv(os.Environ()), f)
		for k, v := range m {
			if v == nil {
				str := ""
				v = &str
			}
			os.Setenv(k, *v)
		}
	}

	return nil
}

func getVariablesFromInputStream(d store.Store, ais *antlr.InputStream) map[string]*string {
	lexer := fileparser.NewEnvLangFileLexer(ais)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := fileparser.NewEnvLangFileParser(stream)
	listener := handlers.NewEnvLangFileListener(d)
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.EnvFile())
	return listener.GetVariables()
}
