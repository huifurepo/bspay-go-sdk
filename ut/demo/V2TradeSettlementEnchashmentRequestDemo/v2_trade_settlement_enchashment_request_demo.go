/**
 * 取现接口 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeSettlementEnchashmentRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeSettlementEnchashmentRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeSettlementEnchashmentRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 取现金额
        CashAmt:"0.01",
        // 取现方ID号
        HuifuId:"6666000021291985",
        // 到账日期类型
        IntoAcctDateType:"T0",
        // 取现卡序列号
        TokenNo:"10004053462",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeSettlementEnchashmentRequest(dgReq)
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
    extendInfoMap["notify_url"] = "http://www.gangcai.com"
    // 备注
    // extendInfoMap["remark"] = ""
    // 账户号
    // extendInfoMap["acct_id"] = ""
    return extendInfoMap
}

