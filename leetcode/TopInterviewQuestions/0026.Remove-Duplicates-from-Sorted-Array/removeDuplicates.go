package topinterviewquestions

func removeDuplicates(nums []int) int {
	n := len(nums)

	i := 0
	for j := 0; j < n; j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
