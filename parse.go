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

func ParseCommand(cmd string, args string) Command {
	if cmd == "print" {
		return parsePrint(args)
	} else if cmd == "printc" {
		return parsePrintc(args)
	} else {
		return &PrintCommand{Msg: fmt.Sprintf("Error: Unknown command: %s", cmd)}
	}
}
