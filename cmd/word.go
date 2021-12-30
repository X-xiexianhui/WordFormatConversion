/*
 * @Description:初始化Word子命令
 * @version:
 * @Author: XieXianhui
 * @Date: 2021-12-29 22:18:00
 * @LastEditors: XieXianhui
 * @LastEditTime: 2021-12-30 17:45:43
 */

package cmd

import (
	"log"
	"strings"
	"tour/internal/word"

	"github.com/spf13/cobra"
)

// 定义word子命令
const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeCamelCaseToUnderscore
)

var str string
var mode int8

// 对单词子命令进行设置和集成
var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部单词转大写",
	"2：全部单词转小写",
	"3：下划线单词转大写驼峰单词",
	"4：下划线单词转小写驼峰单词",
	"5；驼峰单词转下划线单词",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var context string
		switch mode {
		case ModeUpper:
			context = word.ToUpper(str)
		case ModeLower:
			context = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			context = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			context = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			context = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行help查看帮助文档")
		}
		log.Printf("输出结果：%s", context)
	},
}

/**
 * @Description:命令参数设置和初始化
 * @name:init
 * @param {*}
 * @return {*}
 */
func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换模式")
}
