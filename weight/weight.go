package weight

import (
	. "github.com/PavelYadrov/interface/shuffler"
)

type WeightShuffalable interface {
	Shuffler
	Weight(i int) int
}

func WeightShuffle(w WeightShuffalable) {
	total := 0

	for i := 0; i < w.Len(); i++ {
		total += w.Weight(i)
	}
}
