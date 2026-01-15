package presentation

// Slice of (indexed) generator and exponent pairs
// By convention, generator indices are nonnegative integers. The library does not strictly enforce this, but all built-in constructions only use nonnegative indices.
// We index from 0 because computers can't count right
type WordSlice [][2]int

// This struct is treated as immutable
// Word.id permits set-like behavior in word presentations
type Word struct {
	slice WordSlice
	id   string //is always equal to WordID(Word.slice)
}

//Constructor for a new Word based on a WordSlice
//Most functions on Words call NewWord on the output of its corresponding WordSlice version of the function
func NewWord(w WordSlice) Word {
	return Word{
		slice: w,
		id: WordID(w),
	}
}

func EmptyWordSlice() WordSlice { return WordSlice{} }
func EmptyWord() Word {return NewWord(EmptyWordSlice())}

func ConcatWordSlice(a, b WordSlice) WordSlice { return append(append(WordSlice{}, a...), b...) } //double appends for immutability
func ConcatWord(v, w Word) Word { return NewWord(ConcatWordSlice(v.slice, w.slice)) }

// checks if two WordSlices are equal 
func EqualWordSlice(u, v WordSlice) bool {
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

//checks if two Words are equal by comparing their unique IDs
func EqualWord(u, v Word) bool {
	return u.id == v.id
}

//invert a WordSlice
func InvWordSlice(w WordSlice) WordSlice {
	n := len(w)
	v := WordSlice{}
	for i := 1; i <= n; i++ { //list comprehensions at home
		v = append(v, [2]int{w[n-i][0], -w[n-i][1]})
	}
	return v
}

//invert a Word
func InvWord(w Word) Word {
	return NewWord(InvWordSlice(w.slice))
}

//Free reduction of a WordSlice
func ReduceWordSlice(w WordSlice) WordSlice {
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

func ReduceWord(w Word) Word {
	return NewWord(ReduceWordSlice(w.slice))
}

// checks if self is a subword of other
func IsSubWordSlice(self, other WordSlice) bool {
	sub := ReduceWordSlice(self)
	whole := ReduceWordSlice(other)
	if len(whole) < len(sub) {
		return false
	}
	for i := 0; i+len(sub) <= len(whole); i++ {
		if EqualWordSlice(sub, whole[i:i+len(sub)]) {
			return true //match found
		}
	}
	return false //all subwords don't match
}

func IsSubWord(self, other Word) bool {
	return IsSubWordSlice(self.slice, other.slice) 
}

//ShortLexWordSlice reports whether a < b in shortlex order.
func ShortLexWordSlice(a, b WordSlice) bool {
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

func ShortLexWord(a, b Word) bool {
	return ShortLexWordSlice(a.slice, b.slice)
}

//find highest generator index in a Word w
func MaxGenWordSlice(w WordSlice) int {
	gens := 0
	for _, u := range w {
		if u[0] > gens {
			gens = u[0]
		}
	}
	return gens
}

func MaxGenWord(w Word) int {
	return MaxGenWordSlice(w.slice)
}

// reduction to shortlex order that ignores commutativty used for abelian groups only
// second argument should be the largest generator index in w (any generator index larger than gens will result in a panic so this function is not exported!)
// GroupPresentation functions use G.gen for gens so not to waste resources on an extra loop
func abelianReduceWordSlice(w WordSlice, gens int) WordSlice {
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

func abelianReduceWord(w Word, gens int) Word {
	return NewWord(abelianReduceWordSlice(w.slice, gens))
}

// abelianReduce but safe
func AbelianReduceWordSlice(w WordSlice) WordSlice {
	gens := MaxGenWordSlice(w)
	return abelianReduceWordSlice(w, gens)
}

func AbelianReduceWord(w Word) Word {
	return NewWord(AbelianReduceWordSlice(w.slice))
}
