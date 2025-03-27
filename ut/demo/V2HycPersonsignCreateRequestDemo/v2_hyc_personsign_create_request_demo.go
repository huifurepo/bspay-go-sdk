/**
 * 个人签约发起 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2HycPersonsignCreateRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2HycPersonsignCreateRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2HycPersonsignCreateRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 用户汇付id
        HuifuId:"6666000145962643",
        // 落地公司机构号
        MinorAgentId:"L20231113140106443",
        // 乐接活请求参数jsonObject格式 合作平台为乐接活时必传
        // LjhData:getDa04d656159b4ec89e2b3b57e7164683(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2HycPersonsignCreateRequest(dgReq)
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
    // 合作平台
    // extendInfoMap["lg_platform_type"] = ""
    // 是否发送签约短信
    extendInfoMap["send_sms_flag"] = "Y"
    // 签约结果通知地址
    extendInfoMap["asyn_url"] = ""
    return extendInfoMap
}

func getDa04d656159b4ec89e2b3b57e7164683() string {
    dto := make(map[string]interface{})
    // 合同模板id合作平台为乐接活时必填 数字格式
    // dto["contract_template_id"] = "test"
    // 任务模板id合作平台为乐接活时必填 数字格式
    // dto["task_template_id"] = "test"
    // 税源地id合作平台为乐接活时必填 数字格式
    // dto["tax_area_id"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

