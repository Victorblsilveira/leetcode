package twosum

func twosum3(nums []int, target int) []int {
	indices := make(map[int]int, len(nums))

	for i, num := range nums {
		complement := target - num

		idxs, ok := indices[complement]
		if ok {
			return []int{idxs, i}
		}

		indices[num] = i
	}

	return nil
}
