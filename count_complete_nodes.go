package leetcodesolutionsgo
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func countNodes(root *TreeNode) int {
    count:=0
    dfs(root,&count)  
    return (count)
  
}

func dfs(root *TreeNode,count *int){
    if root==nil{
        return
    }
    *count++
    dfs(root.Right,count)
    dfs(root.Left,count)
}