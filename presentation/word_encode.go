package presentation

import (
	"fmt"
	"strconv"
	"strings"
)

// encodes a Word into a canonical string form
// A word is encoded as g:e,g:e,... (for readability) where g is a generator index and e is an integer exponent, e.g. 0:2,1:-3 for x_0^2 x_1^{-3}
func WordID(w [][2]int) string {
	var b strings.Builder
	for i, p := range w {
		if i > 0 {
			b.WriteByte(',') // separator between pairs
		}
		b.WriteString(strconv.Itoa(p[0]))
		b.WriteByte(':')
		b.WriteString(strconv.Itoa(p[1]))
	}
	return b.String()
}

func ParseWordID(id string) (WordSlice, error) {
	if id == "" {
		return nil, nil // empty word
	}
	parts := strings.Split(id, ",")
	w := make(WordSlice, 0, len(parts))
	for _, part := range parts {
		sub := strings.Split(part, ":")
		if len(sub) != 2 {
			return nil, fmt.Errorf("invalid word id segment %v", part)
		}
		g, err := strconv.Atoi(sub[0])
		if err != nil {
			return nil, fmt.Errorf("invalid generator %v: %v", sub[0], err)
		}
		e, err := strconv.Atoi(sub[1])
		if err != nil {
			return nil, fmt.Errorf("invalid exponent %v: %v", sub[1], err)
		}
		w = append(w, [2]int{g, e})
	}
	return w, nil
}
