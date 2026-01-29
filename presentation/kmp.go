package presentation

// Here are some helper functions that use the KMP prefix function on slices
// This will allow multiple functions to be have O(n) time complexity rather than O(n^2) or On(n log n)

// For each i := range w finds the longest prefix of w[i] that is also a suffix
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