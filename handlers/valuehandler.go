package handlers

import (
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ulyssessouza/envlang/parsers/valueparser"
)

var _ valueparser.EnvLangValueListener = &envLangValueListener{}

type envLangValueListener struct {
	valueparser.BaseEnvLangValueListener

	env    func(k string) (string, bool)
	result string
}

func (l *envLangValueListener) ExitVariable(c *valueparser.VariableContext) {
	log.Printf("ExitVariable: %#v\n", c.GetText())
}

func (l *envLangValueListener) ExitStrictVar(c *valueparser.StrictVarContext) {
	log.Printf("ExitStrictVar: %#v\n", c.GetText())

	varName := ""
	log.Printf("ExitStrictVar: %#v\n", c.GetText())

	log.Printf("\tVarName: '")
	for _, t := range c.AllTEXT_NO_SPACE() {
		log.Printf("%s", t.GetText())
		varName += t.GetText()
	}
	log.Println("'")

	varVal, ok := l.env(varName)
	if !ok {
		varVal = ""
	}

	l.result += varVal
}

func (l *envLangValueListener) ExitSimpleVar(c *valueparser.SimpleVarContext) {
	varName := ""
	log.Printf("ExitSimpleVar: %#v\n", c.GetText())

	log.Printf("\tVarName: '")
	for _, t := range c.AllTEXT_NO_SPACE() {
		log.Printf("%s", t.GetText())
		varName += t.GetText()
	}
	log.Println("'")

	varVal, ok := l.env(varName)
	if !ok {
		varVal = ""
	}

	l.result += varVal
}

func (l *envLangValueListener) ExitText(c *valueparser.TextContext) {
	log.Printf("ExitText: %#v\n", c.GetText())
	l.result += c.GetText()
}

func (l *envLangValueListener) ExitSpace(c *valueparser.SpaceContext) {
	log.Printf("ExitSpace: %#v\n", c.GetText())
	l.result += c.GetText()
}

func (l *envLangValueListener) ExitContent(c *valueparser.ContentContext) {
	// log.Printf("ExitContent: %#v\n", c.GetText())
}

var listener = &envLangValueListener{
	env: func(k string) (string, bool) {
		m := map[string]string{
			"MYVALVAR":  "foo",
			"SECONDVAR": "bar",
		}
		r, ok := m[k]
		return r, ok
	},
}

func getValue(s string) string {
	is := antlr.NewInputStream(s)
	lexer := valueparser.NewEnvLangValueLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := valueparser.NewEnvLangValueParser(stream)
	parser.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Dqstring())
	result := listener.result
	listener.result = ""

	return result
}
