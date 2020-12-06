package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToPrefix(t *testing.T) {
	print := ParseCommand("print Hello world")
	printc := ParseCommand("printc   12 h")

	assert.Equal(t, print, &PrintCommand{Msg: "Hello world"})
	assert.Equal(t, printc, &PrintcCommand{Count: 12, Str: "h"})
}
