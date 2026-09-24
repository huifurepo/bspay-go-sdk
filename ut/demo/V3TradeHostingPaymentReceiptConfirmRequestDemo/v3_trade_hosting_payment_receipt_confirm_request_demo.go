/**
 * 确认收货提醒 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V3TradeHostingPaymentReceiptConfirmRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V3TradeHostingPaymentReceiptConfirmRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V3TradeHostingPaymentReceiptConfirmRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000109133323",
        // 快递签收时间
        ReceivedTime:"1670829139",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V3TradeHostingPaymentReceiptConfirmRequest(dgReq)
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
    // 原支付交易对应的微信订单号
    extendInfoMap["transaction_id"] = "fake-transid-20221209132531-44"
    // 支付下单商户的商户号
    extendInfoMap["merchant_id"] = "fake-mchid-123"
    // 二级商户号
    // extendInfoMap["sub_merchant_id"] = ""
    // 商户系统内部订单号
    extendInfoMap["merchant_trade_no"] = "fake-tradeno-20221209132531-44"
    return extendInfoMap
}

