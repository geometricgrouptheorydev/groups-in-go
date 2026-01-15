package presentation_test

import (
	"testing"

	p "github.com/geometricgrouptheorydev/groups-in-go/presentation"
)

// testing IsValidWord
func TestIsValidWord(t *testing.T) {
	type Word = p.WordSlice
	G, err := p.NewFreeGroup(6)
	if err != nil {
		t.Fatal("can't make a new free group (rank 6) smh!")
	}

	tests := []struct {
		name    string
		w       Word
		wantErr bool
	}{
		{
			name:    "valid word",
			w:       Word{{0, 6}, {1, 3}, {2, 4}, {3, 12}, {4, 17}, {5, 78}},
			wantErr: false,
		},
		{
			name:    "invalid word",
			w:       Word{{0, 6}, {1, 3}, {2, 4}, {3, 12}, {4, 17}, {5, 78}, {6, 1}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := G.IsValidWord(tt.w)
			if (err != nil) != tt.wantErr {
				t.Fatalf("wanted error %v got error %v", tt.wantErr, err)
			}
		})
	}
}

//the following double up as a test for the Group[T any] interface

// testing equality method
func TestEqual(t *testing.T) {
	type Word = p.WordSlice
	G, err := p.NewFreeGroup(7)
	if err != nil {
		t.Fatal("can't make a new free group (rank 7) smh!")
	}

	tests := []struct {
		name string
		v    Word
		w    Word
		want bool
	}{
		{
			name: "equal",
			v:    Word{{6, 3}, {4, -2}, {3, -6}, {0, 1}, {1, -1}},
			w:    Word{{6, 3}, {4, -2}, {3, -6}, {0, 1}, {1, -1}},
			want: true,
		},
		{
			name: "different exponent",
			v:    Word{{6, 3}, {4, -2}, {3, -6}, {0, 1}, {1, -1}},
			w:    Word{{6, 3}, {4, -2}, {3, -5}, {0, 1}, {1, -1}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := G.Equal(tt.v, tt.w)
			if got != tt.want {
				t.Fatalf("G.Equal(%v, %v) = %v want %v", tt.v, tt.w, got, tt.want)
			}
		})
	}
}

// testing the operations
func TestMu(t *testing.T) {
	G, err := p.NewFreeGroup(8)
	if err != nil {
		t.Fatal("can't make a new free group (rank 8) smh!")
	}

	type Word = p.WordSlice
	tests := []struct {
		name string
		v    Word
		w    Word
		vw   Word
	}{
		{
			name: "no reduction",
			v:    Word{{1, 2}, {3, 4}},
			w:    Word{{5, 6}, {7, 8}},
			vw:   Word{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
		},
		{
			name: "some reduction",
			v:    Word{{1, 2}, {3, 4}},
			w:    Word{{3, -2}, {5, 6}},
			vw:   Word{{1, 2}, {3, 2}, {5, 6}},
		},
		{
			name: "complete reduction",
			v:    Word{{1, 2}, {3, 4}},
			w:    Word{{3, -4}, {1, -2}},
			vw:   Word{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := G.Mu(tt.v, tt.w)
			if !G.Equal(tt.vw, got) {
				t.Fatalf("G.Mu(%v, %v) = %v want %v", tt.v, tt.w, got, tt.vw)
			}

		})
	}
}
