package presentation

// Slice of (indexed) generator and exponent pairs
// By convention, generator indices are nonnegative integers. The library does not strictly enforce this, but all built-in constructions only use nonnegative indices.
// We index from 0 because computers can't count right
// Can't be used in WordSets; use NewWord
type RawWord [][2]int

// This struct is treated as immutable
// id permits set-like behavior in word presentations
type Word struct {
	seq RawWord //seq is short for sequence
	id  string  //is always equal to WordID(Word.slice)
}

// Constructor for a new Word based on a RawWord
// Most functions on Words call NewWord on the output of its corresponding RawWord version of the function
func NewWord(w [][2]int) Word {
	return Word{
		seq: w,
		id:  WordID(w),
	}
}

func EmptyRawWord() RawWord { return RawWord{} }
func EmptyWord() Word       { return NewWord(EmptyRawWord()) }

func Len(w Word) int { return len(w.seq) } //length of the word

func ConcatRawWord(a, b RawWord) RawWord { return append(append(RawWord{}, a...), b...) } //double appends for immutability
func ConcatWord(v, w Word) Word          { return NewWord(ConcatRawWord(v.seq, w.seq)) }

// checks if two RawWords are equal
func EqualRawWord(u, v RawWord) bool {
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

// checks if two Words are equal by comparing their unique IDs
func EqualWord(u, v Word) bool {
	return u.id == v.id
}

// invert a RawWord
func InvRawWord(w RawWord) RawWord {
	n := len(w)
	v := RawWord{}
	for i := 1; i <= n; i++ { //list comprehensions at home
		v = append(v, [2]int{w[n-i][0], -w[n-i][1]})
	}
	return v
}

// invert a Word
func InvWord(w Word) Word {
	return NewWord(InvRawWord(w.seq))
}

// Free reduction of a RawWord
func ReduceRawWord(w RawWord) RawWord {
	r := make(RawWord, 0, len(w)) //r stands for reduced
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
	return NewWord(ReduceRawWord(w.seq))
}

// Cyclic Reduction
func CyclicReduceRawWord(w RawWord) RawWord {
	r := ReduceRawWord(w) //r for reduced
	lim := len(r)
	for i := 0; i < lim; i++ {
		s := r[i]
		last := r[lim - 1]
		if s[0] == last[0] { //conjugate by a power of s[0]
			r[i][1] += last[1]
			r = r[:len(r) - 1]
			lim--
		} else {
			break //already cyclically reduced
		}
	}
	c := make(RawWord, 0, len(r))
	for _, s := range r {
		if s[1] != 0 { //remove 0 exponents
			c = append(c, s)
		}
	}
	return c
}

func CyclicReduceWord(w Word) Word {
	c := CyclicReduceRawWord(w.seq)
	return NewWord(c)
}

// checks if self is a subword of other
func IsSubRawWord(self, other RawWord) bool {
	sub := ReduceRawWord(self)
	whole := ReduceRawWord(other)
	if len(whole) < len(sub) {
		return false
	}
	for i := 0; i+len(sub) <= len(whole); i++ {
		if EqualRawWord(sub, whole[i:i+len(sub)]) {
			return true //match found
		}
	}
	return false //all subwords don't match
}

func IsSubWord(self, other Word) bool {
	return IsSubRawWord(self.seq, other.seq)
}

// ShortLexRawWord reports whether a < b in shortlex order.
func ShortLexRawWord(a, b RawWord) bool {
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
	return ShortLexRawWord(a.seq, b.seq)
}

// find highest generator index in a Word w
func MaxGenRawWord(w RawWord) int {
	gens := 0
	for _, u := range w {
		if u[0] > gens {
			gens = u[0]
		}
	}
	return gens
}

func MaxGenWord(w Word) int {
	return MaxGenRawWord(w.seq)
}

// reduction to shortlex order that ignores commutativty used for abelian groups only
// second argument should be the largest generator index in w (any generator index larger than gens will result in a panic so this function is not exported!)
// GroupPresentation functions use G.gen for gens so not to waste resources on an extra loop
func abelianReduceRawWord(w RawWord, gens int) RawWord {
	exps := make([]int, gens)
	for _, u := range w {
		exps[u[0]] += u[1]
	}
	reduced := make(RawWord, 0, gens)
	for i := range exps {
		if exps[i] != 0 {
			reduced = append(reduced, [2]int{i, exps[i]})
		}
	}
	return reduced
}

func abelianReduceWord(w Word, gens int) Word {
	return NewWord(abelianReduceRawWord(w.seq, gens))
}

// abelianReduce but safe
func AbelianReduceRawWord(w RawWord) RawWord {
	gens := MaxGenRawWord(w)
	return abelianReduceRawWord(w, gens)
}

func AbelianReduceWord(w Word) Word {
	return NewWord(AbelianReduceRawWord(w.seq))
}
