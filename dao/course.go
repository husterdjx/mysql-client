package dao

import "myclient/model"

func (rs *RdbService) CreateCourse(course *model.Course) error {
	return rs.tx.Table("course").Create(&course).Error
}

func (rs *RdbService) GetCourseByName(course *model.Course) error {
	err := rs.tx.Table("course").Where("Cname = ?", course.Cname).First(&course).Error
	if err != nil {
		return err
	}
	return nil
}

func (rs *RdbService) GetCourseByNo(course *model.Course) error {
	err := rs.tx.Table("course").Where("Cno = ?", course.Cno).First(&course).Error
	if err != nil {
		return err
	}
	return nil
}

func (rs *RdbService) GetAllCourse() ([]model.Course, error) {
	courses := make([]model.Course, 0)
	err := rs.tx.Table("course").Find(&courses).Error
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (rs *RdbService) UpdateCourse(course *model.Course) error {
	err := rs.tx.Table("course").Save(&course).Error
	if err != nil {
		return err
	}
	return nil
}

func (rs *RdbService) DeleteCourse(course *model.Course) error { // delete by id
	err := rs.tx.Table("course").Delete(course).Error
	if err != nil {
		return err
	}
	return nil
}

/*删除没有选课的课程信息*/
func (rs *RdbService) AdjustCourse() error {
	err := rs.tx.Exec("delete from course where course.Cno not in (select Cno from sc)").Error
	if err != nil {
		return err
	}
	return nil
}
