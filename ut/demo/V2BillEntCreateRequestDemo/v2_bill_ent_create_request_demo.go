/**
 * 创建企业账单 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2BillEntCreateRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2BillEntCreateRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2BillEntCreateRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000003100615",
        // 付款人
        PayerId:"P2024082286286456",
        // 账单名称
        BillName:"账单名称是水电费吧",
        // 账单金额
        BillAmt:"0.02",
        // 可支持的付款方式
        SupportPayType:"wx,alipay,online_b2c,online_b2b",
        // 账单截止日期
        BillEndDate:"20990909",
        // 收款人信息
        PayeeInfo:get0117bed87ef6433bAc7f831d09c77a76(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2BillEntCreateRequest(dgReq)
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
    // 账单说明
    extendInfoMap["bill_remark"] = "您本次 SaaS 服务周期为[开始日期]至[结束日期]。费用包括基础服务套餐[X]元，高级功能模块[X]元，总计[X]元。"
    // 汇总信息
    extendInfoMap["bill_summary_info"] = get00cfe2af4f494258A34084b55e2d10f9()
    // 更多信息
    extendInfoMap["bill_extend_info"] = get0d92fec96d1c42aaA519B4fe44b194f4()
    // 账单推送方式
    extendInfoMap["push_type"] = "EMAIL"
    // 抄送邮箱
    extendInfoMap["copy_email"] = "xxx@163.com,xxxx@163.com"
    // 备注信息
    extendInfoMap["remark"] = "I_remark"
    // 账单信息异步通知地址
    extendInfoMap["notify_url"] = "https://spin-test.cloudpnr.com/trade/billing/pcredit/status"
    // 回调地址
    extendInfoMap["front_url"] = "https://spin-test.cloudpnr.com/trade/billing/pcredit/status"
    return extendInfoMap
}

func get00cfe2af4f494258A34084b55e2d10f9() string {
    dto := make(map[string]interface{})
    // 字段名
    dto["extend_name"] = "账单金额"
    // 字段值
    dto["extend_value"] = "128.00"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get0d92fec96d1c42aaA519B4fe44b194f4() string {
    dto := make(map[string]interface{})
    // 字段名
    dto["extend_name"] = "备注"
    // 字段值
    dto["extend_value"] = "额外额外"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get0117bed87ef6433bAc7f831d09c77a76() string {
    dto := make(map[string]interface{})
    // 收款联系人姓名
    dto["payee_name"] = "黄云"
    // 收款联系人手机payee_mobile_no、payee_tel、payee_email 三选一必填
    dto["payee_mobile_no"] = "13456787678"
    // 收款联系人座机payee_mobile_no、payee_tel、payee_email 三选一必填
    dto["payee_tel"] = "0211234444"
    // 收款联系人邮箱payee_mobile_no、payee_tel、payee_email 三选一必填
    dto["payee_email"] = "lieecho@163.com"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

