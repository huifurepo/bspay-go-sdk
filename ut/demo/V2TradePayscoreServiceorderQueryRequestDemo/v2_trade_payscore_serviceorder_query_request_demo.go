/**
 * 查询支付分订单 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradePayscoreServiceorderQueryRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradePayscoreServiceorderQueryRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradePayscoreServiceorderQueryRequest{
        // 汇付商户号
        HuifuId:"6666000108854952",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradePayscoreServiceorderQueryRequest(dgReq)
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
    // 汇付服务订单号
    // extendInfoMap["out_order_no"] = ""
    // 创建服务订单返回的汇付全局流水号
    // extendInfoMap["org_hf_seq_id"] = ""
    // 服务订单创建请求流水号
    // extendInfoMap["org_req_seq_id"] = ""
    return extendInfoMap
}

