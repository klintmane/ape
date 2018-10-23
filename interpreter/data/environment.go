package data

type Environment struct {
	store map[string]Data
}

func NewEnvironment() *Environment {
	s := make(map[string]Data)
	return &Environment{store: s}
}

func (e *Environment) Get(name string) (Data, bool) {
	data, ok := e.store[name]
	return data, ok
}

func (e *Environment) Set(name string, val Data) Data {
	e.store[name] = val
	return val
}
