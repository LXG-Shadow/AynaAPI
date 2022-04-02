package model

type User struct {
	BaseModel
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
}

type Permission struct {
	BaseModel
	UserID int
	Level  int
	User   User `gorm:"foreignKey:UserID;references:ID"`
}
