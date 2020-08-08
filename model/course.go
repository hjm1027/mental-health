package model

func (course *CourseModel) TableName() string {
	return "course"
}

func (course *CourseModel) GetInfo() error {
	d := DB.Self.Where("id = ?", course.Id).First(course)
	return d.Error
}
