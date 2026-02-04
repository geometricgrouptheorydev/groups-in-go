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

func PowRawWord(n int, w RawWord) RawWord {
	switch {
	case n > 0:
		pow := w
		for range n - 1 {
			pow = ConcatRawWord(pow, w)
		}
		return ReduceRawWord(pow)
	case n < 0:
		return PowRawWord(-n, InvRawWord(w))
	}
	return EmptyRawWord()
}

// checks if two RawWords are equal
func EqualRawWord(u, v RawWord) bool {
	return equalSlices(u, v)
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

// Conjugates w by a
func ConjugateRawWord(w RawWord, a RawWord) RawWord {
	return ReduceRawWord(ConcatRawWord(InvRawWord(a), ConcatRawWord(w, a)))
}

func ConjugateWord(w Word, a Word) Word {
	return NewWord(ConjugateRawWord(w.seq, a.seq))
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
// first output is the cyclically reduced word
// the second output is the RawWord we conjugate w by to get the cyclic reduction
func CyclicReduceRawWord(w RawWord) (RawWord, RawWord) {
	reduced := ReduceRawWord(w) //r for reduced
	conjugatedBy := make(RawWord, 0, len(reduced))
	lim := len(reduced)
	for i := 0; i < lim; i++ {
		s := reduced[i]
		last := reduced[lim - 1]
		if s[0] == last[0] { //conjugate by a power of s[0]
			reduced[i][1] += last[1]
			reduced = reduced[:len(reduced) - 1]
			conjugatedBy = append(conjugatedBy, last)
			lim--
		} else {
			break //already cyclically reduced
		}
	}
	cyclicallyReduced := make(RawWord, 0, len(reduced))
	for _, s := range reduced {
		if s[1] != 0 { //remove 0 exponents
			cyclicallyReduced = append(cyclicallyReduced, s)
		}
	}
	return cyclicallyReduced, reverseSlice(conjugatedBy)
}

func CyclicReduceWord(w Word) (Word, Word) {
	cyclicallyReduced, conjugatedBy := CyclicReduceRawWord(w.seq)
	return NewWord(cyclicallyReduced), NewWord((conjugatedBy))
}

// Helper for using KMP algorithms
// Returns the same word but unreduced with all exponents 1 or -1
func expandRawWord(w RawWord) RawWord{
	expSum := 0 // Sum of exponents to determine how much memory we need to allocate
	for _, u := range w {
		expSum += abs(u[1])
	}
	expanded := make(RawWord, 0, expSum)
	for _, u := range w {
		for range abs(u[1]) {
				expanded = append(expanded, [2]int{u[0], sign(u[1])})
		}
	}
	return expanded
}

// Checks if self is a subword of other in O(n) using the KMP prefix function
func IsSubRawWord(self, other RawWord) bool {
	sub := expandRawWord(self)
	whole := expandRawWord(other)
	return KMPCheckSubword(sub, whole)
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

//checks if w is the power of another subword
func CheckIfPowerRawWord(w RawWord) bool {
	// If w has a root, so does its cyclic reduction
	// For a cyclically reduced word, every root is a prefix of it
	r, _ := CyclicReduceRawWord(w)
	c := expandRawWord(r)
	return KMPCheckRepeats(c)
}

func CheckIfPowerWord(w Word) bool {
	return CheckIfPowerRawWord(w.seq)
}

//struct that tells the roots of a word and the number of repeated concatenation needed to get the original word
type rootsOfRawWord struct{
	root RawWord
	exp int
}

func FindRootsRawWord(w RawWord) []rootsOfRawWord{
	reduced, conjugatedBy := CyclicReduceRawWord(w)
	expanded := expandRawWord(reduced)
	reducedRoots := KMPFindRepeats(expanded)
	result := make([]rootsOfRawWord, len(reducedRoots))
	for i := range reducedRoots {
		result[i].exp = reducedRoots[i].Reps
		result[i].root = ConjugateRawWord(reducedRoots[i].Sub, conjugatedBy)
	}
	return result
}

type rootsOfWord struct{
	root Word
	exp int
}

// we don't use the RawWord version of the function as that would require another loop
func FindRootsWord(w Word) []rootsOfWord{
	reduced, conjugatedBy := CyclicReduceRawWord(w.seq)
	expanded := expandRawWord(reduced)
	reducedRoots := KMPFindRepeats(expanded)
	result := make([]rootsOfWord, len(reducedRoots))
	for i := range reducedRoots {
		result[i].exp = reducedRoots[i].Reps
		result[i].root = NewWord(ConjugateRawWord(reducedRoots[i].Sub, conjugatedBy))
	}
	return result
}

// A root of a word w is some subword v such that w = v^k for some positive k
// Such a root is deemed non-trivial if k >= 2
// By the Fine-Wilff theorem, for our purposes all words with a non-trivial root will have a smallest root that is a root of all other roots
// We call this the primitive root, the first output of this function
// The second output gives the k for that primitive root
// The third output is true exactly when the primitive root is non-trivial
func FindPrimitiveRootRawWord(w RawWord) (RawWord, int, bool)  {
	reduced, conj := CyclicReduceRawWord(w)
	root, exp, ok := KMPFindPrimitiveRoot(expandRawWord(reduced))
	if !ok {return reduced, 1, false}
	return ConjugateRawWord(ReduceRawWord(root), conj), exp, true
}

func FindPrimitiveRootWord(w Word) (Word, int, bool) {
	root, exp, ok := FindPrimitiveRootRawWord(w.seq)
	return NewWord(root), exp, ok
}