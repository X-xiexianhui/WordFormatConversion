/*
 * @Description:初始化cmd
 * @version:
 * @Author: XieXianhui
 * @Date: 2021-12-29 22:23:21
 * @LastEditors: XieXianhui
 * @LastEditTime: 2021-12-30 17:45:56
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
 * @Description:命令行初始化
 * @name:init
 * @param {*}
 * @return {*}
 */
func init() {
	rootCmd.AddCommand(wordCmd)
}
