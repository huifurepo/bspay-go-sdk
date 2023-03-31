/**
 * 汇付入账查询 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeOnlinepaymentTransferRemittanceorderRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeOnlinepaymentTransferRemittanceorderRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeOnlinepaymentTransferRemittanceorderRequest{
        // 商户号
        HuifuId:"6666000003100615",
        // 原请求开始日期
        OrgReqStartDate:"20230110",
        // 原请求结束日期
        OrgReqEndDate:"20230110",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeOnlinepaymentTransferRemittanceorderRequest(dgReq)
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
    // 原请求流水号
    extendInfoMap["org_req_seq_id"] = "20230110155433defaultit655128593"
    // 原请求日期
    extendInfoMap["org_req_date"] = "20230110"
    // 原汇款订单号
    extendInfoMap["org_remittance_order_id"] = "20230110155433defaultit655128591"
    // 每页条数
    extendInfoMap["page_size"] = "1"
    // 分页页码
    extendInfoMap["page_no"] = "1"
    return extendInfoMap
}

