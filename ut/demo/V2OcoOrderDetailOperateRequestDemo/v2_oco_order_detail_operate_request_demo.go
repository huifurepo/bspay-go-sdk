/**
 * 全渠道订单分账明细操作 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2OcoOrderDetailOperateRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2OcoOrderDetailOperateRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2OcoOrderDetailOperateRequest{
        // 请求流水号
        // ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        // ReqDate:tool.GetCurrentDate(),
        // 商户号
        // HuifuId:"test",
        // 分账数据源
        // BusiSource:"test",
        // 业务订单号
        // OcoOrderId:"test",
        // 操作类型
        // OperateType:"test",
        // 支付方式枚举：BALANCE-余额支付 EFP-全域资金付款；备注：当operate_type&#x3D;SPLIT 立即分账时，pay_type必填，且若为退款，按原交易类型原路返回；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：BALANCE&lt;/font&gt;
        // PayType:"test",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2OcoOrderDetailOperateRequest(dgReq)
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
    // 分账接收方编号
    // extendInfoMap["in_huifu_id"] = ""
    return extendInfoMap
}

