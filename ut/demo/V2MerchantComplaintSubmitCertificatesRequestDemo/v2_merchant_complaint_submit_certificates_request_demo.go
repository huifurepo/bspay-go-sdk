/**
 * 支付宝申诉提交凭证 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantComplaintSubmitCertificatesRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantComplaintSubmitCertificatesRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantComplaintSubmitCertificatesRequest{
        // 请求汇付流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求汇付时间
        ReqDate:tool.GetCurrentDate(),
        // 支付宝推送流水号
        RiskBizId:"b1e11c97badf1ba025399ee0332d8fb1-ISV",
        // 申诉解限的唯一ID
        RelievingId:"653739ab36362810b7203b304d6f3883",
        // 解限风险类型
        RelieveRiskType:"SMID_MERCHANT",
        // 提交的凭证数据
        RelieveCertDataList:getRelieveCertDataList(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantComplaintSubmitCertificatesRequest(dgReq)
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

func getRelieveCertDataList() string {
    dto := make(map[string]interface{})
    // 凭证的唯一ID
    dto["request_id"] = "1efc8c73afd64fc1b1fc50a834a54be0"
    // 凭证类型
    dto["type"] = "IMAGE"
    // 凭证code
    dto["code"] = "904"
    // 凭证的内容
    dto["info_data"] = "edd2d893-d3c2-342b-9ded-993913effce9"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

