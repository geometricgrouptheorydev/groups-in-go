package presentation

import "maps"

//helper function to (shallow) copy maps
func copyMap[K comparable, V any](m map[K]V) map[K]V {
	copy := make(map[K]V, len(m))
	maps.Copy(copy, m)
	return copy
}

//compare maps
func equalMaps[K comparable, V any](m1, m2 map[K]V) bool {
	if len(m1) != len(m2) {
		return false //different lengths
	}
	for key := range m1 {
		_, ok := m2[key]
		if !ok { 
			return false //A had an id B does not
		}
	}
	return true
}