// Package list contains some functions to help with
// the built-in "container/list" package.
package list

import (
	ls "container/list"
)

// List is the exported list with some more methods
type List struct {
	*ls.List
	Less func(interface{}, interface{}) bool
}

// New creates a list
func New() List {
	l := List{}
	l.List = ls.New()
	return l
}

// Map applies a function to each item and returns the new list
func (l List) Map(f func(interface{}) interface{}) List {
	for e := l.Front(); e != nil; e = e.Next() {
		e.Value = f(e.Value)
	}
	return l
}

// Filter removes some elements given a criterion
func (l List) Filter(f func(interface{}) bool) List {
	for e := l.Front(); e != nil; e = e.Next() {
		if f(e.Value) {
			e = e.Prev()
			l.Remove(e.Next())
		}
	}
	return l
}

// PushSorted inserts a new element sorted given a sort function
func (l List) PushSorted(e interface{}, less func(interface{}, interface{}) bool) {
	for x := l.Front(); x != nil; x = x.Next() {
		if less(e, x.Value) {
			l.InsertBefore(e, x)
			break
		}
	}
}

// PushSort inserts a new element sorted given the defined sort function
func (l List) PushSort(e interface{}) {
	if l.Less == nil {
		panic("Less function is not defined")
	}
	l.PushSorted(e, l.Less)
}
