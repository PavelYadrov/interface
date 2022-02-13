package shuffler

import (
	"math/rand"
)

type Shuffler interface {
	Len() int
	Swap(i, j int)
}

func Shuffle(s Shuffler) {
	for i := 0; i < s.Len(); i++ {
		j := rand.Intn(s.Len() - i)
		s.Swap(i, j)
	}
}
