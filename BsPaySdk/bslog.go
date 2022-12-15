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

func BspayPrintln(v ...interface{}) {
	log.Println("<BspayLog> " + fmt.Sprint(v...))
}
