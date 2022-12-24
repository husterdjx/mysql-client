package dao

import "myclient/model"

func (rs *RdbService) CreateSC(sc *model.SC) error {
	return rs.tx.Table("sc").Create(&sc).Error
}

func (rs *RdbService) GetAllSC() ([]model.SC, error) {
	scs := make([]model.SC, 0)
	err := rs.tx.Table("sc").Find(&scs).Error
	if err != nil {
		return nil, err
	}
	return scs, nil
}

func (rs *RdbService) UpdateSC(sc *model.SC) error {
	err := rs.tx.Table("sc").Save(&sc).Error
	if err != nil {
		return err
	}
	return nil
}

func (rs *RdbService) DeleteSC(sc *model.SC) error { // delete by id
	err := rs.tx.Table("sc").Delete(sc).Error
	if err != nil {
		return err
	}
	return nil
}

/*输入学号，返回该学生sc表中有关的信息*/
func (rs *RdbService) GetSCBySno(sno string) ([]model.SC, error) {
	scs := make([]model.SC, 0)
	err := rs.tx.Table("sc").Where("Sno = ?", sno).Find(&scs).Error
	if err != nil {
		return nil, err
	}
	return scs, nil
}
