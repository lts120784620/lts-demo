package binaryTree

import (
	"fmt"
	"lts-demo/test/common"
)

// 求二叉树的路径之和是否等于sum值
func HasPathSum(root *TreeNode, sum int) bool {
	res := false
	if root == nil {
		return res
	}
	s := common.NewStack()
	s.Push(root)
	t := common.NewStack()
	var total int
	t.Push(root.Val)
	for !s.IsEmpty() {
		node := s.Pop().(*TreeNode)
		total = t.Pop().(int)

		if node.Right != nil {
			s.Push(node.Right)
			t.Push(total + node.Right.Val)
		}
		if node.Left != nil {
			s.Push(node.Left)
			t.Push(total + node.Left.Val)
		}
		if node.Left == nil && node.Right == nil {
			if total == sum {
				return true
			}
		}
	}
	return res
}

// 判断两颗二叉树的叶子节点序列是否相同(非递归)
// 思路:两个栈，分别存放两个树节点，判断叶子节点
func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	res := false
	if root1 == nil || root2 == nil {
		return false
	}
	s1 := common.NewStack()
	s2 := common.NewStack()
	s1.Push(root1)
	s2.Push(root2)
	for !s1.IsEmpty() && !s2.IsEmpty() {
		node1 := s1.Pop().(*TreeNode)
		node2 := s2.Pop().(*TreeNode)
		for node1.Left != nil || node1.Right != nil {
			if node1.Right != nil {
				s1.Push(node1.Right)
			}
			if node1.Left != nil {
				s1.Push(node1.Left)
			}
			node1 = s1.Pop().(*TreeNode)
		}
		for node2.Left != nil || node2.Right != nil {
			if node2.Right != nil {
				s2.Push(node2.Right)
			}
			if node2.Left != nil {
				s2.Push(node2.Left)
			}
			node2 = s2.Pop().(*TreeNode)
		}
		if node1.Val != node2.Val {
			return false
		} else {
			res = true
		}
	}
	return res
}

// 中序遍历节点，生成一个只有右节点的树
// 思路：中序遍历数组，重新生成树
func IncresingBST(root *TreeNode) *TreeNode {
	res := &TreeNode{}
	if root == nil {
		return res
	}
	var arr []int
	midleOrder(root, &arr)
	next := &TreeNode{}
	res.Right = next
	for _, i := range arr {
		n := &TreeNode{Val: i}
		next.Right = n
		next = next.Right
	}
	return res.Right.Right
}

func midleOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	midleOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	midleOrder(root.Right, arr)
}

// 求二叉搜索树任意两个节点的差绝对值最小
// 思路：二叉搜索树的中序遍历结果为从小到的排列，在计算相邻元素最小的绝对值
func GetMinimumDifference(root *TreeNode) int {
	res := 9999
	if root == nil {
		return res
	}
	var arr []int
	midleOrder(root, &arr)
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i]-arr[i-1] < res {
			res = arr[i] - arr[i-1]
		}
	}
	return res
}

// 求两颗二叉树是否完全相同
// 思路：同时遍历两个二叉树，注意为空的时候
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	return preOrder(p,q)
}

func preOrder(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil{
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val{
		return false
	}

	return preOrder(p.Left, q.Left) && preOrder(p.Right, q.Right)
}

// 二叉搜索树结点最小距离
// 思路：中序遍历，找到最小差值，非递归遍历↓
func MinDiffInBST(root *TreeNode) int {
	res := 100
	if root == nil{
		return 0
	}
	s := common.NewStack()
	node := root
	last := -1
	for !s.IsEmpty() || node != nil{
		if node != nil{
			s.Push(node)
			node = node.Left
		}else{
			node = s.Pop().(*TreeNode)
			if last <0 {
				last = node.Val
			}else{
				if node.Val - last < res{
					res = node.Val - last
				}
				last = node.Val
			}
			fmt.Println(node.Val)
			node= node.Right
		}
	}
	return res
}

// 求二叉搜索树中是否存在两个节点的和 等于给定的值
// 思路：中序遍历得到有序数组，两个指针快速找到值
func findTarget(root *TreeNode,k int) bool {
	res := false
	arr := []int{}
	midleOrder(root,&arr)
	p := 0
	q := len(arr) - 1
	for p != q{
		if arr[p] + arr[q] == k {
			res = true
			break
		}
		if arr[p] + arr[q] > k{
			q--
			continue
		}
		if arr[p] + arr[q] < k{
			p++
			continue
		}
	}
	return res
}

