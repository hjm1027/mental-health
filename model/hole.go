package model

func (hole *HoleModel) TableName() string {
	return "hole"
}

func (hole *HoleLikeModel) TableName() string {
	return "hole_like"
}

func (hole *HoleFavoriteModel) TableName() string {
	return "hole_favorite"
}

func (hole *HoleReadModel) TableName() string {
	return "hole_read"
}

func (hole *HoleModel) New() error {
	d := DB.Self.Create(hole)
	return d.Error
}

func (hole *HoleModel) GetById() error {
	d := DB.Self.Where("id = ?", hole.Id).First(hole)
	return d.Error
}

func (hole *HoleModel) HasLiked(userId uint32) (uint32, bool) {
	var data HoleLikeModel
	d := DB.Self.Where("user_id = ? AND hole_id = ? ", userId, hole.Id).First(&data)
	return data.Id, !d.RecordNotFound()
}

func (hole *HoleModel) HasFavorited(userId uint32) (uint32, bool) {
	var data HoleFavoriteModel
	d := DB.Self.Where("user_id = ? AND hole_id = ? ", userId, hole.Id).First(&data)
	return data.Id, !d.RecordNotFound()
}

func (hole *HoleModel) HasRead(userId uint32) (uint32, bool) {
	var data HoleReadModel
	d := DB.Self.Where("user_id = ? AND hole_id = ? ", userId, hole.Id).First(&data)
	return data.Id, !d.RecordNotFound()
}

func (hole *HoleReadModel) Read() error {
	d := DB.Self.Create(hole)
	data := HoleModel{Id: hole.HoleId}

	err := data.GetById()
	if err != nil {
		return d.Error
	}
	//fmt.Println(data)
	data.ReadNum += 1
	d = DB.Self.Save(&data)
	return d.Error
}

// 判断问题是否已经被当前用户点赞
func HasLiked(userId uint32, holeId uint32) (uint32, bool) {
	var data HoleLikeModel
	d := DB.Self.Where("user_id = ? AND hole_id = ? ", userId, holeId).First(&data)
	return data.Id, !d.RecordNotFound()
}

// 点赞问题
func Like(userId uint32, holeId uint32) error {
	var data = HoleLikeModel{
		HoleId: holeId,
		UserId: userId,
	}
	d := DB.Self.Create(&data)
	return d.Error
}

// 取消点赞
func Unlike(id uint32) error {
	var data = HoleLikeModel{Id: id}
	d := DB.Self.Delete(&data)
	return d.Error
}

// 判断问题是否已经被当前用户收藏
func HasFavorited(userId uint32, holeId uint32) (uint32, bool) {
	var data HoleFavoriteModel
	d := DB.Self.Where("user_id = ? AND hole_id = ? ", userId, holeId).First(&data)
	return data.Id, !d.RecordNotFound()
}

// 收藏问题
func Favorite(userId uint32, holeId uint32) error {
	var data = HoleFavoriteModel{
		HoleId: holeId,
		UserId: userId,
	}
	d := DB.Self.Create(&data)
	return d.Error
}

// 取消收藏
func Unfavorite(id uint32) error {
	var data = HoleFavoriteModel{Id: id}
	d := DB.Self.Delete(&data)
	return d.Error
}

//获取收藏夹
func GetHoleCollectionsByUserId(userId uint32, limit, page uint32) (*[]HoleFavoriteModel, error) {
	var data []HoleFavoriteModel
	d := DB.Self.Where("user_id = ?", userId).Order("id DESC").Limit(limit).Offset((page - 1) * limit).Find(&data)
	if d.RecordNotFound() {
		return nil, nil
	}
	return &data, d.Error
}

// 获取问题列表
func GetHoleList(userId uint32, limit, page uint32) (*[]HoleModel, error) {
	var data []HoleModel
	d := DB.Self.Table("hole").Order("id DESC").Limit(limit).Offset((page - 1) * limit).Find(&data)
	if d.RecordNotFound() {
		return nil, nil
	}
	return &data, d.Error
}

/*--------------------------------------Comment operation/*--------------------------------------*/
func (hole *ParentCommentModel) TableName() string {
	return "parent_comment"
}

func (hole *SubCommentModel) TableName() string {
	return "sub_comment"
}

func (hole *CommentLikeModel) TableName() string {
	return "comment_like"
}

// 创建父评论
func (comment *ParentCommentModel) New() (uint32, error) {
	d := DB.Self.Create(comment)
	id := comment.Id
	return id, d.Error
}

func (hole *HoleModel) UpdateCommentNum() error {
	hole.CommentNum += 1
	d := DB.Self.Save(&hole)
	return d.Error
}

// Get a parent comment by its id.
func (comment *ParentCommentModel) GetById() error {
	d := DB.Self.First(comment, "id = ?", comment.Id)
	return d.Error
}

// Like a comment by the current user.
func CommentLiking(userId uint32, commentId string) error {
	var data = &CommentLikeModel{
		UserId:    userId,
		CommentId: commentId,
	}
	d := DB.Self.Create(data)
	return d.Error
}

// Cancel liking a comment by the like-record id.
func CommentCancelLiking(id uint32) error {
	var data = CommentLikeModel{Id: id}
	d := DB.Self.Delete(&data)
	return d.Error
}

func CommentHasLiked(userId uint32, commentId string) (uint32, bool) {
	var data CommentLikeModel
	d := DB.Self.Where("user_id = ? AND comment_id = ?", userId, commentId).Find(&data)
	return data.Id, !d.RecordNotFound()
}

// Get comment's total like amount by commentId.
func GetCommentLikeSum(commentId string) (uint32, error) {
	//var data CommentLikeModel
	//Find()和Count()不要连用
	var count uint32
	d := DB.Self.Table("comment_like").Where("comment_id = ?", commentId).Count(&count)
	//fmt.Println(count)
	return count, d.Error
}

// Create a new subComment.
func (comment *SubCommentModel) New() (uint32, error) {
	d := DB.Self.Create(comment)
	id := comment.Id
	return id, d.Error
}

// Update parentComment's the total number of subComment
func (comment *ParentCommentModel) UpdateSubCommentNum(n int) error {
	comment.SubCommentNum += 1
	d := DB.Self.Save(&comment)
	return d.Error
}
