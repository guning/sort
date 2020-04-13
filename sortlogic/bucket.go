package sortlogic

const BUCKETNUM int = 5

func Bucket(a []int) []int {

	min := a[0]
	max := a[0]
	for _, v := range a {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}

	buckets := make([][]int, selectBucket(max, min)+1)

	for _, v := range a {
		index := selectBucket(v, min)
		buckets[index] = append(buckets[index], v)
	}

	var res []int
	for _, v := range buckets {
		if len(v) > 0 {
			newV := Insert(v)
			res = combine(res, newV)
		}
	}
	return res
}

func combine(a, b []int) []int {
	for _, v := range b {
		a = append(a, v)
	}
	return a
}

func selectBucket(v, min int) int {
	return (v - min) - ((v-min)%BUCKETNUM)/BUCKETNUM
}
