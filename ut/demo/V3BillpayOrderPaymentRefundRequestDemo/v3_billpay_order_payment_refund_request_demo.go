/**
 * 账单退款接口 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V3BillpayOrderPaymentRefundRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V3BillpayOrderPaymentRefundRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V3BillpayOrderPaymentRefundRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000108432796",
        // 账单编号
        BillNo:"BN2026052236198530",
        // 退款金额
        RefAmt:"100.00",
        // 大额转账支付账户信息数据jsonObject格式；银行大额转账支付交易的退款申请,付款方账户类型为对公时必填
        BankInfoData:get39a3f6ccD87e496fB6ac4a9211667d6f(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V3BillpayOrderPaymentRefundRequest(dgReq)
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
    // 退款原因
    extendInfoMap["reason"] = "退货"
    // 异步通知地址
    // extendInfoMap["notify_url"] = ""
    return extendInfoMap
}

func get39a3f6ccD87e496fB6ac4a9211667d6f() string {
    dto := make(map[string]interface{})
    // 省份
    dto["province"] = "0013"
    // 地区
    dto["area"] = "1301"
    // 银行编号
    dto["bank_code"] = "01040000"
    // 联行号
    dto["correspondent_code"] = "102290026507"
    // 付款方账户类型
    dto["card_acct_type"] = "E"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

