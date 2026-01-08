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
		G.classes = addClasses(G.classes, oneRelatorGroupClasses)
	}
	
	//cyclicity and abelianity
	if G.gen == 1 {
		G.classes = addClasses(G.classes, cyclcGroupClasses)
	} else if initCheckCommutativityRelators(G) {
		G.classes = addClasses(G.classes, abelianGroupClasses)
	}
	return G
}

// initCheckCommutativityRelators reports whether all [x_i, x_j] are in rel.
func initCheckCommutativityRelators(G GroupPresentation) bool {
    for i := 0; i < G.gen; i++ {
        for j := 0; j < i; j++ {
            found := false
            target1 := Word{{i, 1}, {j, 1}, {i, -1}, {j, -1}}
            target2 := Word{{i, -1}, {j, -1}, {i, 1}, {j, 1}}
            for _, r := range G.rel {
				if len(r) != 4 {
					continue
				} else if EqualWord(r, target1) || EqualWord(r, target2) {
                    found = true
                    break
                }
            }
            if !found {
                return false
            }
        }
    }
    return true
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