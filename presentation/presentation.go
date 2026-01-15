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
	gen     int            //generators
	rel     []WordSlice    //relations
	classes map[Class]bool //true means the group is in that class, false means it is not, and if a class is not a map key it means we don't know
}

func TrivialPresentation() GroupPresentation {
	return GroupPresentation{classes: trivialGroupClasses}
}

// Arguments: number of generators and the relations. Invalid presentations return an error.
func NewGroupPresentation(generators int, relations []WordSlice) (GroupPresentation, error) {
	if generators < 0 {
		return GroupPresentation{}, ErrInvalidNumGenerators
	} else if len(relations) == 0 {
		return NewFreeGroup(generators) //we'll deal with this case separately because free groups are so cool
	}
	reducedRelations := make([]WordSlice, len(relations)) //we reduce relations provided if possible
	for _, r := range relations {
		for _, p := range r {
			if p[0] < 0 || p[0] >= generators {
				return GroupPresentation{}, ErrInvalidRelation
			}
			reduced := Reduce(r)
			if len(reduced) > 0 {
				reducedRelations = append(reducedRelations, reduced) //empty words are not to be added in rel!
			}
		}
	}
	return initAddProperties(GroupPresentation{gen: generators, rel: reducedRelations, classes: make(map[Class]bool)})
}

// helper to initialize the group presentation with its properties
// includes only checks that are O(n^2) or better to match with NewGroupPresentation
func initAddProperties(G GroupPresentation) (GroupPresentation, error) {
	//one-relator and freedom
	if len(G.rel) == 1 {
		err := G.addClasses(oneRelatorGroupClasses)
		return G, err
	}
	//cyclicity
	if G.gen == 1 {
		G.addClasses(cyclicGroupClasses)
	}
	return G, nil
}

// WARNING: just like for the Group[T any] interface, words are not by default checked if they're actually in the group before the operation is applied
// this will lead to varying behavior for out-of-range generators depending on G.classes
// the following O(n) function lets you check this automatically at an O(n) cost, useful for long words that are hard to check manually
func (G GroupPresentation) IsValidWord(w WordSlice) error {
	for i, u := range w {
		if u[0] >= G.gen || u[0] < 0 {
			return fmt.Errorf("invalid generator %v at word index %v", u[0], i)
		}
	}
	return nil
}

// we want to implement group.Group
// again, use the above IsValidWord method if you wan to verify the inputted words are valid first
type Group[T any] = groups.Group[T]

func (G GroupPresentation) Mu(v WordSlice, w WordSlice) WordSlice {
	return Reduce(Concat(v, w))
}

func (G GroupPresentation) Inv(v WordSlice) WordSlice {
	return Inv(v)
}

func (G GroupPresentation) Id() WordSlice {
	return WordSlice{}
}

// WARNING: if the word problem is not solvable for your particular presentation, false does not guarantee inequality
func (G GroupPresentation) Equal(v WordSlice, w WordSlice) bool {
	return len(G.Mu(v, G.Inv(w))) == 0 //checks if vw^-1 is the empty word
}
