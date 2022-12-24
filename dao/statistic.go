package dao

import "myclient/model"

type Result struct {
	Sdept string
	Grade float64
}

/*统计各个系学生的某课程平均成绩，按系分开*/
func (rs *RdbService) GetAvgGrade(cno string) ([]Result, error) {
	ret := make([]Result, 0)
	err := rs.tx.Raw("select sdept, avg(grade) as avg_grade from sc, student where sc.Sno = student.Sno and sc.Cno = ? group by sdept", cno).Scan(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}

/*统计各个系学生的某课程最好成绩，按系分开*/
func (rs *RdbService) GetMaxGrade(cno string) ([]Result, error) {
	ret := make([]Result, 0)
	err := rs.tx.Raw("select sdept, max(grade) as avg_grade from sc, student where sc.Sno = student.Sno and sc.Cno = ? group by sdept", cno).Scan(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}

/*统计各个系学生的某课程最差成绩，按系分开*/
func (rs *RdbService) GetMinGrade(cno string) ([]Result, error) {
	ret := make([]Result, 0)
	err := rs.tx.Raw("select sdept, min(grade) as avg_grade from sc, student where sc.Sno = student.Sno and sc.Cno = ? group by sdept", cno).Scan(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}

/*统计各个系学生的某课程优秀(90分及以上)率，按系分开*/
func (rs *RdbService) GetExcellentNum(cno string) ([]Result, error) {
	ret := make([]Result, 0)
	err := rs.tx.Raw("select sdept, count(*) as avg_grade from sc, student where sc.Sno = student.Sno and sc.Cno = ? and grade >= 90 group by sdept", cno).Scan(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (rs *RdbService) GetNonExcellentNum(cno string) (map[string]int64, error) {
	num := make([]Result, 0)
	err := rs.tx.Raw("select sdept, count(*) as avg_grade from sc, student where sc.Sno = student.Sno and sc.Cno = ? and grade < 90 group by sdept", cno).Scan(&num).Error
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
func (rs *RdbService) GetFailNum(cno string) ([]Result, error) {
	ret := make([]Result, 0)
	err := rs.tx.Raw("select sdept, count(*) as avg_grade from sc, student where sc.Sno = student.Sno and sc.Cno = ? and grade < 60 group by sdept", cno).Scan(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}

/*按系对学生成绩进行排名，同时返回每个学生的student所有信息、该课程对应的course所有信息和对应的sc表中的所有信息*/
func (rs *RdbService) GetRankByDept(cno string) ([]model.SC, error) {
	ret := make([]model.SC, 0)
	err := rs.tx.Raw("select * from sc, student, course where sc.Sno = student.Sno and sc.Cno = course.Cno and sc.Cno = ? order by grade desc", cno).Scan(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}
