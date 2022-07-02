package problem235

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}

	for {
		if p.Val <= root.Val && root.Val <= q.Val {
			return root
		}

		if root.Val < p.Val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
}
