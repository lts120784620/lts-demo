package tree

/**
No.
描述：
	n叉树结构
思路：
    1、
时间：
    1、2021/10/7
*/

type Node struct {
	Val      int
	Children []*Node
}

func NewTreeByArrays(arr []interface{}) *Node {
	//if arr == nil || len(arr) == 0 {
	//	return nil
	//}
	//return getChildren(arr, &Node{Val: arr[0].(int)})
	return nil
}
