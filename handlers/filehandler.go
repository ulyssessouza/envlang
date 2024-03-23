package handlers

import (
	"fmt"
	"regexp"
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
	var valuePtr *string

	id := strings.TrimSpace(c.Identifier().GetText())
	if id == "" {
		return
	}

	// TODO: Implement this logic in the grammar file
	id, _ = strings.CutPrefix(id, "export ")
	id = strings.TrimSpace(id)
	if strings.HasPrefix(id, "#") {
		return
	}
	re := regexp.MustCompile(`^[0-9a-zA-Z_\-.]+$`)
	if !re.MatchString(id) {
		return
	}
	// TODO: END

	hasAssign := true
	if c.ASSIGN() == nil || c.ASSIGN().GetText() == "" {
		hasAssign = false
		gotFromDAO, ok := l.d.Get(id)
		if ok && gotFromDAO != nil {
			l.d.Put(id, gotFromDAO)
			return
		}
	}
	if hasAssign && c.Value() == nil {
		v := ""
		valuePtr = &v
	}

	toTrim := true
	if c.Value() != nil {
		v := strings.TrimSpace(c.Value().GetText())
		switch c.Value().GetStr().GetTokenType() {
		case fileparser.EnvLangFileParserSQSTRING: // Not evaluating variables inside single quoted values
			v = v[1 : len(v)-1] // Removing quotes
			valuePtr = &v
		case fileparser.EnvLangFileParserDQSTRING:
			v = v[1 : len(v)-1] // Removing quotes
			toTrim = false
			fallthrough
		case fileparser.EnvLangFileParserTEXT:
			if strings.Index(v, "#") != 0 {
				v = strings.SplitN(v, "#", pair)[0]
				if toTrim {
					v = strings.TrimSpace(v)
				}
			}
			v = GetValue(l.d, v)
			valuePtr = &v
		default:
			panic(fmt.Sprintf("unexpected string: %s", c.Value().GetStr().GetText()))
		}
	}

	l.d.Put(id, valuePtr)
}
