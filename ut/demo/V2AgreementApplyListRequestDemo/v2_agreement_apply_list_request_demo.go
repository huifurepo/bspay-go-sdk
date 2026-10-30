/**
 * 协议申请列表查询 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2AgreementApplyListRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2AgreementApplyListRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2AgreementApplyListRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 汇付商户号
        HuifuId:"6666000103334211",
        // 协议类型
        AgreementType:"ELECT",
        // 当前页
        PageNum:"1",
        // 分页大小
        PageSize:"10",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2AgreementApplyListRequest(dgReq)
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
    // 协议编号
    extendInfoMap["agreement_id"] = "202505210201439711"
    // 场景类型
    extendInfoMap["scene_code"] = "MER_KYC"
    // 协议状态
    extendInfoMap["sign_status"] = "SIGN_SUCCESS"
    // 申请单号
    extendInfoMap["order_id"] = "2025052118403054"
    // 商户名称
    // extendInfoMap["mer_name"] = ""
    // 产品号
    // extendInfoMap["product_id"] = ""
    return extendInfoMap
}

