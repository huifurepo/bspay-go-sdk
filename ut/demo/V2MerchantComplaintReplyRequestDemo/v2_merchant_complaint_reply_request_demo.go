/**
 * 回复用户 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantComplaintReplyRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantComplaintReplyRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantComplaintReplyRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
        // 微信投诉单号
        ComplaintId:"200000020221020220032603511",
        // 被诉商户微信号
        ComplaintedMchid:"535295270",
        // 回复内容
        ResponseContent:"该问题请联系商家处理，谢谢。",
        // 微信商户号
        MchId:"1502073961",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantComplaintReplyRequest(dgReq)
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
    // 文件列表
    // extendInfoMap["file_info"] = get42b0f39dAdb944a8B720B6b3a55f3c41()
    // 跳转链接
    extendInfoMap["jump_url"] = ""
    // 跳转链接文案
    extendInfoMap["jump_url_text"] = ""
    return extendInfoMap
}

func get42b0f39dAdb944a8B720B6b3a55f3c41() string {
    dto := make(map[string]interface{})
    // 回复图片1
    // dto["response_pic1"] = ""
    // 回复图片2
    // dto["response_pic2"] = ""
    // 回复图片3
    // dto["response_pic3"] = ""
    // 回复图片4
    // dto["response_pic4"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

