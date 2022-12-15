/**
 * 支付宝申诉请求凭证 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantComplaintRequestCertificatesRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantComplaintRequestCertificatesRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantComplaintRequestCertificatesRequest{
        // 请求汇付流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求汇付时间
        ReqDate:tool.GetCurrentDate(),
        // 支付宝推送流水号
        RiskBizId:"04a3d978689542ed6ba0295fbc2dc137-BANK",
        // 商户类型
        MerchantType:"个体工商户",
        // 商户经营模式
        OperationType:"线上",
        // 收款应用场景
        PaymentScene:"通过小程序收款",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantComplaintRequestCertificatesRequest(dgReq)
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

