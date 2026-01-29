package groups

//We set up the Group interface, which will let users use groups formed from outisde libraries
//Group axioms to be verified mathematically
//Consistency between the methods below to be verified mathematically too
type Group[T any] interface {
	//mu is a common greek letter to represent binary operations in algebra
	Mu(T, T) T 
	//Identity, common math abbreviation
	Id() T 
	//Inverse
	Inv(T) T 
	//as we're allowing T be be uncomparable for Go standards, so we need to define our own equality method.
	//WARNING: Equal(x,y) = false does not guarantee x and y to be not equal in the group due to the word problem being unsolvable in general
	//but Equal(x,y) = true should always guarantee x = y in the group
	Equal(T, T) bool 
}

//conjugates x by y. we take the y^-1xy convention
func Conj[T any](G Group[T], x T, y T) T {
	return G.Mu(G.Mu(G.Inv(y), x), y)
}

//takes the nth power of x in G
func Pow[T any](G Group[T], x T, n int) T {
	result := G.Id()
	switch {
	case n > 0:
		for range n {
			result = G.Mu(result, x)
		}
	case n < 0:
		return Pow(G, G.Inv(x), -n)
	}
	return result
}

//Products of more than two elements at once
func Prod[T any](G Group[T], elem []T) T {
	result := G.Id()
	for _, x := range elem {
		G.Mu(result, x)
	}
	return result
}

//Commutator of x and y
//we take the x^-1y^-1xy convention
func Comm[T any](G Group[T], x T, y T) T {
	return Prod(G, []T{G.Inv(x), G.Inv(y), x, y})
}