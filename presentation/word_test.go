package presentation_test

import (
	"testing"
	"github.com/geometricgrouptheorydev/groups-in-go/presentation"
)

func TestEqualWord(t* testing.T) {
	type Word = presentation.Word

	tests := []struct {
        name string
        first Word
        second Word
		want bool
    }{
        {
            name: "unequal lengths",
            first:   Word{{0,1},{4,3}},
            second: Word{{0,1},{4,3},{6,7}},
			want: false,
        },
        {
            name: "different exponent",
            first:  Word{{0,1},{4,3},{6,7}},
            second: Word{{0,1},{4,3},{6,8}},
			want: false,
        },
        {
            name: "swapped",
            first:  Word{{0, 1}, {1, 2}},
            second: Word{{1, -2}, {0, -1}},
			want: false,
        },
		{
            name: "different generators",
            first:  Word{{0, 1}, {1, 2}},
            second: Word{{0, 1}, {2, 2}},
			want: false,
        },
		{
            name: "actually the same",
            first:  Word{{0, 1}, {1, 2}, {3, -4}, {4, -2}, {3, 5}},
            second: Word{{0, 1}, {1, 2}, {3, -4}, {4, -2}, {3, 5}},
			want: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := presentation.EqualWord(tt.first, tt.second)
            if got != tt.want {
                t.Fatalf("EqualWord(%v, %v) = %v, want %v", tt.first, tt.second, got, tt.want)
            }
        })
    }
}

func TestInverse(t *testing.T) {
    type Word = presentation.Word

    tests := []struct {
        name string
        in   Word
        want Word
    }{
        {
            name: "empty",
            in:   Word{},
            want: Word{},
        },
        {
            name: "single letter",
            in:   Word{{0, 1}},
            want: Word{{0, -1}},
        },
        {
            name: "two letters",
            in:   Word{{0, 1}, {1, 2}},
            want: Word{{1, -2}, {0, -1}},
        },
		{
            name: "five letters",
            in:   Word{{0, 1}, {1, 2}, {3, -4}, {4, -2}, {3, 5}},
            want: Word{{3, -5}, {4, 2}, {3,4}, {1, -2}, {0, -1}},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := presentation.Inv(tt.in)
            if !presentation.EqualWord(got, tt.want) {
                t.Fatalf("Inverse(%v) = %v, want %v", tt.in, got, tt.want)
            }
        })
    }
}

func TestReduce(t *testing.T) {
	type Word = presentation.Word 

	tests := []struct{
		name string
		in Word
		want Word
	}{
		{
			name: "no reduction",
			in: Word{{1,2},{3,4},{5,6}},
			want: Word{{1,2},{3,4},{5,6}},
		},
		{
			name: "yes reduction",
			in: Word{{1,2},{1,-2},{3,5},{3,6},{2,1},{2,-4},{2,2}},
			want: Word{{3,11},{2,-1}},
		},
		{
			name: "complete cancellation",
			in: Word{{0,7},{3,4},{2,-4},{2,4},{9,-6},{9,6},{3,-4},{0,-7}},
			want: presentation.EmptyWord(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := presentation.Reduce(tt.in)
			if !presentation.EqualWord(got, tt.want) {
				t.Fatalf("Reduce(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestIsSubword(t *testing.T) {
	type Word = presentation.Word

	tests := []struct{
		name string
		whole Word
		sub Word
		want bool
	}{
		{
			name: "subword",
			whole: Word{{1,2},{2,-3},{6,-7},{3,1},{4,7}},
			sub: Word{{6,-7},{3,1}},
			want: true,
		},
		{
			name: "not subword",
			whole: Word{{1,2},{2,-3},{6,-7},{3,1},{4,7}},
			sub: Word{{6,-7},{2,-3}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := presentation.IsSubword(tt.sub, tt.whole)
			if got != tt.want {
				t.Fatalf("IsSubword(%v, %v) = %v, want %v", tt.sub, tt.whole, got, tt.want)
			}
		})
	}
}

func TestShortLex(t *testing.T) {
	type Word = presentation.Word

	tests := []struct{
		name string
		smaller Word
		bigger Word
		want bool
	}{
		{
			name: "unequal length words true",
			bigger: Word{{1,2},{2,-3},{6,-7},{3,1},{4,7}},
			smaller: Word{{6,-7},{3,1}},
			want: true,
		},
		{
			name: "equal length words true",
			smaller: Word{{1,2},{2,-3},{6,-7},{3,1},{4,7}},
			bigger: Word{{6,-2},{2,-3},{6,-7},{3,1},{4,7}},
			want: true,
		},
		{
			name: "unequal length words false",
			bigger: Word{{1,2},{2,-3},{6,-7}},
			smaller: Word{{5,-8},{3,1},{4,3},{17,8}},
			want: false,
		},
		{
			name: "same words",
			bigger: Word{{1,2},{2,-3},{6,-7}},
			smaller: Word{{1,2},{2,-3},{6,-7}},
			want: false,
		},
		{
			name: "equal length words false",
			bigger: Word{{5,-8},{3,1},{4,3},{17,8}},
			smaller: Word{{5,-8},{3,1},{4,3},{17,9}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := presentation.ShortLex(tt.smaller, tt.bigger)
			if got != tt.want {
				t.Fatalf("ShortLex(%v, %v) = %v, want %v", tt.smaller, tt.bigger, got, tt.want)
			}
		})
	}
}
