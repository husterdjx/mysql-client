package dao

type Result struct {
	Sdept string
	Grade int
}

type ResultI struct {
	Sdept string  `gorm:"column:sdept"`
	Grade float64 `gorm:"column:avg_grade"`
}

/*统计各个系学生的某课程平均成绩，按系分开*/
func (rs *RdbService) GetAvgGrade(cno string) ([]ResultI, error) {
	ret := make([]ResultI, 0)
	err := rs.tx.Raw("select sdept, avg(grade) as avg_grade from sc, student where sc.Sno = student.Sno and sc.Cno = ? group by sdept", cno).Scan(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}

type ResultMax struct {
	Sdept string `gorm:"column:sdept"`
	Grade int    `gorm:"column:max_grade"`
}

/*统计各个系学生的某课程最好成绩，按系分开*/
func (rs *RdbService) GetMaxGrade(cno string) ([]ResultMax, error) {
	ret := make([]ResultMax, 0)
	err := rs.tx.Raw("select sdept, max(grade) as max_grade from sc, student where sc.Sno = student.Sno and sc.Cno = ? group by sdept", cno).Scan(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}

type ResultMin struct {
	Sdept string `gorm:"column:sdept"`
	Grade int    `gorm:"column:min_grade"`
}

/*统计各个系学生的某课程最差成绩，按系分开*/
func (rs *RdbService) GetMinGrade(cno string) ([]ResultMin, error) {
	ret := make([]ResultMin, 0)
	err := rs.tx.Raw("select sdept, min(grade) as min_grade from sc, student where sc.Sno = student.Sno and sc.Cno = ? group by sdept", cno).Scan(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}

type ResultNum struct {
	Sdept string `gorm:"column:sdept"`
	Grade int    `gorm:"column:num"`
}

/*统计各个系学生的某课程优秀(90分及以上)率，按系分开*/
func (rs *RdbService) GetExcellentNum(cno string) ([]ResultNum, error) {
	ret := make([]ResultNum, 0)
	err := rs.tx.Raw("select sdept, count(*) as num from sc, student where sc.Sno = student.Sno and sc.Cno = ? and grade >= 90 group by sdept", cno).Scan(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (rs *RdbService) GetNonExcellentNum(cno string) (map[string]int64, error) {
	num := make([]ResultNum, 0)
	err := rs.tx.Raw("select sdept, count(*) as num from sc, student where sc.Sno = student.Sno and sc.Cno = ? and grade < 90 group by sdept", cno).Scan(&num).Error
	if err != nil {
		return nil, err
	}
	ret := make(map[string]int64)
	for _, v := range num {
		ret[v.Sdept] = int64(v.Grade)
	}
	return ret, nil
}

/*统计各个系学生的某课程不及格(60分以下)人数，按系分开*/
func (rs *RdbService) GetFailNum(cno string) ([]ResultNum, error) {
	ret := make([]ResultNum, 0)
	err := rs.tx.Raw("select sdept, count(*) as num from sc, student where sc.Sno = student.Sno and sc.Cno = ? and grade < 60 group by sdept", cno).Scan(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}

type Rank struct {
	Sno         string `gorm:"column:sno"` // 为了能够直接使用gorm的Scan方法，需要将字段名全部小写
	Sname       string `gorm:"column:sname"`
	Cno         string `gorm:"column:cno"`
	Grade       int    `gorm:"column:grade"`
	Ssex        string `gorm:"column:ssex"`
	Sage        int    `gorm:"column:sage"`
	Sdept       string `gorm:"column:sdept"`
	Cname       string `gorm:"column:cname"`
	Cpno        string `gorm:"column:cpno"`
	Ccredit     int    `gorm:"column:ccredit"`
	Scholarship string `gorm:"column:scholarship"`
}

/*按系对学生成绩进行排名，同时返回每个学生的student所有信息、该课程对应的course所有信息和对应的sc表中的所有信息*/
func (rs *RdbService) GetRankByDept(cno string) ([]Rank, error) {
	ret := make([]Rank, 0)
	err := rs.tx.Raw("select * from sc, student, course where sc.Sno = student.Sno and sc.Cno = course.Cno and sc.Cno = ? order by grade desc", cno).Scan(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}
