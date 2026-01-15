package presentation_test

import (
	"testing"

	p "github.com/geometricgrouptheorydev/groups-in-go/presentation"
)

//testing word validity checker

type GroupPresentation = p.GroupPresentation

func TestNewGroupPresentation(t *testing.T) {
	tests := []struct {
		name        string
		gen         int
		rel         []Word
		wantErr     bool
		wantRel     map[string]Word
		wantClasses map[p.Class]bool
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := p.NewGroupPresentation(tt.gen, tt.rel)
			if (err != nil) != tt.wantErr {
				t.Fatalf("wanted error %v got error %v", tt.wantErr, err)
			}
		})
	}
}
