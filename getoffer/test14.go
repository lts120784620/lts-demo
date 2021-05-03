package main

import "fmt"

func main() {
	data := []int{2,7,4,6,1,3,8,5}
	recorderOddEven(data)
	fmt.Println(data)
}


func recorderOddEven(data []int) {
	length := len(data)
	p1:= 0
	p2 := length -1
	for p1 < p2{

		for p1 < p2 && data[p1] % 2 != 0 {
			p1+=1
		}

		for p1 < p2 && data[p2] % 2 == 0 {
			p2-=1
		}

		t := data[p1]
		data[p1] = data[p2]
		data[p2] = t

	}
}