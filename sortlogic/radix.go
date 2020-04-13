package sortlogic

func Radix(a []int) []int {
	max := a[0]
	for _, v := range a {
		if max < v {
			max = v
		}
	}
	bitNum := getBit(max)
	for i := 1; i <= bitNum; i++ {
		dict := make([][]int, 10)
		for _, v := range a {
			index := v % (10 * i)
			dict[index] = append(dict[index], v)
		}
		cnt := 0
		for _, v := range dict {
			for _, vv := range v {
				a[cnt] = vv
				cnt++
			}
		}
	}
	return a
}

func getBit(num int) int {
	if num == 0 {
		return 0
	}
	return 1 + getBit(num/10)
}
