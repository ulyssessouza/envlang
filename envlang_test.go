package envlang

import (
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"gotest.tools/v3/assert"

	"github.com/ulyssessouza/envlang/dao"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func TestGetFromReader(t *testing.T) {
	expected := map[string]*string{
		"A": strPtr("aaa"),
	}
	d := dao.NewDefaultDaoFromMap(nil)
	assert.DeepEqual(t, expected, GetVariablesFromInputStream(d, strings.NewReader(`A=aaa`)))
}

//nolint:funlen
func TestGetValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]*string
		envState map[string]*string
	}{
		{
			"Simple",
			`A=aaa`,
			map[string]*string{"A": strPtr("aaa")},
			nil,
		},
		{
			"SimpleWithSpaceBeforeAssign",
			`A =aaa`,
			map[string]*string{"A": strPtr("aaa")},
			nil,
		},
		{
			"SimpleWithSpaceAfterAssign",
			`A= aaa`,
			map[string]*string{"A": strPtr("aaa")},
			nil,
		},
		{
			"SimpleWithDoubleQuotes",
			`A="aaa"`,
			map[string]*string{"A": strPtr("aaa")},
			nil,
		},
		{
			"SimpleWithSingleQuotes",
			`A='aaa'`,
			map[string]*string{"A": strPtr("aaa")},
			nil,
		},
		{
			"SimpleWithDoubleQuotesAndSpaces",
			`A=aaa bbb ccc`,
			map[string]*string{"A": strPtr("aaa bbb ccc")},
			nil,
		},
		{
			"SimpleWithSingleQuotesAndSpaces",
			`A='aaa bbb ccc'`,
			map[string]*string{"A": strPtr("aaa bbb ccc")},
			nil,
		},
		{
			"SimpleWithDoubleQuotesAndSpaces",
			`A="aaa bbb ccc"`,
			map[string]*string{"A": strPtr("aaa bbb ccc")},
			nil,
		},
		{
			"MultiLine",
			`
A="my
multi
line
entry"
`,
			map[string]*string{"A": strPtr("my\nmulti\nline\nentry")},
			nil,
		},
		{
			"VariableWithEquals",
			`A=`,
			map[string]*string{"A": strPtr("")},
			nil,
		},
		{
			"VariableOnly",
			`A`,
			map[string]*string{"A": nil},
			nil,
		},
		{
			"SimpleWithSpaces",
			`
A = aaa 
`,
			map[string]*string{"A": strPtr("aaa")},
			nil,
		},
		{
			"SimpleWithSimpleVariable",
			`
A=$VAR_FROM_DAO
`,
			map[string]*string{
				"A": strPtr("aaa"),
			},
			map[string]*string{
				"VAR_FROM_DAO": strPtr("aaa"),
			},
		},
		{
			"SimpleWithSimpleVariableWithSpaces",
			`
A = $VAR_FROM_DAO
		`,
			map[string]*string{
				"A": strPtr("aaa"),
			},
			map[string]*string{
				"VAR_FROM_DAO": strPtr("aaa"),
			},
		},
		{
			"SimpleWithStrictVariable",
			`
A=${VAR_FROM_DAO}
`,
			map[string]*string{
				"A": strPtr("aaa"),
			},
			map[string]*string{
				"VAR_FROM_DAO": strPtr("aaa"),
			},
		},
		{
			"SimpleWithStrictVariableWithSpaces",
			`
A = ${VAR_FROM_DAO}
`,
			map[string]*string{
				"A": strPtr("aaa"),
			},
			map[string]*string{
				"VAR_FROM_DAO": strPtr("aaa"),
			},
		},
		{
			"SimpleWithStrictVariableWithSpacesAndInternalSpaces",
			`
A = ${ VAR_FROM_DAO }
`,
			map[string]*string{
				"A": strPtr("aaa"),
			},
			map[string]*string{
				"VAR_FROM_DAO": strPtr("aaa"),
			},
		},

		{
			"DoubleQuotedWithMixedValue",
			`
A = "aaa ${B} ccc "
`,
			map[string]*string{
				"A": strPtr("aaa bbb  ccc "),
			},
			map[string]*string{
				"B": strPtr("bbb "),
			},
		},
		{
			"SpecialWithVariable",
			`SPECIAL1 = "{{{ ${A} }}}"`,
			map[string]*string{
				"SPECIAL1": strPtr("{{{ aaa }}}"),
			},
			map[string]*string{
				"A": strPtr("aaa"),
			},
		},
		{
			"SpecialWithPesoSign",
			`SPECIAL3 = "{{{ $ }}}"`,
			map[string]*string{
				"SPECIAL3": strPtr("{{{ $ }}}"),
			},
			nil,
		},
		{
			"VariableWithDefaultForEmpty",
			`VAR_DEFAULT_UNSET_OR_EMPTY = "${EMPTY_VAR:-eee}"`,
			map[string]*string{
				"VAR_DEFAULT_UNSET_OR_EMPTY": strPtr("eee"),
			},
			map[string]*string{
				"EMPTY_VAR": strPtr(""),
			},
		},
		{
			"VariableWithDefaultForUnset",
			`VAR_DEFAULT_UNSET_OR_EMPTY = "${UNSET_VAR-uuu}"`,
			map[string]*string{
				"VAR_DEFAULT_UNSET_OR_EMPTY": strPtr("uuu"),
			},
			nil,
		},
		{
			"VariableWithDefaultForUnset",
			`VAR_DEFAULT_UNSET_OR_EMPTY = "${EMPTY_VAR-uuu}"`,
			map[string]*string{
				"VAR_DEFAULT_UNSET_OR_EMPTY": strPtr(""),
			},
			map[string]*string{
				"EMPTY_VAR": strPtr(""),
			},
		},
		{
			"PrefixExport",
			`export A = aaa"`,
			map[string]*string{
				"A": strPtr("aaa"),
			},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := dao.NewDefaultDaoFromMap(tt.envState)
			assert.DeepEqual(t, tt.expected, GetVariables(d, tt.input))
		})
	}
}

func TestFull(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	is := `
INVALID LINE

# my comment
	# my tab comment

VAR_TO_BE_LOADED_FROM_OS_ENV

A=aaa # my inline comment on A
B="bbb" # my inline comment on B
C =ccc
D= ddd
E='eee' # my inline comment on E
F=
G

H="my_value"

I = bar baz

export EXPORTED_VAR = exported_value

J = "foo bar"

L="my
multi
line
entry"

MYVAR = "before ${I} after $J opa ${UNKNOWNVAR}"

M = "foo ${A} bar"

N1=41

N2="42"

N3=43AS3sA43

N4=44AS4sA44

SIMPLE_SPACING = aaa ${B} ccc

export OPTION_B='\n'

SPECIAL1 = "{{{ ${A} }}}"
SPECIAL2 = "{{{ $A }}}"
SPECIAL3 = "{{{ $ }}}"
SPECIAL4 = "{{{ $ $ $}}}"
SPECIAL5 = "{{{ $$ }}}"
SPECIAL6 = "{{{ $$$ }}}"

EMPTY_VAR=""

VAR_DEFAULT_UNSET = "${UNSET_VAR-uuu}"
VAR_DEFAULT_UNSET_OR_EMPTY = "${EMPTY_VAR-eee}"
VAR_DEFAULT_EMPTY = "${EMPTY_VAR:-eee}"

export EQUALS='postgres://localhost:5432/database?sslmode=disable'
`
	expected := map[string]*string{
		"A":                            strPtr("aaa"),
		"B":                            strPtr("bbb"),
		"C":                            strPtr("ccc"),
		"D":                            strPtr("ddd"),
		"E":                            strPtr("eee"),
		"F":                            strPtr(""),
		"G":                            nil,
		"H":                            strPtr("my_value"),
		"I":                            strPtr("bar baz"),
		"J":                            strPtr("foo bar"),
		"L":                            strPtr("my\nmulti\nline\nentry"),
		"M":                            strPtr("foo aaa bar"),
		"MYVAR":                        strPtr("before bar baz after foo bar opa "),
		"N1":                           strPtr("41"),
		"N2":                           strPtr("42"),
		"N3":                           strPtr("43AS3sA43"),
		"N4":                           strPtr("44AS4sA44"),
		"SIMPLE_SPACING":               strPtr("aaa bbb ccc"),
		"SPECIAL1":                     strPtr("{{{ aaa }}}"),
		"SPECIAL2":                     strPtr("{{{ aaa }}}"),
		"SPECIAL3":                     strPtr("{{{ $ }}}"),
		"SPECIAL4":                     strPtr("{{{ $ $ $}}}"),
		"SPECIAL5":                     strPtr("{{{ $$ }}}"),
		"SPECIAL6":                     strPtr("{{{ $$$ }}}"),
		"EMPTY_VAR":                    strPtr(""),
		"VAR_DEFAULT_UNSET":            strPtr("uuu"),
		"VAR_DEFAULT_UNSET_OR_EMPTY":   strPtr(""),
		"VAR_DEFAULT_EMPTY":            strPtr("eee"),
		"EXPORTED_VAR":                 strPtr("exported_value"),
		"OPTION_B":                     strPtr("\\n"),
		"EQUALS":                       strPtr("postgres://localhost:5432/database?sslmode=disable"),
		"VAR_TO_BE_LOADED_FROM_OS_ENV": strPtr("loaded_from_os_env"),
	}

	d := dao.NewDefaultDaoFromMap(map[string]*string{
		"VAR_TO_BE_LOADED_FROM_OS_ENV": strPtr("loaded_from_os_env"),
	})
	assert.DeepEqual(t, expected, GetVariables(d, is))
}

func strPtr(s string) *string {
	return &s
}
