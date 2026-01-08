package presentation

import (
	"errors"
	"github.com/geometricgrouptheorydev/groups-in-go/groups"
)

var (
    ErrInvalidNumGenerators = errors.New("presentation: invalid number of generators")
    ErrInvalidRelation      = errors.New("presentation: relation uses out-of-range generator index")
)

type GroupPresentation struct {
	gen int
	rel []Word
	classes map[Class]struct{}
}

func TrivialPresentation() GroupPresentation{
	return GroupPresentation{}
}

//Arguments: number of generators and the relations. Invalid presentations return an error.
func NewGroupPresentation(generators int, relations []Word) (GroupPresentation, error) {
	if generators < 0 {
		return GroupPresentation{}, ErrInvalidNumGenerators
	} else if len(relations) == 0 {
		return NewFreeGroup(generators)
	}
	for _, r := range relations {
		for _, p := range r {
			if p[0] < 0 || p[0] >= generators {
				return GroupPresentation{}, ErrInvalidRelation
			}
		}
	}
	return GroupPresentation{gen: generators, rel: relations}, nil
}

//we want to implement group
type Group[T any] = groups.Group[T]

func (G GroupPresentation) Mu(v Word, w Word) Word {
	return Reduce(Concat(v,w))
}

func (G GroupPresentation) Inv(v Word) Word {
	return Inv(v)
}

func (G GroupPresentation) Id() Word {
	return Word{}
}