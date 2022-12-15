/**
 * 手续费试算 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeFeecalcRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeFeecalcRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeFeecalcRequest{
        // 商户号
        HuifuId:"6666000116584429",
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 交易类型
        TradeType:"ENCASHMENT",
        // 交易金额
        TransAmt:"1000.00",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeFeecalcRequest(dgReq)
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
    // 网银交易类型
    extendInfoMap["online_trans_type"] = ""
    // 付款方银行编号
    extendInfoMap["bank_id"] = "01020000"
    // 卡类型
    extendInfoMap["card_type"] = "D"
    // 渠道号
    extendInfoMap["channel_no"] = "10000001"
    // 数字货币银行编号
    extendInfoMap["digital_bank_no"] = "01002"
    // 取现到账类型
    extendInfoMap["encash_type"] = "T0"
    // 场景类型
    extendInfoMap["pay_scene"] = "01"
    return extendInfoMap
}

