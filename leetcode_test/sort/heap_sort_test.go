package sort

import (
	"fmt"
	"testing"

	"lts-demo/leetcode_test/unit"
)

/**
No.
描述：
	堆排序
思路：
    1、堆排序通常用来解决topK问题，分为两个步骤
	2、首先，建堆，将无序的序列转换成堆，即满足父节点大于左右孩子的树结构，正序排列建大顶堆，倒序排列建小顶堆
	3、其次，遍历无序序列，每次调堆完成后，将堆定数据与队尾数据交换，即最大值放在后面，继续调整n-1位。
	4、时间复杂度O(n*logn)
时间：
    1、2021/9/18
*/

func TestHeapSort(t *testing.T) {
	arr := unit.RandomArrays(10)
	fmt.Println(arr)
	fmt.Println(">>>>>>>>>>>>>>>>")
	heapSort(arr)
	fmt.Println(arr)
}

func heapSort(arr []int) {
	length := len(arr)
	// 初始化堆结构，这样才能调整整个堆，变成大顶堆
	// ps:为什么不能省去建堆步骤，直接遍历调堆，因为调的范围不一样，可以理解为建堆是从叶子到跟的完全调整，为了节省步骤，n--的调整也可以用额外的fullFixHeap
	last := len(arr)/2 - 1
	for i := last; i >= 0; i-- {
		fixHeap(arr, i, length)
	}
	// 后续只要调整n--的序列即可
	for i := length - 1; i > 0; i-- {
		unit.ArraysSwap(arr, i, 0)
		// 交换完最大节点，重新调堆，这里是从0开始调，因为跟的以下可能都变了
		length--
		fixHeap(arr, 0, length)
	}

}

// 只调整当前root节点之下，以及受影响的子树，也就是说只会调整一个小子堆
func fixHeap(arr []int, root, len int) {
	l, r := 2*root+1, 2*root+2
	// 大顶堆交换左孩子
	if l < len && arr[l] > arr[root] {
		unit.ArraysSwap(arr, l, root)
		fixHeap(arr, l, len)
	}
	// 大顶堆交换右孩子
	if r < len && arr[r] > arr[root] {
		unit.ArraysSwap(arr, r, root)
		fixHeap(arr, r, len)
	}
}

// 全部调整威力加强版
func fullFixHeap(arr []int, length int) {
	last := length/2 - 1
	for i := last; i >= 0; i-- {
		fixHeap(arr, i, length)
	}
}
