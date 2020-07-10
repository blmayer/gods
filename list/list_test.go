package list

import (
	"testing"
)

var (
	vals = []int{1, 6, 9, 13, 4}
	l    = New()
)

func init() {
	load()
}

func load() {
	l = New()
	for _, v := range vals {
		l.PushBack(v)
	}
}

func TestGroupBy(t *testing.T) {
	m := l.GroupBy(func(i interface{}) interface{} { return i.(int) > 7 })

	if len(m) != 2 {
		t.Error("map returned wrong keys\n")
	}
	if m[true].Len() != 2 {
		t.Error("true contains different number of elements")
	}
	if m[false].Len() != 3 {
		t.Error("false contains different number of elements")
	}

	for k, v := range m {
		t.Logf("key %v value %+v\n", k, v.List)
		for e := v.Front(); e != nil; e = e.Next() {
			t.Log(e.Value)
			if e.Value == nil {
				t.Error("nil element")
			}
		}
	}
}

func TestMap(t *testing.T) {
	m := l.Map(func(i interface{}) interface{} { return i.(int) + 10 })

	i := 0
	for e := m.Front(); e != nil; e = e.Next() {
		if e.Value != vals[i]+10 {
			t.Errorf("%d is not %d\n", i, e.Value)
		}
		i++
	}
}

func TestFilter(t *testing.T) {
	l.Filter(func(i interface{}) bool { return i.(int) > 17 })

	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(int) > 17 {
			t.Errorf("%d is > 17\n", e.Value)
		}
		i++
	}
	for e := l.Front(); e != nil; e = e.Next() {
		t.Log(e.Value)
	}
}

func TestPushSorted(t *testing.T) {
	l = New()
	load()

	l.PushSorted(10, func(i, j interface{}) bool { return i.(int) < j.(int) })
	l.PushSorted(20, func(i, j interface{}) bool { return i.(int) < j.(int) })
	l.PushSorted(0, func(i, j interface{}) bool { return i.(int) < j.(int) })

	newVals := []int{0, 1, 6, 9, 10, 13, 4, 20}
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(int) != newVals[i] {
			t.Errorf("%d is not %d\n", e.Value, newVals[i])
		}
		i++
	}
	for e := l.Front(); e != nil; e = e.Next() {
		t.Log(e.Value)
	}
}

func TestPushSort(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	// The following must panic
	l.PushSort(10)
	l.PushSort(20)
	l.PushSort(0)

	l = New()
	load()

	l.Less = func(i, j interface{}) bool { return i.(int) < j.(int) }
	l.PushSort(10)
	l.PushSort(20)
	l.PushSort(0)

	newVals := []int{0, 1, 6, 9, 10, 13, 4, 20}
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(int) != newVals[i] {
			t.Errorf("%d is not %d\n", e.Value, newVals[i])
		}
		i++
	}
	for e := l.Front(); e != nil; e = e.Next() {
		t.Log(e.Value)
	}
}
