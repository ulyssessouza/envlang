package main

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"gotest.tools/v3/assert"

	"github.com/ulyssessouza/envlang/dao"
)

func TestFull(t *testing.T) {
	log.SetLevel(log.PanicLevel)
	is := `


A=aaa
B="bbb"
C =ccc
D= ddd
E='eee'
F=
G

H="my_value"


I = bar baz



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


SPECIAL1 = "{{{ ${A} }}}"
SPECIAL2 = "{{{ $A }}}"

`
	expected := map[string]*string{
		"A":        strPtr("aaa"),
		"B":        strPtr("bbb"),
		"C":        strPtr("ccc"),
		"D":        strPtr("ddd"),
		"E":        strPtr("'eee'"),
		"F":        strPtr(""),
		"G":        nil,
		"H":        strPtr("my_value"),
		"I":        strPtr("bar baz"),
		"J":        strPtr("foo bar"),
		"L":        strPtr("my\nmulti\nline\nentry"),
		"M":        strPtr("foo aaa bar"),
		"MYVAR":    strPtr("before bar baz after foo bar opa "),
		"N1":       strPtr("41"),
		"N2":       strPtr("42"),
		"N3":       strPtr("43AS3sA43"),
		"N4":       strPtr("44AS4sA44"),
		"SPECIAL1": strPtr("{{{ aaa }}}"),
		"SPECIAL2": strPtr("{{{ aaa }}}"),
	}

	d := dao.NewDefaultDao()
	assert.DeepEqual(t, expected, GetVariables(d, is))
}

func strPtr(s string) *string {
	return &s
}
