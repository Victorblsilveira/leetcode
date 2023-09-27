package twosum

type Number struct {
	value int
	index int
}

func twosum(nums []int, target int) []int {
	numbers := map[int]Number{}

	for i, num := range nums {
		numbers[num] = Number{value: num, index: i}
	}

	for i, num := range nums {
		complement := target - num

		number, ok := numbers[complement]
		if !ok {
			continue
		}

		// Same element cannot be used twice
		if number.index == i {
			continue
		}

		return []int{i, number.index}
	}

	return []int{}
}
