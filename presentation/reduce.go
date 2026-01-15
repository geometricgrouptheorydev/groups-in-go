package presentation

//In this file, we'll have the word reduction methods based on the different group presenetation classes
//reductions ordered in levels of power

// the higher the priority the better the reduction algorithm for computation!
var reduceCLassPriority = []Class{Trivial, Cyclic, FreeAbelian, Abelian, Free, OneRelator}

func (G GroupPresentation) Reduce(w WordSlice) (WordSlice, error) {
	err := G.IsValidWord(w) //we need to use the O(n) IsValidWord method because some cases panic on invalid words
	if err != nil {
		return WordSlice{}, err
	}
	for _, c := range reduceCLassPriority {
		if val, ok := G.classes[c]; val && ok {
			switch c {
			case Trivial:
				return WordSlice{}, nil
			case Cyclic:
				return G.handleReduceCyclic(w), nil
			case FreeAbelian:
				return G.handleReduceFreeAbelian(w), nil
			case Abelian:
				return G.handleReduceAbelian(w), nil
			case Free:
				return Reduce(w), nil //plain old word reduction
			case OneRelator:
				return G.handleReduceOneRelator(w), nil
			}
		}
	}
	return WordSlice{}, nil
}

// O(n)
func (G GroupPresentation) handleReduceCyclic(w WordSlice) WordSlice {
	G.SimplifyCyclicPresentation()
	wordExp := Reduce(w)[0][1]
	relExp := G.rel[0][0][1]
	newExp := wordExp % relExp
	if newExp < 0 {
		newExp = -newExp
	} else if newExp == 0 {
		return WordSlice{}
	}
	return WordSlice{{0, newExp}}
}

// O(n)
func (G GroupPresentation) handleReduceFreeAbelian(w WordSlice) WordSlice {
	return abelianReduce(w, G.gen)
}

func (G GroupPresentation) handleReduceAbelian(w WordSlice) WordSlice {
	return w
}

func (G GroupPresentation) handleReduceOneRelator(w WordSlice) WordSlice {
	return w
}
