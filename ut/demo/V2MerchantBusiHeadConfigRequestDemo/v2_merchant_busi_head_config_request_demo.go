/**
 * 开通下级商户权限配置接口 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantBusiHeadConfigRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantBusiHeadConfigRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantBusiHeadConfigRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
        // 汇付客户Id
        HuifuId:"6666000108854952",
        // 产品编号
        // ProductId:"test",
        // 直属渠道号
        // UpperHuifuId:"test",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantBusiHeadConfigRequest(dgReq)
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
    // 支付宝配置对象
    // extendInfoMap["ali_conf_list"] = getSubmerAliConfList()
    // 微信配置对象
    // extendInfoMap["wx_conf_list"] = getSubmerWxConfList()
    // 银联二维码配置对象
    // extendInfoMap["union_conf_list"] = getSubmerUnionConfList()
    // 银联卡配置对象
    // extendInfoMap["bank_card_config"] = getSubmerBankCardConfig()
    // 分账配置对象
    // extendInfoMap["split_config"] = getSubmerSplitConfig()
    // 微信直连配置对象
    // extendInfoMap["wx_zl_conf_list"] = getSubmerWxZlConfList()
    // 支付宝直连配置对象
    // extendInfoMap["ali_zl_conf"] = getSubmerAliZlConf()
    // 线上配置对象
    // extendInfoMap["online_fee_conf_list"] = getSubmerOnlineFeeConfList()
    // 余额支付配置对象
    // extendInfoMap["balance_pay_config"] = getSubmerBalancePayConfig()
    // 补贴支付配置对象
    // extendInfoMap["combine_pay_config"] = getSubmerCombinePayConfig()
    // 银行大额转账配置对象
    // extendInfoMap["bank_big_amt_pay_config"] = getSubmerBankBigAmtPayConfig()
    // 全域资金管理配置对象（华通银行）
    // extendInfoMap["out_order_funds_config"] = getSubmerOutOrderFundsConfig()
    // 全域资金管理配置(新网银行)
    // extendInfoMap["out_order_funds_new_net_config"] = getOutOrderFundsNewNetConfig()
    // 结算配置对象
    // extendInfoMap["settle_config_list"] = getSubmerSettleConfigList()
    // 取现配置对象
    // extendInfoMap["cash_config_list"] = getSubmerCashConfigList()
    // 外扣配置对象
    // extendInfoMap["out_fee_config"] = getSubmerOutFeeConfig()
    // 允许开通支付宝预授权
    // extendInfoMap["alipay_pre_auth_flag"] = ""
    // 允许开通微信预授权
    // extendInfoMap["wechatpay_pre_auth_flag"] = ""
    // 允许开通商户定时结算
    // extendInfoMap["mer_timing_settle_flag"] = ""
    // 允许开通商户优先结算
    // extendInfoMap["mer_prior_settle_flag"] = ""
    // 允许使用上级商户经营信息
    // extendInfoMap["use_upper_mer_auth_flag"] = ""
    // 允许使用上级商户号发起AT交易
    // extendInfoMap["use_upper_mer_at_trans_flag"] = ""
    return extendInfoMap
}

func getSubmerAliConfList() string {
    dto := make(map[string]interface{})
    // 支付场景
    // dto["pay_scene"] = "test"
    // 手续费
    // dto["fee_rate"] = "test"
    // 允许开通该业务
    // dto["open_flag"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getSubmerWxConfList() string {
    dto := make(map[string]interface{})
    // 支付场景
    // dto["pay_scene"] = "test"
    // 手续费
    // dto["fee_rate"] = "test"
    // 允许开通该场景业务
    // dto["open_flag"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getSubmerUnionConfList() string {
    dto := make(map[string]interface{})
    // 借记卡手续费1000以上(%)
    // dto["debit_fee_rate_up"] = "test"
    // 银联二维码业务贷记卡手续费1000以上(%)
    // dto["credit_fee_rate_up"] = "test"
    // 借记卡手续费1000以下(%)
    // dto["debit_fee_rate_down"] = "test"
    // 银联二维码业务贷记卡手续费1000以下(%)
    // dto["credit_fee_rate_down"] = "test"
    // 允许开通银联二维码业务
    // dto["open_flag"] = "test"
    // 银联业务手续费类型
    // dto["charge_cate_code"] = ""
    // 借记卡封顶1000以上（元）
    // dto["debit_fee_limit_up"] = ""
    // 借记卡封顶1000以下（元）
    // dto["debit_fee_limit_down"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getSubmerBankCardConfig() string {
    dto := make(map[string]interface{})
    // 借记卡手续费（%）
    // dto["debit_fee_rate"] = "test"
    // 贷记卡手续费（%）
    // dto["credit_fee_limit"] = "test"
    // 允许开通银行卡业务
    // dto["open_flag"] = "test"
    // 借记卡封顶值
    // dto["debit_fee_limit"] = ""
    // 银联手机闪付借记卡手续费1000以上（%）
    // dto["cloud_debit_fee_rate_up"] = ""
    // 银联手机闪付借记卡封顶1000以上（元）
    // dto["cloud_debit_fee_limit_up"] = ""
    // 银联手机闪付贷记卡手续费1000以上（%）
    // dto["cloud_credit_fee_rate_up"] = ""
    // 银联手机闪付借记卡手续费1000以下（%）
    // dto["cloud_debit_fee_rate_down"] = ""
    // 银联手机闪付借记卡封顶1000以下（元）
    // dto["cloud_debit_fee_limit_down"] = ""
    // 银联手机闪付贷记卡手续费1000以下（%）
    // dto["cloud_credit_fee_rate_down"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getSubmerSplitConfig() string {
    dto := make(map[string]interface{})
    // 手续费（%）
    // dto["fee_rate"] = "test"
    // 允许开通分账业务
    // dto["open_flag"] = "test"
    // 固定手续费(元)
    // dto["fee_fix_amt"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getSubmerWxZlConfList() string {
    dto := make(map[string]interface{})
    // 支付场景
    // dto["pay_scene"] = "test"
    // 手续费（%）
    // dto["fee_rate"] = "test"
    // 允许开通微信直连业务
    // dto["open_flag"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getSubmerAliZlConf() string {
    dto := make(map[string]interface{})
    // 手续费（%）
    // dto["fee_rate"] = "test"
    // 允许开通支付宝直连业务
    // dto["open_flag"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getSubmerOnlineFeeConfList() string {
    dto := make(map[string]interface{})
    // 业务类型
    // dto["bus_type"] = "test"
    // 借贷记标识
    // dto["dc_flag"] = "test"
    // 固定手续费(元)
    // dto["fee_fix_amt"] = "test"
    // 银行号
    // dto["bank_code"] = "test"
    // 手续费（%）
    // dto["fee_rate"] = "test"
    // 允许开通线上支付业务
    // dto["open_flag"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getSubmerBalancePayConfig() string {
    dto := make(map[string]interface{})
    // 手续费（%）
    // dto["fee_rate"] = "test"
    // 固定手续费(元)
    // dto["fee_fix_amt"] = "test"
    // 允许开通余额支付业务
    // dto["open_flag"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getSubmerCombinePayConfig() string {
    dto := make(map[string]interface{})
    // 手续费（%）
    // dto["fee_rate"] = "test"
    // 固定手续费(元)
    // dto["fee_fix_amt"] = "test"
    // 允许开通补贴支付业务
    // dto["open_flag"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getSubmerBankBigAmtPayConfig() string {
    dto := make(map[string]interface{})
    // 手续费（%）
    // dto["fee_rate"] = "test"
    // 固定手续费(元)
    // dto["fee_fix_amt"] = "test"
    // 允许开通大额转账业务
    // dto["open_flag"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getSubmerOutOrderFundsConfig() string {
    dto := make(map[string]interface{})
    // 手续费（%）
    // dto["fee_rate"] = "test"
    // 开户手续费（元）
    // dto["open_fee_fix_amt"] = "test"
    // 保底手续费(元)
    // dto["fee_fix_amt"] = "test"
    // 允许开通全域资金业务（华通）
    // dto["open_flag"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getOutOrderFundsNewNetConfig() string {
    dto := make(map[string]interface{})
    // 手续费（%）
    // dto["fee_rate"] = "test"
    // 保底手续费(元)
    // dto["fee_min_amt"] = "test"
    // 对公固定手续费(元)
    // dto["public_fee_fix_amt"] = "test"
    // 对私固定手续费(元)
    // dto["private_fee_fix_amt"] = "test"
    // 允许开通全域资金业务(新网)
    // dto["open_flag"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getSubmerSettleConfigList() string {
    dto := make(map[string]interface{})
    // 业务类型
    // dto["bus_type"] = "test"
    // 手续费（%）
    // dto["fee_rate"] = "test"
    // 固定手续费(元)
    // dto["fee_fix_amt"] = "test"
    // 允许开通结算配置
    // dto["open_flag"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getSubmerCashConfigList() string {
    dto := make(map[string]interface{})
    // 业务类型
    // dto["bus_type"] = "test"
    // 手续费（%）
    // dto["fee_rate"] = "test"
    // 固定手续费(元)
    // dto["fee_fix_amt"] = "test"
    // 允许开通取现配置
    // dto["open_flag"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getSubmerOutFeeConfig() string {
    dto := make(map[string]interface{})
    // 支持结算手续费外扣
    // dto["settle_out_fee_flag"] = "test"
    // 支持交易手续费外扣
    // dto["trans_fee_out_flag"] = "test"
    // 支持取现手续费外扣
    // dto["cash_out_fee_flag"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

