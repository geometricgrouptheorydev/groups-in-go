package presentation_test

import (
	"testing"

	"github.com/geometricgrouptheorydev/groups-in-go/presentation"
)

type WordSlice = presentation.WordSlice

func TestEqualWord(t *testing.T) {
	tests := []struct {
		name   string
		first  WordSlice
		second WordSlice
		want   bool
	}{
		{
			name:   "unequal lengths",
			first:  WordSlice{{0, 1}, {4, 3}},
			second: WordSlice{{0, 1}, {4, 3}, {6, 7}},
			want:   false,
		},
		{
			name:   "different exponent",
			first:  WordSlice{{0, 1}, {4, 3}, {6, 7}},
			second: WordSlice{{0, 1}, {4, 3}, {6, 8}},
			want:   false,
		},
		{
			name:   "swapped",
			first:  WordSlice{{0, 1}, {1, 2}},
			second: WordSlice{{1, -2}, {0, -1}},
			want:   false,
		},
		{
			name:   "different generators",
			first:  WordSlice{{0, 1}, {1, 2}},
			second: WordSlice{{0, 1}, {2, 2}},
			want:   false,
		},
		{
			name:   "actually the same",
			first:  WordSlice{{0, 1}, {1, 2}, {3, -4}, {4, -2}, {3, 5}},
			second: WordSlice{{0, 1}, {1, 2}, {3, -4}, {4, -2}, {3, 5}},
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := presentation.EqualWordSlice(tt.first, tt.second)
			if got != tt.want {
				t.Fatalf("EqualWord(%v, %v) = %v, want %v", tt.first, tt.second, got, tt.want)
			}
		})
	}
}

func TestInverse(t *testing.T) {
	tests := []struct {
		name string
		in   WordSlice
		want WordSlice
	}{
		{
			name: "empty",
			in:   WordSlice{},
			want: WordSlice{},
		},
		{
			name: "single letter",
			in:   WordSlice{{0, 1}},
			want: WordSlice{{0, -1}},
		},
		{
			name: "two letters",
			in:   WordSlice{{0, 1}, {1, 2}},
			want: WordSlice{{1, -2}, {0, -1}},
		},
		{
			name: "five letters",
			in:   WordSlice{{0, 1}, {1, 2}, {3, -4}, {4, -2}, {3, 5}},
			want: WordSlice{{3, -5}, {4, 2}, {3, 4}, {1, -2}, {0, -1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := presentation.InvWordSlice(tt.in)
			if !presentation.EqualWordSlice(got, tt.want) {
				t.Fatalf("Inverse(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name string
		in   WordSlice
		want WordSlice
	}{
		{
			name: "no reduction",
			in:   WordSlice{{1, 2}, {3, 4}, {5, 6}},
			want: WordSlice{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name: "yes reduction",
			in:   WordSlice{{1, 2}, {1, -2}, {3, 5}, {3, 6}, {2, 1}, {2, -4}, {2, 2}},
			want: WordSlice{{3, 11}, {2, -1}},
		},
		{
			name: "complete cancellation",
			in:   WordSlice{{0, 7}, {3, 4}, {2, -4}, {2, 4}, {9, -6}, {9, 6}, {3, -4}, {0, -7}},
			want: presentation.EmptyWordSlice(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := presentation.ReduceWordSlice(tt.in)
			if !presentation.EqualWordSlice(got, tt.want) {
				t.Fatalf("Reduce(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestIsSubword(t *testing.T) {
	tests := []struct {
		name  string
		whole WordSlice
		sub   WordSlice
		want  bool
	}{
		{
			name:  "subword",
			whole: WordSlice{{1, 2}, {2, -3}, {6, -7}, {3, 1}, {4, 7}},
			sub:   WordSlice{{6, -7}, {3, 1}},
			want:  true,
		},
		{
			name:  "not subword",
			whole: WordSlice{{1, 2}, {2, -3}, {6, -7}, {3, 1}, {4, 7}},
			sub:   WordSlice{{6, -7}, {2, -3}},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := presentation.IsSubWordSlice(tt.sub, tt.whole)
			if got != tt.want {
				t.Fatalf("IsSubword(%v, %v) = %v, want %v", tt.sub, tt.whole, got, tt.want)
			}
		})
	}
}

func TestShortLex(t *testing.T) {
	tests := []struct {
		name    string
		smaller WordSlice
		bigger  WordSlice
		want    bool
	}{
		{
			name:    "unequal length words true",
			bigger:  WordSlice{{1, 2}, {2, -3}, {6, -7}, {3, 1}, {4, 7}},
			smaller: WordSlice{{6, -7}, {3, 1}},
			want:    true,
		},
		{
			name:    "equal length words true",
			smaller: WordSlice{{1, 2}, {2, -3}, {6, -7}, {3, 1}, {4, 7}},
			bigger:  WordSlice{{6, -2}, {2, -3}, {6, -7}, {3, 1}, {4, 7}},
			want:    true,
		},
		{
			name:    "unequal length words false",
			bigger:  WordSlice{{1, 2}, {2, -3}, {6, -7}},
			smaller: WordSlice{{5, -8}, {3, 1}, {4, 3}, {17, 8}},
			want:    false,
		},
		{
			name:    "same words",
			bigger:  WordSlice{{1, 2}, {2, -3}, {6, -7}},
			smaller: WordSlice{{1, 2}, {2, -3}, {6, -7}},
			want:    false,
		},
		{
			name:    "equal length words false",
			bigger:  WordSlice{{5, -8}, {3, 1}, {4, 3}, {17, 8}},
			smaller: WordSlice{{5, -8}, {3, 1}, {4, 3}, {17, 9}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := presentation.ShortLexWordSlice(tt.smaller, tt.bigger)
			if got != tt.want {
				t.Fatalf("ShortLex(%v, %v) = %v, want %v", tt.smaller, tt.bigger, got, tt.want)
			}
		})
	}
}
