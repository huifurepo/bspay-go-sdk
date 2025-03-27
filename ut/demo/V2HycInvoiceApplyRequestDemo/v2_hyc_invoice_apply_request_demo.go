/**
 * 申请开票 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2HycInvoiceApplyRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2HycInvoiceApplyRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2HycInvoiceApplyRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户汇付id
        HuifuId:"6666000109133323",
        // 交易流水列表
        BatchList:get7087791b2c6441789d3fAa1688a130fb(),
        // 接收人手机号
        ReceiveMobile:"13945641357",
        // 接收人姓名
        ReceiveName:"黄二",
        // 快递地址
        CourierAddress:"长江大街161号上海长江经济园",
        // 开票类目
        InvoiceCategory:"信息技术服务*软件测试服务",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2HycInvoiceApplyRequest(dgReq)
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
    // 异步地址
    extendInfoMap["asyn_url"] = "https://www.baidu.com"
    // 购方税号
    extendInfoMap["invoice_tax_no"] = "91310230MA1JTWAK98"
    // 购方公司名称
    extendInfoMap["invoice_name"] = "上海汇涵信息科技服务有限公司"
    // 购方公司地址
    extendInfoMap["invoice_address"] = "长江大街161号上海长江经济园"
    // 购方公司银行账号
    extendInfoMap["invoice_card_num"] = "631252533"
    // 购方银行名称
    extendInfoMap["invoice_bank_name"] = "中国民生银行股份有限公司"
    // 购方联系电话
    extendInfoMap["invoice_phone"] = "400-820-2819"
    // 发票类型
    extendInfoMap["invoice_type"] = "1"
    // 备注
    extendInfoMap["remarks"] = ""
    return extendInfoMap
}

func get7087791b2c6441789d3fAa1688a130fb() string {
    dto := make(map[string]interface{})
    // 交易流水号
    dto["trans_seq_id"] = "SSPC8d4406cff4584b2391e113eaa32432bb"
    // 交易日期
    dto["trans_date"] = "20240112"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

