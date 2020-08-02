package model

// 必须要写一个TableName函数，返回table的名字，否则gorm读取不到表。
func (u *MoodModel) TableName() string {
	return "mood"
}

// 添加心情
func (mood *MoodModel) New() error {
	d := DB.Self.Create(mood)
	return d.Error
}

//获取一个月的心情指数
func GetMoodScore(userId uint32, year uint32, month uint8) ([]MoodScoreItem, error) {
	var data []MoodScoreItem
	query := DB.Self.Table("mood").Select("day, score").Where("user_id = ? AND year = ? AND month = ?", userId, year, month)
	//fmt.Println("This is query : ", query)
	if err := query.Scan(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

//获取一个月的心情笔记
func GetMoodNote(userId uint32, year uint32, month uint8) ([]MoodNoteItem, error) {
	var data []MoodNoteItem
	query := DB.Self.Table("mood").Select("date, note").Where("user_id = ? AND year = ? AND month = ?", userId, year, month)
	if err := query.Scan(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
