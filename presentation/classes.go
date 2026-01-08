package presentation

//here we have all the supported group classes so far
type Class string

const (
	Trivial Class = "trivial"
	Free Class = "free"
	OneRelator Class = "one_relator"
	Abelian Class = "abelian"
	Cyclic Class = "cyclic"
)

var supportedClasses = map[Class]struct{}{
	Trivial: {},
	Free: {},
	OneRelator: {},
	Abelian: {},
	Cyclic: {},
}

var FreeGroupClasses = map[Class]struct{}{
	Free: {},
}