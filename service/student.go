package service

import (
	"fmt"
	"myclient/dao"
	"myclient/model"
)

func AddStudent() {
	fmt.Println("请输入学生学号：")
	var sno string
	fmt.Scanln(&sno)
	fmt.Println("请输入学生姓名：")
	var sname string
	fmt.Scanln(&sname)
	fmt.Println("请输入学生性别：")
	var ssex string
	fmt.Scanln(&ssex)
	fmt.Println("请输入学生年龄：")
	var sage int16
	fmt.Scanln(&sage)
	fmt.Println("请输入学生院系：")
	var sdept string
	fmt.Scanln(&sdept)
	fmt.Println("请输入学生是否有奖学金：")
	var scholarship string
	fmt.Scanln(&scholarship)
	stu := model.Student{
		Sno:         sno,
		Sname:       sname,
		Ssex:        ssex,
		Sage:        sage,
		Sdept:       sdept,
		Scholarship: scholarship,
	}
	fmt.Println(stu)
	err := dao.RS.CreateStudent([]*model.Student{&stu})
	if err != nil {
		fmt.Printf("添加失败:%v\n", err)
		return
	}
	fmt.Println("添加成功")
}

func DeleteStudent() {
	fmt.Println("请输入学生学号：")
	var sno string
	fmt.Scanln(&sno)
	stu := model.Student{
		Sno: sno,
	}
	err := dao.RS.DeleteStudent(&stu)
	if err != nil {
		fmt.Printf("删除失败:%v\n", err)
		return
	}
	fmt.Println("删除成功")
}

func UpdateStudent() {
	fmt.Println("请输入学生学号：")
	var sno string
	fmt.Scanln(&sno)
	fmt.Println("请输入学生姓名：")
	var sname string
	fmt.Scanln(&sname)
	fmt.Println("请输入学生性别：")
	var ssex string
	fmt.Scanln(&ssex)
	fmt.Println("请输入学生年龄：")
	var sage int16
	fmt.Scanln(&sage)
	fmt.Println("请输入学生院系：")
	var sdept string
	fmt.Scanln(&sdept)
	fmt.Println("请输入学生是否有奖学金：")
	var scholarship string
	fmt.Scanln(&scholarship)
	stu := model.Student{
		Sno:         sno,
		Sname:       sname,
		Ssex:        ssex,
		Sage:        sage,
		Sdept:       sdept,
		Scholarship: scholarship,
	}
	err := dao.RS.UpdateStudent(&stu)
	if err != nil {
		fmt.Printf("修改失败:%v\n", err)
		return
	}
	fmt.Println("修改成功")
}

func QueryStudent() {
	fmt.Println("请输入学生学号：")
	var sno string
	fmt.Scanln(&sno)
	stu := model.Student{
		Sno: sno,
	}
	err := dao.RS.GetStudentByNo(&stu)
	if err != nil {
		fmt.Printf("查询失败:%v\n", err)
		return
	}
	fmt.Println("查询成功")
}
