/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
	var res []int
    if root == nil {
		return []int{}
	}
    // We declare dfs first so it can call itself recursively
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node != nil {
            dfs(node.Left)
            dfs(node.Right)
			res = append(res, node.Val)
        }
    }
    
    dfs(root)
    return res
}
