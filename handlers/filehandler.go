package handlers

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ulyssessouza/envlang/gen/fileparser"
	"github.com/ulyssessouza/envlang/store"
)

var _ fileparser.EnvLangFileListener = &EnvLangFileListener{}

type EnvLangFileListener struct {
	*fileparser.BaseEnvLangFileListener

	d store.Store
}

func NewEnvLangFileListener(d store.Store) *EnvLangFileListener {
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

	// Handle export prefix (grammar can't handle it due to lexer limitations)
	id, _ = strings.CutPrefix(id, "export ")
	id = strings.TrimSpace(id)
	if strings.HasPrefix(id, "#") {
		return
	}

	// Validate identifier format
	re := regexp.MustCompile(`^[0-9a-zA-Z_\-.]+$`)
	if !re.MatchString(id) {
		return
	}

	hasAssign := true
	if c.ASSIGN() == nil || c.ASSIGN().GetText() == "" {
		hasAssign = false
		gotFromStore, ok := l.d.Get(id)
		if ok && gotFromStore != nil {
			l.d.Put(id, gotFromStore)
			return
		}
	}
	if hasAssign && c.Value() == nil {
		v := ""
		valuePtr = &v
	}

	if c.Value() != nil {
		v := strings.TrimSpace(c.Value().GetText())

		// Check which type of value we have
		switch {
		case c.Value().SQSTRING() != nil:
			// Single quoted - no variable expansion
			v = v[1 : len(v)-1] // Remove quotes
			valuePtr = &v
		case c.Value().DQSTRING() != nil:
			// Double quoted - expand variables
			v = v[1 : len(v)-1] // Remove quotes
			v = GetValue(l.d, v)
			valuePtr = &v
		case c.Value().IDENTIFIER() != nil || c.Value().UNQUOTED_VALUE() != nil:
			// Identifier or unquoted value - expand variables
			// Grammar handles comments, so no need to check for #
			// Trim trailing quotes that may appear in malformed input
			v = strings.TrimRight(v, `"'`)
			v = GetValue(l.d, v)
			valuePtr = &v
		default:
			panic(fmt.Sprintf("unexpected value type: %s", v))
		}
	}

	l.d.Put(id, valuePtr)
}
