package presentation

//In this file, we'll have the word reduction methods based on the different group presenetation classes
//reductions ordered in levels of power

//the higher the priority the better the reduction algorithm for computation!
var reduceCLassPriority = []Class{Trivial, Cyclic, FreeAbelian, Abelian, Free, OneRelator}

func (G GroupPresentation) Reduce(w Word) Word {
	for _, c := range reduceCLassPriority {
		if val, ok := G.classes[c]; val && ok {
			switch c{
			case Trivial:
				return Word{}
			case Cyclic:
				return G.handleReduceCyclic(w)
			case FreeAbelian:
				return G.handleReduceFreeAbelian(w)
			case Abelian:
				return G.handleReduceAbelian(w)
			case Free:
				return G.handleReduceFree(w)
			case OneRelator:
				return G.handleReduceOneRelator(w)
			}
		}
	}
	return Word{}
}

//O(n)
func (G GroupPresentation) handleReduceCyclic(w Word) Word{
	G.SimplifyCyclicPresentation()
	wordExp := Reduce(w)[0][1]
	relExp := G.rel[0][0][1]
	newExp := wordExp % relExp
	if newExp < 0 {
		newExp = -newExp
	} else if newExp == 0 {
		return Word{}
	}
	return Word{{0, newExp}}
}

func (G GroupPresentation) handleReduceFreeAbelian(w Word) Word{
	return w
}

func (G GroupPresentation) handleReduceAbelian(w Word) Word{
	return w
}

func (G GroupPresentation) handleReduceFree(w Word) Word {
	return w
}

func (G GroupPresentation) handleReduceOneRelator(w Word) Word{
	return w
}