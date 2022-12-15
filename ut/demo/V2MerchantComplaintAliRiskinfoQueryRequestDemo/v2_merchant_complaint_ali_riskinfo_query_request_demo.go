/**
 * 支付宝投诉查询 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantComplaintAliRiskinfoQueryRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantComplaintAliRiskinfoQueryRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantComplaintAliRiskinfoQueryRequest{
        // 请求汇付流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求汇付时间
        ReqDate:tool.GetCurrentDate(),
        // 开始日期
        BeginDate:"20221115",
        // 结束日期
        EndDate:"20221115",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantComplaintAliRiskinfoQueryRequest(dgReq)
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
    // 分页开始位置
    extendInfoMap["offset"] = "1"
    // 分页大小
    extendInfoMap["limit"] = "5"
    // 通道风险类型
    extendInfoMap["risk_type"] = ""
    // 汇付商户号
    extendInfoMap["huifu_id"] = ""
    // 支付宝推送流水号
    extendInfoMap["risk_biz_id"] = ""
    // 是否可申诉
    extendInfoMap["support_appeal"] = ""
    return extendInfoMap
}

