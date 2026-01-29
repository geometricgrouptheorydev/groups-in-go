package presentation

import "maps"

//helper function to (shallow) copy maps
func copyMap[K comparable, V any](m map[K]V) map[K]V {
	copy := make(map[K]V, len(m))
	maps.Copy(copy, m)
	return copy
}

//calculates absolute value of an int
func abs(x int) int {
	if x < 0 { return -x } else { return x }
}

//calculates sign of na int
func sign(x int) int {
	if x < 0 { return -1 } else if x > 0 { return 1 } else { return 0 }
}