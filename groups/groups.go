package groups

//We set up the Group interface, which will let users use groups formed from outisde libraries
//Group axioms to be verified mathematically
type Group[T any] interface {
	Mu(T, T) T //mu is a common greek letter to represent binary operations in algebra
	Id() T //Identity, common math abbreviation
	Inv(T) T //Inverse
}

//conjugates x by y
func Conj[T any](G Group[T], x T, y T) T {
	return G.Mu(G.Mu(G.Inv(y), x), y)
}

//takes the nth power of x in G
func Pow[T any](G Group[T], x T, n int) T {
	result := G.Id()
	switch {
	case n > 0:
		for i := 0; i < n; i++ {
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
func Comm[T any](G Group[T], x T, y T) T {
	return Prod(G, []T{x, y, G.Inv(x), G.Inv(y)})
}