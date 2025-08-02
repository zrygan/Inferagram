package grammar

type FormalSymbol struct {
	Name       string
	IsTerminal bool
}

type FormalProduction struct {
	Left  FormalSymbol
	Right []FormalSymbol
}

type Grammar struct {
	Start        FormalSymbol
	NonTerminals map[string]FormalSymbol
	Terminals    map[string]FormalSymbol
	Productions  map[string][]FormalProduction
}
