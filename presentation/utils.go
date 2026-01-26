package presentation

import "maps"

//helper function to (shallow) copy maps
func copyMap[K comparable, V any](m map[K]V) map[K]V {
	copy := make(map[K]V, len(m))
	maps.Copy(copy, m)
	return copy
}