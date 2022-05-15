package service

import (
	"github.com/RaymondCode/simple-demo/repository"
)

func QueryUser(userId int64, token string) (User, error) {
	rawUser, err := repository.QueryUserById(userId)
	if err != nil {
		return User{}, err
	}
	user := User{
		Id:            rawUser.Id,
		Name:          rawUser.Username,
		FollowCount:   rawUser.FollowCount,
		FollowerCount: rawUser.FollowerCount,
		IsFollow:      false,
	}
	return user, nil
}
