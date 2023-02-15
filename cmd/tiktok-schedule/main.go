package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/hjk-cloud/tiktok/config"
)

func main() {
	// 删除前days天的目录
	days := time.Duration(1)
	earliest := time.Now().Local().Add(-days * 24 * time.Hour)
	year, month, day := earliest.Year(), int(earliest.Month()), earliest.Day()
	fmt.Println(config.Config.StaticDir, "###删到", year, month, day)

	yfiles, err := ioutil.ReadDir(config.Config.StaticDir)
	if err != nil {
		log.Fatal(err)
	}
	// 所有年
	for _, yf := range yfiles {
		// fmt.Println(yf.Name())
		y, err := strconv.Atoi(yf.Name())
		if err != nil {
			log.Println(err)
			continue
		}
		ypath := path.Join(config.Config.StaticDir, yf.Name())
		if y < year {
			fmt.Println("删除", ypath)
			os.RemoveAll(ypath)
			continue
		}
		if y > year {
			continue
		}
		mfiles, err := ioutil.ReadDir(ypath)
		if err != nil {
			log.Println(err)
			break
		}
		// 所有月
		for _, mf := range mfiles {
			// fmt.Println(mf.Name())
			m, err := strconv.Atoi(mf.Name())
			if err != nil {
				log.Println(err)
				continue
			}
			mpath := path.Join(ypath, mf.Name())
			if m < month {
				fmt.Println("删除", mpath)
				os.RemoveAll(mpath)
				continue
			}
			if m > month {
				continue
			}
			dfiles, err := ioutil.ReadDir(mpath)
			if err != nil {
				log.Println(err)
				break
			}
			// 所有日
			for _, df := range dfiles {
				// fmt.Println(mf.Name())
				d, err := strconv.Atoi(df.Name())
				if err != nil {
					log.Println(err)
					continue
				}
				dpath := path.Join(mpath, df.Name())
				if d <= day {
					fmt.Println("删除", dpath)
					os.RemoveAll(dpath)
					continue
				}
			}
		}
	}

	// os.RemoveAll(dir)
	// if err := os.MkdirAll(dir, 0755); err != nil {
	// 	return "", "", err
	// }
	time.Sleep(30 * time.Second)
}

/*
mkdir -p 2022/11/30 2022/12/1 2023/2/2 2023/2/3 2023/2/4 2023/2/5 2023/2/6 2023/2/7 2023/2/8 2023/2/9
*/
