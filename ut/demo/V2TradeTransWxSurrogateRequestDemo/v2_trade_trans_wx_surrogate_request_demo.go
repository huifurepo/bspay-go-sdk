/**
 * 微信代发 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeTransWxSurrogateRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeTransWxSurrogateRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeTransWxSurrogateRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 出账商户号
        OutHuifuId:"6666000000041651",
        // 代发金额
        TransAmt:"9.00",
        // 收款用户openid
        OpenId:"o-MYE42l80oelYMDE34nYD456Xoy",
        // 微信收款用户姓名
        UserName:"王大锤",
        // 代发备注
        Remark:"测试用",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeTransWxSurrogateRequest(dgReq)
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
    // 账户类型
    extendInfoMap["out_acct_type"] = "05"
    // 异步通知地址
    extendInfoMap["notify_url"] = "http://www.gangcai.com"
    // 子商户号
    // extendInfoMap["sub_mch_id"] = ""
    // 子商户应用Id
    // extendInfoMap["sub_app_id"] = ""
    return extendInfoMap
}

