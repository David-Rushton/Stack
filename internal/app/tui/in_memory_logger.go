package tui

import (
	"fmt"
	"sync"
)

type InMemoryLog struct {
	mu    sync.Mutex
	lines []string
}

func (w *InMemoryLog) Write(b []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	// DEBUG: Remove later.
	fmt.Print(string(b))

	w.lines = append(w.lines, string(b))
	return len(b), nil
}
