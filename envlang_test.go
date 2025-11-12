package envlang

import (
	"os"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"gotest.tools/v3/assert"

	"github.com/ulyssessouza/envlang/handlers"
	"github.com/ulyssessouza/envlang/store"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func TestGetFromReader(t *testing.T) {
	expected := map[string]*string{
		"A": strPtr("aaa"),
	}
	d := store.NewDefaultStoreFromMap(nil)
	assert.DeepEqual(t, expected, GetVariablesFromInputStream(d, strings.NewReader(`A=aaa`)))
}

func TestLoad(t *testing.T) {
	const envlangOsLoadVariable = "ENVLANG_TEST_OSLOAD_VARIABLE"
	err := Load("./fixtures/load.env")
	assert.NilError(t, err)
	osLoaded, ok := os.LookupEnv(envlangOsLoadVariable)
	assert.Assert(t, ok)
	assert.Equal(t, osLoaded, "ENVLANG_TEST_OSLOAD_VALUE")
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
A=$VAR_FROM_STORE
`,
			map[string]*string{
				"A": strPtr("aaa"),
			},
			map[string]*string{
				"VAR_FROM_STORE": strPtr("aaa"),
			},
		},
		{
			"SimpleWithSimpleVariableWithSpaces",
			`
A = $VAR_FROM_STORE
		`,
			map[string]*string{
				"A": strPtr("aaa"),
			},
			map[string]*string{
				"VAR_FROM_STORE": strPtr("aaa"),
			},
		},
		{
			"SimpleWithStrictVariable",
			`
A=${VAR_FROM_STORE}
`,
			map[string]*string{
				"A": strPtr("aaa"),
			},
			map[string]*string{
				"VAR_FROM_STORE": strPtr("aaa"),
			},
		},
		{
			"SimpleWithStrictVariableWithSpaces",
			`
A = ${VAR_FROM_STORE}
`,
			map[string]*string{
				"A": strPtr("aaa"),
			},
			map[string]*string{
				"VAR_FROM_STORE": strPtr("aaa"),
			},
		},
		{
			"SimpleWithStrictVariableWithSpacesAndInternalSpaces",
			`
A = ${ VAR_FROM_STORE }
`,
			map[string]*string{
				"A": strPtr("aaa"),
			},
			map[string]*string{
				"VAR_FROM_STORE": strPtr("aaa"),
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
			"VariableWithAssignForUnsetOrEmpty",
			`VAR_ASSIGN_UNSET_OR_EMPTY = "${EMPTY_VAR:=assigned}"`,
			map[string]*string{
				"VAR_ASSIGN_UNSET_OR_EMPTY": strPtr("assigned"),
				"EMPTY_VAR":                 strPtr("assigned"),
			},
			map[string]*string{
				"EMPTY_VAR": strPtr(""),
			},
		},
		{
			"VariableWithAssignForUnset",
			`VAR_ASSIGN_UNSET = "${UNSET_VAR=assigned}"`,
			map[string]*string{
				"VAR_ASSIGN_UNSET": strPtr("assigned"),
				"UNSET_VAR":        strPtr("assigned"),
			},
			nil,
		},
		{
			"VariableWithAssignKeepsEmpty",
			`VAR_ASSIGN_EMPTY = "${EMPTY_VAR=assigned}"`,
			map[string]*string{
				"VAR_ASSIGN_EMPTY": strPtr(""),
			},
			map[string]*string{
				"EMPTY_VAR": strPtr(""),
			},
		},
		{
			"VariableWithAssignKeepsExisting",
			`VAR_ASSIGN_EXISTING = "${EXISTING_VAR:=assigned}"`,
			map[string]*string{
				"VAR_ASSIGN_EXISTING": strPtr("existing"),
			},
			map[string]*string{
				"EXISTING_VAR": strPtr("existing"),
			},
		},
		{
			"VariableWithAlternateForSet",
			`VAR_ALTERNATE_SET = "${EXISTING_VAR:+alternate}"`,
			map[string]*string{
				"VAR_ALTERNATE_SET": strPtr("alternate"),
			},
			map[string]*string{
				"EXISTING_VAR": strPtr("existing"),
			},
		},
		{
			"VariableWithAlternateForUnset",
			`VAR_ALTERNATE_UNSET = "${UNSET_VAR:+alternate}"`,
			map[string]*string{
				"VAR_ALTERNATE_UNSET": strPtr(""),
			},
			nil,
		},
		{
			"VariableWithAlternateForEmpty",
			`VAR_ALTERNATE_EMPTY = "${EMPTY_VAR:+alternate}"`,
			map[string]*string{
				"VAR_ALTERNATE_EMPTY": strPtr(""),
			},
			map[string]*string{
				"EMPTY_VAR": strPtr(""),
			},
		},
		{
			"VariableWithAlternateForSetEvenEmpty",
			`VAR_ALTERNATE_SET_EMPTY = "${EMPTY_VAR+alternate}"`,
			map[string]*string{
				"VAR_ALTERNATE_SET_EMPTY": strPtr("alternate"),
			},
			map[string]*string{
				"EMPTY_VAR": strPtr(""),
			},
		},
		{
			"VariableLengthOfString",
			`VAR_LENGTH = "${#EXISTING_VAR}"`,
			map[string]*string{
				"VAR_LENGTH": strPtr("5"),
			},
			map[string]*string{
				"EXISTING_VAR": strPtr("hello"),
			},
		},
		{
			"VariableLengthOfEmpty",
			`VAR_LENGTH = "${#EMPTY_VAR}"`,
			map[string]*string{
				"VAR_LENGTH": strPtr("0"),
			},
			map[string]*string{
				"EMPTY_VAR": strPtr(""),
			},
		},
		{
			"VariableLengthOfUnset",
			`VAR_LENGTH = "${#UNSET_VAR}"`,
			map[string]*string{
				"VAR_LENGTH": strPtr("0"),
			},
			nil,
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
			d := store.NewDefaultStoreFromMap(tt.envState)
			assert.DeepEqual(t, tt.expected, GetVariables(d, tt.input))
		})
	}
}

func TestGetValueWithError(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		envState      map[string]*string
		expectedError string
		expectedVar   string
	}{
		{
			"ErrorForUnsetWithMessage",
			`VAR_ERROR = "${UNSET_VAR:?custom error message}"`,
			nil,
			"UNSET_VAR: custom error message",
			"UNSET_VAR",
		},
		{
			"ErrorForEmptyWithMessage",
			`VAR_ERROR = "${EMPTY_VAR:?variable is empty}"`,
			map[string]*string{
				"EMPTY_VAR": strPtr(""),
			},
			"EMPTY_VAR: variable is empty",
			"EMPTY_VAR",
		},
		{
			"ErrorForUnsetWithoutMessage",
			`VAR_ERROR = "${UNSET_VAR?}"`,
			nil,
			"UNSET_VAR: parameter not set",
			"UNSET_VAR",
		},
		{
			"ErrorOperatorKeepsEmpty",
			`VAR_ERROR = "${EMPTY_VAR?not unset}"`,
			map[string]*string{
				"EMPTY_VAR": strPtr(""),
			},
			"",
			"",
		},
		{
			"ErrorOperatorKeepsValue",
			`VAR_NO_ERROR = "${EXISTING_VAR:?should not error}"`,
			map[string]*string{
				"EXISTING_VAR": strPtr("existing"),
			},
			"",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := store.NewDefaultStoreFromMap(tt.envState)

			if tt.expectedError == "" {
				// Should not panic
				result := GetVariables(d, tt.input)
				assert.Assert(t, result != nil)
			} else {
				// Should panic with ParameterExpansionError
				defer func() {
					r := recover()
					assert.Assert(t, r != nil, "expected panic but got none")

					err, ok := r.(*handlers.ParameterExpansionError)
					assert.Assert(t, ok, "expected ParameterExpansionError")
					assert.Equal(t, tt.expectedVar, err.VarName)
					assert.Equal(t, tt.expectedError, err.Error())
				}()

				GetVariables(d, tt.input)
				t.Fatal("expected panic but execution continued")
			}
		})
	}
}

//nolint:funlen
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

SPECIAL_CHAR_A=unquoted phrase special ã char
SPECIAL_CHAR_C=unquoted phrase special ç char
SPECIAL_CHAR_E=unquoted phrase special è char

VAR-WITH-DASHES="dashes"
VAR.WITH.DOTS="dots"
VAR_WITH_UNDERSCORES="underscores"

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
		"VAR-WITH-DASHES":              strPtr("dashes"),
		"VAR.WITH.DOTS":                strPtr("dots"),
		"VAR_WITH_UNDERSCORES":         strPtr("underscores"),
		"SPECIAL_CHAR_A":               strPtr("unquoted phrase special ã char"),
		"SPECIAL_CHAR_C":               strPtr("unquoted phrase special ç char"),
		"SPECIAL_CHAR_E":               strPtr("unquoted phrase special è char"),
	}

	d := store.NewDefaultStoreFromMap(map[string]*string{
		"VAR_TO_BE_LOADED_FROM_OS_ENV": strPtr("loaded_from_os_env"),
	})
	assert.DeepEqual(t, expected, GetVariables(d, is))
}

func strPtr(s string) *string {
	return &s
}
