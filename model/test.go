package model

func (test *TestModel) TableName() string {
	return "test"
}

// 添加测评
func (test *TestModel) New() error {
	d := DB.Self.Create(test)
	return d.Error
}

func (test *TestModel) GetById() error {
	d := DB.Self.Where("id = ?", test.Id).First(test)
	return d.Error
}

func GetList(limit, page uint32) (*[]TestModel, error) {
	var tests []TestModel
	d := DB.Self.Table("test").Order("id DESC").Limit(limit).Offset((page - 1) * limit).Find(&tests)
	if d.RecordNotFound() {
		return &tests, nil
	}
	return &tests, d.Error
}
