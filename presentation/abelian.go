package presentation

// CheckCommutativityRelators reports whether all [x_i, x_j] are in rel (so abelian) and whether there are only and all [x_i,x_j] relators (so free abelian)
// WARNING: false returns do not automatically mean that the group is not abelian/free abelian
// This check is not in initAddProperties because of its time complexity of O(n^2m) which will make it slower for larger presentations
// This method also updates G's classes accordingly via addClasses, returning an error if there is one
func (G *GroupPresentation) CheckCommutativityRelators() (bool, bool, error) {
	//we check if these properties are already recorded so not to waste time (this is O(1))
	val, ok := G.classes[FreeAbelian]
	if val && ok {
		return true, true, nil
	} else if ok && !val {
		val2, ok2 := G.classes[Abelian]
		if val2 && ok2 {
			return true, false, nil
		} else if ok2 && !val2 {
			return false, false, nil
		}
	}

	//tackling trivial cases in O(1) time
	//classes already added in construction of these groups
	switch G.gen {
	case 0:
		return true, true, nil
	case 1:
		if len(G.rel) == 0 {
			return true, true, nil
		} else {
			return true, false, nil //useless relations are already discared with in NewGroupPresentation so we have a finite cyclic group here
		}
	}

	onlyCommutativityRelators := true
	hasAllCommutativityRelators := true
	foundCount := 0 //counts how many commutativity relators we found. if this ends up being less than len(G.rel), then we know there are non-commutativity relators
	for i := range G.gen {
		for j := range i {
			found := false
			target1 := RawWord{{i, 1}, {j, 1}, {i, -1}, {j, -1}}
			target2 := RawWord{{i, -1}, {j, -1}, {i, 1}, {j, 1}}
			for _, r := range G.rel {
				if len(r.seq) != 4 {
					onlyCommutativityRelators = false //r is not a commutativity relator for sure
					continue
				} else if EqualRawWord(r.seq, target1) || EqualRawWord(r.seq, target2) {
					found = true
					foundCount++
					break
				}
			}
			if !found {
				hasAllCommutativityRelators = false
				onlyCommutativityRelators = false
			}
		}
	}
	if foundCount < len(G.rel) {
		onlyCommutativityRelators = false
	}

	var err error
	if onlyCommutativityRelators {
		if G.gen == 2 {
			err = G.addClasses(freeAbelianGroup2GeneratorClasses)
		} else {
			err = G.addClasses(freeAbelianGroupMultipleGeneratorClasses)
		}
	} else if hasAllCommutativityRelators {
		err = G.addClasses(abelianGroupClasses)
	}
	return hasAllCommutativityRelators, onlyCommutativityRelators, err
}

// creating a new free abelian group
func NewFreeAbelianGroup(rank int) (*GroupPresentation, error) {
	classes := freeAbelianGroupMultipleGeneratorClasses //we'll change this if needed below
	if rank < 0 {
		return nil, ErrInvalidNumGenerators
	} else if rank == 0 {
		return NewFreeGroup(0) //return trivial group
	} else if rank == 1 {
		return NewFreeGroup(1) //return infinite cyclic group
	} else if rank == 2 {
		classes = freeAbelianGroup2GeneratorClasses
	}

	G := &GroupPresentation{
		gen:     rank,
		rel:     make(WordSet), //we add the relators in a double loop below
		classes: classes,
	}
	for i := range rank {
		for j := range i {
			r := NewWord(RawWord{{i, -1}, {j, -1}, {i, 1}, {j, 1}})
			G.rel.Add(r)
		}
	}
	return G, nil
}
