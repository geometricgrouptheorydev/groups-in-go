package presentation

import "errors"

// helper gcd function using the Euclidian algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	if a < 0 {
		return -a
	} //a%b keeps the sign of a in Go but the gcd should be positive
	return a
}

// helper gcd for for more than 2 numbers
// we use the identity gcd(a,b,c) = gcd(gcd(a,b),c)
func MultiGCD(nums []int) int {
	gcd := 0
	previous := 0
	for _, x := range nums {
		if gcd == 0 {
			gcd = x
			previous = x
		} else {
			gcd = GCD(previous, x)
		}
	}
	return gcd
}

func (G *GroupPresentation) SimplifyCyclicPresentation() error {
	if val, ok := G.classes[Cyclic]; !ok || !val {
		return errors.New("This Group is not cyclic")
	}
	exps := make([]int, 0, len(G.rel)) //we'll extract the exponent of each relation
	//each relation is already in the form Word{{0,n}} due to the word reduction in NewGroupPresentation
	for _, r := range G.rel {
		exps = append(exps, r.word[0][1])
	}
	combinedRel := NewWord(WordSlice{{0, MultiGCD(exps)}})
	G.rel[combinedRel.id] = combinedRel
	return nil
}
