package advent20201223

import "fmt"

type Node struct {
	ID   int
	Next *Node
}

func GenerateRing(seed []int, ct int) (Node, map[int]Node) {
	nodeMap := make(map[int]Node)
	firstNode := Node{ID: seed[0]}
	firstNodeAddress := &firstNode
	var appendTo Node
	appendTo = firstNode
	var value int
	for i := 1; i < ct; i++ {
		if i < len(seed) {
			value = seed[i]
		} else {
			value = i
		}
		newNode := Node{ID: value}
		appendTo.Next = &newNode
		fmt.Printf("%p %v \n", &appendTo, appendTo)

		appendTo = newNode
	}
	fmt.Println(appendTo)
	appendTo.Next = firstNodeAddress
	fmt.Println(appendTo)
	return *(appendTo.Next), nodeMap
}
