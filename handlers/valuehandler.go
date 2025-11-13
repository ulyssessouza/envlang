package handlers

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ulyssessouza/envlang/gen/valueparser"
	"github.com/ulyssessouza/envlang/logger"
)

func (l *envLangValueListener) ExitDqstring(c *valueparser.DqstringContext) {
	fullText := c.GetText()
	logger.Debugf("ExitDqstring: %s", fullText)
	if len(c.GetChildren()) == 0 {
		logger.Debugf("ExitDqstring in if: %s", c.GetText())
		l.append(c.GetText())
		return
	}
}

func (l *envLangValueListener) ExitContent(c *valueparser.ContentContext) {
	fullText := c.GetText()
	logger.Debugf("ExitContent: %s", fullText)

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
	logger.Debugf("ExitVariable: %s", fullText)

	// Check which type of variable token we have
	switch {
	case c.SIMPLE_VAR() != nil:
		l.handleSimpleVar(fullText)
	case c.SIMPLE_STRICT_VAR() != nil:
		l.handleSimpleStrictVar(fullText)
	case c.STRICT_VAR_WITH_ASSIGN_IF_UNSET_OR_EMPTY() != nil:
		l.handleAssignIfUnsetOrEmpty(fullText)
	case c.STRICT_VAR_WITH_ASSIGN_IF_UNSET() != nil:
		l.handleAssignIfUnset(fullText)
	case c.STRICT_VAR_WITH_ALTERNATE_IF_SET_AND_NOT_EMPTY() != nil:
		l.handleAlternateIfSetAndNotEmpty(fullText)
	case c.STRICT_VAR_WITH_ALTERNATE_IF_SET() != nil:
		l.handleAlternateIfSet(fullText)
	case c.STRICT_VAR_WITH_ERROR_IF_UNSET_OR_EMPTY() != nil:
		l.handleErrorIfUnsetOrEmpty(fullText)
	case c.STRICT_VAR_WITH_ERROR_IF_UNSET() != nil:
		l.handleErrorIfUnset(fullText)
	case c.STRICT_VAR_LENGTH() != nil:
		l.handleVarLength(fullText)
	case c.STRICT_VAR_REMOVE_LONGEST_PREFIX() != nil:
		l.handleRemoveLongestPrefix(fullText)
	case c.STRICT_VAR_REMOVE_SHORTEST_PREFIX() != nil:
		l.handleRemoveShortestPrefix(fullText)
	case c.STRICT_VAR_REMOVE_LONGEST_SUFFIX() != nil:
		l.handleRemoveLongestSuffix(fullText)
	case c.STRICT_VAR_REMOVE_SHORTEST_SUFFIX() != nil:
		l.handleRemoveShortestSuffix(fullText)
	case c.STRICT_VAR_WITH_DEFAULT_IF_UNSET_OR_EMPTY() != nil:
		l.handleDefaultIfUnsetOrEmpty(fullText)
	case c.STRICT_VAR_WITH_DEFAULT_IF_UNSET() != nil:
		l.handleDefaultIfUnset(fullText)
	case c.DOLLAR() != nil:
		l.append(fullText)
	default:
		logger.Debugf("unexpected variable token: %s", fullText)
	}
}

func (l *envLangValueListener) handleSimpleVar(fullText string) {
	vName := fullText[1:] // Remove $
	value, ok := l.d.Get(vName)
	if ok && value != nil {
		l.append(*value)
	}
}

func (l *envLangValueListener) handleSimpleStrictVar(fullText string) {
	vName := strings.TrimSpace(fullText[2 : len(fullText)-1]) // Remove ${ and }
	value, ok := l.d.Get(vName)
	if ok && value != nil {
		l.append(*value)
	}
}

func (l *envLangValueListener) handleAssignIfUnsetOrEmpty(fullText string) {
	vName, defaultValue := l.getNameAndDefault(fullText, ":=")
	value, ok := l.d.Get(vName)
	if !ok || value == nil || *value == "" {
		l.d.Put(vName, &defaultValue)
		l.append(defaultValue)
		return
	}
	l.append(*value)
}

func (l *envLangValueListener) handleAssignIfUnset(fullText string) {
	vName, defaultValue := l.getNameAndDefault(fullText, "=")
	value, ok := l.d.Get(vName)
	if !ok || value == nil {
		l.d.Put(vName, &defaultValue)
		l.append(defaultValue)
		return
	}
	l.append(*value)
}

func (l *envLangValueListener) handleAlternateIfSetAndNotEmpty(fullText string) {
	vName, alternateValue := l.getNameAndDefault(fullText, ":+")
	value, ok := l.d.Get(vName)
	if ok && value != nil && *value != "" {
		l.append(alternateValue)
	}
}

func (l *envLangValueListener) handleAlternateIfSet(fullText string) {
	vName, alternateValue := l.getNameAndDefault(fullText, "+")
	value, ok := l.d.Get(vName)
	if ok && value != nil {
		l.append(alternateValue)
	}
}

func (l *envLangValueListener) handleErrorIfUnsetOrEmpty(fullText string) {
	vName, errorMessage := l.getNameAndDefault(fullText, ":?")
	value, ok := l.d.Get(vName)
	if !ok || value == nil || *value == "" {
		panic(&ParameterExpansionError{
			VarName: vName,
			Message: errorMessage,
		})
	}
	l.append(*value)
}

func (l *envLangValueListener) handleErrorIfUnset(fullText string) {
	vName, errorMessage := l.getNameAndDefault(fullText, "?")
	value, ok := l.d.Get(vName)
	if !ok || value == nil {
		panic(&ParameterExpansionError{
			VarName: vName,
			Message: errorMessage,
		})
	}
	l.append(*value)
}

func (l *envLangValueListener) handleVarLength(fullText string) {
	vName := l.getVarNameFromLength(fullText)
	value, ok := l.d.Get(vName)
	length := 0
	if ok && value != nil {
		length = len(*value)
	}
	l.append(fmt.Sprintf("%d", length))
}

func (l *envLangValueListener) handleRemoveLongestPrefix(fullText string) {
	vName, pattern := l.getNameAndDefault(fullText, "##")
	value, ok := l.d.Get(vName)
	if ok && value != nil {
		l.append(removeLongestPrefixMatch(*value, pattern))
	}
}

func (l *envLangValueListener) handleRemoveShortestPrefix(fullText string) {
	vName, pattern := l.getNameAndDefault(fullText, "#")
	value, ok := l.d.Get(vName)
	if ok && value != nil {
		l.append(removeShortestPrefixMatch(*value, pattern))
	}
}

func (l *envLangValueListener) handleRemoveLongestSuffix(fullText string) {
	vName, pattern := l.getNameAndDefault(fullText, "%%")
	value, ok := l.d.Get(vName)
	if ok && value != nil {
		l.append(removeLongestSuffixMatch(*value, pattern))
	}
}

func (l *envLangValueListener) handleRemoveShortestSuffix(fullText string) {
	vName, pattern := l.getNameAndDefault(fullText, "%")
	value, ok := l.d.Get(vName)
	if ok && value != nil {
		l.append(removeShortestSuffixMatch(*value, pattern))
	}
}

func (l *envLangValueListener) handleDefaultIfUnsetOrEmpty(fullText string) {
	vName, defaultValue := l.getNameAndDefault(fullText, ":-")
	value, ok := l.d.Get(vName)
	if !ok || value == nil || *value == "" {
		l.append(defaultValue)
		return
	}
	l.append(*value)
}

func (l *envLangValueListener) handleDefaultIfUnset(fullText string) {
	vName, defaultValue := l.getNameAndDefault(fullText, "-")
	value, ok := l.d.Get(vName)
	if !ok || value == nil {
		l.append(defaultValue)
		return
	}
	l.append(*value)
}

func (l *envLangValueListener) getNameAndDefault(text string, splitter string) (string, string) {
	logger.Debugf("Name with Default: %s", text)

	// Remove ${ and }
	vName := strings.TrimSpace(text[2 : len(text)-1])
	parts := strings.SplitN(vName, splitter, pair)
	if len(parts) < pair {
		return strings.TrimSpace(parts[0]), ""
	}
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
}

func (l *envLangValueListener) getVarNameFromLength(text string) string {
	logger.Debugf("Variable Length: %s", text)

	// Remove ${ and }
	content := strings.TrimSpace(text[2 : len(text)-1])
	// Remove the # prefix
	if strings.HasPrefix(content, "#") {
		content = strings.TrimSpace(content[1:])
	}
	return content
}

// removeShortestPrefixMatch removes the shortest match of pattern from the beginning of value.
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

// removeLongestPrefixMatch removes the longest match of pattern from the beginning of value.
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

// removeShortestSuffixMatch removes the shortest match of pattern from the end of value.
func removeShortestSuffixMatch(value, pattern string) string {
	if pattern == "" {
		return value
	}

	// Simple literal match (no wildcards)
	if !strings.Contains(pattern, "*") && !strings.Contains(pattern, "?") {
		if strings.HasSuffix(value, pattern) {
			return value[:len(value)-len(pattern)]
		}
		return value
	}

	// Pattern with wildcards - find shortest match from the end
	// Try matching from shortest to longest suffix
	for i := len(value); i >= 0; i-- {
		suffix := value[i:]
		if matchPattern(suffix, pattern) {
			return value[:i]
		}
	}

	return value
}

// removeLongestSuffixMatch removes the longest match of pattern from the end of value.
func removeLongestSuffixMatch(value, pattern string) string {
	if pattern == "" {
		return value
	}

	// Simple literal match (no wildcards)
	if !strings.Contains(pattern, "*") && !strings.Contains(pattern, "?") {
		if strings.HasSuffix(value, pattern) {
			return value[:len(value)-len(pattern)]
		}
		return value
	}

	// Pattern with wildcards - find longest match from the end
	// Try matching from longest to shortest suffix
	for i := 0; i <= len(value); i++ {
		suffix := value[i:]
		if matchPattern(suffix, pattern) {
			return value[:i]
		}
	}

	return value
}

// matchPattern performs simple glob-style pattern matching.
// Supports * (matches any sequence) and ? (matches single character).
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
