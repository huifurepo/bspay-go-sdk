/*
 * @Author: huaze.fan
 * @LastEditTime: 2022-11-14 16:11:29
 * @LastEditors: huaze.fan
 * @Description: Good Good Study！Day Day Up
 */
package BsPaySdk

import (
	"fmt"
	"log"
)

type Logger interface {
	Println(v ...any)
}

type defaultLogger struct{}

func (d *defaultLogger) Println(v ...any) {
	log.Println("<BspayLog> " + fmt.Sprint(v...))
}

var logger Logger = &defaultLogger{}

// SetLogger 设置日志记录器
func SetLogger(l Logger) {
	logger = l
}

func BspayPrintln(v ...interface{}) {
	logger.Println(v...)
}
