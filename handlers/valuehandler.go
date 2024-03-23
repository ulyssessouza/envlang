package handlers

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"
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

	varToken := c.GetVar_()
	switch varToken.GetTokenType() {
	case valueparser.EnvLangValueParserSIMPLE_VAR:
		vName := varToken.GetText()[1:]
		value, ok := l.d.Get(vName)
		if ok && value != nil {
			l.append(*value)
		}
	case valueparser.EnvLangValueParserSIMPLE_STRICT_VAR:
		vName := strings.TrimSpace(varToken.GetText()[2 : len(varToken.GetText())-1])
		value, ok := l.d.Get(vName)
		if ok && value != nil {
			l.append(*value)
		}
	case valueparser.EnvLangValueParserSTRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY:
		vName, defaultValue := l.getNameAndDefault(varToken)
		value, ok := l.d.Get(vName)
		if !ok || value == nil || *value == "" {
			l.append(defaultValue)
			return
		}
		l.append(*value)
	case valueparser.EnvLangValueParserSTRICT_VAR_WITH_DEFAULT_IF_UNSET:
		vName, defaultValue := l.getNameAndDefault(varToken)
		value, ok := l.d.Get(vName)
		if !ok {
			l.append(defaultValue)
			return
		}
		l.append(*value)
	default:
		log.Debugln("unexpected token: " + varToken.GetText())
	}
}

func (l *envLangValueListener) getNameAndDefault(token antlr.Token) (string, string) {
	log.Debugf("Name with Default: %s", token.GetText())
	splitter := ""
	switch token.GetTokenType() {
	case valueparser.EnvLangValueParserSTRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY:
		splitter = ":-"
	case valueparser.EnvLangValueParserSTRICT_VAR_WITH_DEFAULT_IF_UNSET:
		splitter = "-"
	default:
		log.Debugln("unexpected token: " + token.GetText())
	}

	vName := strings.TrimSpace(token.GetText()[2 : len(token.GetText())-1])
	parts := strings.SplitN(vName, splitter, pair)
	if len(parts) < pair {
		return parts[0], ""
	}
	return parts[0], parts[1]
}
