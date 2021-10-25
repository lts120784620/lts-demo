package array

import (
	"fmt"
	"sort"
	"testing"
)

/**
No.853. 车队
描述：
		N 辆车沿着一条车道驶向位于target英里之外的共同目的地。
	每辆车i以恒定的速度speed[i]（英里/小时），从初始位置position[i]（英里） 沿车道驶向目的地。
	一辆车永远不会超过前面的另一辆车，但它可以追上去，并与前车以相同的速度紧接着行驶。
	此时，我们会忽略这两辆车之间的距离，也就是说，它们被假定处于相同的位置。
	车队是一些由行驶在相同位置、具有相同速度的车组成的非空集合。注意，一辆车也可以是一个车队。
	即便一辆车在目的地才赶上了一个车队，它们仍然会被视作是同一个车队。
	会有多少车队到达目的地?

	示例：
		输入：target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
		[2,4,12,7,9]
		[2,4,1 ,1,3]
		输出：3
	解释：
		从 10 和 8 开始的车会组成一个车队，它们在 12 处相遇。
		从 0 处开始的车无法追上其它车，所以它自己就是一个车队。
		从 5 和 3 开始的车会组成一个车队，它们在 6 处相遇。
		请注意，在到达目的地之前没有其它车会遇到这些车队，所以答案是 3
思路：
    1、一种思路是，计算出每辆车到达终点的时间，因为距离是排好序的，因此如果序列中出现递增部分，则为新车队
时间：
    1、2021/10/22
*/

func TestCarFleet(t *testing.T) {
	target := 10
	position := []int{0, 4, 2}
	speed := []int{2, 1, 3}
	// 10,6,8
	// 2,1,3
	// 5,6,2

	// 10,8,6
	// 5,2,6
	fmt.Println(carFleet(target, position, speed))
}

func carFleet(target int, position []int, speed []int) int {
	// 定义车，保留数组对应的时间和位置
	type car struct {
		// 距离终点的位置，target - position
		distance int
		// 速度
		speed int
		// 花费的时间
		spent float64
	}

	cars := make([]*car, len(position))
	for i, p := range position {
		c := &car{distance: target - p, speed: speed[i], spent: float64(target-p) / float64(speed[i])}
		cars[i] = c
	}
	// 从大到小排列，也就是说如果距离远且花费时间更少，那一定能追上
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].distance > cars[j].distance
	})

	res := 0
	// 上一个值
	last := -1.0

	// 细节：一定从最远开始向前找，如果出现当前大于前车，说明追不上了
	for i := len(cars) - 1; i >= 0; i-- {
		if cars[i].spent > last {
			res++
			last = cars[i].spent
		}
	}
	return res
}
