package presentation

//In this file, we'll have the word reduction methods based on the different group presenetation classes
//reductions ordered in levels of power

//the higher the priority the better the reduction algorithm for computation!
var reduceCLassPriority = []Class{Trivial, Cyclic, FreeAbelian, Abelian, Free, OneRelator}

var reduceHandles = map[Class]func(GroupPresentation, Word) Word{
	Trivial: GroupPresentation.handleReduceTrivial,
	Cyclic: GroupPresentation.handleReduceCyclic,
	Abelian: GroupPresentation.handleReduceAbelian,
	Free: GroupPresentation.handleReduceFree,
	OneRelator: GroupPresentation.handleReduceOneRelator,
}

func (G GroupPresentation) Reduce(w Word) Word {
	for _, c := range reduceCLassPriority {
		if val, ok := G.classes[c]; val && ok {
			switch c{
			case Trivial:
				return G.handleReduceTrivial(w)
			case Cyclic:
				return G.handleReduceCyclic(w)
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

func (G GroupPresentation) handleReduceTrivial(w Word) Word{
	return Word{}
}

func (G GroupPresentation) handleReduceCyclic(w Word) Word{
	return Word{}
}

func (G GroupPresentation) handleReduceFreeAbelian(w Word) Word{
	return Word{}
}

func (G GroupPresentation) handleReduceAbelian(w Word) Word{
	return Word{}
}

func (G GroupPresentation) handleReduceFree(w Word) Word {
	return Word{}
}

func (G GroupPresentation) handleReduceOneRelator(w Word) Word{
	return Word{}
}