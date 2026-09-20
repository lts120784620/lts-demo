package daily20260919

import (
	"reflect"
	"sort"
	"testing"
)

/*
LeetCode 56. 合并区间（中等）
链接：https://leetcode.cn/problems/merge-intervals/

题目要求：
给定若干区间 intervals，合并所有重叠区间，并返回互不重叠且覆盖原输入的区间。

复杂度目标：时间 O(n log n)，额外空间按输出之外 O(log n) 或取决于排序实现。
复习要求：先确定排序依据，再用一次线性扫描维护当前合并区间。
建议限时：25 分钟。
提示触发点：独立思考 12 分钟后，仍然需要反复扫描所有已合并区间寻找重叠时，
可索取一级提示。

完成评价（2026-09-20）：
  - 测试结果：通过（5 个测试用例）
  - 按什么字段排序：按区间左端点升序
  - 当前区间与最后一个已合并区间在什么条件下重叠：新区间左端点 n0 <= 当前右边界 p1
  - 重叠时如何更新右边界：p1 = max(p1, n1)
  - 时间复杂度及原因：O(n log n)，排序为 O(n log n)，扫描为 O(n)
  - 空间复杂度：O(n) 用于结果；排序额外空间取决于实现
  - 核心模式与不变量：[p0, p1] 始终表示已扫描区间合并后尚未写入结果的最后一段
  - 提示使用：四级
  - 主要错误类型：实现、边界
  - 本次掌握度（0～3）：1
  - 下次复习日期：2026-09-22
  - 备注：初版在分段时提前追加新区间且返回值错误；完整修正后通过全部测试，耗时未记录
*/
func mergeIntervalsReview(intervals [][]int) [][]int {
	// 按照0位排序，后面只要比较第1位和下行的第0位就可以了
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	// 比较第1位和下行的第0位就可以了，还要注意，如果没有出现边界，则不能append到res
	p0, p1 := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		n0, n1 := intervals[i][0], intervals[i][1]
		if n0 <= p1 {
			p0, p1 = min(p0, n0), max(p1, n1)
		} else {
			res = append(res, []int{p0, p1})
			p0, p1 = n0, n1
		}
	}
	res = append(res, []int{p0, p1})
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func TestMergeIntervalsReview(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{name: "basic", intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, want: [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{name: "touching", intervals: [][]int{{1, 4}, {4, 5}}, want: [][]int{{1, 5}}},
		{name: "contained", intervals: [][]int{{1, 10}, {2, 3}, {4, 8}}, want: [][]int{{1, 10}}},
		{name: "unsorted", intervals: [][]int{{8, 10}, {1, 3}, {2, 6}}, want: [][]int{{1, 6}, {8, 10}}},
		{name: "single", intervals: [][]int{{1, 2}}, want: [][]int{{1, 2}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeIntervalsReview(tt.intervals)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("mergeIntervalsReview(%v) = %v, want %v", tt.intervals, got, tt.want)
			}
		})
	}
}
