package presentation

//CheckForCommutativityRelators reports whether all [x_i, x_j] are in rel.
//This check is not in initAddProperties because of its time complexity of O(n^2m) which will make it slower for larger presentations
//A more comprehensive check TBA
func (G GroupPresentation) CheckForCommutativityRelators() bool {
	if _, ok := G.classes[Abelian]; ok {
		return true
	}
    for i := 0; i < G.gen; i++ {
        for j := 0; j < i; j++ {
            found := false
            target1 := Word{{i, 1}, {j, 1}, {i, -1}, {j, -1}}
            target2 := Word{{i, -1}, {j, -1}, {i, 1}, {j, 1}}
            for _, r := range G.rel {
				if len(r) != 4 {
					continue
				} else if EqualWord(r, target1) || EqualWord(r, target2) {
                    found = true
                    break
                }
            }
            if !found {
                return false
            }
        }
    }
	G.addClasses(abelianGroupClasses)
    return true
}
