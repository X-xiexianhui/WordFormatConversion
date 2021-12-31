/*
 * @Description:时间工具
 * @version: 1.0
 * @Author: XieXianhui
 * @Date: 2021-12-31 20:23:50
 * @LastEditors: XieXianhui
 * @LastEditTime: 2021-12-31 21:27:13
 */
package timer

import (
	"time"
)

/**
 * @Description:获取当前本地时间
 * @param {*}
 * @return {time.Time}
 */
func GetNowTime() time.Time {
	return time.Now()
}

/**
 * @Description: 推算时间
 * @param {time.Time} currencyTimer
 * @param {string} d
 * @return {time.Time,error}
 */
func GetCalculateTime(currencyTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currencyTimer.Add(duration), nil
}
