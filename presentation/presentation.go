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
	gen int //generators
	rel []Word //relations
	classes map[Class]struct{} //classes we know the group is part of
	negClasses map[Class]struct{} //classes we know the group is not part of
}

func TrivialPresentation() GroupPresentation{
	return GroupPresentation{classes: map[Class]struct{}{Trivial: {}}}
}

//Arguments: number of generators and the relations. Invalid presentations return an error.
func NewGroupPresentation(generators int, relations []Word) (GroupPresentation, error) {
	if generators < 0 {
		return GroupPresentation{}, ErrInvalidNumGenerators
	} else if len(relations) == 0 {
		return NewFreeGroup(generators) //we'll deal with this case separately because free groups are so cool
	}
	for _, r := range relations {
		for _, p := range r {
			if p[0] < 0 || p[0] >= generators {
				return GroupPresentation{}, ErrInvalidRelation
			}
		}
	}
	return initAddProperties(GroupPresentation{gen: generators, rel: relations}), nil
}

//helper to initialize the group presentation with its properties
func initAddProperties(G GroupPresentation) GroupPresentation {
	//one-relator
	if len(G.rel) == 1 {
		G.addClasses(oneRelatorGroupClasses)
	}
	//cyclicity
	if G.gen == 1 {
		G.addClasses(cyclicGroupClasses)
	}
	return G
}

//we want to implement group.Group
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