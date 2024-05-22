package leetcodesolutionsgo
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSymmetric(root *TreeNode) bool {
	return dfs(root.Left,root.Right)
}

func dfs(leftNode *TreeNode,rightNode *TreeNode) bool{
	 if leftNode==nil && rightNode == nil{
		 return true
	 }
	 if leftNode==nil || rightNode == nil{
		 return false
	 }
	 return  (leftNode.Val==rightNode.Val) && dfs(leftNode.Left,rightNode.Right) && dfs(leftNode.Right,rightNode.Left)

 }
