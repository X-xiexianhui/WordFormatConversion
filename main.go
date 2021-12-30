/*
 * @Description:
 * @version:
 * @Author: XieXianhui
 * @Date: 2021-12-25 22:06:35
 * @LastEditors: XieXianhui
 * @LastEditTime: 2021-12-29 22:51:11
 */
package main

import (
	"log"
	"tour/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err:%v", err)
	}
}
