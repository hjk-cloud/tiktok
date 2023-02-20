package dto

type FollowActionDTO struct {
	Token      string
	ToUserId   int64
	ActionType bool
}

type FollowRelationDTO struct {
	Token  string
	UserId int64
}
