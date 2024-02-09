package main

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/ulyssessouza/envlang/dao"
	"github.com/ulyssessouza/envlang/gen/fileparser"
	"github.com/ulyssessouza/envlang/handlers"
)

func GetVariables(d dao.EnvLangDao, s string) map[string]*string {
	is := antlr.NewInputStream(s)
	lexer := fileparser.NewEnvLangFileLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := fileparser.NewEnvLangFileParser(stream)
	listener := handlers.NewEnvLangFileListener(d)
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.EnvFile())
	return listener.GetVariables()
}
