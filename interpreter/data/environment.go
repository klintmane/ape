package data

type Environment struct {
	store map[string]Data
	outer *Environment
}

func NewEnvironmentClosure(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Environment {
	s := make(map[string]Data)
	return &Environment{store: s, outer: nil}
}

func (e *Environment) Get(name string) (Data, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Data) Data {
	e.store[name] = val
	return val
}
