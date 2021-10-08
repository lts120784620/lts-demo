package binaryTree

import (
	"fmt"
	"math"
	"testing"
)

/**
No.124 二叉树中的最大路径和
描述：
	路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。
	该路径 至少包含一个 节点，且不一定经过根节点。路径和 是路径中各节点值的总和。给你一个二叉树的根节点 root ，返回其 最大路径和.
思路：
    1、和543的思路差不多，通过递归法后续遍历的方式，取l和r中较大的值+root的值向上传递，因为跟节点不一定经过，所以要用全局变量sum来记录总和，当l+r+root>sum时进行替换。ps:注意处理负数值
时间：
    1、2021/10/6
*/

func TestMaxPathSum(t *testing.T) {
	r := NewTreeByArrays([]interface{}{2, -1})
	fmt.Println(maxPathSum(r))
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, sum := pathSum(root, math.MinInt64)
	return sum
}

func pathSum(root *TreeNode, sum int) (int, int) {
	if root == nil {
		return 0, sum
	}
	l, sum := pathSum(root.Left, sum)
	// 处理负数节点
	if l < 0 {
		l = 0
	}
	r, sum := pathSum(root.Right, sum)
	if r < 0 {
		r = 0
	}
	if l+r+root.Val > sum {
		sum = l + r + root.Val
	}

	if l > r {
		return l + root.Val, sum
	} else {
		return r + root.Val, sum
	}
}
