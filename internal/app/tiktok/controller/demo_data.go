package controller

import (
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
)

var DemoVideos = []vo.Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "http://42.192.195.241:9001/video/2023/2/7/12/cd05492c-af9d-41b6-bced-640848318e1b_1_bandicam 2023-02-05 10-33-42-723 lcwk331_1.mp4",
		CoverUrl:      "http://42.192.195.241:9001/image/2023/2/7/12/cd05492c-af9d-41b6-bced-640848318e1b_1_bandicam 2023-02-05 10-33-42-723 lcwk331_1-cover.png",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []vo.Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = vo.User{
	Id:            219526910052663296,
	Name:          "xhm0215",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}

var ToDemoUser = vo.User{
	Id:            219526947737960448,
	Name:          "xhm02151",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
