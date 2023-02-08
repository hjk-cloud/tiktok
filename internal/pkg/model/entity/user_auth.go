package entity

type UserAuth struct {
	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

func (UserAuth) TableName() string {
	return "t_user_auth"
}

