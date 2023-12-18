package models

type User struct {
	Id       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
}

type UsernameFollower struct {
	Username         string `json:"username" validate:"required"`
	FollowerUsername string `json:"followerUsername" validate:"required"`
}

type Follower struct {
	Id         uint `gorm:"primaryKey"`
	UserID     uint
	FollowerID uint
}

type CountID struct {
	ID int
}
