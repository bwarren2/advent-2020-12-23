package advent20201223

import "fmt"

type Node struct {
	ID   int
	Next *Node
}

func GenerateRing(seed []int, ct int) (Node, map[int]*Node) {
	nodeMap := make(map[int]*Node)
	nodes := make([]Node, ct+1)
	for i := 1; i <= ct; i++ {
		var value int
		if i <= len(seed) {
			value = seed[i-1]
		} else {
			value = i
		}
		nodes[i] = Node{ID: value}
		if i > 1 {
			nodes[i-1].Next = &nodes[i]
		}
	}
	for i := 1; i <= ct; i++ {
		nodeMap[nodes[i].ID] = &nodes[i]
	}
	nodes[ct].Next = &nodes[1]
	return nodes[1], nodeMap
}

func IndexFor(given int, banlist []int, numNodes int) int {
	fmt.Println(given, banlist, numNodes)
	for i := 1; i < len(banlist)+2; i++ {
		candidate := given - i
		if candidate <= 0 {
			candidate += numNodes
		}
		if Contains(banlist, candidate) {
			continue
		}
		return candidate
	}
	return -1
}

func Contains(lst []int, target int) bool {
	for _, v := range lst {
		if v == target {
			return true
		}
	}
	return false
}

func PrintRing(node Node, terms int) {
	position := node
	for i := 0; i < terms; i++ {
		fmt.Print(position.ID, " ")
		position = *position.Next
	}
	fmt.Println()
}
