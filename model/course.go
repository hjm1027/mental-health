package model

import "fmt"

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

//更新点赞数
func (course *CourseModel) UpdateLikeNum(n int) error {
	if n == 1 {
		course.LikeNum += 1
	} else if n == -1 {
		course.LikeNum -= 1
	}

	d := DB.Self.Save(&course)
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

//更新收藏数
func (course *CourseModel) UpdateFavoriteNum(n int) error {
	if n == 1 {
		course.FavoriteNum += 1
	} else if n == -1 {
		course.FavoriteNum -= 1
	}

	d := DB.Self.Save(&course)
	return d.Error
}

//获取所有点赞课程
func GetCourseLikeCollectionsByUserId(userId uint32, limit, page uint32) (*[]CourseLikeModel, error) {
	var data []CourseLikeModel
	d := DB.Self.Where("user_id = ?", userId).Order("id DESC").Limit(limit).Offset((page - 1) * limit).Find(&data)
	if d.RecordNotFound() {
		return nil, nil
	}
	return &data, d.Error
}

//获取所有收藏课程
func GetCourseFavoriteCollectionsByUserId(userId uint32, limit, page uint32) (*[]CourseFavoriteModel, error) {
	var data []CourseFavoriteModel
	d := DB.Self.Where("user_id = ?", userId).Order("id DESC").Limit(limit).Offset((page - 1) * limit).Find(&data)
	if d.RecordNotFound() {
		return nil, nil
	}
	return &data, d.Error
}

/*------------------------------------------search operation------------------------------------------*/
//根据关键词搜索
func AgainstAndMatchCourses(kw string, page, limit uint64, t string) ([]CourseModel, error) {
	courses := &[]CourseModel{}
	where := "MATCH (course.name,course.source) AGAINST ('" + kw + "') "
	d := DB.Self.Debug().Table("course").
		Where(where).Order(t).Limit(limit).Offset((page - 1) * limit).Find(courses)
	if d.RecordNotFound() {
		return nil, nil
	}
	return *courses, nil
}

// 获取所有课程
func AllCourses(page, limit uint64, t string) ([]CourseModel, error) {
	//fmt.Println(t)
	courses := &[]CourseModel{}
	d := DB.Self.Table("course").Order(t).Limit(limit).Offset((page - 1) * limit).Find(&courses)
	if d.RecordNotFound() {
		return nil, nil
	}
	return *courses, nil
}

// 获取所有问题
func AllHoles(page, limit uint64, t uint8) ([]HoleModel, error) {
	//fmt.Println(t)
	holes := &[]HoleModel{}
	if t != 0 {
		d := DB.Self.Table("hole").Where("type = ?", t).Limit(limit).Offset((page - 1) * limit).Find(&holes)
		if d.RecordNotFound() {
			return nil, nil
		}
	} else {
		d := DB.Self.Table("hole").Limit(limit).Offset((page - 1) * limit).Find(&holes)
		if d.RecordNotFound() {
			return nil, nil
		}
	}
	return *holes, nil
}

//根据关键词搜索
func SearchHoles(kw string, page, limit uint64, t uint8) ([]HoleModel, error) {
	holes := &[]HoleModel{}
	where := "MATCH (hole.name,hole.content) AGAINST ('" + kw + "') "
	d := DB.Self.Debug().Table("hole").
		Where(where).Limit(limit).Offset((page - 1) * limit).Find(holes)
	if d.RecordNotFound() {
		fmt.Println("wu")
		return nil, nil
	}
	return *holes, nil
}
