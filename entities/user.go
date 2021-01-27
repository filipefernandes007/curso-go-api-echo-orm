package entities

type User struct {
	ID       int64
	Username string `gorm:"size:50"`
}
