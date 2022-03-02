package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ulyssessouza/envlang/handlers"
)

func main() {
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

MYVAR = "before ${MYVALVAR} after $SECONDVAR opa ${UNKNOWNVAR}"

M = "foo${A} bar"

N1=41

N2="42"

N3="ASSSASA43"


`

	// is := `N2="ASSSASA42"`
	disableLogs()

	for k, v := range handlers.GetVariables(is) {
		val := "<nil>"
		if v != nil {
			val = *v
		}
		fmt.Printf(
			"%10q >>> %q\n",
			k,
			val,
		)
	}
}

func disableLogs() {
	f, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)
}
