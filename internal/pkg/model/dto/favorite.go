package dto

type FavoriteActionDTO struct {
	Token      string
	VideoId    int64
	ActionType bool
}

type FavoriteListDTO struct {
	UserId int64
	Token  string
}
