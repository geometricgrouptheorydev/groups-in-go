package presentation

import "errors"

var ErrInternalGroupClassTruth = errors.New("error: truth conflict for a group class")

//here we have all the supported group classes so far
type Class string

const (
	Trivial Class = "trivial"
	Free Class = "free"
	OneRelator Class = "one_relator"
	Abelian Class = "abelian"
	Cyclic Class = "cyclic"
	Finite Class = "finite"
)

var supportedClasses = map[Class]struct{}{
	Trivial: {},
	Free: {},
	OneRelator: {},
	Abelian: {},
	Cyclic: {},
	Finite: {},
}

//helper to copy class maps definied here without mutating them. 
//fallback to reset group classes upon error to be added
func (G GroupPresentation) addClasses(newClasses map[Class]bool) error {
	for c := range newClasses { 
		if _, ok := G.classes[c]; !ok {
			G.classes[c] = newClasses[c] 
		} else if G.classes[c] != newClasses[c] {
			return ErrInternalGroupClassTruth
		}
	}
	return nil
}

//trivial group
var trivialGroupClasses = map[Class]bool{
	Trivial: true,
	Free: true,
	OneRelator: false,
	Abelian: true,
	Cyclic: true,
	Finite: true,
}

//free on one generator
var freeGroupOneGeneratorClasses = map[Class]bool{
	Trivial: false,
	Free: true,
	OneRelator: false,
	Abelian: true,
	Cyclic: true,
	Finite: false,
}

//free with multiple generators
var freeGroupMultipleGeneratorClasses = map[Class]bool{
	Trivial: false,
	Free: true,
	OneRelator: false,
	Abelian: false,
	Cyclic: false,
	Finite: false,
}

//one relator groups
var oneRelatorGroupClasses = map[Class]bool{
	OneRelator: true,
}

//cyclic groups
var cyclicGroupClasses = map[Class]bool{
	Cyclic: true,
	Abelian: true,
}

var abelianGroupClasses = map[Class]bool{
	Abelian: true,
}