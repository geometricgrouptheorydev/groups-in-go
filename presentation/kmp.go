package presentation

// Here are some helper functions that use the KMP prefix function on slices
// This will allow multiple functions to be have O(n) time complexity rather than O(n^2)

// For each i := range w finds the length of the longest prefix of w[i] that is also a suffix
func KMPPrefixFunction[T comparable](w []T) []int {
	pi := make([]int, len(w)) // pi[0] is always 0 so we won't change that in the loop
	for i := 1; i < len(w); i++ {
		// Initialize j with the value from the previous position
        j := pi[i - 1]
        // Continue updating j until a match is found or j becomes 0
        for j > 0 && w[i] != w[j]{
            j = pi[j - 1]
		}
        // If a match is found, increment the length of the common prefix/suffix
        if w[i] == w[j] {
            j++
		}
        // Update the Prefix Function value for the current position
        pi[i] = j
	}
	return pi
}

// Returns the indices where each occurence of sub appears in whole
func KMPSearchSubword[T comparable](sub, whole []T) []int {
	// Take care of the trivial cases
	if len(sub) == 0 {
		everywhere := make([]int, len(whole))
		for i := range whole {
			everywhere[i] = i
		}
		return everywhere
	} else if len(sub) > len(whole) {
		nowhere := make([]int, 0)
		return nowhere
	}
	occurences := make([]int, 0)
	pi := KMPPrefixFunction(sub)
	j := 0 // current match length
	for i := range whole { 
		for j > 0 && whole[i] != sub[j] {
			j = pi[j-1] 
		} 
		if whole[i] == sub[j] { 
			j++ 
		} 
		if j == len(sub) { 
			occurences = append(occurences, i - j + 1) 	 
			j = pi[j-1]
		} 
	}
	return occurences
}

// returns true iff sub is in whole
func KMPCheckSubword[T comparable](sub, whole []T) bool {
	// Take care of the trivial cases
	if len(sub) == 0 {
		return true
	} else if len(sub) > len(whole) {
		return false
	}
	pi := KMPPrefixFunction(sub)
	j := 0 // current match length
	for i := range whole { 
		for j > 0 && whole[i] != sub[j] {
			j = pi[j-1] 
		} 
		if whole[i] == sub[j] { 
			j++ 
		} 
		if j == len(sub) { 
			return true
		} 
	}
	return false
}

// lists w as a subslice repeated several times, ordered from largest to lowest
// the trivial (w, 1) is always returned
func KMPFindRepeats[T comparable](w []T) []struct{
	Sub []T
	Reps int
} {
repeats := []struct{
	Sub []T
	Reps int
}{
	{
		Sub: w,
		Reps: 1,
	},
}
pi := KMPPrefixFunction(w)
n := len(w)
for n > 0 {
	// Initialize to the longest known prefix that is also a suffix (also known as a border)
	// if w is v repeated, then v is repeated
	// we check each border in order of decreasing length in the next loop iterations (borders of borders are borders)
	n = pi[n-1]
	// k is the number of positions between the start of the prefix and the start of its repetition as a suffix
	// therefore k is the period of the slice
	k := len(w) - n
	// len(w) needs to be a multiple of k to have a chance of w being w[:n] repeated
	// k also needs to be smaller than w lest we don't have a period at all!
	// this suffices because w[i] = w[i + k]
	if k < len(w) && len(w) % k == 0 {
		rpt := struct{
			Sub []T
			Reps int
		}{
			Sub: w[:n],
			Reps: len(w) / n,
		}
		repeats = append(repeats, rpt)
	}
}
return repeats
}

// Checks if w is a slice that is a repeated subslice
func KMPCheckRepeats[T comparable](w []T) bool {
pi := KMPPrefixFunction(w)
n := len(w)
for n > 0 {
	// Initialize to the longest known prefix that is also a suffix (also known as a border)
	// if w is v repeated, then v is repeated
	// we check each border in order of decreasing length in the next loop iterations (borders of borders are borders)
	n = pi[n-1]
	// k is the number of positions between the start of the prefix and the start of its repetition as a suffix
	// therefore k is the period of the slice
	k := len(w) - n
	// len(w) needs to be a multiple of k to have a chance of w being w[:n] repeated
	// k also needs to be smaller than w lest we don't have a period at all!
	// this suffices because w[i] = w[i + k]
	if k < len(w) && len(w) % k == 0 {
		return true
	}
}
return false
}

// A root of a word w is some subword v such that w = v^k for some positive k
// Such a root is deemed non-trivial if k >= 2
// By the Fine-Wilff theorem, for our purposes all words with a non-trivial root will have a smallest root that is a root of all other roots
// We call this the primitive root, the first output of this function
// The second output gives the k for that primitive root
// The third output is true exactly when the primitive root is non-trivial
func KMPFindPrimitiveRoot[T comparable](w []T) ([]T, int, bool){
	n := len(w)
	if n == 0 {
		return []T{}, 0, false
	}

	pi := KMPPrefixFunction(w)
	r := n - pi[n-1] //primitive root length

	// If r divides n and r < n, the word is (non-trivially) periodic 
	if n%r == 0 {
		exp := n/r
		return w[:r], exp, exp > 1
	} else {
		return w, 1, false
	}
}