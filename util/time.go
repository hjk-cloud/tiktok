package util

import "time"

// 秒级时间戳转time
func UnixSecondToTime(second int64) time.Time {
	return time.Unix(second, 0)
}

// 毫秒级时间戳转time
func UnixMilliToTime(milli int64) time.Time {
	return time.Unix(milli/1000, (milli%1000)*(1000*1000))
}

// 纳秒级时间戳转time
func UnixNanoToTime(nano int64) time.Time {
	return time.Unix(nano/(1000*1000*1000), nano%(1000*1000*1000))
}
