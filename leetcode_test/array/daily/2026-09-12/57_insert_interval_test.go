package daily20260912

import (
	"reflect"
	"testing"
)

/*
LeetCode 57. 插入区间（中等）
链接：https://leetcode.cn/problems/insert-interval/

题目要求：
给定一个已经按照左端点升序排列、且内部互不重叠的闭区间列表 intervals，
再给定一个闭区间 newInterval。将 newInterval 插入 intervals，并在必要时合并重叠区间，
使结果仍然有序且互不重叠。

两个区间只要至少共享一个点，就视为重叠。

示例：

	intervals = [[1,3],[6,9]], newInterval = [2,5]
	返回 [[1,5],[6,9]]

	intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
	返回 [[1,2],[3,10],[12,16]]

约束：

	0 <= len(intervals) <= 10^4
	len(intervals[i]) == 2
	0 <= start <= end <= 10^5
	intervals 已按左端点升序排列，且内部互不重叠
	len(newInterval) == 2

复杂度目标：时间 O(n)。可以创建新的结果数组，不要求原地修改。
今日训练点：利用已有顺序完成一次遍历，并维护“尚未最终确定的插入区间”。
建议限时：35 分钟。
提示触发点：独立思考 15 分钟后，仍然无法区分当前区间在新区间左侧、与新区间重叠、
位于新区间右侧这三种关系时，可索取一级提示。

完成后请补充：
  - 时间复杂度：
  - 空间复杂度：
  - 遍历过程中 newInterval 表示什么：
  - 为什么本题不需要再次排序：
*/
func insert(intervals [][]int, newInterval []int) [][]int {
	// 大概思路是遍历左右、换行遍历，当n1,n0存在可替换的则合并，写入新数组
	res := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		p0, p1 := newInterval[0], newInterval[1]
		n0, n1 := intervals[i][0], intervals[i][1]
		col := []int{n0, n1}
		if n1 < p0 {
			// n 在p的左边，append后面
			res = append(res, col)
		} else if n0 > p1 {
			// n 在 p的右边
			res = append(res, []int{p0, p1})
			res = append(res, intervals[i:]...)
			return res
		} else {
			// 扩大范围，但是先不要往里填充，注意关键的一步是扩大newInterval的范围
			newInterval[0] = min(n0, p0)
			newInterval[1] = max(n1, p1)
		}
	}
	res = append(res, newInterval)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestInsertInterval(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{
			name:        "merge-one",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			want:        [][]int{{1, 5}, {6, 9}},
		},
		{
			name:        "merge-many",
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			want:        [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{name: "insert-before", intervals: [][]int{{3, 5}, {7, 9}}, newInterval: []int{0, 1}, want: [][]int{{0, 1}, {3, 5}, {7, 9}}},
		{name: "insert-after", intervals: [][]int{{1, 2}, {4, 5}}, newInterval: []int{7, 8}, want: [][]int{{1, 2}, {4, 5}, {7, 8}}},
		{name: "contained", intervals: [][]int{{1, 5}}, newInterval: []int{2, 3}, want: [][]int{{1, 5}}},
		{name: "contains-all", intervals: [][]int{{2, 3}, {5, 7}}, newInterval: []int{1, 8}, want: [][]int{{1, 8}}},
		{name: "touching", intervals: [][]int{{1, 2}, {5, 7}}, newInterval: []int{2, 5}, want: [][]int{{1, 7}}},
		{name: "empty", intervals: [][]int{}, newInterval: []int{4, 6}, want: [][]int{{4, 6}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.intervals, tt.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("insert(%v, %v) = %v, want %v", tt.intervals, tt.newInterval, got, tt.want)
			}
		})
	}
}
