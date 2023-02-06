package controller

var DemoVideos = []Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "http://192.168.1.2:8080/static/2023/2/5/21/e1c218f4-3288-4f17-9117-0f7f5e1407ac_1_tiktok.mp4",
		CoverUrl:      "http://192.168.1.2:8080/static/2023/2/5/21/e1c218f4-3288-4f17-9117-0f7f5e1407ac_1_tiktok-cover.png",
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
