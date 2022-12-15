/**
 * 微信用户标识查询接口 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeWxusermarkQueryRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeWxusermarkQueryRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeWxusermarkQueryRequest{
        // 商户号
        HuifuId:"6666000003100616",
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 支付授权码
        AuthCode:"130636925881320560",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeWxusermarkQueryRequest(dgReq)
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
    // 子商户公众账号ID
    extendInfoMap["sub_appid"] = "oQOa46X2FxRqEy6F4YmwIRCrA7Mk"
    // 渠道号
    extendInfoMap["channel_no"] = ""
    // 场景类型
    extendInfoMap["pay_scene"] = ""
    return extendInfoMap
}

