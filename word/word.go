package word

import "sync"

// We give each Word encoutered a unique pointer via this thread-safe map.
var WordCache sync.Map

// Word represents an element of a free group using a tree-based structure.
// This allows for compact representation of large powers and efficient 
// boundary reduction.
type Word struct {
	// gen determines the node type:
	// >= 0: A leaf node representing the generator at this index in the group's generating set.
	// -1  : A product node representing (left * right)^exp.
	// -2  : The identity element (empty word).
	gen int

	// exp is the exponent. For leaves, it can be negative (inverses). For products, it must be positive (for uniqueness purposes)
	// For product nodes, it allows compact repetition like (ab^2)^1000.
	exp int

	// Tree pointers for product nodes. Both are nil for leaf and identity nodes.
	left  *Word
	right *Word

	// len is the total word length (sum of absolute exponents of all generators).
	len int

	// Metadata for O(1) free reduction checks at the boundaries.
	startGen int
	startExp int
	endGen   int
	endExp   int
}