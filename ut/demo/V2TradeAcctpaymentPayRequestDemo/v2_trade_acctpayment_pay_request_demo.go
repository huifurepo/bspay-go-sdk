/**
 * 余额支付 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeAcctpaymentPayRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeAcctpaymentPayRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeAcctpaymentPayRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 出款方商户号
        OutHuifuId:"6666000109133323",
        // 支付金额
        OrdAmt:"0.01",
        // 分账对象
        AcctSplitBunch:get887321ddC9244df89a700d1bbcc6fe17(),
        // 安全信息
        RiskCheckData:get6d96310288604dbfBe6a8eff149352ec(),
        // 资金类型资金类型。支付渠道为中信E管家时，资金类型必填（[详见说明](https://paas.huifu.com/open/doc/api/#/yuer/api_zxegjzllx)）
        // FundType:"test",
        // 手续费承担方标识余额支付手续费承担方标识；商户余额支付扣收规则为接口指定承担方时必填！枚举值：&lt;br/&gt;OUT：出款方；&lt;br/&gt;IN：分账接受方。&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：IN&lt;/font&gt;
        // TransFeeTakeFlag:"test",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeAcctpaymentPayRequest(dgReq)
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
    // 发起方商户号
    // extendInfoMap["huifu_id"] = ""
    // 商品描述
    // extendInfoMap["good_desc"] = ""
    // 备注
    // extendInfoMap["remark"] = ""
    // 是否延迟交易
    // extendInfoMap["delay_acct_flag"] = ""
    // 出款方账户号
    // extendInfoMap["out_acct_id"] = ""
    // 支付渠道
    // extendInfoMap["acct_channel"] = ""
    // 灵活用工标志
    // extendInfoMap["hyc_flag"] = ""
    // 灵活用工平台
    // extendInfoMap["lg_platform_type"] = ""
    // 代发模式
    // extendInfoMap["salary_modle_type"] = ""
    // 落地公司商户号
    // extendInfoMap["bmember_id"] = ""
    // 乐接活请求参数集合
    // extendInfoMap["ljh_data"] = get9dbad888434a466785f39d12cfb2ecbc()
    // 异步通知地址
    // extendInfoMap["notify_url"] = ""
    return extendInfoMap
}

func getF90f4f31607e4aa5934d9d69fbb958d2() interface{} {
    dto := make(map[string]interface{})
    // 分账接收方ID
    dto["huifu_id"] = "6666000109133323"
    // 分账金额
    dto["div_amt"] = "0.01"
    // 账户号
    // dto["acct_id"] = ""
    // 分账百分比%
    // dto["percentage_div"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get887321ddC9244df89a700d1bbcc6fe17() string {
    dto := make(map[string]interface{})
    // 分账明细
    dto["acct_infos"] = getF90f4f31607e4aa5934d9d69fbb958d2()
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get6d96310288604dbfBe6a8eff149352ec() string {
    dto := make(map[string]interface{})
    // 转账原因
    dto["transfer_type"] = "04"
    // 产品子类
    dto["sub_product"] = "1"
    // 纬度
    // dto["latitude"] = ""
    // 经度
    // dto["longitude"] = ""
    // 基站地址
    // dto["base_station"] = ""
    // IP地址
    // dto["ip_addr"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get9dbad888434a466785f39d12cfb2ecbc() string {
    dto := make(map[string]interface{})
    // 税源地ID
    // dto["tax_area_id"] = ""
    // 任务模板ID
    // dto["template_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

