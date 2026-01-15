package presentation

//emulating a mathematical set of relations with a map[string]Word with key always the word.id
type WordSet map[string]Word 

//build a WordSet from a []Word
//if len(l) == 0 then this is equivalent to make(WordSet)
func NewWordSet(l []Word) WordSet {
	A := make(WordSet)
	for _, w := range l {
		A.Add(w)
	}
	return A
}

//checks if two WordSets are equal
func EqualWordSet(A, B WordSet) bool {
	if len(A) != len(B) {
		return false //different lengths
	}
	for id := range A {
		_, ok := B[id]
		if !ok { 
			return false //A had an id B does not
		}
	}
	return true
}

//check whether a word is in a WordSet
func (A WordSet) Has(w Word) bool {
	_, ok := A[w.id]
	return ok
}

//adds a Word to a Wordset
func (A WordSet) Add(w Word) {
	A[w.id] = w
}


//removes a Word from a Wordset
func (A WordSet) Remove(w Word) {
	delete(A, w.id)
}

//returns a new WordSet that's the union of the inputted two
func Union(A, B WordSet) WordSet {
	C := make(WordSet, len(A) + len(B))
	for _, w := range A {
		C.Add(w)
	}
	for _, w := range B {
		C.Add(w)
	}
	return C
}