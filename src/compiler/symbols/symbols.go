package symbols

type Scope string

const (
	GlobalScope Scope = "GLOBAL"
)

type Symbol struct {
	Scope Scope
	Index int
	Name  string
}

type SymbolTable struct {
	store           map[string]Symbol
	definitionCount int
}

func New() *SymbolTable {
	s := make(map[string]Symbol)
	return &SymbolTable{store: s}
}

func (s *SymbolTable) Define(name string) Symbol {
	symbol := Symbol{Name: name, Index: s.definitionCount, Scope: GlobalScope}
	s.store[name] = symbol
	s.definitionCount++
	return symbol
}

func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.store[name]
	return obj, ok
}
