package advent20201223_test

import (
	advent "advent20201223"
	"testing"
)

func TestPart2(t *testing.T) {
	// firstFew := []int{6, 2, 4, 3, 9, 7, 1, 5, 8}
	firstFew := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	iterations := 3
	numNodes := 9
	firstNode, nodeMap := advent.GenerateRing(firstFew, numNodes)
	currentNode := &firstNode
	// advent.PrintRing(currentNode, numNodes)
	for i := 0; i < iterations; i++ {
		// fmt.Printf("%v %v", &firstNode, currentNode)
		advent.PrintRing(*currentNode, numNodes)
		snip := (*currentNode).Next
		currentNode.Next = (*currentNode).Next.Next.Next.Next
		advent.PrintRing(*currentNode, numNodes)
		nextThree := []int{snip.ID, snip.Next.ID, snip.Next.Next.ID}
		// fmt.Println(nextThree)
		insertAfter := advent.IndexFor(currentNode.ID, nextThree, len(nodeMap))

		// fmt.Println(insertAfter)
		leftNode := *(nodeMap[insertAfter])
		rightNodePointer := leftNode.Next
		leftNode.Next = snip
		leftNode.Next.Next.Next.Next = rightNodePointer
		currentNode = currentNode.Next
	}
	t.Fail()
}

// func TestPart2(t *testing.T) {
// 	fmt.Println(advent.Part2("sample.txt"))
// 	fmt.Println(advent.Part2("input.txt"))
// 	t.Fail()
// }
