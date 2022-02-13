package main

import (
	"fmt"

	. "github.com/PavelYadrov/interface/shuffler"
	. "github.com/PavelYadrov/interface/types"
)

func main() {
	myType := Something{Slice: []int{1, 2, 3, 4, 5, 6}}
	//myType1 := Something{Slice: []int{6,5,4,3,2,1}}
	//fmt.Print(myType1)
	fmt.Print("\n")
	fmt.Print(myType)
	fmt.Print("\n")
	Shuffle(myType)
	fmt.Print(myType)
	fmt.Print("\n")

	whatIsThis(myType)
	fmt.Print(myType)

	fmt.Print("\n")
	kek(myType)
	fmt.Print(myType)
}

func whatIsThis(i interface{}) {
	//fmt.Printf("%T\n", i)
	switch val := i.(type) {
	case Shuffler:
		Shuffle(val)
	}
}

func kek(ss Something) {
	ss.Swap(1, 5)
}
