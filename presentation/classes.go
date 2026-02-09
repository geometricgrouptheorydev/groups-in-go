package presentation

import "fmt"

//here we have all the supported group classes so far
type Class string

const (
	Trivial Class = "trivial"
	Free Class = "free"
	FreeAbelian Class = "free_abelian"
	OneRelator Class = "one_relator"
	Abelian Class = "abelian"
	Cyclic Class = "cyclic"
	Finite Class = "finite"
	Dehn Class = "dehn"
)

//helper to copy class maps definied here without mutating them. 
//fallback to reset group classes upon error to be added
func (G *GroupPresentation) addClasses(newClasses map[Class]bool) error {
	for c := range newClasses { 
		if _, ok := G.classes[c]; !ok {
			G.classes[c] = newClasses[c] 
		} else if G.classes[c] != newClasses[c] {
			return fmt.Errorf("error: internal truth value conflict for group class %v", c)
		}
	}
	return nil
}

//for manually adding a property, positive or negative, but warning if that changed another value
func (G * GroupPresentation) AddClass(c Class, t bool) error {	
		oldVal, ok := G.classes[c]
		G.classes[c] = t 
		if ok && oldVal != t {
			return fmt.Errorf("error: internal truth value changed for group class %v", c)
		} else {
			return nil
		}
}

//this is for group properties that depend on a particular presentation rather than being 
func (G *GroupPresentation) RemoveClass(c Class) {
	delete(G.classes, c)
}


//below are internal arguments used in addClasses calls elsewhere
//not to be mutated

//trivial group
var trivialGroupClasses = map[Class]bool{
	Trivial: true,
	Free: true,
	FreeAbelian: true,
	OneRelator: false,
	Abelian: true,
	Cyclic: true,
	Finite: true,
}

//free on one generator
var freeGroupOneGeneratorClasses = map[Class]bool{
	Trivial: false,
	Free: true,
	FreeAbelian: true,
	OneRelator: false,
	Abelian: true,
	Cyclic: true,
	Finite: false,
}

//free with multiple generators
var freeGroupMultipleGeneratorClasses = map[Class]bool{
	Trivial: false,
	Free: true,
	FreeAbelian: false,
	OneRelator: false,
	Abelian: false,
	Cyclic: false,
	Finite: false,
}

//free abelian with 3+ generators
var freeAbelianGroupMultipleGeneratorClasses = map[Class]bool{
	Trivial: false,
	Free: false,
	FreeAbelian: true,
	OneRelator: false,
	Abelian: true,
	Cyclic: false,
	Finite: false,
}

//free abelian with 2 generators
var freeAbelianGroup2GeneratorClasses = map[Class]bool{
	Trivial: false,
	Free: false,
	FreeAbelian: true,
	OneRelator: true,
	Abelian: true,
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

//abelian groups
var abelianGroupClasses = map[Class]bool{
	Abelian: true,
}