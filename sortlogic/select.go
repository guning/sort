package sortlogic

func Select(a []int) []int {
	for i := 0; i < len(a); i++ {
		//找最小的下标
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		if i != min {
			tmp := a[i]
			a[i] = a[min]
			a[min] = tmp
		}
	}
	return a
}
