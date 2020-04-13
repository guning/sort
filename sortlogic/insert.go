package sortlogic

func Insert(a []int) []int {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && a[j-1] > a[j]; j-- {
			tmp := a[j]
			a[j] = a[j-1]
			a[j-1] = tmp
		}
	}
	return a
}
