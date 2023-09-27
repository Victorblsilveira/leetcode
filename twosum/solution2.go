package twosum

func twosum2(nums []int, target int) []int {
	indices := map[int][]int{}

	for i, num := range nums {
		complement := target - num

		idxs, ok := indices[num]
		if ok {
			return append(idxs, i)
		}

		idxs, ok = indices[complement]
		if !ok {
			indices[complement] = []int{i}
		}
	}

	return []int{}
}
