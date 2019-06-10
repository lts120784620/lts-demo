package binaryTree

import (
	"fmt"
	"lts-demo/test/common"
	"strconv"
	"math"
)

// 无厘头的题目
func Tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	if t.Right == nil && t.Left == nil {
		return strconv.Itoa(t.Val)
	}
	if t.Left != nil && t.Right == nil {
		return strconv.Itoa(t.Val) + "(" + Tree2str(t.Left) + ")"
	} else if t.Left == nil && t.Right != nil {
		return strconv.Itoa(t.Val) + "()" + "(" + Tree2str(t.Right) + ")"
	}
	return strconv.Itoa(t.Val) + "(" + Tree2str(t.Left) + ")" + "(" + Tree2str(t.Right) + ")"
}

// 计算二叉搜索树在L R之间的的节点的和
// 遍历，判断，累加
func RangeSumBST(root *TreeNode, L int, R int) int {
	res := 0
	if root == nil {
		return res
	}
	s := common.NewStack()
	node := root
	for !s.IsEmpty() || node != nil {
		if node != nil {
			s.Push(node)
			node = node.Left
		} else {
			node = s.Pop().(*TreeNode)
			if node.Val >= L && node.Val <= R {
				res += node.Val
			}
			node = node.Right
		}
	}
	return res
}

// 寻找二叉树中所有左叶子节点的和
// 思路：判断左节点的左右节点是否为空，做累加
func SumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Right != nil {
		fmt.Println(root.Left.Val)
		return root.Left.Val + SumOfLeftLeaves(root.Right)
	}

	return SumOfLeftLeaves(root.Left) + SumOfLeftLeaves(root.Right)
}

//　找到二叉树的路径总和三
// 遍历所有节点，每个节点到找到等于sum的值（第二种，可以通过建立map，快速定位遍历过的节点，会耗些内存）
func PathSum(root *TreeNode, sum int) int {
	res := 0
	_preOrder(root, sum, &res)
	return res
}

func pathOrder(root *TreeNode, sum int, res *int) {
	if root == nil {
		return
	}
	if sum == root.Val {
		*res += 1
		fmt.Println(root.Val)
	}
	pathOrder(root.Left, sum - root.Val, res)
	pathOrder(root.Right, sum - root.Val, res)
}

func _preOrder(root *TreeNode, sum int, res *int) {
	if root == nil {
		return
	}
	pathOrder(root, sum, res)
	_preOrder(root.Left, sum, res)
	_preOrder(root.Right, sum, res)
}

// 思路:后续遍历,左右中
func FindTilt(root *TreeNode) int {
	res := 0
	postOrder(root,&res)
	return res
}

func postOrder(root *TreeNode, sum *int) int{
	if root == nil {
		return 0
	}

	l := postOrder(root.Left, sum)
	r := postOrder(root.Right, sum)
	*sum+=int(math.Abs(float64(l-r)))
	return l + r + root.Val
}