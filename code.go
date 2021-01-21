package advent20201223

// Part1 solves part1
func Permute(input []int, times int) []int {
	for i := 0; i < times; i++ {
		index := i % len(input)
		current := index
		currentValue := input[index]
		// fmt.Println("Input:", input)
		// fmt.Println("Index:", index)
		// fmt.Println("Current:", current)
		removed := []int{}
		for j := 1; j < 4; j++ {
			removed = append(removed, input[(i+j)%len(input)])
		}
		// fmt.Println("pick up:", removed)
		input = OtherNumbers(input, removed)
		// fmt.Println("remaining:", input)
		insertIndex := IndexForCurrentCup(input, currentValue)
		// fmt.Println("insert at:", insertIndex)
		piece1 := append([]int{}, input[:insertIndex]...)
		piece2 := append([]int{}, input[insertIndex:]...)
		input = append(piece1, removed...)
		input = append(input, piece2...)
		// fmt.Println("Produces:", input)
		input = RotateToAssert(input, current, currentValue)
		// fmt.Println("Rotated:", input)
	}
	return input
}

func OtherNumbers(input, exclude []int) (result []int) {
	for _, value := range input {
		if !Contains(exclude, value) {
			result = append(result, value)
		}
	}
	return
}

func Contains(lst []int, target int) bool {
	for _, value := range lst {
		if value == target {
			return true
		}
	}
	return false
}

func IndexForCurrentCup(input []int, value int) int {
	for i := 1; i < 4; i++ {
		target := value - i
		if target <= 0 {
			target += len(input) + 3
		}
		for idx, value := range input {
			if value == target {
				return idx + 1
			}
		}
	}
	return 0
}

func RotateToAssert(input []int, idx, value int) []int {
	for input[idx] != value {
		input = append(input[1:], input[0])
	}
	return input
}

func IndexOf(data []int, target int) int {
	for k, v := range data {
		if target == v {
			return k
		}
	}
	return -1
}
