package presentation_test

import (
	"testing"

	p "github.com/geometricgrouptheorydev/groups-in-go/presentation"
)

type Word = p.Word

// testing IsValidWord
func TestIsValidWord(t *testing.T) {
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
			w:       p.NewWord(RawWord{{0, 6}, {1, 3}, {2, 4}, {3, 12}, {4, 17}, {5, 78}}),
			wantErr: false,
		},
		{
			name:    "invalid word",
			w:       p.NewWord(RawWord{{0, 6}, {1, 3}, {2, 4}, {3, 12}, {4, 17}, {5, 78}, {6, 1}}),
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
			v:    p.NewWord(RawWord{{6, 3}, {4, -2}, {3, -6}, {0, 1}, {1, -1}}),
			w:    p.NewWord(RawWord{{6, 3}, {4, -2}, {3, -6}, {0, 1}, {1, -1}}),
			want: true,
		},
		{
			name: "different exponent",
			v:    p.NewWord(RawWord{{6, 3}, {4, -2}, {3, -6}, {0, 1}, {1, -1}}),
			w:    p.NewWord(RawWord{{6, 3}, {4, -2}, {3, -5}, {0, 1}, {1, -1}}),
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

	tests := []struct {
		name string
		v    Word
		w    Word
		vw   Word
	}{
		{
			name: "no reduction",
			v:    p.NewWord(RawWord{{1, 2}, {3, 4}}),
			w:    p.NewWord(RawWord{{5, 6}, {7, 8}}),
			vw:   p.NewWord(RawWord{{1, 2}, {3, 4}, {5, 6}, {7, 8}}),
		},
		{
			name: "some reduction",
			v:    p.NewWord(RawWord{{1, 2}, {3, 4}}),
			w:    p.NewWord(RawWord{{3, -2}, {5, 6}}),
			vw:   p.NewWord(RawWord{{1, 2}, {3, 2}, {5, 6}}),
		},
		{
			name: "complete reduction",
			v:    p.NewWord(RawWord{{1, 2}, {3, 4}}),
			w:    p.NewWord(RawWord{{3, -4}, {1, -2}}),
			vw:   p.EmptyWord(),
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
