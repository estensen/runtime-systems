package sort

import (
	"fmt"
	"math/rand"
	"sort"
)

func Sort() {
	numberList := []int{}
	for n := 0; n < 100000; n++ {
		numberList = append(numberList, rand.Intn(1000))
	}
	sort.Ints(numberList)
	fmt.Println("sorted")
}
