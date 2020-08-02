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

func GetMoodScore(userId uint32, year uint32, month uint8) ([]MoodScoreItem, error) {
	var data []MoodScoreItem
	query := DB.Self.Table("mood").Select("day, score").Where("user_id = ? AND year = ? AND month = ?", userId, year, month)
	//fmt.Println("This is query : ", query)
	if err := query.Scan(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
