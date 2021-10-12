package binaryTree

import (
	"fmt"
	"testing"
)

/**
No.652. 寻找重复的子树
描述：
	给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
	两棵树重复是指它们具有相同的结构以及相同的结点值。
	示例 1：
			1
		   / \
		  2   3
		 /   / \
		4   2   4
		   /
		  4
	下面是两个重复的子树：

		  2
		 /
		4
	和
		4
	因此，你需要以列表的形式返回上述重复子树的根结点。
思路：
    1、
时间：
    1、2021/10/11
*/

func TestFindDuplicateSubtrees(t *testing.T) {
	r := NewBinTreeByArrays([]interface{}{1, 2, 3, 4, nil, 2, 4, nil, nil, nil, nil, 4, nil, nil, nil})
	a := findDuplicateSubtrees(r)
	for _, v := range a {
		fmt.Println(v.String())
	}
}

var m = make(map[string]int, 0)
var res = []*TreeNode{}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {

	if root == nil {
		return res
	}

	reverseSubtrees(root)

	return res
}

func reverseSubtrees(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	left := reverseSubtrees(root.Left)
	right := reverseSubtrees(root.Right)

	sub := fmt.Sprintf("%s#%s#%d", left, right, root.Val)
	if cnt, ok := m[sub]; ok && cnt == 1 {
		res = append(res, root)
	}
	m[sub] += 1

	return sub
}
