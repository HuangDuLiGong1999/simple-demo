package repository

import "fmt"

type User struct {
	Id            int64
	UserName      string
	PassWord      string
	FollowCount   int64
	FollowerCount int64
}

func (User) TableName() string {
	return "user_info"
}

func QueryUserById(id int64) (*User, error) {
	var user User
	err := db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		fmt.Println("Error in repository::QueryUserById")
		return nil, err
	}
	return &user, nil
}
