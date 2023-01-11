/**
 * 查询投诉单列表及详情 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantComplaintListInfoQueryRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantComplaintListInfoQueryRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantComplaintListInfoQueryRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
        // 开始日期
        BeginDate:"2022-10-20",
        // 结束日期
        EndDate:"2022-10-20",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantComplaintListInfoQueryRequest(dgReq)
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
    extendInfoMap["offset"] = ""
    // 分页大小
    extendInfoMap["limit"] = ""
    // 被诉的汇付商户ID
    extendInfoMap["huifu_id"] = ""
    // 被诉的商户名称
    extendInfoMap["reg_name"] = ""
    // 微信订单号
    extendInfoMap["transaction_id"] = ""
    // 微信投诉单号
    extendInfoMap["complaint_id"] = ""
    // 投诉状态
    extendInfoMap["complaint_state"] = ""
    // 用户投诉次数
    extendInfoMap["user_complaint_times"] = ""
    // 是否有待回复的用户留言
    extendInfoMap["incoming_user_response"] = "0"
    return extendInfoMap
}

