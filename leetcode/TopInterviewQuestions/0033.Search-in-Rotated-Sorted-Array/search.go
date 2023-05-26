package topinterviewquestions

func search(nums []int, target int) int {
	n := len(nums)
	l, h := 0, n-1

	for l < h {
		m := (l + h) / 2
		if nums[m] > nums[h] {
			l = m + 1
		} else {
			h = m
		}
	}

	r := l
	l = 0
	h = n - 1
	for l <= h {
		m := (l + h) / 2
		rm := (m + r) % n
		if nums[rm] == target {
			return rm
		}

		if nums[rm] < target {
			l = m + 1
		} else {
			h = m - 1
		}

	}

	return -1
}
