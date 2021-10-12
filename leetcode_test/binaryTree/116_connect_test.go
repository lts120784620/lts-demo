package binaryTree

import (
	"fmt"
	"testing"
)

/**
No.116. 填充每个节点的下一个右侧节点指针
描述：
	给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
		struct Node {
		  int val;
		  Node *left;
		  Node *right;
		  Node *next;
		}
		填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
	初始状态下，所有 next 指针都被设置为 NULL。
	进阶：
		你只能使用常量级额外空间。
		使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
思路：
    1、层级遍历：比较简单，按层遍历后指向右节点即可
	2、递归法：太牛逼了
时间：
    1、2021/10/11
*/

func TestConnect(t *testing.T) {
	r := NewTreeByArrays([]interface{}{1, 2, 3, 4, 5, 6, 7})
	connect(r)
	fmt.Println()
}

// 迭代法，层级遍历，利用队列(数组模拟队列)
func connect(root *Node) *Node {
	q := []*Node{}
	// 尾插
	q = append(q, root)
	for len(q) != 0 {
		// 记录每层数量
		width := len(q)
		var right *Node
		for i := 0; i < width; i++ {
			// 头取
			node := q[0]
			q = q[1:]
			if node == nil {
				continue
			}
			// 指向右边
			node.Next = right
			right = node
			// 从右到左塞进去
			q = append(q, node.Right)
			q = append(q, node.Left)
		}
	}
	return root
}

// 递归法：前序递归遍历左右节点，左->右。非公共子树的右节点->左节点即可
func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectBetween(root.Left, root.Right)
	return root
}

func connectBetween(left, right *Node) {
	if left == nil || right == nil {
		return
	}
	// 连接
	left.Next = right
	connectBetween(left.Left, left.Right)
	connectBetween(right.Left, right.Right)
	// 两个子树右左相连
	connectBetween(left.Right, right.Left)
}
