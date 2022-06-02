package repository

type Group struct {
	VideoRepository
	UserRepository
	CommentRepository
	// VideoRepository
	// ...
}

var GroupApp = new(Group)
