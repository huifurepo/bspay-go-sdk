package BsPaySdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

/**
* post 请求
 */
func DoPostReq(url string, params map[string]interface{}, msc *MerchSysConfig) (*http.Response, error) {
	// 构建参数字符串
	signMap := DeepCopy(DeleteEmptyValue(params)).(map[string]interface{})
	delete(signMap, "extend_infos")
	if params["extend_infos"] != nil {
		signMap = AddMapValue(params["extend_infos"].(map[string]interface{}), signMap)
	}
	paramText, _ := FormatSignSrcText(signMap)

	sign, _ := RsaSign(paramText, msc)

	respData := make(map[string]interface{})
	respData["sign"] = sign
	respData["sys_id"] = msc.SysId
	respData["product_id"] = msc.ProductId
	respData["data"] = signMap

	respDataText, _ := FormatSignSrcText(respData)
	BspayPrintln("request url: " + url)
	BspayPrintln("request data: " + respDataText)
	request, _ := http.NewRequest("POST", url, strings.NewReader(respDataText))
	request.Header.Add("Content-Type", "application/json")

	return http.DefaultClient.Do(request)
}

func DoUploadFile(url string, params map[string]interface{}, filedirname string, msc *MerchSysConfig) (*http.Response, error) {
	// 构建参数字符串

	// 打开文件句柄
	file, openErr := os.Open(filedirname)
	if openErr != nil {
		BspayPrintln("open file error: " + openErr.Error())
		return nil, openErr
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	var fileName string
	if str, ok := params["picture"].(string); ok {
		fileName = str
	} else {
		fileName = filepath.Base(filedirname)
	}
	// 参数写入文件
	part, writerErr := writer.CreateFormFile("file", fileName)
	if writerErr != nil {
		BspayPrintln("writer file error: " + writerErr.Error())
		return nil, writerErr
	}
	_, _ = io.Copy(part, file)
	// 写入其余参数
	dataStr, _ := FormatSignSrcText(params)
	_ = writer.WriteField("data", dataStr)
	BspayPrintln("request data: " + dataStr)

	respData := make(map[string]interface{})
	respData["sys_id"] = msc.SysId
	respData["product_id"] = msc.ProductId

	for key, val := range respData {
		if str, ok := val.(string); ok {
			_ = writer.WriteField(key, str)
		}
	}
	writerErr = writer.Close()
	if writerErr != nil {
		BspayPrintln("writer file error: " + writerErr.Error())
		return nil, writerErr
	}
	request, _ := http.NewRequest("POST", url, body)

	request.Header.Set("Content-Type", writer.FormDataContentType())

	return http.DefaultClient.Do(request)
}

/*
将参数格式话成验签字符串。
@param: 	paramMap - 数据集。
@return:	满足服务端请求格式的加签字符串，无法进行处理时返回 error。
*/
func FormatSignSrcText(paramMap map[string]interface{}) (string, error) {
	postResult, postErr := json.Marshal(paramMap)
	if postErr != nil {
		return "", postErr
	}
	content := string(postResult)
	// 防止 < > & 序列化被转 unecode
	content = strings.ReplaceAll(content, "\\u003c", "<")
	content = strings.ReplaceAll(content, "\\u003e", ">")
	content = strings.ReplaceAll(content, "\\u0026", "&")
	return content, nil
}

/**
* api 应答数据统一处理
 */
func HandleResponse(resp *http.Response, msc *MerchSysConfig, needVerfySign bool) (map[string]interface{}, error) {
	if resp == nil {
		BspayPrintln("network error")
		return nil, errors.New("network error")
	}

	// 数据读取完成后自动关闭链接。----- copy 网络请求代码时千万别漏了 by Charles.Huang
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil && (body == nil || string(body) == "") { // 读取数据失败返回
		return nil, errors.New(resp.Status)
	}

	BspayPrintln("response data: " + string(body))

	// 返回map
	respDataMap := make(map[string]interface{})
	if !needVerfySign {
		respDataMap["data"] = string(body)
		return respDataMap, nil
	}

	// json 序列化响应数据
	tmpJsonErr1 := json.Unmarshal(body, &respDataMap)
	if tmpJsonErr1 != nil {
		BspayPrintln("response json unmarshal error: " + tmpJsonErr1.Error())
		return nil, tmpJsonErr1
	}

	// 验签
	contentString, _ := FormatSignSrcText(respDataMap["data"].(map[string]interface{}))
	sign, _ := RsaSignVerify(respDataMap["sign"].(string), contentString, msc)
	if !sign {
		BspayPrintln("check signature error")
		return nil, errors.New("check signature error")
	}

	return respDataMap, nil
}

func PostRequest(reqUrl string, reqParam map[string]interface{}, msc *MerchSysConfig, needVerfySign ...bool) (map[string]interface{}, error) {
	var resp *http.Response
	var err error

	resp, err = DoPostReq(reqUrl, reqParam, msc)

	if err != nil { // 网络问题返回
		BspayPrintln("post request error: " + err.Error())
		return nil, err
	}
	if resp.StatusCode == 200 {
		// 处理服务端返回数据
		var needTag = true
		if needVerfySign != nil {
			needTag = needVerfySign[0]
		}
		data, errRes := HandleResponse(resp, msc, needTag)
		if errRes != nil { // 非正常情况
			BspayPrintln("post request error: " + errRes.Error())
			return nil, errRes
		}
		return data, nil
	}

	BspayPrintln("post request error: " + resp.Status)
	return nil, errors.New(resp.Status)
}

/**
* 上传文件统一方法
 */
func UploadBsPay(reqUrl string, reqParam map[string]interface{}, filepath string, msc *MerchSysConfig) (map[string]interface{}, error) {
	var resp *http.Response
	var err error

	resp, err = DoUploadFile(reqUrl, reqParam, filepath, msc)

	if err != nil { // 网络问题返回
		BspayPrintln("upload request error: " + err.Error())
		return nil, err
	}

	if resp.StatusCode == 200 {
		// 处理服务端返回数据
		data, errRes := HandleResponse(resp, msc, false)

		if errRes != nil { // 非正常情况
			BspayPrintln("upload request error: " + errRes.Error())
			return nil, errRes
		}
		return data, nil
	}

	BspayPrintln("upload request error: " + resp.Status)
	return nil, errors.New(resp.Status)
}
