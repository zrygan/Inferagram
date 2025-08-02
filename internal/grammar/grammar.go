package grammar

type Symbol struct {
	Name       string
	IsTerminal bool
}

type Production struct {
	Left  Symbol
	Right []Symbol
}

type Grammar struct {
	Start        Symbol
	NonTerminals map[string]Symbol
	Terminals    map[string]Symbol
	Productions  map[string][]Production
}
