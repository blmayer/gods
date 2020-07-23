package dict

type Dict struct {
	Keys []string
	m    map[string]interface{}
}

func New() Dict {
	m := Dict{}
	m.Keys = make([]string, 0, 0)
	m.m = make(map[string]interface{})
	return m
}

func (m Dict) Get(key string) interface{} {
	return m.m[key]
}

func (m *Dict) Set(key string, value interface{}) {
	if m.m[key] == nil {
		m.Keys = append(m.Keys, key)
	}

	m.m[key] = value
}

func (m *Dict) Remove(key string) {
	l := len(m.Keys)
	for i := 0; i < l; i++ {
		if m.Keys[i] == key {
			m.Keys[i] = m.Keys[l-1]
			m.Keys = m.Keys[:l-1]
			break
		}
	}

	delete(m.m, key)
}

func (m Dict) Dict(f func(interface{}) interface{}) Dict {
	for _, k := range m.Keys {
		m.m[k] = f(m.m[k])
	}
	return m
}

// Filter removes some elements given a criterion
func (m *Dict) Filter(f func(interface{}) bool) Dict {
	for _, k := range m.Keys {
		if f(m.m[k]) {
			m.Remove(k)
		}
	}
	return *m
}
