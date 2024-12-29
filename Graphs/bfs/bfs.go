package bfs

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

//        5
//    4      7
//  1   3  6   8
//

func bfs(root *Node) *Node {
	if root == nil {
		return root
	}
	// HashMap of key int and value of Array of Nodes
	bfsMap := make(map[int][]*Node)
	bfsMap[0] = []*Node{root}

	// Visited Array to track Visited Nodes
	visited := []*Node{}

	i := 0

	for {

		arr, ok := bfsMap[i]
		if !ok {
			break
		}
		for _, node := range arr {

			visited = append(visited, node)
			if node.Left != nil {

				_, ok := bfsMap[i+1]
				if !ok {
					bfsMap[i+1] = []*Node{}
				}

				bfsMap[i+1] = append(bfsMap[i+1], node.Left)
				bfsMap[i+1] = append(bfsMap[i+1], node.Right)

			}

		}
		i += 1
	}
	return root
}
