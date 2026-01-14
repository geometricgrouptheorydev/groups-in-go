package presentation_test

import (
	"testing"
	p "github.com/geometricgrouptheorydev/groups-in-go/presentation"
)


func TestWordID(t *testing.T) {
	tests := []struct{
		name string
		word [][2]int
		want string
	}{
		{
			name: "some word",
			word: [][2]int{{0,2},{1,-3},{3,4}},
			want: "0:2,1:-3,3:4",
		},
		{
			name: "empty word",
			word: make([][2]int, 0),
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := p.WordID(tt.word)
			if got != tt.want {
				t.Fatalf("WordID(%v) = %v want %v", tt.word, got, tt.want)
			}
		})
	}
}

func TestParseWordID(t *testing.T) {
	tests := []struct{
		name string
		wordID string
		want p.Word
		wantErr bool
	}{
		{
			name: "some word",
			want: [][2]int{{0,2},{1,-3},{3,4}},
			wordID: "0:2,1:-3,3:4",
			wantErr: false,
		},
		{
			name: "empty word",
			want: make([][2]int, 0),
			wordID: "",
			wantErr: false,
		},
		{
			name: "invalid word 1",
			want: nil,
			wordID: "abcdefg,sde",
			wantErr: true,
		},
		{
			name: "invalid word 2",
			want: nil,
			wordID: "1:2,a:3",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.ParseWordID(tt.wordID)
			if (err != nil) != tt.wantErr {
				t.Fatalf("got error %v but didn't want one", err)
			}
			if !p.EqualWord(got, tt.want) {
				t.Fatalf("ParseWordID(%v) = %v want %v", tt.wordID, got, tt.want)
			}
		})
	}
}