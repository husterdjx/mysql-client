package service

import (
	"fmt"
	"myclient/dao"
	"myclient/model"
)

func AddSc() {
	fmt.Println("请输入学生学号：")
	var sno string
	fmt.Scanln(&sno)
	fmt.Println("请输入课程号：")
	var cno string
	fmt.Scanln(&cno)
	fmt.Println("请输入成绩：")
	var score int16
	fmt.Scanln(&score)
	sc := model.SC{
		Sno:   sno,
		Cno:   cno,
		Grade: score,
	}
	err := dao.RS.CreateSC(&sc)
	if err != nil {
		fmt.Printf("添加失败:%v\n", err)
		return
	}
	fmt.Println("添加成功")
}

func DeleteSc() {
	fmt.Println("请输入学生学号：")
	var sno string
	fmt.Scanln(&sno)
	fmt.Println("请输入课程号：")
	var cno string
	fmt.Scanln(&cno)
	sc := model.SC{
		Sno: sno,
		Cno: cno,
	}
	err := dao.RS.DeleteSC(&sc)
	if err != nil {
		fmt.Printf("删除失败:%v\n", err)
		return
	}
	fmt.Println("删除成功")
}

func UpdateSc() {
	fmt.Println("请输入学生学号：")
	var sno string
	fmt.Scanln(&sno)
	fmt.Println("请输入课程号：")
	var cno string
	fmt.Scanln(&cno)
	fmt.Println("请输入成绩：")
	var score int16
	fmt.Scanln(&score)
	sc := model.SC{
		Sno:   sno,
		Cno:   cno,
		Grade: score,
	}
	err := dao.RS.UpdateSC(&sc)
	if err != nil {
		fmt.Printf("更新失败:%v\n", err)
		return
	}
	fmt.Println("更新成功")
}

func QuerySc() {
	fmt.Println("请输入学生学号：")
	var sno string
	fmt.Scanln(&sno)
	scs, err := dao.RS.GetSCBySno(sno)
	if err != nil {
		fmt.Printf("查询失败:%v\n", err)
		return
	}
	fmt.Println("查询成功，结果如下：")
	for _, sc := range scs {
		stu := model.Student{Sno: sc.Sno}
		err := dao.RS.GetStudentByNo(&stu)
		if err != nil {
			fmt.Printf("查询失败:%v\n", err)
			continue
		}
		fmt.Println(sc)
	}
}

func QueryAvg() {
	fmt.Println("请输入课程号：")
	var cno string
	fmt.Scanln(&cno)
	avg, err := dao.RS.GetAvgGrade(cno)
	if err != nil {
		fmt.Printf("查询失败:%v\n", err)
		return
	}
	fmt.Println("查询成功，结果如下：")
	fmt.Println(avg)
}

func QueryMax() {
	fmt.Println("请输入课程号：")
	var cno string
	fmt.Scanln(&cno)
	max, err := dao.RS.GetMaxGrade(cno)
	if err != nil {
		fmt.Printf("查询失败:%v\n", err)
		return
	}
	fmt.Println("查询成功，结果如下：")
	fmt.Println(max)
}

func QueryMin() {
	fmt.Println("请输入课程号：")
	var cno string
	fmt.Scanln(&cno)
	min, err := dao.RS.GetMinGrade(cno)
	if err != nil {
		fmt.Printf("查询失败:%v\n", err)
		return
	}
	fmt.Println("查询成功，结果如下：")
	fmt.Println(min)
}

func QueryFail() {
	fmt.Println("请输入课程号：")
	var cno string
	fmt.Scanln(&cno)
	fails, err := dao.RS.GetFailNum(cno)
	if err != nil {
		fmt.Printf("查询失败:%v\n", err)
		return
	}
	fmt.Println("查询成功，结果如下：")
	for _, fail := range fails {
		fmt.Println(fail)
	}
}

func QueryExcellentRate() {
	fmt.Println("请输入课程号：")
	var cno string
	fmt.Scanln(&cno)
	excellent, err := dao.RS.GetExcellentNum(cno)
	if err != nil {
		fmt.Printf("查询失败:%v\n", err)
		return
	}
	nonexcellent, err := dao.RS.GetNonExcellentNum(cno)
	if err != nil {
		fmt.Printf("查询失败:%v\n", err)
		return
	}
	fmt.Println("查询成功，结果如下：")
	for _, ex := range excellent {
		allnum := int64(ex.Grade) + nonexcellent[ex.Sdept]
		fmt.Println(ex.Sdept, "优秀率为：", float64(ex.Grade)/float64(allnum))
	}
}
