/**
 * 不明来账处理结果查询 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradePaymentZxeUnknownincomeDisposequeryRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradePaymentZxeUnknownincomeDisposequeryRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradePaymentZxeUnknownincomeDisposequeryRequest{
        // 商户号
        HuifuId:"6666000109133323",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradePaymentZxeUnknownincomeDisposequeryRequest(dgReq)
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
    // 原请求流水号
    extendInfoMap["org_req_seq_id"] = "20240925test100001"
    // 原请求日期
    extendInfoMap["org_req_date"] = "20240925"
    // 原全局流水号
    extendInfoMap["org_hf_seq_id"] = ""
    return extendInfoMap
}

