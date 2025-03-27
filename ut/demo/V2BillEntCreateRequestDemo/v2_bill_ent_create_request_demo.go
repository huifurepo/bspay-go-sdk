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
        PayeeInfo:get7cb71210B6e945638ed7D26989882224(),
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
    // extendInfoMap["bill_summary_info"] = getB7fe106a1ade4660865e83374ad40fbf()
    // 更多信息
    // extendInfoMap["bill_extend_info"] = get31ba075559984fd7B7d0C256e889b5c0()
    // 账单推送方式
    extendInfoMap["push_type"] = "EMAIL"
    // 抄送邮箱
    extendInfoMap["copy_email"] = "xuemei.ren@huifu.com,guowen.jiang@huifu.com"
    // 备注信息
    extendInfoMap["remark"] = "I_remark"
    // 账单信息异步通知地址
    extendInfoMap["notify_url"] = "https://spin-test.cloudpnr.com/trade/billing/pcredit/status"
    // 回调地址
    extendInfoMap["front_url"] = "https://spin-test.cloudpnr.com/trade/billing/pcredit/status"
    return extendInfoMap
}

func getB7fe106a1ade4660865e83374ad40fbf() string {
    dto := make(map[string]interface{})
    // 字段名
    // dto["extend_name"] = "test"
    // 字段值
    // dto["extend_value"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get31ba075559984fd7B7d0C256e889b5c0() string {
    dto := make(map[string]interface{})
    // 字段名
    // dto["extend_name"] = "test"
    // 字段值
    // dto["extend_value"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get7cb71210B6e945638ed7D26989882224() string {
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

