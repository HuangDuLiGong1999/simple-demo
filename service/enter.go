package service

type Group struct {
	FavoriteService
	FeedService
	UserService
	PublishService
	CommentService
	//VideoService
	// ...
}

var GroupApp = new(Group)
