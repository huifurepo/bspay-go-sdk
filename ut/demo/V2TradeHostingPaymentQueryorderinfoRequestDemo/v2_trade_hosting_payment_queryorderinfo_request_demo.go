/**
 * 托管交易查询 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeHostingPaymentQueryorderinfoRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeHostingPaymentQueryorderinfoRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeHostingPaymentQueryorderinfoRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号开户自动生成；商户号不填时必填party_order_id；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000123123123&lt;/font&gt;
        HuifuId:"6666000109133323",
        // 原交易请求日期请求格式：yyyyMMdd；该字段不填时必填party_order_id；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20221023&lt;/font&gt;
        OrgReqDate:"20231020",
        // 原交易请求流水号该字段不填时必填party_order_id；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：rQ2021121311173944&lt;/font&gt;
        OrgReqSeqId:"202310201652361987182512",
        // 用户账单上的商户订单号该字段不填时，商户号、原交易请求日期、原交易请求流水号必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：03232109190255105603561&lt;/font&gt;
        // PartyOrderId:"test",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeHostingPaymentQueryorderinfoRequest(dgReq)
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
    return extendInfoMap
}

