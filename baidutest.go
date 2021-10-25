package main

import "fmt"

/**
No.
描述：

思路：
    1、
时间：
    1、2021/10/21
*/

//d := [10,9,6,4,2,1]
//数组d[i]表示到终点距离
//s := [10, 3, 5, 2,4,1]
//s[i] 表示i量车的速度
//最终车队数

func main() {
	d := []float64{10, 9, 6, 4, 2, 1}
	s := []float64{10, 3, 5, 2, 4, 1}
	fmt.Println(getgroups(d, s))
}

func getgroups(dis []float64, speed []float64) int {
	//
	group := 0
	// 远
	for i := 0; i < len(dis); i++ {
		last := -1
		// 近
		for j := i + 1; j < len(dis); j++ {
			d := dis[i] - dis[j]
			s := speed[i] - speed[j]

			if d/s <= dis[j]/speed[j] && i != last {
				group++
				last = j
				break
			}
		}
	}
	return group
}
