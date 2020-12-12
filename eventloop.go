package main

import (
	"sync"
)

type queue struct {
	sync.Mutex
	queue      []Command
	waiting    bool
	newRequest chan struct{}
}

func (q *queue) push(cmd Command) {
	q.Lock()
	defer q.Unlock()
	q.queue = append(q.queue, cmd)
	if q.waiting {
		q.waiting = false
		q.newRequest <- struct{}{}
	}
}

func (q *queue) pop() Command {
	q.Lock()
	defer q.Unlock()
	if q.empty() {
		q.waiting = true
		q.Unlock()
		<-q.newRequest
		q.Lock()
	}

	cmd := q.queue[0]
	q.queue[0] = nil
	q.queue = q.queue[1:]
	return cmd
}

func (q *queue) empty() bool {
	return len(q.queue) == 0
}

func newQueue() queue {
	return queue{
		queue:      []Command{},
		waiting:    false,
		newRequest: make(chan struct{}),
	}
}

type Loop struct {
	queue   queue
	running bool
	done    chan struct{}
}

func (l *Loop) Start() {
	l.queue = newQueue()
	l.running = true
	l.done = make(chan struct{})
	go func() {
		for l.running || !l.queue.empty() {
			cmd := l.queue.pop()
			cmd.Execute(l)
		}
		l.done <- struct{}{}
	}()
}

func (l *Loop) Post(cmd Command) {
	l.queue.push(cmd)
}

func (l *Loop) AwaitFinish() {
	l.queue.push(CommandFunc(func(h Handler) {
		l.running = false
	}))
	<-l.done
}
