package test

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"testing"
)

func TestRegister(t *testing.T) {
	password := "password"

	// 定义盐
	salt := "tiktok_salt"

	// 基于Argon2id生成密码的散列值
	key := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)

	// 将key编码为base64字符串
	data := base64.StdEncoding.EncodeToString(key)

	fmt.Println(data)
	fmt.Println(len(data))
}
