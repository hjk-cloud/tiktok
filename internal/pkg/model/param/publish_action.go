package param

import "mime/multipart"

type PublishActionParam struct {
	Token  string
	Title  string
	Data   *multipart.FileHeader
	UserId int64
}
