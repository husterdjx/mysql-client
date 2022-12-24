package dao

import "myclient/model"

func (rs *RdbService) CreateStudent(stu []*model.Student) error {
	return rs.tx.Table("student").Create(&stu).Error
}

func (rs *RdbService) GetStudentByName(stu *model.Student) error {
	err := rs.tx.Table("student").Where("Sname = ?", stu.Sname).First(&stu).Error
	if err != nil {
		return err
	}
	return nil
}

func (rs *RdbService) GetStudentByNo(stu *model.Student) error {
	err := rs.tx.Table("student").Where("Sno = ?", stu.Sno).First(&stu).Error
	if err != nil {
		return err
	}
	return nil
}

func (rs *RdbService) GetAllStudent() ([]model.Student, error) {
	students := make([]model.Student, 0)
	err := rs.tx.Table("student").Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (rs *RdbService) UpdateStudent(stu *model.Student) error {
	err := rs.tx.Table("student").Save(&stu).Error
	if err != nil {
		return err
	}
	return nil
}

func (rs *RdbService) DeleteStudent(stu *model.Student) error { // delete by id
	err := rs.tx.Table("student").Delete(stu).Error
	if err != nil {
		return err
	}
	return nil
}
