package repository

type Group struct {
	VideoRepository
	UserRepository
	CommentRepository
	RelationRepository
	// ...
}

var GroupApp = new(Group)
