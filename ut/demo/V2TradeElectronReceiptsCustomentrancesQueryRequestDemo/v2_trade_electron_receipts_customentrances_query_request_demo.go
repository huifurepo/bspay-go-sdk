/**
 * 查询小票自定义入口 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeElectronReceiptsCustomentrancesQueryRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeElectronReceiptsCustomentrancesQueryRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeElectronReceiptsCustomentrancesQueryRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000103334211",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeElectronReceiptsCustomentrancesQueryRequest(dgReq)
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
    // 票据信息
    extendInfoMap["receipt_data"] = getReceiptDataRucan()
    return extendInfoMap
}

func getWxReceiptDataRucan() interface{} {
    dto := make(map[string]interface{})
    // 品牌ID
    dto["brand_id"] = "11"

    return dto;
}

func getReceiptDataRucan() string {
    dto := make(map[string]interface{})
    // 三方通道类型
    dto["third_channel_type"] = "T"
    // 微信票据信息
    dto["wx_receipt_data"] = getWxReceiptDataRucan()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

