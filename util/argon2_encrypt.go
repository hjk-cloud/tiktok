package util

import (
	"encoding/base64"
	"golang.org/x/crypto/argon2"
)

func Argon2Encrypt(password string) string {
	// 定义盐
	salt := "some salt"

	// 基于Argon2id生成密码的散列值
	key := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)

	// 将key编码为base64字符串
	password = base64.StdEncoding.EncodeToString(key)

	return password
}
