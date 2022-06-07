package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
var ErrNoWorkers = errors.New("no workers provided")

type Task func() error

type counter struct {
	mu            sync.Mutex
	maxErrors     int
	currentErrors int
	ignoreErrors  bool
}

func createCounter(maxErrors int) *counter {
	ignoreErros := maxErrors <= 0
	return &counter{
		maxErrors:    maxErrors,
		ignoreErrors: ignoreErros,
	}
}

func (c *counter) addError() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.currentErrors++
}

func (c *counter) hasMaxErrors() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return !c.ignoreErrors && c.currentErrors >= c.maxErrors
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n < 1 {
		return ErrNoWorkers
	}

	c := createCounter(m)

	var wg sync.WaitGroup
	wg.Add(n)

	tasksCh := make(chan Task)

	for i := 0; i < n; i++ {
		go runner(&wg, tasksCh, c)
	}

	for _, task := range tasks {
		if c.hasMaxErrors() {
			break
		}
		tasksCh <- task
	}

	close(tasksCh)

	wg.Wait()

	if c.hasMaxErrors() {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func runner(wg *sync.WaitGroup, tc <-chan Task, c *counter) {
	defer wg.Done()

	for t := range tc {
		if t() != nil {
			c.addError()
		}
	}
}
