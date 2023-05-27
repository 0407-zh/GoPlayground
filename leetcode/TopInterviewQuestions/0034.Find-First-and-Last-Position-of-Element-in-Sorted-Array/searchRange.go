package topinterviewquestions

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	start := binarySearch(nums, target, true)
	end := binarySearch(nums, target, false)

	return []int{start, end}
}

func binarySearch(nums []int, target int, findFirst bool) int {
	left, right := 0, len(nums)-1
	index := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] > target || (findFirst && nums[mid] == target) {
			right = mid - 1
		} else {
			left = mid + 1
		}

		if nums[mid] == target {
			index = mid
		}
	}

	return index
}
