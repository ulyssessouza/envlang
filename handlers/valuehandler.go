package handlers

import (
	"fmt"
	"regexp"
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

	// Variables are handled by ExitVariable
	if c.Variable() != nil {
		return
	}

	// Handle escaped characters
	if c.EscapedChar() != nil {
		escaped := c.GetText()
		if len(escaped) == 2 && escaped[0] == '\\' {
			switch escaped[1] {
			case 'n':
				l.append("\n")
			case 't':
				l.append("\t")
			case 'r':
				l.append("\r")
			case '\\':
				l.append("\\")
			case '$':
				l.append("$")
			case '"':
				l.append("\"")
			default:
				// Unknown escape sequence, append as-is
				l.append(escaped)
			}
			return
		}
	}

	// Regular text, whitespace, or newlines
	l.append(c.GetText())
}

func (l *envLangValueListener) ExitVariable(c *valueparser.VariableContext) {
	fullText := c.GetText()
	log.Debugf("ExitVariable: %s", fullText)

	// Check which type of variable token we have
	switch {
	case c.SIMPLE_VAR() != nil:
		vName := fullText[1:] // Remove $
		value, ok := l.d.Get(vName)
		if ok && value != nil {
			l.append(*value)
		}
	case c.SIMPLE_STRICT_VAR() != nil:
		vName := strings.TrimSpace(fullText[2 : len(fullText)-1]) // Remove ${ and }
		value, ok := l.d.Get(vName)
		if ok && value != nil {
			l.append(*value)
		}
	case c.STRICT_VAR_WITH_ASSIGN_IF_UNSET_OR_EMPTY() != nil:
		vName, defaultValue := l.getNameAndDefault(fullText, ":=")
		value, ok := l.d.Get(vName)
		if !ok || value == nil || *value == "" {
			l.d.Put(vName, &defaultValue)
			l.append(defaultValue)
			return
		}
		l.append(*value)
	case c.STRICT_VAR_WITH_ASSIGN_IF_UNSET() != nil:
		vName, defaultValue := l.getNameAndDefault(fullText, "=")
		value, ok := l.d.Get(vName)
		if !ok || value == nil {
			l.d.Put(vName, &defaultValue)
			l.append(defaultValue)
			return
		}
		l.append(*value)
	case c.STRICT_VAR_WITH_ALTERNATE_IF_SET_AND_NOT_EMPTY() != nil:
		vName, alternateValue := l.getNameAndDefault(fullText, ":+")
		value, ok := l.d.Get(vName)
		if ok && value != nil && *value != "" {
			l.append(alternateValue)
			return
		}
	case c.STRICT_VAR_WITH_ALTERNATE_IF_SET() != nil:
		vName, alternateValue := l.getNameAndDefault(fullText, "+")
		value, ok := l.d.Get(vName)
		if ok && value != nil {
			l.append(alternateValue)
			return
		}
	case c.STRICT_VAR_WITH_ERROR_IF_UNSET_OR_EMPTY() != nil:
		vName, errorMessage := l.getNameAndDefault(fullText, ":?")
		value, ok := l.d.Get(vName)
		if !ok || value == nil || *value == "" {
			panic(&ParameterExpansionError{
				VarName: vName,
				Message: errorMessage,
			})
		}
		l.append(*value)
	case c.STRICT_VAR_WITH_ERROR_IF_UNSET() != nil:
		vName, errorMessage := l.getNameAndDefault(fullText, "?")
		value, ok := l.d.Get(vName)
		if !ok || value == nil {
			panic(&ParameterExpansionError{
				VarName: vName,
				Message: errorMessage,
			})
		}
		l.append(*value)
	case c.STRICT_VAR_LENGTH() != nil:
		vName := l.getVarNameFromLength(fullText)
		value, ok := l.d.Get(vName)
		length := 0
		if ok && value != nil {
			length = len(*value)
		}
		l.append(fmt.Sprintf("%d", length))
	case c.STRICT_VAR_REMOVE_LONGEST_PREFIX() != nil:
		vName, pattern := l.getNameAndDefault(fullText, "##")
		value, ok := l.d.Get(vName)
		if ok && value != nil {
			l.append(removeLongestPrefixMatch(*value, pattern))
		}
	case c.STRICT_VAR_REMOVE_SHORTEST_PREFIX() != nil:
		vName, pattern := l.getNameAndDefault(fullText, "#")
		value, ok := l.d.Get(vName)
		if ok && value != nil {
			l.append(removeShortestPrefixMatch(*value, pattern))
		}
	case c.STRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY() != nil:
		vName, defaultValue := l.getNameAndDefault(fullText, ":-")
		value, ok := l.d.Get(vName)
		if !ok || value == nil || *value == "" {
			l.append(defaultValue)
			return
		}
		l.append(*value)
	case c.STRICT_VAR_WITH_DEFAULT_IF_UNSET() != nil:
		vName, defaultValue := l.getNameAndDefault(fullText, "-")
		value, ok := l.d.Get(vName)
		if !ok || value == nil {
			l.append(defaultValue)
			return
		}
		l.append(*value)
	case c.DOLLAR() != nil:
		// Lone dollar sign - append as-is (may include following char)
		l.append(fullText)
	default:
		log.Debugln("unexpected variable token: " + fullText)
	}
}

func (l *envLangValueListener) getNameAndDefault(text string, splitter string) (string, string) {
	log.Debugf("Name with Default: %s", text)

	// Remove ${ and }
	vName := strings.TrimSpace(text[2 : len(text)-1])
	parts := strings.SplitN(vName, splitter, pair)
	if len(parts) < pair {
		return strings.TrimSpace(parts[0]), ""
	}
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
}

func (l *envLangValueListener) getVarNameFromLength(text string) string {
	log.Debugf("Variable Length: %s", text)

	// Remove ${ and }
	content := strings.TrimSpace(text[2 : len(text)-1])
	// Remove the # prefix
	if strings.HasPrefix(content, "#") {
		content = strings.TrimSpace(content[1:])
	}
	return content
}

// removeShortestPrefixMatch removes the shortest match of pattern from the beginning of value
func removeShortestPrefixMatch(value, pattern string) string {
	if pattern == "" {
		return value
	}

	// Simple literal match (no wildcards)
	if !strings.Contains(pattern, "*") && !strings.Contains(pattern, "?") {
		if strings.HasPrefix(value, pattern) {
			return value[len(pattern):]
		}
		return value
	}

	// Pattern with wildcards - find shortest match
	// Try matching from shortest to longest prefix
	for i := 0; i <= len(value); i++ {
		prefix := value[:i]
		if matchPattern(prefix, pattern) {
			return value[i:]
		}
	}

	return value
}

// removeLongestPrefixMatch removes the longest match of pattern from the beginning of value
func removeLongestPrefixMatch(value, pattern string) string {
	if pattern == "" {
		return value
	}

	// Simple literal match (no wildcards)
	if !strings.Contains(pattern, "*") && !strings.Contains(pattern, "?") {
		if strings.HasPrefix(value, pattern) {
			return value[len(pattern):]
		}
		return value
	}

	// Pattern with wildcards - find longest match
	// Try matching from longest to shortest prefix
	for i := len(value); i >= 0; i-- {
		prefix := value[:i]
		if matchPattern(prefix, pattern) {
			return value[i:]
		}
	}

	return value
}

// matchPattern performs simple glob-style pattern matching
// Supports * (matches any sequence) and ? (matches single character)
func matchPattern(text, pattern string) bool {
	// Convert glob pattern to regex
	regexPattern := "^"
	for i := 0; i < len(pattern); i++ {
		switch pattern[i] {
		case '*':
			regexPattern += ".*"
		case '?':
			regexPattern += "."
		case '.', '+', '(', ')', '[', ']', '{', '}', '^', '$', '|', '\\':
			// Escape regex special characters
			regexPattern += "\\" + string(pattern[i])
		default:
			regexPattern += string(pattern[i])
		}
	}
	regexPattern += "$"

	matched, _ := regexp.MatchString(regexPattern, text)
	return matched
}
