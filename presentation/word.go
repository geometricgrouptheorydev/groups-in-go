package presentation

// Slice of (indexed) generator and exponent pairs
// By convention, generator indices are nonnegative integers. The library does not strictly enforce this, but all built-in constructions only use nonnegative indices.
// We index from 0 because computers can't count right
type WordSlice [][2]int

func EmptyWord() WordSlice            { return WordSlice{} }
func Single(gen, exp int) WordSlice   { return WordSlice{{gen, exp}} }
func Concat(a, b WordSlice) WordSlice { return append(append(WordSlice{}, a...), b...) } //double appends for immutability

// checks if two words are equal
func EqualWord(u, v WordSlice) bool {
	if len(u) != len(v) {
		return false
	}
	for i := range u {
		if u[i][0] != v[i][0] || u[i][1] != v[i][1] {
			return false
		}
	}
	return true
}

func Inv(w WordSlice) WordSlice {
	n := len(w)
	v := WordSlice{}
	for i := 1; i <= n; i++ { //list comprehensions at home
		v = append(v, [2]int{w[n-i][0], -w[n-i][1]})
	}
	return v
}

func Reduce(w WordSlice) WordSlice {
	r := make(WordSlice, 0, len(w)) //r stands for reversed
	for _, s := range w {
		if len(r) > 0 && r[len(r)-1][1] == 0 {
			continue //ignore 0 exponents
		} else if len(r) > 0 && r[len(r)-1][0] == s[0] {
			if s[1]+r[len(r)-1][1] == 0 {
				r = r[:len(r)-1] //remove 0 exponent
			} else {
				r[len(r)-1] = [2]int{s[0], s[1] + r[len(r)-1][1]} //combine exponents
			}
		} else {
			r = append(r, s)
		}
	}
	return r
}

// checks if self is a subword of other
func IsSubword(self, other WordSlice) bool {
	sub := Reduce(self)
	whole := Reduce(other)
	if len(whole) < len(sub) {
		return false
	}
	for i := range whole {
		if EqualWord(sub, whole[i:i+len(sub)]) {
			return true //match found
		}
	}
	return false //all subwords don't match
}

// ShortLex reports whether a < b in shortlex order.
func ShortLex(a, b WordSlice) bool {
	if len(a) != len(b) {
		return len(a) < len(b)
	}
	// same length: lexicographic on (gen, exp)
	for i := range a {
		if a[i][0] != b[i][0] {
			return a[i][0] < b[i][0]
		}
		if a[i][1] != b[i][1] {
			return a[i][1] < b[i][1]
		}
	}
	return false //equal
}

// find higher generator index in a Word w
func MaxGen(w WordSlice) int {
	gens := 0
	for _, u := range w {
		if u[0] > gens {
			gens = u[0]
		}
	}
	return gens
}

// reduction to shortlex order that ignores commutativty used for abelian groups only
// second argument should be the largest generator index in w (any generator index larger than gens will result in a panic so this function is not exported!)
// GroupPresentation functions use G.gen for gens so not to waste resources on an extra loop
func abelianReduce(w WordSlice, gens int) WordSlice {
	exps := make([]int, gens)
	for _, u := range w {
		exps[u[0]] += u[1]
	}
	reduced := make(WordSlice, 0, gens)
	for i := range exps {
		if exps[i] != 0 {
			reduced = append(reduced, [2]int{i, exps[i]})
		}
	}
	return reduced
}

// abelianReduce but safe
func AbelianReduce(w WordSlice) WordSlice {
	gens := MaxGen(w)
	return abelianReduce(w, gens)
}
