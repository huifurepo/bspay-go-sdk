/**
 * 更新花呗分期方案 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2PcreditSolutionModifyRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2PcreditSolutionModifyRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2PcreditSolutionModifyRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
        // 汇付客户Id
        HuifuId:"6666000110468104",
        // 创建成功后返回的贴息活动方案id
        SolutionId:"1515",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2PcreditSolutionModifyRequest(dgReq)
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
    // 开发者的应用ID
    extendInfoMap["app_id"] = "2019090666961966"
    // 花呗分期商家贴息活动名称
    // extendInfoMap["activity_name"] = ""
    // 活动开始时间
    extendInfoMap["start_time"] = "2019-07-08 00:00:00"
    // 活动结束时间
    extendInfoMap["end_time"] = "2029-07-10 00:00:00"
    return extendInfoMap
}

