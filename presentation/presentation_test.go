package presentation_test

import (
	"testing"

	p "github.com/geometricgrouptheorydev/groups-in-go/presentation"
)

//testing word validity checker

type GroupPresentation = p.GroupPresentation

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

func TestNewGroupPresentation(t *testing.T) {
	tests := []struct {
		name        string
		gen         int
		rel         p.WordSet
		wantErr     bool
		wantRel     p.WordSet
		wantClasses map[p.Class]bool
	}{
		{
			name: "infinite cyclic group",
			gen: 1,
			rel: make(p.WordSet),
			wantErr: false,
			wantRel: make(p.WordSet),
			wantClasses: map[p.Class]bool{
				p.Trivial: false,
				p.Free: true,
				p.FreeAbelian: true,
				p.OneRelator: false,
				p.Abelian: true,
				p.Cyclic: true,
				p.Finite: false,
				},
		},
		{
			name: "one non-trivial relator",
			gen: 4,
			rel: p.NewWordSet([]p.Word{
				p.NewWord([][2]int{{0,3},{1,4},{1,-2},{1,-2},{0,-3}}), //this reduces to empty word
				p.NewWord([][2]int{{3,2},{2,-3},{2,1}}),
				}),
			wantErr: false,
			wantRel: p.NewWordSet([]p.Word{
				p.NewWord([][2]int{{3,2},{2,-2}}),
			}),
			wantClasses: map[p.Class]bool{
				p.OneRelator: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G, err := p.NewGroupPresentation(tt.gen, tt.rel)
			if (err != nil) != tt.wantErr {
				t.Fatalf("wanted error %v got error %v", tt.wantErr, err)
			}
			if !p.EqualWordSet(G.Relations(), tt.wantRel) {
				t.Fatalf("relator set mismatch: got %v wanted %v", G.Relations(), tt.wantRel)
			}
			if !equalMaps(G.Classes(), tt.wantClasses) {
				t.Fatalf("relator set mismatch: got %v wanted %v", G.Classes(), tt.wantClasses)
			}
		})
	}
}
