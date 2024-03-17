package handlers

import (
	"fmt"
	"strings"

	"github.com/ulyssessouza/envlang/dao"
	"github.com/ulyssessouza/envlang/gen/fileparser"
)

var _ fileparser.EnvLangFileListener = &EnvLangFileListener{}

type EnvLangFileListener struct {
	*fileparser.BaseEnvLangFileListener

	d dao.EnvLangDao
}

func NewEnvLangFileListener(d dao.EnvLangDao) *EnvLangFileListener {
	return &EnvLangFileListener{
		d: d,
	}
}

func (l *EnvLangFileListener) GetVariables() map[string]*string {
	return l.d.ExportMap()
}

func (l *EnvLangFileListener) ExitEntry(c *fileparser.EntryContext) {
	var valuePtr *string = nil

	id := strings.TrimSpace(c.Identifier().GetText())
	if id == "" {
		return
	}
	hasAssign := true
	if c.ASSIGN() == nil || c.ASSIGN().GetText() == "" {
		hasAssign = false
	}
	if hasAssign && c.Value() == nil {
		v := ""
		valuePtr = &v
	}

	if c.Value() != nil {
		v := strings.TrimSpace(c.Value().GetText())
		switch c.Value().GetStr().GetTokenType() {
		case fileparser.EnvLangFileParserSQSTRING: // Not evaluating variables inside single quoted values
			v = v[1 : len(v)-1] // Removing quotes
			valuePtr = &v
		case fileparser.EnvLangFileParserDQSTRING:
			v = v[1 : len(v)-1] // Removing quotes
			fallthrough
		case fileparser.EnvLangFileParserTEXT:
			v := GetValue(l.d, v)
			valuePtr = &v
		default:
			panic(fmt.Sprintf("unexpected string: %s", c.Value().GetStr().GetText()))
		}
	}

	l.d.Put(id, valuePtr)
}
