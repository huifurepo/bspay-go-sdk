/**
 * 申诉单列表查询 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantAppealListQueryRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantAppealListQueryRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantAppealListQueryRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 分页条数
        PageSize:"20",
        // 申诉创建开始日期
        BeginDate:"20240701",
        // 申诉创建结束日期
        EndDate:"20240731",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantAppealListQueryRequest(dgReq)
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
    // 分页页码
    extendInfoMap["page_num"] = "1"
    // 协查单号
    extendInfoMap["assist_id"] = ""
    // 渠道/代理/商户/用户编号
    extendInfoMap["huifu_id"] = "6666000108285670"
    // 商户名称
    extendInfoMap["mer_name"] = ""
    // 申诉状态
    extendInfoMap["appeal_node"] = ""
    // 审核结论
    extendInfoMap["audit_result"] = ""
    // 运营处理状态
    extendInfoMap["operation_status"] = ""
    // 汇付处置等级
    extendInfoMap["handle_degree"] = ""
    return extendInfoMap
}

