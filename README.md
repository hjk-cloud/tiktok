# tiktok 

## 抖音项目服务端简单示例

具体功能内容参考飞书说明文档

工程无其他依赖，直接编译运行即可

```shell
# tiktok服务端
cd ./cmd/tiktok
go build && ./tiktok

# tiktok定时任务删除/web下到期视频和图片文件
cd ./cmd/tiktok-schedule
go build && ./tiktok-schedule
```

### 完成功能

1.可以进行注册以及登录，发布视频；

2.对视频点赞/取消点赞

3.关注其他用户/取关其他用户

4.对于互关用户，可以进行聊天

5.查看自己投稿的视频、点赞的视频、粉丝、已关注用户

