package problem110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isBalanced(root *TreeNode) bool {
	ok, _ := isBalancedWithHeight(root)
	return ok
}

func isBalancedWithHeight(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}

	leftBalanced, leftHeight := isBalancedWithHeight(node.Left)
	rigthBalanced, rigthHeight := isBalancedWithHeight(node.Right)
	height := max(leftHeight, rigthHeight) + 1
	if !leftBalanced || !rigthBalanced {
		return false, height
	}
	if abs(leftHeight-rigthHeight) <= 1 {
		return true, height
	} else {
		return false, height
	}
}
