/*
 * @Author: huaze.fan
 * @LastEditTime: 2022-11-14 16:26:03
 * @LastEditors: huaze.fan
 * @Description: Good Good Study！Day Day Up
 */
package BsPaySdk

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/fatih/structs"
)

/**
 * 去掉参数中value为 nil 的字段
 */
func DeleteEmptyValue(src map[string]interface{}) map[string]interface{} {
	resultMap := make(map[string]interface{})

	for key, value := range src {
		if key != "" && value != nil {
			resultMap[key] = value
		}
	}
	return resultMap
}

/**
 * 将 map 添加到另外 map 中
 */
func AddMapValue(src map[string]interface{}, original map[string]interface{}) map[string]interface{} {
	for key, value := range src {
		if key != "" && value != nil {
			original[key] = value
		}
	}
	return original
}

func ToString(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	default:
		r, _ := json.Marshal(v)
		return string(r)
	}
}

// map 深拷贝
func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}

		return newSlice
	}

	return value
}

// struct 转 map
func ToMap(s interface{}) map[string]interface{} {
	m := structs.Map(s)
	return m
}

/**
读取商户配置文件
*/
func ReadMerchConfig(configPath string) (*MerchSysConfig, error) {
	// 读文件
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		BspayPrintln("read merch config file failed ---> path: " + configPath + " ----> error: " + err.Error())
		return nil, errors.New("read merch config file failed")
	}

	// 映射 json 结构
	tmpMsc := MerchSysConfig{}
	jsonError := json.Unmarshal(data, &tmpMsc)
	if jsonError != nil {
		BspayPrintln("read merch config file failed ---> path: " + configPath + " ----> error: " + jsonError.Error())
		return nil, errors.New("read merch config file failed")
	}

	// 检查 4个 必要配置
	if tmpMsc.ProductId == "" || tmpMsc.SysId == "" ||
		tmpMsc.RsaHuifuPublicKey == "" || tmpMsc.RsaMerchPrivateKey == "" {

		BspayPrintln("read merch config file failed ---> path: " + configPath)
		return nil, errors.New("read merch config file failed")
	}

	return &tmpMsc, nil
}
