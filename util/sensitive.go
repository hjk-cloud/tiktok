package util

import (
	"errors"
	"fmt"
	"os"

	"github.com/importcjj/sensitive"
)

func CheckSensitive(text string) error {
	workDir, _ := os.Getwd() //获取目录对应的路径
	filter := sensitive.New()
	filter.LoadWordDict(workDir + "/../../util/sensitive-dict.txt")
	// fmt.Println(workDir + "/../../util/sensitive-dict.txt")
	// filter.LoadNetWordDict("https://raw.githubusercontent.com/importcjj/sensitive/master/dict/dict.txt")
	ok, err := filter.Validate(text)
	// fmt.Println("敏感词", ok, err, text)
	if !ok {
		return errors.New(fmt.Sprintf("含有敏感词：%s", err))
	}
	return nil
}
