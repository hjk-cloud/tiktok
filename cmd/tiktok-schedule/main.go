package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/hjk-cloud/tiktok/config"
)

func main() {
	// 删除前hours天的视频和封面所在的目录
	hours := time.Duration(config.Config.TmpFileExpireHours)
	earliest := time.Now().Add(-hours * time.Hour)
	year, month, day := earliest.Year(), int(earliest.Month()), earliest.Day()
	log.Println(config.Config.StaticDir, "###删到", year, month, day)

	yfiles, err := ioutil.ReadDir(config.Config.StaticDir)
	if err != nil {
		log.Fatal(err)
	}
	// 所有年
	for _, yf := range yfiles {
		y, err := strconv.Atoi(yf.Name())
		if err != nil {
			log.Println(err)
			continue
		}
		ypath := path.Join(config.Config.StaticDir, yf.Name())
		if y < year {
			log.Println("删除", ypath)
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
			m, err := strconv.Atoi(mf.Name())
			if err != nil {
				log.Println(err)
				continue
			}
			mpath := path.Join(ypath, mf.Name())
			if m < month {
				log.Println("删除", mpath)
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
				d, err := strconv.Atoi(df.Name())
				if err != nil {
					log.Println(err)
					continue
				}
				dpath := path.Join(mpath, df.Name())
				if d <= day {
					log.Println("删除", dpath)
					os.RemoveAll(dpath)
					continue
				}
			}
		}
	}
}
