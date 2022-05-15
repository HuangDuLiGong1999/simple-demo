package repository

type User struct {
	Id            int64
	Username      string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
}

func (User) TableName() string {
	return "user_info"
}

func QueryUserById(id int64) (*User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
