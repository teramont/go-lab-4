package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type putCmd struct {
	a     int
	index int
	arr   []int
}

type putAsyncCmd struct {
	a     int
	index int
	arr   []int
}

func (cmd *putCmd) Execute(handler Handler) {
	cmd.arr[cmd.index] = cmd.a
}

func (cmd *putAsyncCmd) Execute(handler Handler) {
	handler.Post(&putCmd{a: cmd.a, index: cmd.index, arr: cmd.arr})
}

func TestLoop(t *testing.T) {
	loop := Loop{}
	loop.Start()
	var arr [5]int
	loop.Post(&putCmd{a: 17, index: 0, arr: arr[:]})
	loop.Post(&putAsyncCmd{a: 1, index: 2, arr: arr[:]})
	loop.Post(&putAsyncCmd{a: 2, index: 3, arr: arr[:]})
	loop.Post(&putCmd{a: 10, index: 1, arr: arr[:]})
	loop.Post(&putAsyncCmd{a: 3, index: 4, arr: arr[:]})

	loop.AwaitFinish()
	assert.Equal(t, arr, [5]int{17, 10, 1, 2, 3})
}
