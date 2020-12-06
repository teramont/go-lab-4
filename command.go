package main

import (
	"fmt"
	"strings"
)

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command)
}

type PrintCommand struct {
	Msg string
}

func (cmd *PrintCommand) Execute(handler Handler) {
	fmt.Println(cmd.Msg)
}

type PrintcCommand struct {
	Str   string
	Count int
}

func (cmd *PrintcCommand) Execute(handler Handler) {
	msg := strings.Repeat(cmd.Str, cmd.Count)
	print := PrintCommand{msg}
	handler.Post(&print)
}
