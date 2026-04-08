package presentation_test

import (
	"testing"

	p "github.com/geometricgrouptheorydev/groups-in-go/presentation"
)

func TestReplaceRawSubWordFirstMatch(t *testing.T) {
	tests := []struct {
		name  string
		whole RawWord
		sub   RawWord
		replacement RawWord
		want  RawWord
	}{
		{
			name:  "subword",
			whole: RawWord{{1, 2}, {2, -3}, {6, -7}, {3, 1}, {4, 7}, {6, -7}, {3, 1}},
			sub:   RawWord{{6, -7}, {3, 1}},
			replacement: RawWord{{2, 3},{4, -1}},
			want:  RawWord{{1, 2}, {4, 6}, {6, -7}, {3, 1}},
		},
		{
			name:  "not subword",
			whole: RawWord{{1, 2}, {2, -3}, {6, -7}, {3, 1}, {4, 7}},
			sub:   RawWord{{6, -7}, {2, -3}},
			replacement: RawWord{},
			want:  RawWord{{1, 2}, {2, -3}, {6, -7}, {3, 1}, {4, 7}},
		},
		{
			name: "empty subword",
			whole: RawWord{{4,6},{9,2},{3,-4}},
			sub: RawWord{},
			replacement: RawWord{{1,1}},
			want: RawWord{{1,1},{4,6},{9,2},{3,-4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.whole.ReplaceRawSubWordFirstMatch(tt.sub, tt.replacement)
			if !p.EqualRawWord(tt.want, got) {
				t.Fatalf("%v.ReplaceSubWordFirstMatch(%v, %v) = %v, expected %v", tt.whole, tt.sub, tt.replacement, got, tt.want)
			}
		})
	}
}