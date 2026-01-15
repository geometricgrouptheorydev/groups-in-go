package presentation

func NewFreeGroup(rank int) (GroupPresentation, error) {
	if rank < 0 {
		return GroupPresentation{}, ErrInvalidNumGenerators
	}
	G := GroupPresentation{
		gen:     rank,
		rel:     make(map[string]Word),
		classes: make(map[Class]bool),
	}
	switch rank {
	case 0:
		G.addClasses(trivialGroupClasses)
	case 1:
		G.addClasses(freeGroupOneGeneratorClasses)
	default:
		G.addClasses(freeGroupMultipleGeneratorClasses)
	}
	return G, nil
}
