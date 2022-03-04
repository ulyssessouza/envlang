package handlers

import (
	"fmt"
	"strings"

	"github.com/ulyssessouza/envlang/dao"
	"github.com/ulyssessouza/envlang/gen/fileparser"
)

var _ fileparser.EnvLangFileListener = &envLangFileListener{}

type envLangFileListener struct {
	*fileparser.BaseEnvLangFileListener

	d dao.EnvLangDao
}

func NewEnvLangFileListener(d dao.EnvLangDao) *envLangFileListener {
	return &envLangFileListener{
		d: d,
	}
}

func (l envLangFileListener) GetVariables() map[string]*string {
	return l.d.ExportMap()
}

func (l *envLangFileListener) ExitEntry(c *fileparser.EntryContext) {
	var valuePtr *string = nil

	id := strings.TrimSpace(c.Identifier().GetText())
	hasAssign := true
	if c.Assign() == nil || c.Assign().IsEmpty() {
		hasAssign = false
	}
	if hasAssign && c.Value() == nil {
		v := ""
		valuePtr = &v
	}

	if c.Value() != nil {
		v := c.Value().GetText()
		switch c.Value().GetOp().GetTokenType() {
		case fileparser.EnvLangFileParserSQSTRING: // Not evaluating variables inside single quoted values
			valuePtr = &v
		case fileparser.EnvLangFileParserTEXT:
			v := GetValue(l.d, v)
			valuePtr = &v
		case fileparser.EnvLangFileParserDQSTRING:
			v := GetValue(l.d, v[1:len(v)-1])
			valuePtr = &v
		default:
			panic(fmt.Sprintf("unexpected op: %s", c.Value().GetOp().GetText()))
		}
	}

	l.d.Put(id, valuePtr)
}
