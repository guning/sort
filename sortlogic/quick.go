package sortlogic

func Quick(a []int) []int {
	depart(&a, 0, len(a)-1)
	return a
}

func depart(a *[]int, st, ed int) {
	if ed-st < 1 {
		return
	}
	rawSt := st
	rawEd := ed
	for st != ed {
		for ed > st {
			if (*a)[st] > (*a)[ed] {
				(*a)[st], (*a)[ed] = (*a)[ed], (*a)[st]
				st++
				break
			}
			ed--
		}
		for st < ed {
			if (*a)[st] > (*a)[ed] {
				(*a)[st], (*a)[ed] = (*a)[ed], (*a)[st]
				ed--
				break
			}
			st++
		}
	}
	depart(a, rawSt, st-1)
	depart(a, st+1, rawEd)
}
