package main

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
	"gotest.tools/v3/assert"

	"github.com/ulyssessouza/envlang/dao"
)

func TestFull(t *testing.T) {
	// disableLogs()
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


`
	expected := map[string]*string{
		"A":     strPtr("aaa"),
		"B":     strPtr("bbb"),
		"C":     strPtr("ccc"),
		"D":     strPtr("ddd"),
		"E":     strPtr("'eee'"),
		"F":     strPtr(""),
		"G":     nil,
		"H":     strPtr("my_value"),
		"I":     strPtr("bar baz"),
		"J":     strPtr("foo bar"),
		"L":     strPtr("my\nmulti\nline\nentry"),
		"M":     strPtr("foo aaa bar"),
		"MYVAR": strPtr("before bar baz after foo bar opa "),
		"N1":    strPtr("41"),
		"N2":    strPtr("42"),
		"N3":    strPtr("43AS3sA43"),
		"N4":    strPtr("44AS4sA44"),
	}

	d := dao.NewDefaultEnvLangDao()
	assert.DeepEqual(t, expected, GetVariables(d, is))
}

func disableLogs() {
	f, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(f)
}

func strPtr(s string) *string {
	return &s
}
