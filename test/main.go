package main

import "time"

func main() {
	i:=0
	for true{
		i++
		time.Sleep(1000 * 1000 * 1000)
		println("文件传输中", i ,"，请勿离开。")
		println("souf------ <<<<<<<< ------------ >>>>>>>")
	}
}
