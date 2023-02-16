package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"

	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
)

type FuncCollection map[string]reflect.Value

// http://127.0.0.1:8088/gosdk?method=V2MerchantBasicdataQueryRequest&env=prod&sysId=&productId=&rsaPrivateKey=&rsaPublicKey=
// post body jsonStr

func GoSDKHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// 调用SDK方法
	// 1. 数据初始化
	queryVal := r.URL.Query()
	isProd := false

	if queryVal.Get("env") == "prod" {
		isProd = true
	}
	sdk1, _ := BsPaySdk.NewBsPay(isProd, "../config/config.json")

	if queryVal.Get("sysId") != "" {
		sdk1.Msc.SysId = queryVal.Get("sysId")
	}
	if queryVal.Get("productId") != "" {
		sdk1.Msc.ProductId = queryVal.Get("productId")
	}
	if queryVal.Get("rsaPrivateKey") != "" {
		sdk1.Msc.RsaMerchPrivateKey = queryVal.Get("rsaPrivateKey")
	}
	if queryVal.Get("rsaPublicKey") != "" {
		sdk1.Msc.RsaHuifuPublicKey = queryVal.Get("rsaPublicKey")
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	bodyData := buf.String()
	fmt.Println("bodyData:", bodyData)
	// 取出extendInfos对象,放入bodyData
	bodyMap := make(map[string]interface{})
	json.Unmarshal([]byte(bodyData), &bodyMap)
	extendInfos := bodyMap["extendInfos"]
	delete(bodyMap, "extendInfos")

	if extendInfos != nil {
		extendMap := extendInfos.(map[string]interface{})
		for k, v := range extendMap {
			bodyMap[k] = v
		}
	}

	newBodyByte, _ := json.Marshal(bodyMap)
	newBodyData := string(newBodyByte)
	fmt.Println("newBodyData:", newBodyData)

	var rfv []reflect.Value
	var err error
	method := queryVal.Get("method")
	if method == "V2SupplementaryPictureRequest" {
		filePath, _ := filepath.Abs("./IMG_4955.JPG")
		rfv, err = CallFunc(sdk1, "Str"+method, newBodyData, filePath)
	} else {
		rfv, err = CallFunc(sdk1, "Str"+method, newBodyData)
	}
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
	fmt.Println("callReflect resp:", rfv[0], ", err:", rfv[1])

	resp, _ := json.Marshal(rfv[0].Interface())
	fmt.Println("resp value:", string(resp))
	if rfv[1].Interface() != nil {
		fmt.Fprintf(w, "%s", rfv[1].Interface())
	} else {
		re, _ := json.Marshal(string(resp))
		fmt.Fprintf(w, "%s", re)
	}
}

func main() {
	http.HandleFunc("/gosdk", GoSDKHandler)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Println("http listen error:", err)
	}
}

func CallFunc(router *BsPaySdk.BsPay, tableName string, args ...interface{}) (result []reflect.Value, err error) {
	FuncMap := make(FuncCollection, 0)
	rf := reflect.ValueOf(router)
	rft := rf.Type()
	funcNum := rf.NumMethod()
	for i := 0; i < funcNum; i++ {
		mName := rft.Method(i).Name
		FuncMap[mName] = rf.Method(i)
	}
	parameter := make([]reflect.Value, len(args))
	for k, arg := range args {
		parameter[k] = reflect.ValueOf(arg)
	}
	result = FuncMap[tableName].Call(parameter)
	return
}
