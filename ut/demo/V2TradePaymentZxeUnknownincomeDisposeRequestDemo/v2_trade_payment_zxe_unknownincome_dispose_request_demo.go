/**
 * 不明来账处理 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradePaymentZxeUnknownincomeDisposeRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradePaymentZxeUnknownincomeDisposeRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradePaymentZxeUnknownincomeDisposeRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000109133323",
        // 银行侧交易流水号
        BankSerialNo:"FRSC202409252NEA000121452600000",
        // 操作类型
        OperateType:"0",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradePaymentZxeUnknownincomeDisposeRequest(dgReq)
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
    // 异步通知地址
    extendInfoMap["notify_url"] = "https://mock.uutool.cn/fat69kri74k"
    return extendInfoMap
}

