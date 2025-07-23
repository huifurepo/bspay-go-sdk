/**
 * 终端云MIS交易 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeCloudmisDeviceInformationMisRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeCloudmisDeviceInformationMisRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeCloudmisDeviceInformationMisRequest{
        // 请求流水号
        ReqId:"20240313115926539uf7cqcmwxl30",
        // 终端设备号
        DeviceId:"310000011015000063677",
        // 商户号
        HuifuId:"6666000141203565",
        // 交易信息
        JsonData:"{\"transAmount\":\"11\",\"interfaceType\":\"SALE\",\"thirdOrderId\":\"20240313115926539uf7cqcmwxl30\" }",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeCloudmisDeviceInformationMisRequest(dgReq)
  	if err != nil {
		fmt.Println("请求异常:", err)
	}

	fmt.Println("返回数据:", resp)
}

/**
 * 非必填字段
 */
func getExtendInfos() map[string]interface{} {
    // 设置非必填字段
    extendInfoMap := make(map[string]interface{})
    return extendInfoMap
}

