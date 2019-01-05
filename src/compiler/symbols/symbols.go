package symbols

type Scope string

const (
	GlobalScope Scope = "GLOBAL"
	LocalScope  Scope = "LOCAL"
)

type Symbol struct {
	Scope Scope
	Index int
	Name  string
}

type SymbolTable struct {
	Outer           *SymbolTable
	store           map[string]Symbol
	definitionCount int
}

func New() *SymbolTable {
	s := make(map[string]Symbol)
	return &SymbolTable{store: s}
}

func NewEnclosed(outer *SymbolTable) *SymbolTable {
	s := New()
	s.Outer = outer
	return s
}

func (s *SymbolTable) Define(name string) Symbol {
	symbol := Symbol{Name: name, Index: s.definitionCount, Scope: GlobalScope}

	if s.Outer == nil {
		symbol.Scope = GlobalScope
	} else {
		symbol.Scope = LocalScope
	}

	s.store[name] = symbol
	s.definitionCount++
	return symbol
}

func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.store[name]
	if !ok && s.Outer != nil {
		obj, ok = s.Outer.Resolve(name)
		return obj, ok
	}
	return obj, ok
}
