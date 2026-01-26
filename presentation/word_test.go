package presentation_test

import (
	"testing"

	"github.com/geometricgrouptheorydev/groups-in-go/presentation"
)

type RawWord = presentation.RawWord

func TestEqualWord(t *testing.T) {
	tests := []struct {
		name   string
		first  RawWord
		second RawWord
		want   bool
	}{
		{
			name:   "unequal lengths",
			first:  RawWord{{0, 1}, {4, 3}},
			second: RawWord{{0, 1}, {4, 3}, {6, 7}},
			want:   false,
		},
		{
			name:   "different exponent",
			first:  RawWord{{0, 1}, {4, 3}, {6, 7}},
			second: RawWord{{0, 1}, {4, 3}, {6, 8}},
			want:   false,
		},
		{
			name:   "swapped",
			first:  RawWord{{0, 1}, {1, 2}},
			second: RawWord{{1, -2}, {0, -1}},
			want:   false,
		},
		{
			name:   "different generators",
			first:  RawWord{{0, 1}, {1, 2}},
			second: RawWord{{0, 1}, {2, 2}},
			want:   false,
		},
		{
			name:   "actually the same",
			first:  RawWord{{0, 1}, {1, 2}, {3, -4}, {4, -2}, {3, 5}},
			second: RawWord{{0, 1}, {1, 2}, {3, -4}, {4, -2}, {3, 5}},
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := presentation.EqualRawWord(tt.first, tt.second)
			if got != tt.want {
				t.Fatalf("EqualWord(%v, %v) = %v, want %v", tt.first, tt.second, got, tt.want)
			}
		})
	}
}

func TestInverse(t *testing.T) {
	tests := []struct {
		name string
		in   RawWord
		want RawWord
	}{
		{
			name: "empty",
			in:   RawWord{},
			want: RawWord{},
		},
		{
			name: "single letter",
			in:   RawWord{{0, 1}},
			want: RawWord{{0, -1}},
		},
		{
			name: "two letters",
			in:   RawWord{{0, 1}, {1, 2}},
			want: RawWord{{1, -2}, {0, -1}},
		},
		{
			name: "five letters",
			in:   RawWord{{0, 1}, {1, 2}, {3, -4}, {4, -2}, {3, 5}},
			want: RawWord{{3, -5}, {4, 2}, {3, 4}, {1, -2}, {0, -1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := presentation.InvRawWord(tt.in)
			if !presentation.EqualRawWord(got, tt.want) {
				t.Fatalf("Inverse(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name string
		in   RawWord
		want RawWord
	}{
		{
			name: "no reduction",
			in:   RawWord{{1, 2}, {3, 4}, {5, 6}},
			want: RawWord{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name: "yes reduction",
			in:   RawWord{{1, 2}, {1, -2}, {3, 5}, {3, 6}, {2, 1}, {2, -4}, {2, 2}},
			want: RawWord{{3, 11}, {2, -1}},
		},
		{
			name: "complete cancellation",
			in:   RawWord{{0, 7}, {3, 4}, {2, -4}, {2, 4}, {9, -6}, {9, 6}, {3, -4}, {0, -7}},
			want: presentation.EmptyRawWord(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := presentation.ReduceRawWord(tt.in)
			if !presentation.EqualRawWord(got, tt.want) {
				t.Fatalf("Reduce(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestCyclicReduce(t *testing.T) {
	tests := []struct {
		name string
		in   RawWord
		want RawWord
	}{
		{
			name: "no reduction",
			in:   RawWord{{1, 2}, {3, 4}, {5, 6}},
			want: RawWord{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name: "non-cyclic reduction",
			in:   RawWord{{1, 2}, {1, -2}, {3, 5}, {3, 6}, {2, 1}, {2, -4}, {2, 2}},
			want: RawWord{{3, 11}, {2, -1}},
		},
		{
			name: "simple cyclic reduction",
			in:   RawWord{{0, 7}, {3, 4}, {2, -4}, {0, -7}},
			want: RawWord{{3, 4}, {2, -4}},
		},
		{
			name: "bigger cyclic reduction",
			in:   RawWord{{5, 5}, {4, -2}, {0, 8}, {3, 4}, {2, -4}, {0, -7}, {4, 2}, {5, -5}},
			want: RawWord{{0, 1}, {3, 4}, {2, -4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := presentation.CyclicReduceRawWord(tt.in)
			if !presentation.EqualRawWord(got, tt.want) {
				t.Fatalf("Reduce(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestIsSubword(t *testing.T) {
	tests := []struct {
		name  string
		whole RawWord
		sub   RawWord
		want  bool
	}{
		{
			name:  "subword",
			whole: RawWord{{1, 2}, {2, -3}, {6, -7}, {3, 1}, {4, 7}},
			sub:   RawWord{{6, -7}, {3, 1}},
			want:  true,
		},
		{
			name:  "not subword",
			whole: RawWord{{1, 2}, {2, -3}, {6, -7}, {3, 1}, {4, 7}},
			sub:   RawWord{{6, -7}, {2, -3}},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := presentation.IsSubRawWord(tt.sub, tt.whole)
			if got != tt.want {
				t.Fatalf("IsSubword(%v, %v) = %v, want %v", tt.sub, tt.whole, got, tt.want)
			}
		})
	}
}

func TestShortLex(t *testing.T) {
	tests := []struct {
		name    string
		smaller RawWord
		bigger  RawWord
		want    bool
	}{
		{
			name:    "unequal length words true",
			bigger:  RawWord{{1, 2}, {2, -3}, {6, -7}, {3, 1}, {4, 7}},
			smaller: RawWord{{6, -7}, {3, 1}},
			want:    true,
		},
		{
			name:    "equal length words true",
			smaller: RawWord{{1, 2}, {2, -3}, {6, -7}, {3, 1}, {4, 7}},
			bigger:  RawWord{{6, -2}, {2, -3}, {6, -7}, {3, 1}, {4, 7}},
			want:    true,
		},
		{
			name:    "unequal length words false",
			bigger:  RawWord{{1, 2}, {2, -3}, {6, -7}},
			smaller: RawWord{{5, -8}, {3, 1}, {4, 3}, {17, 8}},
			want:    false,
		},
		{
			name:    "same words",
			bigger:  RawWord{{1, 2}, {2, -3}, {6, -7}},
			smaller: RawWord{{1, 2}, {2, -3}, {6, -7}},
			want:    false,
		},
		{
			name:    "equal length words false",
			bigger:  RawWord{{5, -8}, {3, 1}, {4, 3}, {17, 8}},
			smaller: RawWord{{5, -8}, {3, 1}, {4, 3}, {17, 9}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := presentation.ShortLexRawWord(tt.smaller, tt.bigger)
			if got != tt.want {
				t.Fatalf("ShortLex(%v, %v) = %v, want %v", tt.smaller, tt.bigger, got, tt.want)
			}
		})
	}
}
