package main

import (
	"fmt"
	"strconv"
	"strings"
)

func errorCmd(Msg string) Command {
	return &PrintCommand{Msg}
}

func parsePrint(args string) Command {
	return &PrintCommand{Msg: args}
}

func parsePrintc(args string) Command {
	parts := strings.Fields(args)
	if len(parts) != 2 {
		msg := fmt.Sprintf("Error in 'printc %s': expected 2 arguments, found %d", args, len(parts))
		return errorCmd(msg)
	}

	count, err := strconv.Atoi(parts[0])
	if err != nil {
		msg := fmt.Sprintf("Error while parsing number '%s': %s", parts[0], err.Error())
		return errorCmd(msg)
	}
	return &PrintcCommand{Count: count, Str: parts[1]}
}

func ParseCommand(line string) Command {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return errorCmd("Empty line")
	}
	cmd := parts[0]
	start := len(cmd)
	args := strings.TrimPrefix(line[start:], " ")
	if cmd == "print" {
		return parsePrint(args)
	} else if cmd == "printc" {
		return parsePrintc(args)
	} else {
		return errorCmd(fmt.Sprintf("Error: Unknown command: %s", cmd))
	}
}
