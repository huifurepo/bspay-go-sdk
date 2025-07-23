/**
 * 云MIS订单详情查询接口 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeCloudmisOrderDetailRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeCloudmisOrderDetailRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeCloudmisOrderDetailRequest{
        // 请求流水号
        // ReqId:"test",
        // 原MIS请求商户号
        OrgHuifuId:"6666000141203565",
        // 原MIS请求终端号
        OrgDeviceId:"310000011015000063677",
        // 原MIS请求日期
        OrgReqDate:"20240715",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeCloudmisOrderDetailRequest(dgReq)
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
    // 原MIS请求流水号
    extendInfoMap["org_req_id"] = "reqId20240624091729005"
    // 原MIS请求jsonData中的三方单号
    extendInfoMap["org_third_order_id"] = "20240313115926539uf7cqcmwxl30"
    return extendInfoMap
}

