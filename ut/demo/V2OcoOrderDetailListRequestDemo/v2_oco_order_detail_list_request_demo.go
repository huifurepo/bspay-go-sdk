/**
 * 全渠道订单分账接收方查询 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2OcoOrderDetailListRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2OcoOrderDetailListRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2OcoOrderDetailListRequest{
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
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2OcoOrderDetailListRequest(dgReq)
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

