/**
 * 更新退款审批结果 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantComplaintUpdateRefundprogressRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantComplaintUpdateRefundprogressRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantComplaintUpdateRefundprogressRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
        // 投诉单号
        ComplaintId:"200000020221020220032600930",
        // 审批动作
        Action:"APPROVE",
        // 微信商户号
        MchId:"1502074862",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantComplaintUpdateRefundprogressRequest(dgReq)
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
    // 预计发起退款时间
    extendInfoMap["launch_refund_day"] = ""
    // 拒绝退款原因
    extendInfoMap["reject_reason"] = ""
    // 文件列表
    extendInfoMap["file_info"] = "{\"reject_media_pic1\":\"a8a096a3-0dd4-3b0e-886c-9afb20d23b1a\",\"reject_media_pic2\":\"a8a096a3-0dd4-3b0e-886c-9afb20d23b2a\",\"reject_media_pic3\":\"a8a096a3-0dd4-3b0e-886c-9afb20d23b3a\",\"reject_media_pic4\":\"a8a096a3-0dd4-3b0e-886c-9afb20d23b4a\"}"
    // 备注
    extendInfoMap["remark"] = "我是备注1111101"
    return extendInfoMap
}

