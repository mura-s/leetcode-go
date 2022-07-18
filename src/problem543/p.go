package problem543

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func countDiameter(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	leftDiameter, leftDepth := countDiameter(node.Left)
	rightDiameter, rigthDepth := countDiameter(node.Right)
	maxDiameter := max(max(leftDiameter, rightDiameter), leftDepth+rigthDepth)
	maxDepth := max(leftDepth, rigthDepth)
	return maxDiameter, maxDepth + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans, _ := countDiameter(root)
	return ans
}
