<!--
 * @Description: 
 * @version: 1.0
 * @Author: XieXianhui
 * @Date: 2021-12-31 21:28:43
 * @LastEditors: XieXianhui
 * @LastEditTime: 2021-12-31 22:12:40
-->
# 格式转换命令行工具
## 命令格式
###
#### 1.help：
    go run main.go help word
#### 2.单词转换：
    go run main.go word -s=abcd -m=1
    输出结果：ABCD
#### 3.获取本地时间:
    go run main.go time now
#### 4.时间计算
    go run main.go time calc -c="2021-12-31 21:30:00" -d=5m
    输出结果：2021-12-31 21:35:00,1640986500
