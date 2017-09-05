package structures

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	first := LinkedList{data: 1}
	second := LinkedList{data: 2}
	first.next = &second

	if first.next.data != 2 {
		t.Errorf("Expected 2, got %d", first.next.data)
	}
}
