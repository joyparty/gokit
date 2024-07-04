package gokit

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestListOf(t *testing.T) {
	l := NewListOf[int]()

	t.Run("PushBack", func(t *testing.T) {
		l.PushBack(1)
		l.PushBack(2)
		l.PushBack(3)

		if !assertList(l, []int{1, 2, 3}) {
			t.Error("list is not correct")
		}
	})

	t.Run("Back", func(t *testing.T) {
		if l.Back().Value() != 3 {
			t.Error("back is not correct")
		}
	})

	t.Run("Front", func(t *testing.T) {
		if l.Front().Value() != 1 {
			t.Error("front is not correct")
		}
	})

	t.Run("InsertAfter", func(t *testing.T) {
		l.InsertAfter(4, l.Front())

		if !assertList(l, []int{1, 4, 2, 3}) {
			t.Error("list is not correct")
		}
	})

	t.Run("InsertBefore", func(t *testing.T) {
		l.InsertBefore(5, l.Back())

		if !assertList(l, []int{1, 4, 2, 5, 3}) {
			t.Error("list is not correct")
		}
	})

	t.Run("MoveAfter", func(t *testing.T) {
		l.MoveAfter(l.Front(), l.Front().Next())

		if !assertList(l, []int{4, 1, 2, 5, 3}) {
			t.Error("list is not correct")
		}
	})

	t.Run("MoveBefore", func(t *testing.T) {
		l.MoveBefore(l.Front(), l.Back())

		if !assertList(l, []int{1, 2, 5, 4, 3}) {
			t.Error("list is not correct")
		}
	})

	t.Run("MoveToBack", func(t *testing.T) {
		l.MoveToBack(l.Front().Next())

		if !assertList(l, []int{1, 5, 4, 3, 2}) {
			t.Error("list is not correct")
		}
	})

	t.Run("MoveToFront", func(t *testing.T) {
		l.MoveToFront(l.Back().Prev())

		if !assertList(l, []int{3, 1, 5, 4, 2}) {
			t.Error("list is not correct")
		}
	})

	t.Run("PushFront", func(t *testing.T) {
		l.PushFront(6)

		if !assertList(l, []int{6, 3, 1, 5, 4, 2}) {
			t.Error("list is not correct")
		}
	})

	t.Run("PushFrontList", func(t *testing.T) {
		other := NewListOf[int]()
		other.PushBack(7)
		other.PushBack(8)

		l.PushFrontList(other)

		if !assertList(l, []int{7, 8, 6, 3, 1, 5, 4, 2}) {
			t.Error("list is not correct")
		}
	})

	t.Run("PushBackList", func(t *testing.T) {
		other := NewListOf[int]()
		other.PushBack(9)
		other.PushBack(10)

		l.PushBackList(other)

		if !assertList(l, []int{7, 8, 6, 3, 1, 5, 4, 2, 9, 10}) {
			t.Error("list is not correct")
		}
	})

	t.Run("Remove", func(t *testing.T) {
		l.Remove(l.Front().Next().Next())

		if !assertList(l, []int{7, 8, 3, 1, 5, 4, 2, 9, 10}) {
			t.Error("list is not correct")
		}
	})
}

func assertList(actual *ListOf[int], expected []int) bool {
	if actual.Len() != len(expected) {
		return false
	}

	i := 0
	for el := actual.Front(); el != nil; el = el.Next() {
		if el.Value() != expected[i] {
			return false
		}
		i++
	}

	return true
}

func printList(l *ListOf[int]) {
	vals := []string{}
	for el := l.Front(); el != nil; el = el.Next() {
		vals = append(vals, strconv.Itoa(el.Value()))
	}
	fmt.Fprintf(os.Stderr, "%s\n", strings.Join(vals, ", "))
}
