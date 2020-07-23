package dict

import "testing"

var (
	vals = []int{6, 7, 3, 5, 1}
	keys = []string{"g", "h", "c", "e", "a"}
	m    = New()
)

func TestSetGet(t *testing.T) {
	for i, k := range keys {
		m.Set(k, vals[i])
	}

	if m.Keys[0] != keys[0] {
		t.Error("key 0 is wrong")
	}
	if m.Get(keys[2]).(int) != vals[2] {
		t.Error("wrong value for key 3")
	}

	m.Set("c", vals[1])
	if m.Get("c").(int) != vals[1] {
		t.Error("substitution error")
	}
}
