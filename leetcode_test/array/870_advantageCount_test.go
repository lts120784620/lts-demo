package array

import (
	"fmt"
	"sort"
	"testing"
)

/**
No.870. 优势洗牌
描述：
	给定两个大小相等的数组A和B，A 相对于 B 的优势可以用满足A[i] > B[i]的索引 i的数目来描述。返回A的任意排列，使其相对于 B的优势最大化。
	示例 1：
	输入：A = [2,7,11,15], B = [1,10,4,11]
	输出：[2,11,7,15]

	示例 2：
	输入：A = [12,24,8,32], B = [13,25,32,11]
	输出：[24,32,8,12]

思路：
    1、田忌赛马思路，用我最好的去打他最好的，如果打不过，就用最差的去送人头
时间：
    1、2021/10/12
*/

func TestAdvantageCount(t *testing.T) {
	a := []int{12, 24, 8, 32}
	b := []int{13, 25, 32, 11}
	fmt.Println(advantageCount(a, b))
}

func advantageCount(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums2))
	// a排序
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] > nums1[j]
	})

	// b排序，要保存原位置，用来返回nums1位置
	// quickSort
	type Pair struct {
		Key int //原值
		Val int //原下标
	}
	n2 := []*Pair{}
	for i, v := range nums2 {
		n2 = append(n2, &Pair{Key: v, Val: i})
	}
	sort.Slice(n2, func(i, j int) bool {
		return n2[i].Key > n2[j].Key
	})

	// 比较过程，
	last := len(nums1) - 1
	n1best := 0
	n2best := 0
	for n1best <= last {
		n2Pos := n2[n2best].Val
		// 我最好的能打过你最好的
		if nums1[n1best] > n2[n2best].Key {
			res[n2Pos] = nums1[n1best]
			n1best++
		} else {
			// 没有你的好，我就找个最差的送人头
			res[n2Pos] = nums1[last]
			last--
		}
		n2best++
	}

	return res
}
