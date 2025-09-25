/**
 * 银行大额支付差错申请 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeOnlinepaymentTransferBankmistakeApplyRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeOnlinepaymentTransferBankmistakeApplyRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeOnlinepaymentTransferBankmistakeApplyRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000110468104",
        // 交易金额
        TransAmt:"0.01",
        // 订单类型
        OrderType:"REFUND",
        // 原请求流水号order_flag&#x3D;Y时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：2022012514120615009&lt;/font&gt;
        OrgReqSeqId:"202308312345678931",
        // 原请求日期格式:yyyyMMdd；order_flag&#x3D;Y时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20220125&lt;/font&gt;
        OrgReqDate:"20230831",
        // 异步通知地址
        NotifyUrl:"http://www.baidu.com",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeOnlinepaymentTransferBankmistakeApplyRequest(dgReq)
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
    // 下单标识
    // extendInfoMap["order_flag"] = ""
    // 备注
    extendInfoMap["remark"] = "大额支付补入账验证"
    // 银行信息数据
    extendInfoMap["bank_info_data"] = getB60d3b633db64945Aa3d7e2d7d75114e()
    // 延时标记
    // extendInfoMap["delay_acct_flag"] = ""
    // 分账对象
    // extendInfoMap["acct_split_bunch"] = getF5a7bd570f5b47e58d43Af05a3b46050()
    // 实际打款信息
    // extendInfoMap["actual_remit_data"] = getC9fbe6ceA66f4533A4e8716e37ca1292()
    return extendInfoMap
}

func getB60d3b633db64945Aa3d7e2d7d75114e() string {
    dto := make(map[string]interface{})
    // 银行编号
    dto["bank_code"] = "03080000"
    // 对公对私标识
    dto["card_acct_type"] = "P"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get7e4a26656c684722897041d65438a95a() interface{} {
    dto := make(map[string]interface{})
    // 支付金额
    // dto["div_amt"] = ""
    // 商户号
    // dto["huifu_id"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getF5a7bd570f5b47e58d43Af05a3b46050() string {
    dto := make(map[string]interface{})
    // 分账信息列表
    // dto["acct_infos"] = get7e4a26656c684722897041d65438a95a()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getC9fbe6ceA66f4533A4e8716e37ca1292() string {
    dto := make(map[string]interface{})
    // 实际打款日期
    // dto["actual_remit_date"] = "test"
    // 实际打款方姓名
    // dto["actual_remit_name"] = "test"
    // 实际打款金额
    // dto["actual_remit_amt"] = "test"
    // 实际打款方银行卡号
    // dto["actual_remit_card_no"] = "test"
    // 汇款凭证文件ID
    // dto["certificate_file_id"] = "test"
    // 退款卡标识
    // dto["refund_card_flag"] = "test"
    // 实际打款卡号银行名称
    // dto["actual_bank_name"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

