package missle

import (
	"fmt"
)

type WaitQueue struct {
	items []interface{}
}

func (w *WaitQueue) String() string {
	return fmt.Sprintf("WaitQueue(%d): %v", w.Len(), w.items)
}

func (w *WaitQueue) Push(item interface{}) bool {
	if w.Index(item) >= 0 {
		return false
	}
	w.items = append(w.items, item)
	return true
}

func (w *WaitQueue) Pop() interface{} {
	if w.Len() <= 0 {
		return nil
	}

	item := w.items[0]
	w.items = w.items[1:]
	return item
}

func (w *WaitQueue) Len() int {
	return len(w.items)
}

func (w *WaitQueue) Delete(item interface{}) bool {
	index := w.Index(item)
	if index < 0 {
		return false
	}

	w.items = append(w.items[:index], w.items[index+1:]...)
	return true
}

func (w *WaitQueue) Index(item interface{}) int {
	for i := 0; i < len(w.items); i++ {
		if w.items[i] == item {
			return i
		}
	}

	return -1
}
