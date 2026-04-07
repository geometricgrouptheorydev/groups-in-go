package presentation

// Replaces the first instance of sub in w by replacement
func (w RawWord) ReplaceRawSubWordFirstMatch(sub RawWord, replacement RawWord) RawWord {
	// We expand w for now, this will change once the KMP algorithms are exponent-sensitive
	v, s := expandRawWord(w), expandRawWord(sub)
	index, exists := SubExpandedRawWordFirstMatch(s, v)
	if !exists {
		return w //nothing to change
	}
	// We split the slice into 3 (possibly empty parts), with the middle being replacement
	partLeft, partRight := v[:index], v[index+len(s):]
	return ReduceRawWord(ConcatRawWord(partLeft, ConcatRawWord(replacement, partRight)))
}

// Treated as immutable
type RewritingSystem struct {
	LHS []RawWord //subwords to be replaced
	RHS []RawWord //replacements
}

// Concurrency is planned for particularly large rewriting systems
func (R RewritingSystem) Rewrite(w RawWord) RawWord {
	return w
}
