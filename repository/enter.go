package repository

type Group struct {
	VideoRepository
	UserRepository
	CommentRepository
	RelationRepository
	// VideoRepository
	// ...
}

var GroupApp = new(Group)
