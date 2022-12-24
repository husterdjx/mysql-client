package main

import (
	"fmt"
	"myclient/api"
)

func main() {
	ui()
	loop()
}

/*打印一个ASCII art图像，内容为"husterdjx"*/
func ui() {
	fmt.Println("  __  __           _       _   _      _   _             _ ")
	fmt.Println(" |  \\/  |         | |     | | | |    | | | |           | |")
	fmt.Println(" | \\  / | ___   __| | __ _| |_| | ___| |_| |__   ___   | |")
	fmt.Println(" | |\\/| |/ _ \\ / _` |/ _` | __| |/ _ \\ __| '_ \\ / _ \\  | |")
	fmt.Println(" | |  | | (_) | (_| | (_| | |_| |  __/ |_| | | |  __/  |_|")
	fmt.Println(" |_|  |_|\\___/ \\__,_|\\__,_|\\__|_|\\___|\\__|_| |_|\\___|  (_)")
	fmt.Println("                                                          ")
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
