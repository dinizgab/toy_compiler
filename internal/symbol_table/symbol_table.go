package symboltable

type SymbolInformation struct {
	Name string
	Addr int
}

// Maps name of the id to its information
type SymbolTable map[string]SymbolInformation

func New() SymbolTable {
	return make(SymbolTable)
}

func (s SymbolTable) AddEntry(name string, entry SymbolInformation) {
	s[name] = entry
}

func (s SymbolTable) LookUp(name string) (SymbolInformation, bool) {
	entry, found := s[name]
	if !found {
		return SymbolInformation{}, false
	}

	return entry, found
}