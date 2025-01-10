package model

type User struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	AvatarUrl string `json:"avatar_url"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (*User) TableName() string {
	return "user"
}
