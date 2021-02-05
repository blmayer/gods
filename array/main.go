package array

type Array struct {
	a []interface{}
}

func New() Array {
	return Array{}
}

func (a Array) NotIn(b Array, eq func(a, b interface{}) bool) (c Array) {
	for _, i := range a.a {
		for _, j := range b.a {
			if !eq(i, j) {
				c.a = append(c.a, i)
			}
		}
	}
	return
}
