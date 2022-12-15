/**
 * 余额支付查询 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeAcctpaymentPayQueryRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeAcctpaymentPayQueryRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeAcctpaymentPayQueryRequest{
        // 商户号
        HuifuId:"6666000018344461",
        // 原交易请求日期
        OrgReqDate:"20211021",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeAcctpaymentPayQueryRequest(dgReq)
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
    // 原交易请求流水号
    extendInfoMap["org_req_seq_id"] = "202110210012100005"
    // 原交易全局流水号
    extendInfoMap["org_hf_seq_id"] = ""
    return extendInfoMap
}

