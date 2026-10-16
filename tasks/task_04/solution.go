package main

type Stats struct {
	Count         int
	Sum, Min, Max int64
}

func Calc(nums []int64) Stats {
	var result Stats

	i := 0

	for ; i+1 < len(nums); i++ {
		current := nums[i+1] - nums[i]
		result.Count++
		result.Sum += current

		if result.Min > current || result.Count == 1 {
			result.Min = current
		}

		if result.Max < current || result.Count == 1 {
			result.Max = current
		}
	}

	return result
}
