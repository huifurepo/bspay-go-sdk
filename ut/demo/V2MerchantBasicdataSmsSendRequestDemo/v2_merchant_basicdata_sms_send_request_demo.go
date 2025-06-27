/**
 * 商户短信发送 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantBasicdataSmsSendRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantBasicdataSmsSendRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantBasicdataSmsSendRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
        // 商户汇付Id
        HuifuId:"6666000105013599",
        // 手机号verify_type&#x3D;&#39;elecAcctSign&#39;时，手机号为空，系统自动取联系人手机号; &lt;font color&#x3D;&quot;green&quot;&gt;示例值：13911111111&lt;/font&gt;
        Phone:"13917111111",
        // 验证类型
        VerifyType:"settleBankChange",
        // 操作类型verify_type&#x3D;&#39;elecAcctSign&#39;时必填；枚举值：sendSmsCode-发送验证码；identitySmsCode-验证码核实；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：sendSmsCode&lt;/font&gt;
        // OperationType:"test",
        // 验证码verify_type&#x3D;&#39;elecAcctSign&#39;且operation_type&#x3D;&#39;identitySmsCode&#39;时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：123456&lt;/font&gt;
        // VerifyCode:"test",
        // 中信签约流水号verify_type&#x3D;&#39;elecAcctSign&#39;且operation_type&#x3D;&#39;identitySmsCode&#39;时必填；值为中信E管家签约发送短信时返回值；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：EMSSBPG2504284593690058431260676&lt;/font&gt;
        // ElecAcctSignSeqId:"test",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantBasicdataSmsSendRequest(dgReq)
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

