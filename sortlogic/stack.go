package sortlogic

import (
	"math"
)

func Stack(a []int) []int {
	for i := len(a); i > 0; i-- {
		getMaxStack(&a, i)
		//fmt.Println(a)
		tmp := a[0]
		a[0] = a[i-1]
		a[i-1] = tmp
	}
	return a
}

func getMaxStack(a *[]int, num int) {
	deep := int(math.Log2(float64(num)))

	for i := deep; i > 0; i-- {
		tmp := int(math.Pow(2, float64(i-1)))
		for j := 1; j <= tmp; j++ {
			checkMaxStack(a, tmp+j-2, num)
		}
	}

}

func checkMaxStack(a *[]int, f int, num int) {
	c1 := 2*f + 1
	c2 := 2*f + 2
	if c1 < num {
		if (*a)[f] < (*a)[c1] {
			tmp := (*a)[f]
			(*a)[f] = (*a)[c1]
			(*a)[c1] = tmp
		}
	}

	if c2 < num {
		if (*a)[f] < (*a)[c2] {
			tmp := (*a)[f]
			(*a)[f] = (*a)[c2]
			(*a)[c2] = tmp
		}
	}
}
