/*
 * @Description:初始化cmd,用于注册子命令
 * @version:
 * @Author: XieXianhui
 * @Date: 2021-12-29 22:23:21
 * @LastEditors: XieXianhui
 * @LastEditTime: 2021-12-31 21:01:50
 */
package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{}

/**
 * @Description:启动cmd
 * @name:Execute
 * @param {*}
 * @return {error}
 */
func Execute() error {
	return rootCmd.Execute()
}

/**
 * @Description:子命令注册
 * @name:init
 * @param {*}
 * @return {*}
 */
func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
}
