package symbols

type Scope string

const (
	GlobalScope  Scope = "GLOBAL"
	LocalScope   Scope = "LOCAL"
	BuiltinScope Scope = "BUILTIN"
	FreeScope    Scope = "FREE"
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
	Free            []Symbol // Free symbol store
}

func New() *SymbolTable {
	return &SymbolTable{
		store: make(map[string]Symbol),
		Free:  []Symbol{},
	}
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
		if !ok {
			return obj, ok
		}
		if obj.Scope == GlobalScope || obj.Scope == BuiltinScope {
			return obj, ok
		}
		free := s.defineFree(obj)
		return free, true
	}
	return obj, ok
}

func (s *SymbolTable) defineFree(original Symbol) Symbol {
	s.Free = append(s.Free, original)
	symbol := Symbol{Name: original.Name, Index: len(s.Free) - 1}
	symbol.Scope = FreeScope
	s.store[original.Name] = symbol
	return symbol
}
