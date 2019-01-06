package symbols

type Scope string

const (
	GlobalScope  Scope = "GLOBAL"
	LocalScope   Scope = "LOCAL"
	BuiltinScope Scope = "BUILTIN"
)

type Symbol struct {
	Scope Scope
	Index int
	Name  string
}

type SymbolTable struct {
	Outer           *SymbolTable
	store           map[string]Symbol
	DefinitionCount int
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
	symbol := Symbol{Name: name, Index: s.DefinitionCount, Scope: GlobalScope}

	if s.Outer == nil {
		symbol.Scope = GlobalScope
	} else {
		symbol.Scope = LocalScope
	}

	s.store[name] = symbol
	s.DefinitionCount++
	return symbol
}

func (s *SymbolTable) DefineBuiltin(index int, name string) Symbol {
	symbol := Symbol{Name: name, Index: index, Scope: BuiltinScope}
	s.store[name] = symbol
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
