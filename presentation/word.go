package presentation

//Slice of (indexed) generator and exponent pairs
//By convention, generator indices are nonnegative integers. The library does not strictly enforce this, but all built-in constructions only use nonnegative indices.
//We index from 0 because computers can't count right
type Word [][2]int 

func EmptyWord() Word            { return Word{} }
func Single(gen, exp int) Word   { return Word{{gen, exp}} }
func Concat(a, b Word) Word          { return append(append(Word{}, a...), b...) } //double appends for immutability

//checks if two words are equal
func EqualWord(u, v Word) bool {
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

func Inv(w Word) Word {
	n := len(w)
	v := Word{}
	for i := 1; i <= n ; i++ { //list comprehensions at home
		v = append(v, [2]int{w[n-i][0], -w[n-i][1]})
	}
	return v
}

func Reduce(w Word) Word {
	r := Word{} //r stands for reversed
	for _, s := range w {
		if len(r) > 0 && r[len(r)-1][0] == s[0] {
			r[len(r)-1] =  [2]int{s[0], s[1] + r[len(r)-1][1]}
		} else {
			r = append(r, s)
		}
	}
	return r
}

//checks if self is a subword of other
func IsSubword(self, other Word) bool {
	sub :=  Reduce(self)
	whole := Reduce(other)
	if len(whole) < len(sub) { return false }
	for i := range whole {
		if EqualWord(sub, whole[i: i + len(sub)]) {
			return true //match found
		}
	}
	return false //all subwords don't match
}