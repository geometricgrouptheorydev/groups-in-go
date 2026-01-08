package presentation

func NewFreeGroup(rank int) (GroupPresentation, error){
	if rank < 0 {
		return GroupPresentation{}, ErrInvalidNumGenerators
	}
	G := GroupPresentation{
		gen: rank, 
		rel: make([]Word, 0),
		classes: addClasses(make(map[Class]struct{}), freeGroupClasses),
		}
	switch rank {
	case 0:
		G.classes[Trivial] = struct{}{}
	case 1:
		G.classes[Abelian] = struct{}{}
		G.classes[Cyclic] = struct{}{}
	}
	return G, nil
}