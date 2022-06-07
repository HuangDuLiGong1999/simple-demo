package service

type Group struct {
	FavoriteService
	FeedService
	UserService
	RelationService
	PublishService
	CommentService
	// ...
}

var GroupApp = new(Group)
