package presentation

import (
	"errors"
	"fmt"
	"github.com/geometricgrouptheorydev/groups-in-go/groups"
)

var (
    ErrInvalidNumGenerators = errors.New("presentation: invalid number of generators")
    ErrInvalidRelation      = errors.New("presentation: relation uses out-of-range generator index")
)

type GroupPresentation struct {
	gen int //generators
	rel []Word //relations
	classes map[Class]bool //true means the group is in that class, false means it is not, and if a class is not a map key it means we don't know
}

func TrivialPresentation() GroupPresentation{
	return GroupPresentation{classes: trivialGroupClasses}
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

//WARNING: just like for the Group[T any] interface, words are not by default checked if they're actually in the group before the operation is applied
//this will lead to varying behavior for out-of-range generators depending on G.classes
//the following O(n) function lets you check this automatically at an O(n) cost, useful for long words that are hard to check manually
func (G GroupPresentation) IsValidWord(w Word) error{
	for i, u := range w {
		if u[0] >= G.gen || u[0] < 0 {
			return fmt.Errorf("invalid generator %v at word index %v", u[0], i)
		}
	}
	return nil
}

//we want to implement group.Group
//again, use the above IsValidWord method if you wan to verify the inputted words are valid first
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