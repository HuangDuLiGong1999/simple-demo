package service

type Group struct {
	FavoriteService
	FeedService
	UserService
	RelationService
	PublishService
	CommentService
	//VideoService
	// ...
}

var GroupApp = new(Group)
