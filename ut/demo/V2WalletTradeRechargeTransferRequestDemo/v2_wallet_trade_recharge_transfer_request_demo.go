/**
 * 用户补贴 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2WalletTradeRechargeTransferRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2WalletTradeRechargeTransferRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2WalletTradeRechargeTransferRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 出款方商户号
        HuifuId:"6666000107309462",
        // 收款方用户号
        UserHuifuId:"6666000187364826",
        // 转账金额
        TransAmt:"0.01",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2WalletTradeRechargeTransferRequest(dgReq)
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
    // 出款方账户
    extendInfoMap["acct_id"] = "F00598600"
    // 转账描述
    extendInfoMap["description"] = "用户补贴"
    // 备注
    extendInfoMap["remark"] = "备注"
    return extendInfoMap
}

