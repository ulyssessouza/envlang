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

	vName := ""
	var_ := c.GetVar_()
	switch var_.GetTokenType() {
	case valueparser.EnvLangValueParserSIMPLE_VAR:
		vName = var_.GetText()[1:]
	case valueparser.EnvLangValueParserSTRICT_VAR:
		vName = strings.TrimSpace(var_.GetText()[2 : len(var_.GetText())-1])
	default:
		panic("unexpected token: " + var_.GetText())
	}

	value, ok := l.d.Get(vName)
	if ok && value != nil {
		l.append(*value)
	}
}
