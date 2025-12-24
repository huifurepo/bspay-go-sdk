/**
 * 创建批量账单数据 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V3BillpayOrderBatchAddRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V3BillpayOrderBatchAddRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V3BillpayOrderBatchAddRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000123123123",
        // 账单项目编号
        ProjectNo:"BN2025091279190693",
        // 用户资料信息列表
        // UserDocInfoList:get5c2e7c16296940238d680181de1820b9(),
        // 账单收费项信息列表
        // PaymentInfoList:get892c7e04F78848c79b53333c9ca3135a(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V3BillpayOrderBatchAddRequest(dgReq)
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
    // 是否生效
    // extendInfoMap["effective_flag"] = ""
    return extendInfoMap
}

func get5c2e7c16296940238d680181de1820b9() string {
    dto := make(map[string]interface{})
    // 账单表单字段属性ID
    // dto["key_no"] = "test"
    // 账单表单字段属性值
    // dto["key_value"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get892c7e04F78848c79b53333c9ca3135a() string {
    dto := make(map[string]interface{})
    // 账单表单字段属性ID
    // dto["key_no"] = "test"
    // 账单表单字段属性值
    // dto["key_value"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

