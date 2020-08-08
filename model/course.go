package model

func (course *CourseModel) TableName() string {
	return "course"
}

func (course *CourseLikeModel) TableName() string {
	return "course_like"
}

func (course *CourseFavoriteModel) TableName() string {
	return "course_favorite"
}

func (course *CourseModel) GetInfo() error {
	d := DB.Self.Where("id = ?", course.Id).First(course)
	return d.Error
}

// 判断课程是否已经被当前用户点赞
func CourseHasLiked(userId uint32, courseId uint32) (uint32, bool) {
	var data CourseLikeModel
	d := DB.Self.Where("user_id = ? AND course_id = ? ", userId, courseId).First(&data)
	return data.Id, !d.RecordNotFound()
}

// 点赞课程
func CourseLike(userId uint32, courseId uint32) error {
	var data = CourseLikeModel{
		CourseId: courseId,
		UserId:   userId,
	}
	d := DB.Self.Create(&data)
	return d.Error
}

// 取消点赞
func CourseUnlike(id uint32) error {
	var data = CourseLikeModel{Id: id}
	d := DB.Self.Delete(&data)
	return d.Error
}

// 判断课程是否已经被当前用户收藏
func CourseHasFavorited(userId uint32, courseId uint32) (uint32, bool) {
	var data CourseFavoriteModel
	d := DB.Self.Where("user_id = ? AND course_id = ? ", userId, courseId).First(&data)
	return data.Id, !d.RecordNotFound()
}

// 收藏课程
func CourseFavorite(userId uint32, courseId uint32) error {
	var data = CourseFavoriteModel{
		CourseId: courseId,
		UserId:   userId,
	}
	d := DB.Self.Create(&data)
	return d.Error
}

// 取消收藏
func CourseUnfavorite(id uint32) error {
	var data = CourseFavoriteModel{Id: id}
	d := DB.Self.Delete(&data)
	return d.Error
}
