package main

import (
	"io"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ulyssessouza/envlang/dao"
	"github.com/ulyssessouza/envlang/gen/fileparser"
	"github.com/ulyssessouza/envlang/handlers"
)

func GetVariablesFromInputStream(d dao.EnvLangDao, r io.Reader) map[string]*string {
	return getVariablesFromInputStream(d, antlr.NewIoStream(r))
}

func GetVariables(d dao.EnvLangDao, s string) map[string]*string {
	return getVariablesFromInputStream(d, antlr.NewInputStream(s))
}

func getVariablesFromInputStream(d dao.EnvLangDao, ais *antlr.InputStream) map[string]*string {
	lexer := fileparser.NewEnvLangFileLexer(ais)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := fileparser.NewEnvLangFileParser(stream)
	listener := handlers.NewEnvLangFileListener(d)
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.EnvFile())
	return listener.GetVariables()
}
