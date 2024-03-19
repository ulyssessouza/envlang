package handlers

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/ulyssessouza/envlang/gen/valueparser"
)

func (l *envLangValueListener) ExitDqstring(c *valueparser.DqstringContext) {
	fullText := c.GetText()
	log.Debugf("ExitDqstring: %s", fullText)
	if len(c.GetChildren()) == 0 {
		log.Debugf("ExitDqstring in if: %s", c.GetText())
		l.append(c.GetText())
		return
	}
}

func (l *envLangValueListener) ExitContent(c *valueparser.ContentContext) {
	fullText := c.GetText()
	log.Debugf("ExitContent: %s", fullText)
	if c.Variable() == nil {
		l.append(c.GetText())
	}
}

func (l *envLangValueListener) ExitVariable(c *valueparser.VariableContext) {
	fullText := c.GetText()
	log.Debugf("ExitVariable: %s", fullText)

	var vName string
	varr := c.GetVar_()
	switch varr.GetTokenType() {
	case valueparser.EnvLangValueParserSIMPLE_VAR:
		vName = varr.GetText()[1:]
	case valueparser.EnvLangValueParserSTRICT_VAR:
		vName = strings.TrimSpace(varr.GetText()[2 : len(varr.GetText())-1])
	default:
		panic("unexpected token: " + varr.GetText())
	}

	value, ok := l.d.Get(vName)
	if ok && value != nil {
		l.append(*value)
	}
}
