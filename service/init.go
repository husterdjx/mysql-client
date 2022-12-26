package service

import (
	"math/rand"
	"myclient/dao"
	"myclient/model"
)

func RandInt(max int) int16 {
	return int16(rand.Intn(max))
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

/*随机生成"男"或者"女"*/
func RandSex() string {
	if RandInt(2) == 0 {
		return "男"
	} else {
		return "女"
	}
}

/*随机产生英文名*/
func RandName() string {
	return RandStringRunes(10)
}

/*随机产生9位学号*/
func RandSno() string {
	var letterRunes = []rune("0123456789")
	b := make([]rune, 9)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

/*随机产生2位以内的课程号*/
func RandCno() string {
	var letterRunes = []rune("0123456789")
	b := make([]rune, 2)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

/*随机产生院系名，要求在CS、Philosophy、Logic、Religious Studies、Economics、Insurance、Law、Software等选择中随机产生*/
func RandDept() string {
	var depts = []string{"CS", "Philosophy", "Logic", "Religious Studies", "Economics", "Insurance", "Law", "Software"}
	return depts[rand.Intn(len(depts))]
}

/*向数据库注入大量随机信息*/
func InitDB() {
	InitStudent()
	InitCourse()
	InitSc()
}

func InitStudent() {
	// 生成随机学生信息
	for i := 0; i < 100; i++ {
		student := model.Student{
			Sno:         RandSno(),
			Sname:       RandName(),
			Ssex:        RandSex(),
			Sage:        RandInt(10) + 18,
			Sdept:       RandDept(),
			Scholarship: "否",
		}
		// 将学生信息插入数据库
		err := dao.RS.CreateStudent([]*model.Student{&student})
		if err != nil {
			continue
		}
	}
}

func InitCourse() {
	// 生成随机课程信息
	for i := 0; i < 100; i++ {
		course := model.Course{
			Cno:     RandCno(),
			Cname:   RandStringRunes(10),
			Ccredit: RandInt(3) + 1,
		}
		// 将课程信息插入数据库
		err := dao.RS.CreateCourse(&course)
		if err != nil {
			continue
		}
	}
}

func InitSc() {
	// 随机从数据库中选取学生和课程，生成随机成绩
	students, err := dao.RS.GetAllStudent()
	if err != nil {
		return
	}
	courses, err := dao.RS.GetAllCourse()
	if err != nil {
		return
	}
	for i := 0; i < 100; i++ {
		student := students[rand.Intn(len(students))]
		course := courses[rand.Intn(len(courses))]
		sc := model.SC{
			Sno:   student.Sno,
			Cno:   course.Cno,
			Grade: RandInt(100),
		}
		// 将成绩信息插入数据库
		err := dao.RS.CreateSC(&sc)
		if err != nil {
			continue
		}
	}
}
