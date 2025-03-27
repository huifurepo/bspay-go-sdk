/**
 * 钱包支付退款查询 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2WalletTradePayRefundQueryRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2WalletTradePayRefundQueryRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2WalletTradePayRefundQueryRequest{
        // 原退款交易请求日期
        OrgReqDate:"20230816",
        // 原退款交易请求流水号
        // OrgReqSeqId:"test",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2WalletTradePayRefundQueryRequest(dgReq)
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
    // 钱包用户id
    extendInfoMap["user_huifu_id"] = "6666000136655020"
    // 原退款交易全局流水号
    extendInfoMap["org_hf_seq_id"] = "003100TOP1A230816150903P990ac139c0600000"
    return extendInfoMap
}

