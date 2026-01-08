package presentation

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

//helper to copy class maps definied here without mutating them
func addClasses(oldClasses, newClasses map[Class]struct{}) map[Class]struct{} {
	for c := range newClasses { oldClasses[c] = struct{}{} }
	return oldClasses
}

var freeGroupClasses = map[Class]struct{}{
	Free: {},
}

var oneRelatorGroupClasses = map[Class]struct{}{
	OneRelator: {},
}

var cyclcGroupClasses = map[Class]struct{}{
	Cyclic: {},
	Abelian: {},
}

var abelianGroupClasses = map[Class]struct{}{
	Abelian: {},
}