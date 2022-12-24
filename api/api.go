package api

import (
	"fmt"
	"myclient/service"
)

func Student() {
	for {
		fmt.Println("----------------------------------------")
		fmt.Println("1.添加学生信息")
		fmt.Println("2.删除学生信息")
		fmt.Println("3.修改学生信息")
		fmt.Println("4.查询学生信息")
		fmt.Println("5.返回上一级")
		fmt.Println("请输入你的选择：")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			service.AddStudent()
		case 2:
			service.DeleteStudent()
		case 3:
			service.UpdateStudent()
		case 4:
			service.QueryStudent()
		case 5:
			return
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}
}

func Course() {
	for {
		fmt.Println("----------------------------------------")
		fmt.Println("1.添加课程信息")
		fmt.Println("2.删除课程信息")
		fmt.Println("3.修改课程信息")
		fmt.Println("4.查询课程信息")
		fmt.Println("5.删除冗余课程信息")
		fmt.Println("6.返回上一级")
		fmt.Println("请输入你的选择：")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			service.AddCourse()
		case 2:
			service.DeleteCourse()
		case 3:
			service.UpdateCourse()
		case 4:
			service.QueryCourse()
		case 5:
			service.AddCourse()
		case 6:
			return
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}
}

func Sc() {
	for {
		fmt.Println("----------------------------------------")
		fmt.Println("1.添加成绩信息")
		fmt.Println("2.删除成绩信息")
		fmt.Println("3.修改成绩信息")
		fmt.Println("4.输入学号，查询某学生的基本信息和成绩信息")
		fmt.Println("5.查询某科目的各系平均成绩")
		fmt.Println("6.查询某科目的各系最高成绩")
		fmt.Println("7.查询某科目的各系最低成绩")
		fmt.Println("8.查询某科目的各系优秀率")
		fmt.Println("9.查询某科目的各系不及格人数")
		fmt.Println("10.返回上一级")
		fmt.Println("请输入你的选择：")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			service.AddSc()
		case 2:
			service.DeleteSc()
		case 3:
			service.UpdateSc()
		case 4:
			service.QuerySc()
		case 5:
			service.QueryAvg()
		case 6:
			service.QueryMax()
		case 7:
			service.QueryMin()
		case 8:
			service.QueryExcellentRate()
		case 9:
			service.QueryFail()
		case 10:
			return
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}
}
