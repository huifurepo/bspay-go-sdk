/**
 * 账单退款接口 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V3BillpayOrderPaymentRefundRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V3BillpayOrderPaymentRefundRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V3BillpayOrderPaymentRefundRequest{
        // 请求流水号
        // ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        // ReqDate:tool.GetCurrentDate(),
        // 商户号
        // HuifuId:"test",
        // 账单编号
        // BillNo:"test",
        // 退款金额
        // RefAmt:"test",
        // 大额转账支付账户信息数据jsonObject格式；银行大额转账支付交易的退款申请,付款方账户类型为对公时必填
        // BankInfoData:get1d68b0833d9146359c7eAb5f4dbf001b(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V3BillpayOrderPaymentRefundRequest(dgReq)
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
    // 退款原因
    // extendInfoMap["reason"] = ""
    // 异步通知地址
    // extendInfoMap["notify_url"] = ""
    return extendInfoMap
}

func get1d68b0833d9146359c7eAb5f4dbf001b() string {
    dto := make(map[string]interface{})
    // 省份付款方为对公账户时必填,参见[代发省市地区码](https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/opps/api/prod/download_file/area/%E6%96%97%E6%8B%B1%E4%BB%A3%E5%8F%91%E7%9C%81%E4%BB%BD%E5%9C%B0%E5%8C%BA%E7%BC%96%E7%A0%81.xlsx#代发省市地区码);&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0013&lt;/font&gt;
    // dto["province"] = "test"
    // 地区付款方为对公账户时必填，,参见[代发省市地区码](https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/opps/api/prod/download_file/area/%E6%96%97%E6%8B%B1%E4%BB%A3%E5%8F%91%E7%9C%81%E4%BB%BD%E5%9C%B0%E5%8C%BA%E7%BC%96%E7%A0%81.xlsx#代发省市地区码)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1301&lt;/font&gt;
    // dto["area"] = "test"
    // 银行编号付款方为对公账户时必填,参见[银行编码](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_yhbm#银行编码);&lt;font color&#x3D;&quot;green&quot;&gt;示例值：01040000&lt;/font&gt;
    // dto["bank_code"] = "test"
    // 联行号付款方为对公账户时必填,参见[银行支行编码](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_yhzhbm#银行支行编码);&lt;font color&#x3D;&quot;green&quot;&gt;示例值：102290026507&lt;/font&gt;
    // dto["correspondent_code"] = "test"
    // 付款方账户类型
    // dto["card_acct_type"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

