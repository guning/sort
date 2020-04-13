package sortlogic

func Shell(a []int) []int {
	length := len(a)

	for gap := length / 2; gap >= 1; gap = gap / 2 {
		for i := gap; i < length; i++ {
			for j := i; j-gap >= 0 && a[j] < a[j-gap]; j = j - gap {
				tmp := a[j]
				a[j] = a[j-gap]
				a[j-gap] = tmp
			}
		}
	}
	return a
}
