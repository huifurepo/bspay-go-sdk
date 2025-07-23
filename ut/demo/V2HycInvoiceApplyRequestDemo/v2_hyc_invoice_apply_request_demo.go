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
        // 开票类目
        InvoiceCategory:"信息技术服务*软件测试服务",
        // 汇付全局流水号集合
        // HfSeqIds:"test",
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
    // 接收人手机号
    extendInfoMap["receive_mobile"] = "13945641357"
    // 接收人姓名
    extendInfoMap["receive_name"] = "黄二"
    // 快递地址
    extendInfoMap["courier_address"] = "长江大街161号上海长江经济园"
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

