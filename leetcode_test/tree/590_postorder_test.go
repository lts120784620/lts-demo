package tree

import (
	"fmt"
	"testing"
)

/**
No.590 N 叉树的后序遍历
描述：
	给定一个 N 叉树，返回其节点值的 后序遍历 。
	N 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。

	进阶：
	递归法很简单，你可以使用迭代法完成此题吗?
思路：
    1、迭代法，先递归子节点，在塞值
时间：
    1、2021/10/7
*/

func TestPostorder(t *testing.T) {
	r := &Node{Val: 1}
	r.Children = []*Node{
		{Val: 3, Children: []*Node{
			{Val: 5},
			{Val: 6},
		}},
		{Val: 2},
		{Val: 4},
	}
	fmt.Println(postorder(r))
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	for _, c := range root.Children {
		res = append(res, postorder(c)...)
	}
	res = append(res, root.Val)
	return res
}
