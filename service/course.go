package service

import (
	"fmt"
	"myclient/dao"
	"myclient/model"
)

func AddCourse() {
	fmt.Println("请输入课程号：")
	var cno string
	fmt.Scanln(&cno)
	fmt.Println("请输入课程名：")
	var cname string
	fmt.Scanln(&cname)
	fmt.Println("请输入学分：")
	var credit int16
	fmt.Scanln(&credit)
	fmt.Println("请输入课程前置课程：")
	var pre string
	fmt.Scanln(&pre)
	course := model.Course{Cno: cno, Cname: cname, Ccredit: credit, Cpno: pre}
	err := dao.RS.CreateCourse(&course)
	if err != nil {
		fmt.Printf("添加失败:%v\n", err)
		return
	}
	fmt.Println("添加成功")
}

func DeleteCourse() {
	fmt.Println("请输入课程号：")
	var cno string
	fmt.Scanln(&cno)
	course := model.Course{Cno: cno}
	err := dao.RS.DeleteCourse(&course)
	if err != nil {
		fmt.Printf("删除失败:%v\n", err)
		return
	}
	fmt.Println("删除成功")
}

func UpdateCourse() {
	fmt.Println("请输入课程号：")
	var cno string
	fmt.Scanln(&cno)
	fmt.Println("请输入课程名：")
	var cname string
	fmt.Scanln(&cname)
	fmt.Println("请输入学分：")
	var credit int16
	fmt.Scanln(&credit)
	fmt.Println("请输入课程前置课程：")
	var pre string
	fmt.Scanln(&pre)
	course := model.Course{Cno: cno, Cname: cname, Ccredit: credit, Cpno: pre}
	err := dao.RS.UpdateCourse(&course)
	if err != nil {
		fmt.Printf("更新失败:%v\n", err)
		return
	}
	fmt.Println("更新成功")
}

func QueryCourse() {
	fmt.Println("请输入课程号：")
	var cno string
	fmt.Scanln(&cno)
	course := model.Course{Cno: cno}
	err := dao.RS.GetCourseByNo(&course)
	if err != nil {
		fmt.Printf("查询失败:%v\n", err)
		return
	}
	fmt.Println("查询成功")
	fmt.Println(course)
}

func AdjustCourse() {
	err := dao.RS.AdjustCourse()
	if err != nil {
		fmt.Printf("调整失败:%v\n", err)
		return
	}
	fmt.Println("调整成功")
}
