/*
 * @Description:时间子命令
 * @version: 1.0
 * @Author: XieXianhui
 * @Date: 2021-12-31 20:31:36
 * @LastEditors: XieXianhui
 * @LastEditTime: 2021-12-31 21:57:04
 */
package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"
	"tour/internal/timer"

	"github.com/spf13/cobra"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

// 获取当前时间子命令
var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果：%s,%d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

//获取计算时间子命令
var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04:05"
			}
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err：%v", err)
		}
		log.Printf("输出结果：%s,%d", t.Format(layout), t.Unix())
	},
}

/**
 * @Description: 对time命令的now、calc子命令和其相关联的命令行参数进行注册
 * @param {*}
 * @return {*}
 */
func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算的时间，有效单位为时间戳格式或者已格式化后的时间")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "持续时间，有效时间单位为'ns','us','ms','s','m','h'")
}
