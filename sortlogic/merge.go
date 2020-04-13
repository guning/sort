package sortlogic

func Merge(a []int) []int {
	a = merge(a)
	return a
}

func merge(a []int) []int {
	if len(a) == 1 {
		return a
	}
	var res []int
	st := 0
	ed := len(a) - 1
	divide := (ed - st + 1) / 2
	left := merge(a[st:divide])
	right := merge(a[divide:])
	i := 0
	j := 0
	for i < len(left) || j < len(right) {
		if i >= len(left) {
			for t := j; t < len(right); t++ {
				res = append(res, right[t])
			}
			j = len(right)
			break
		}
		if j >= len(right) {
			for t := i; t < len(left); t++ {
				res = append(res, left[t])
			}
			i = len(left)
			break
		}
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	return res
}
