package daily20260906

import (
	"reflect"
	"sort"
	"testing"
)

/*
LeetCode 56. 合并区间（中等）
链接：https://leetcode.cn/problems/merge-intervals/

题目要求：
给定若干闭区间 intervals，其中 intervals[i] = [start, end]。合并所有重叠区间，
返回一组互不重叠、且恰好覆盖输入全部区间的结果。

端点相接也视为重叠，例如 [1,4] 和 [4,5] 应合并成 [1,5]。

示例：

	输入：[[1,3],[2,6],[8,10],[15,18]]
	输出：[[1,6],[8,10],[15,18]]

约束：

	1 <= len(intervals) <= 10^4
	len(intervals[i]) == 2
	0 <= start <= end <= 10^4

今日训练点：通过预处理建立可比较的顺序，并维护当前已经合并的区间。
建议限时：35 分钟。
提示触发点：独立思考 15 分钟后，仍然无法保证每次只需和最近的合并结果比较时，
可索取一级提示。

完成后请补充：
  - 时间复杂度：
  - 空间复杂度：
  - 当前合并区间的不变量：
  - 为什么必须先处理输入顺序：
*/
func merge(intervals [][]int) [][]int {
	// TODO: 在这里实现。
	// 先进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	prei := 0
	pre0 := intervals[0][0]
	pre1 := intervals[0][1]
	res = append(res, []int{pre0, pre1})
	for i := 1; i < len(intervals); i++ {
		n0 := intervals[i][0]
		n1 := intervals[i][1]
		if pre1 >= n0 {
			if pre1 < n1 { // 当前1大于上一层的0
				pre1 = n1
				res[prei][1] = n1 //
			}
		} else {
			pre0 = intervals[i][0]
			pre1 = intervals[i][1]
			res = append(res, []int{pre0, pre1})
			prei++
		}
	}
	return res
}

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name:      "basic",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{name: "touching", intervals: [][]int{{1, 4}, {4, 5}}, want: [][]int{{1, 5}}},
		{name: "unsorted", intervals: [][]int{{4, 7}, {1, 4}}, want: [][]int{{1, 7}}},
		{name: "contained", intervals: [][]int{{1, 4}, {2, 3}}, want: [][]int{{1, 4}}},
		{name: "chain", intervals: [][]int{{1, 4}, {0, 2}, {3, 5}}, want: [][]int{{0, 5}}},
		{name: "separate", intervals: [][]int{{1, 2}, {4, 5}}, want: [][]int{{1, 2}, {4, 5}}},
		{name: "single", intervals: [][]int{{2, 3}}, want: [][]int{{2, 3}}},
		{name: "test1", intervals: [][]int{{1, 4}, {0, 0}}, want: [][]int{{0, 0}, {1, 4}}},
		{name: "test2", intervals: [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}, want: [][]int{{1, 10}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("merge(%v) = %v, want %v", tt.intervals, got, tt.want)
			}
		})
	}
}
