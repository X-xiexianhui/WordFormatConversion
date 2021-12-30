/*
 * @Description:对单词类型进行编码转换
 * @version:1.0
 * @Author: XieXianhui
 * @Date: 2021-12-29 22:58:24
 * @LastEditors: XieXianhui
 * @LastEditTime: 2021-12-30 17:37:27
 */
package word

import (
	"strings"
	"unicode"
)

/**
 * @Description:单词全部转大写，调用string.ToUpper方法
 * @name:ToUpper
 * @param {string} s
 * @return {string}
 */
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

/**
 * @Description:单词全部转小写，调用strings.ToLower方法
 * @name:ToLower
 * @param {string} s
 * @return {string}
 */
func ToLower(s string) string {
	return strings.ToLower(s)
}

/**
 * @Description:下划线单词转大写驼峰单词
 * @name:UnderscoreToUpperCamelCase
 * @param {string} s
 * @return {string}
 */
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

/**
 * @Description: 下划线单词转小写驼峰单词
 * @name:UnderscoreToLowerCamelCase
 * @param {string} s
 * @return {string}
 */
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

/**
 * @Description:驼峰单词转下划线单词
 * @name:CamelCaseToUnderscore
 * @param {string} s
 * @return {string}
 */
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
