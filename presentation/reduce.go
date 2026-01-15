package presentation

//In this file, we'll have the word reduction methods based on the different group presenetation classes
//reductions ordered in levels of power

// the higher the priority the better the reduction algorithm for computation!
var reduceCLassPriority = []Class{Trivial, Cyclic, FreeAbelian, Abelian, Free, OneRelator}

func (G *GroupPresentation) Reduce(w Word) (Word, error) {
	err := G.IsValidWord(w) //we need to use the O(n) IsValidWord method because some cases panic on invalid words
	if err != nil {
		return EmptyWord(), err
	}
	for _, c := range reduceCLassPriority {
		if val, ok := G.classes[c]; val && ok {
			switch c {
			case Trivial:
				return EmptyWord(), nil
			case Cyclic:
				return G.handleReduceCyclic(w), nil
			case FreeAbelian:
				return G.handleReduceFreeAbelian(w), nil
			case Abelian:
				return G.handleReduceAbelian(w), nil
			case Free:
				return ReduceWord(w), nil //plain old word reduction
			case OneRelator:
				return G.handleReduceOneRelator(w), nil
			}
		}
	}
	return EmptyWord(), nil
}

// O(n)
func (G *GroupPresentation) handleReduceCyclic(w Word) Word {
	G.SimplifyCyclicPresentation()
	wordExp := ReduceWordSlice(w.word)[0][1]
	var rel Word
	for _, r := range G.rel {
		rel = r //extracting the only member of G.rel
		break
	}
	relExp := rel.word[0][1]
	newExp := wordExp % relExp
	if newExp < 0 {
		newExp = -newExp
	} else if newExp == 0 {
		return EmptyWord()
	}
	return NewWord(WordSlice{{0, newExp}})
}

// O(n)
func (G *GroupPresentation) handleReduceFreeAbelian(w Word) Word {
	return abelianReduceWord(w, G.gen)
}

func (G *GroupPresentation) handleReduceAbelian(w Word) Word {
	return w
}

func (G *GroupPresentation) handleReduceOneRelator(w Word) Word {
	return w
}
