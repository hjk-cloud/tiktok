根据token获取用户id

    util.JWTAuth(token string) (int64, error)

根据id获取用户信息
    
    service.GetUserInfoById(userId int64) (*do.UserInfo, error)