package dto

type UserAuthDTO struct {
	Username string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserLoginDTO struct {
	UserId int64
	Token  string
}
