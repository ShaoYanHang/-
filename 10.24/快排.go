func quicksort(nums []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	p := nums[i]
	for i < j {
		for i < j && nums[i] < p {
			i++
		}
		for i < j && nums[j] > p {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[j] = p
	quicksort(nums, left, i-1)
	quicksort(nums, i+1, right)
}
