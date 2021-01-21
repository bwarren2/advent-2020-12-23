package advent20201223_test

import (
	advent "advent20201223"
	"fmt"
	"testing"
)

func TestSample(t *testing.T) {
	fmt.Println(advent.Permute([]int{3, 8, 9, 1, 2, 5, 4, 6, 7}, 10))
	t.Fail()
}
func TestPart1(t *testing.T) {
	fmt.Println(advent.Permute([]int{6, 2, 4, 3, 9, 7, 1, 5, 8}, 100))
	t.Fail()
}
func TestPart2(t *testing.T) {
	firstFew := []int{6, 2, 4, 3, 9, 7, 1, 5, 8}
	input := make([]int, 1000000)
	for i := 1; i < 1000001; i++ {
		if i < 10 {
			input[i-1] = firstFew[i-1]
		} else {
			input[i-1] = i
		}
	}
	ints := advent.Permute(input, 1000)
	onePlace := advent.IndexOf(ints, 1)
	fmt.Println(onePlace, ints[onePlace:onePlace+3], ints[:2])
	t.Fail()
}

// func TestPart2(t *testing.T) {
// 	fmt.Println(advent.Part2("sample.txt"))
// 	fmt.Println(advent.Part2("input.txt"))
// 	t.Fail()
// }
