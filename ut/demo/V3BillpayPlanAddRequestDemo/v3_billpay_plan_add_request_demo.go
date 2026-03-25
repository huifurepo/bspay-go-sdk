/**
 * 创建账单计划 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V3BillpayPlanAddRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V3BillpayPlanAddRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V3BillpayPlanAddRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000123123123",
        // 账单项目编号
        // ProjectNo:"test",
        // 账单周期
        PlanCycle:"MONTH",
        // 账单日
        BillDay:"15",
        // 补发当前周期账单标志枚举:Y-是、N-否；指定账单日时，必填；若填写是，则立即生成当前系统时间所在周期的账单； 滚动账单日时，此字段无效
        ReissueBillFlag:"Y",
        // 代扣信息jsonObject格式；账单计划需自动代扣时必填
        WithholdInfoData:get49450184739c47e6Bab527b61aa24f1a(),
        // 用户资料信息列表
        UserDocInfoList:get919c625f5db74bca802a9a31f9647a55(),
        // 账单收费项信息列表
        PaymentInfoList:get64292bd749be4204A99cA4e9c977a623(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V3BillpayPlanAddRequest(dgReq)
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
    // 账单计划有效期
    extendInfoMap["plan_expire_date"] = "20251231"
    return extendInfoMap
}

func get49450184739c47e6Bab527b61aa24f1a() string {
    dto := make(map[string]interface{})
    // 卡令牌
    dto["token_no"] = "CT202412270001"
    // 是否发送代扣前短信通知
    dto["sms_notify_flag"] = "Y"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get919c625f5db74bca802a9a31f9647a55() string {
    dto := make(map[string]interface{})
    // 账单表单字段属性ID
    dto["key_no"] = "userName"
    // 账单表单字段属性值
    dto["key_value"] = "张三"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get64292bd749be4204A99cA4e9c977a623() string {
    dto := make(map[string]interface{})
    // 账单表单字段属性ID
    dto["key_no"] = "propertyFee"
    // 账单表单字段属性值
    dto["key_value"] = "500.00"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

