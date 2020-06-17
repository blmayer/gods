package list

import (
	"testing"
)

var (
	vals = []int{1, 6, 9, 13, 4}
	l    = New()
)

func init() {
	for _, v := range vals {
		l.PushBack(v)
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
