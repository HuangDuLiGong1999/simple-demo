package model

type Comment struct {
	Id         int64  `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true"`
	UserId     int64  `json:"user_id"`
	VideoId    int64  `json:"video_id"`
	User       User   `json:"user" gorm:"foreignKey:user_id;references:id;"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

func (Comment) TableName() string {
	return "comment_info"
}
