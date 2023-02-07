package controller

var DemoVideos = []Video{
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

var DemoComments = []Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
