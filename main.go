package main

import (
	"fmt"
	"myclient/api"
)

func main() {
	loop()
}

func loop() {
	for {
		fmt.Println("--------------------欢迎使用学生信息管理系统--------------------")
		fmt.Println("1.学生信息管理")
		fmt.Println("2.课程信息管理")
		fmt.Println("3.成绩信息管理")
		fmt.Println("4.退出")
		fmt.Println("请输入你的选择：")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			api.Student()
		case 2:
			api.Course()
		case 3:
			api.Sc()
		case 4:
			return
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}
}
