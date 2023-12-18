package models

type Timeline struct {
	Username string `json:"username" validate:"required"`
}

type FollowerPerUser struct {
	Id               int
	FollowerID       int    `gorm:"column:follower_id"`
	FollowerUsername string `gorm:"column:follower_username"`
}

type Feed struct {
	Username  string `gorm:"column:username"`
	Text      string `gorm:"column:text"`
	Timestamp string `gorm:"column:timestamp;type:timestamp;format:2006-01-02 15:04:05"`
}
