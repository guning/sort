package sortlogic

func Count(a []int) []int {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}
	tmp := make([]int, max+1)
	for i := 0; i < len(a); i++ {
		tmp[a[i]]++
	}

	var res []int
	for i := 0; i < len(tmp); i++ {
		for tmp[i] != 0 {
			res = append(res, i)
			tmp[i]--
		}
	}

	return res
}
