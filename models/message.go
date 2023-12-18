package models

type Message struct {
	Id        int
	Username  string `gorm:"column:username"`
	Text      string `gorm:"column:text" validate:"max=250"`
	Timestamp string `gorm:"column:timestamp;default:current_timestamp"`
}
