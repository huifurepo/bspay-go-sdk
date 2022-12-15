/*
 * @Author: huaze.fan
 * @LastEditTime: 2022-11-14 19:18:19
 * @LastEditors: huaze.fan
 * @Description: Good Good Study！Day Day Up
 */
package tool

import (
	"fmt"
	"math/rand"
	"time"
)

// 获取当前时间
func GetCurrentDate() string {
	currentTime := time.Now()
	strTime := currentTime.Format("20060102")
	return strTime
}

// 获取当前时间
func GetCurrentTime() string {
	currentTime := time.Now()
	strTime := currentTime.Format("20060102150405")
	return strTime
}

// 生成随机数
func GetReqSeqId() string {
	strTime := GetCurrentTime()
	rand.Seed(time.Now().UnixNano())
	num := fmt.Sprintf("%d", rand.Intn(1000)+1000)
	return strTime + num
}
