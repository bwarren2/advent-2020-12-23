package advent20201223_test

import (
	advent "advent20201223"
	"fmt"
	"testing"
)

func TestPart2(t *testing.T) {
	firstFew := []int{6, 2, 4, 3, 9, 7, 1, 5, 8}
	firstNode, nodeMap := advent.GenerateRing(firstFew, 9)
	// next := firstNode.Next
	fmt.Println(firstNode)
	// next = next.Next
	for i := 0; i <= 9; i++ {
	}
	fmt.Println(nodeMap)
	t.Fail()
}

// func TestPart2(t *testing.T) {
// 	fmt.Println(advent.Part2("sample.txt"))
// 	fmt.Println(advent.Part2("input.txt"))
// 	t.Fail()
// }
