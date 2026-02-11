package presentation

//treated as immutable
type RewritingSystem struct{
	LHS []RawWord //subwords to be replaced
	RHS []RawWord //replacements
}

func (R RewritingSystem) Rewrite(w RawWord) {
	
}