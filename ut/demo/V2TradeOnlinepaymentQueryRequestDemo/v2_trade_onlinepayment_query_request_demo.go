/**
 * 线上交易查询 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeOnlinepaymentQueryRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeOnlinepaymentQueryRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeOnlinepaymentQueryRequest{
        // 商户号
        HuifuId:"6666000109133323",
        // 原交易请求日期
        OrgReqDate:"20240401",
        // 原交易返回的全局流水号原交易请求流水号、原交易返回的全局流水号至少要送其中一项；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：00290TOP1GR210919004230P853ac13262200000&lt;/font&gt;
        OrgHfSeqId:"00600000000240401100725P864ac13645d00000",
        // 原交易请求流水号原交易请求流水号、原交易返回的全局流水号至少要送其中一项；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：2021091708126665001&lt;/font&gt;
        OrgReqSeqId:"295700155481522176",
        // 原交易支付类型QUICK_PAY：快捷支付、快捷充值(查询快捷交易必填)&lt;br/&gt;ONLINE_PAY：网银支付、网银充值&lt;br/&gt;WAP_PAY：手机WAP支付&lt;br/&gt;UNION_PAY：银联APP统一支付&lt;br/&gt;QUICK_PAY_APPLY：银行卡分期申请&lt;br/&gt;QUICK_PAY_CONFIRM：银行卡分期确认&lt;br/&gt;TRANSFER_ACCT：网银转账&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：TRANSFER_ACCT&lt;/font&gt;&lt;br/&gt;注意：**不支持聚合扫码接口生成的微信、支付宝、银联二维码等交易的查询。**
        PayType:"QUICK_PAY",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeOnlinepaymentQueryRequest(dgReq)
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

