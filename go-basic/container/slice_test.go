package container

import (
	"fmt"
	"testing"
)

/**
No.
描述：
	slice测试
思路：
    1、
时间：
    1、2021/9/24
*/

func TestSlice(t *testing.T) {
	s := []int{7}      //[7]
	s = append(s, 1)   //[7 1]
	s = append(s, 2)   //此时切片s的len=3；cap=3 [7 1 2]
	x := append(s, 10) //此时是对 切片[7 1 2]的操作
	fmt.Println(s, x)
	y := append(s, 20) //此时也是对 切片[7 1 2]的操作
	//x,y两个切片底层的数组指针 指向的是同一块数组存储空间 所以上述语句执行完后 x,y的输出是一样的。
	//但是注意当cap容量发生变化时，数组指针(地址)改变，例如 y := append(s, 20，30)时x,y的输出就会不一样 [7 1 2 10] [7 1 2 20 30]
	fmt.Println(s, y)
	fmt.Println(s, x, y)
	//输出结果：[7 1 2] [7 1 2 20] [7 1 2 20]

	a := []int{1}
	fmt.Println(a)
	add(a)
	fmt.Println(a)
}

func add(arr []int) {
	arr[0] = 9
	arr = append(arr, 9)
	//arr = append(arr, 9,9,9,9,9,9,9)
}
