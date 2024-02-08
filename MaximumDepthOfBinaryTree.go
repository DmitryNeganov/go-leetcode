
 type TreeNode struct {
     Val int
     Right *TreeNode

 }
 
 func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    var left = maxDepth(root.Left)
    var right = maxDepth(root.Right)

    return maxValue(left, right) + 1
}

func maxValue(a, b int) int {
    if (a > b) {
        return a
    } else {
        return b
    }
}