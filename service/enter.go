package service

type Group struct {
	FavoriteService
	FeedService
	UserService
	RelationService
	PublishService
<<<<<<< HEAD
	CommentService
=======
	RelationService
>>>>>>> zzz
	//VideoService
	// ...
}

var GroupApp = new(Group)
