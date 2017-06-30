package main

import (
	"./bubblesort"
	"testing"
)

func TestBubbleSort1(t *testing.T) {
	list := []int{}
	sortedList := []int{}

	if bubblesort.Sort(list); !slicesAreEqual(list, sortedList) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", list, sortedList)
	}
}

func TestBubbleSort2(t *testing.T) {
	list := []int{1}
	sortedList := []int{1}

	if bubblesort.Sort(list); !slicesAreEqual(list, sortedList) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", list, sortedList)
	}
}

func TestBubbleSort3(t *testing.T) {
	list := []int{5, 4, 2, 3, 1, 7, 0}
	sortedList := []int{0, 1, 2, 3, 4, 5, 7}

	if bubblesort.Sort(list); !slicesAreEqual(list, sortedList) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", list, sortedList)
	}
}

func slicesAreEqual(list, sortedList []int) bool {
	for i := range list {
		if list[i] != sortedList[i] {
			return false
		}
	}
	return true
}
