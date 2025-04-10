/**
 * 钱包转账下单 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2WalletTradeTransferRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2WalletTradeTransferRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2WalletTradeTransferRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000135653240",
        // 出款方钱包用户ID
        UserHuifuId:"6666000136655020",
        // 收款方钱包用户ID
        InUserHuifuId:"6666000136655254",
        // 订单金额
        TransAmt:"0.03",
        // 跳转地址
        FrontUrl:"http://www.huifu.com/products-services/",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2WalletTradeTransferRequest(dgReq)
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
    // 密码页面类型
    extendInfoMap["request_type"] = "M"
    // 备注
    extendInfoMap["remark"] = "remark11"
    // 商户扩展信息
    extendInfoMap["mer_priv"] = "mer_priv1"
    // 订单失效时间
    // extendInfoMap["time_expire"] = ""
    // 异步通知地址
    // extendInfoMap["notify_url"] = ""
    return extendInfoMap
}

