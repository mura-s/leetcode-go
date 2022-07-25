package problem104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeDepth struct {
	Node  *TreeNode
	Depth int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := []*NodeDepth{
		{Node: root, Depth: 1},
	}
	for len(queue) != 0 {
		nd := queue[0]
		queue = queue[1:]
		if depth < nd.Depth {
			depth = nd.Depth
		}
		if nd.Node.Left != nil {
			queue = append(queue, &NodeDepth{Node: nd.Node.Left, Depth: nd.Depth + 1})
		}
		if nd.Node.Right != nil {
			queue = append(queue, &NodeDepth{Node: nd.Node.Right, Depth: nd.Depth + 1})
		}
	}
	return depth
}
