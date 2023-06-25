/**
 * 商户业务开通修改 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantBusiModifyRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantBusiModifyRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantBusiModifyRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付客户Id
        HuifuId:"6666000103668046",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantBusiModifyRequest(dgReq)
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
    // 经营简称
    extendInfoMap["short_name"] = ""
    // 税务登记证
    extendInfoMap["tax_reg_pic"] = ""
    // 公司照片一
    extendInfoMap["comp_pic1"] = ""
    // 公司照片二
    extendInfoMap["comp_pic2"] = ""
    // 公司照片三
    extendInfoMap["comp_pic3"] = ""
    // 法人身份证反面
    extendInfoMap["legal_cert_back_pic"] = ""
    // 法人身份证正面
    extendInfoMap["legal_cert_front_pic"] = ""
    // 营业执照图片
    extendInfoMap["license_pic"] = ""
    // 组织机构代码证
    extendInfoMap["org_code_pic"] = ""
    // 商务协议
    extendInfoMap["ba_pic"] = ""
    // 开户许可证
    extendInfoMap["reg_acct_pic"] = ""
    // 结算卡反面
    extendInfoMap["settle_card_back_pic"] = ""
    // 结算卡正面
    extendInfoMap["settle_card_front_pic"] = ""
    // 结算人身份证反面
    extendInfoMap["settle_cert_back_pic"] = ""
    // 结算人身份证正面
    extendInfoMap["settle_cert_front_pic"] = ""
    // 授权委托书
    extendInfoMap["auth_enturst_pic"] = "[http://192.168.85.157:30031/sspm/testVirgo](http://192.168.85.157:30031/sspm/testVirgo)"
    // 协议信息实体
    extendInfoMap["agreement_info"] = getAgreementInfo()
    // 是否交易手续费外扣
    extendInfoMap["out_fee_flag"] = "2"
    // 交易手续费外扣汇付ID
    extendInfoMap["out_fee_huifuid"] = ""
    // 交易手续费外扣时的账户类型
    extendInfoMap["out_fee_acct_type"] = ""
    // 是否开通网银
    extendInfoMap["online_flag"] = ""
    // 是否开通快捷
    extendInfoMap["quick_flag"] = ""
    // 是否开通代扣
    extendInfoMap["withhold_flag"] = ""
    // 延迟入账开关
    extendInfoMap["delay_flag"] = "Y"
    // 开通支付宝预授权
    extendInfoMap["alipay_pre_auth_flag"] = "Y"
    // 开通微信预授权
    // extendInfoMap["wechatpay_pre_auth_flag"] = ""
    // 商户业务类型
    // extendInfoMap["mer_bus_type"] = ""
    // 线上费率配置
    // extendInfoMap["online_fee_conf_list"] = getOnlineFeeConfList()
    // 支付宝配置对象
    extendInfoMap["ali_conf_list"] = getAliConfList()
    // 微信配置对象
    extendInfoMap["wx_conf_list"] = getWxConfList()
    // 银联二维码配置对象
    extendInfoMap["union_conf_list"] = getUnionConfList()
    // 银行卡支付配置信息
    extendInfoMap["bank_card_conf"] = getBankCardConf()
    // 余额支付配置对象
    extendInfoMap["balance_pay_config"] = getBalancePayConfig()
    // 补贴支付
    extendInfoMap["combine_pay_config"] = getCombinePayConfig()
    // 线上手续费承担方配置
    // extendInfoMap["online_pay_fee_conf_list"] = getOnlinePayFeeConfList()
    // 银行大额转账对象
    // extendInfoMap["bank_big_amt_pay_config"] = getBankBigAmtPayConfig()
    // 全域资金管理配置
    // extendInfoMap["out_order_funds_config"] = getOutOrderFundsConfig()
    // 汇总结算配置实体
    // extendInfoMap["collection_settle_config_list"] = getCollectionSettleConfigList()
    // 微信直连配置对象
    // extendInfoMap["wx_zl_conf"] = getWxZlConf()
    // 异步消息接收地址
    extendInfoMap["async_return_url"] = "[http://www.baidu55.com](http://www.baidu55.com/)"
    // 业务开通结果异步消息接收地址
    extendInfoMap["busi_async_return_url"] = ""
    // 交易异步应答地址
    extendInfoMap["recon_resp_addr"] = "[http://192.168.85.157:30031/sspm/testVirgo](http://192.168.85.157:30031/sspm/testVirgo)"
    return extendInfoMap
}

func getAgreementInfo() string {
    dto := make(map[string]interface{})
    // 协议类型
    dto["agreement_type"] = "0"
    // 协议开始日期
    dto["agree_begin_date"] = "20200325"
    // 协议结束日期
    dto["agree_end_date"] = "20400325"
    // 协议模板号
    dto["agreement_model"] = "202106070100000380"
    // 协议模板名称
    dto["agreement_name"] = "电子协议签约模板"
    // 协议号
    dto["agreement_no"] = "202106070100000380"
    // 签约日期
    dto["sign_date"] = "20200325"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getOnlineFeeConfList() string {
    dto := make(map[string]interface{})
    // 业务类型
    // dto["fee_type"] = "test"
    // 银行编码
    // dto["bank_id"] = "test"
    // 借贷标志
    // dto["dc_flag"] = "test"
    // 费率状态
    // dto["stat_flag"] = "test"
    // 手续费（固定/元）
    // dto["fix_amt"] = ""
    // 费率（百分比/%）
    // dto["fee_rate"] = ""
    // 银行名称
    // dto["bank_name"] = ""
    // 银行中文简称
    // dto["bank_short_chn"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getAliConfList() string {
    dto := make(map[string]interface{})
    // 手续费（%）
    dto["fee_rate"] = "3.15"
    // 支付场景
    dto["pay_scene"] = "1"
    // 商户经营类目
    dto["mcc"] = "2016062900190337"
    // 子渠道号
    dto["pay_channel_id"] = "10000001"
    // 拟申请的间联商户等级
    dto["indirect_level"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getWxConfList() string {
    dto := make(map[string]interface{})
    // 手续费（%）
    dto["fee_rate"] = "2.15"
    // 支付场景
    dto["pay_scene"] = "10"
    // 商户经营类目[参见微信支付宝MCC](https://paas.huifu.com/partners/api/#/csfl/api_csfl_zfbmcc)；个体工商户、小微商户必填，企业使用fee_rule_id代替原有mcc
    dto["mcc"] = "5943"
    // 费率规则ID
    dto["fee_rule_id"] = "765"
    // 子渠道号
    dto["pay_channel_id"] = "JP00001"
    // 申请服务
    dto["service_codes"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getUnionConfList() string {
    dto := make(map[string]interface{})
    // 银联二维码1000以上借记卡费率
    dto["debit_fee_rate_up"] = "6"
    // 银联二维码1000以下借记卡费率
    dto["debit_fee_rate_down"] = "2.55"
    // 银联二维码1000以下贷记卡费率
    dto["credit_fee_rate_down"] = "1"
    // 银联二维码1000以上贷记卡费率
    dto["credit_fee_rate_up"] = "6.566"
    // 银行业务手续费类型
    dto["charge_cate_code"] = ""
    // 银联二维码1000以上借记卡费率封顶值
    dto["debit_fee_limit_up"] = "641"
    // 银联二维码1000以下借记卡费率封顶值
    dto["debit_fee_limit_down"] = "11.3"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getBankCardConf() string {
    dto := make(map[string]interface{})
    // 借记卡手续费（%）
    dto["debit_fee_rate"] = "3"
    // 贷记卡手续费（%）
    dto["credit_fee_rate"] = "6"
    // 银行业务手续费类型
    dto["charge_cate_code"] = "02"
    // 借记卡封顶值
    dto["debit_fee_limit"] = "5"
    // 云闪付借记卡手续费1000以上（%）
    dto["cloud_debit_fee_rate_up"] = "7"
    // 云闪付借记卡封顶1000以上(元)
    dto["cloud_debit_fee_limit_up"] = "8.922"
    // 云闪付贷记卡手续费1000以上（%）
    dto["cloud_credit_fee_rate_up"] = "4.86"
    // 云闪付借记卡手续费1000以下（%）
    dto["cloud_debit_fee_rate_down"] = "0"
    // 云闪付借记卡封顶1000以下(元)
    dto["cloud_debit_fee_limit_down"] = "10"
    // 云闪付贷记卡手续费1000以下（%）
    dto["cloud_credit_fee_rate_down"] = "2"
    // 是否开通小额双免
    dto["is_open_small_flag"] = "0"
    // 小额双免单笔限额(元)
    dto["small_free_amt"] = "500"
    // 小额双免手续费（%）
    dto["small_fee_amt"] = "1"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getBalancePayConfig() string {
    dto := make(map[string]interface{})
    // 支付手续费(%)
    dto["fee_rate"] = "2"
    // 支付固定手续费(元)
    dto["fee_fix_amt"] = "1"
    // 费率开关
    dto["switch_state"] = "1"
    // 交易手续费外扣时的账户类型
    // dto["out_fee_acct_type"] = ""
    // 交易手续费外扣汇付ID
    // dto["out_fee_huifuid"] = ""
    // 是否交易手续费外扣
    // dto["out_fee_flag"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getCombinePayConfig() string {
    dto := make(map[string]interface{})
    // 开通状态
    dto["switch_state"] = "0"
    // 支付手续费(%)
    dto["fee_rate"] = "10"
    // 支付固定手续费(元)
    dto["fee_fix_amt"] = "5"
    // 交易手续费外扣时的账户类型
    // dto["out_fee_acct_type"] = ""
    // 交易手续费外扣汇付ID
    // dto["out_fee_huifuid"] = ""
    // 是否交易手续费外扣
    // dto["out_fee_flag"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getOnlinePayFeeConfList() string {
    dto := make(map[string]interface{})
    // 业务类型
    // dto["pay_type"] = ""
    // 交易手续费外扣时的账户类型
    // dto["out_fee_acct_type"] = ""
    // 交易手续费外扣汇付ID
    // dto["out_fee_huifuid"] = ""
    // 是否交易手续费外扣
    // dto["out_fee_flag"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getBankBigAmtPayConfig() string {
    dto := make(map[string]interface{})
    // 开关状态
    // dto["switch_state"] = ""
    // 大额调账标识申请类型
    // dto["biz_type"] = ""
    // 费率（百分比/%）
    // dto["fee_rate"] = ""
    // 交易手续费（固定/元）
    // dto["fee_fix_amt"] = ""
    // 手续费外扣标记
    // dto["out_fee_flag"] = ""
    // 手续费外扣时的汇付ID
    // dto["out_fee_huifuid"] = ""
    // 外扣手续费费承担账户号
    // dto["out_fee_acct_id"] = ""
    // 银行大额转账单笔额度
    // dto["big_amt_limit_per_time"] = ""
    // 银行大额转账单日额度
    // dto["big_amt_limit_per_day"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getOutOrderAcctCard() string {
    dto := make(map[string]interface{})
    // 结算账户名
    // dto["card_name"] = "test"
    // 银行卡号
    // dto["card_no"] = "test"
    // 卡类型
    // dto["card_type"] = "test"
    // 持卡人证件类型00:身份证；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：00&lt;/font&gt;；card_type为1时选填。
    // dto["cert_type"] = "test"
    // 持卡人证件有效期（起始）card_type为1时选填；格式：yyyyMMdd，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20210830&lt;/font&gt;；&lt;br/&gt;若填写cert_no，cert_validity_type，cert_type需同时填写。
    // dto["cert_begin_date"] = "test"
    // 持卡人证件有效期（截止）cert_validity_type变更为0时必填，格式：yyyyMMdd；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20210830&lt;/font&gt;
    // dto["cert_end_date"] = "test"
    // 持卡人证件号码card_type为1时选填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：310112200001018888&lt;/font&gt;；
    // dto["cert_no"] = "test"
    // 银行卡绑定手机号
    // dto["mp"] = "test"
    // 银行所在省
    // dto["prov_id"] = ""
    // 银行所在市
    // dto["area_id"] = ""
    // 银行编码
    // dto["bank_code"] = ""
    // 支行联行号
    // dto["branch_code"] = ""
    // 支行名称
    // dto["branch_name"] = ""
    // 持卡人证件有效期类型
    // dto["cert_validity_type"] = ""
    // 开户许可证核准号
    // dto["open_licence_no"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getOutOrderAcctOpenFees() string {
    dto := make(map[string]interface{})
    // 支付固定手续费(元)
    // dto["fee_fix_amt"] = "test"
    // 开户手续费外扣时的账户类型
    // dto["out_fee_acct_type"] = "test"
    // 开户手续费外扣汇付ID
    // dto["out_fee_huifuid"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getOutOrderFundsConfig() string {
    dto := make(map[string]interface{})
    // 开通状态
    // dto["switch_state"] = "test"
    // 自动入账开关
    // dto["out_order_auto_acct_flag"] = "test"
    // 批次入账时间10:00-10点批次入账&lt;br/&gt;16:00-16点批次入账&lt;/br&gt;开通批次入账时必填 ，多个批次使用逗号分隔；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：10:00,16:00&lt;/font&gt;
    // dto["batch_no"] = "test"
    // 批量入账开关
    // dto["batch_auto_acct_flag"] = ""
    // 支付手续费(%)
    // dto["fee_rate"] = ""
    // 支付固定手续费(元)
    // dto["fee_fix_amt"] = ""
    // 手续费最小值(元)
    // dto["fee_min_amt"] = ""
    // 交易手续费外扣时的账户类型
    // dto["out_fee_acct_type"] = ""
    // 交易手续费外扣标记
    // dto["out_fee_flag"] = ""
    // 交易手续费外扣汇付ID
    // dto["out_fee_huifuid"] = ""
    // 全域资金开户使用的银行卡信息
    // dto["out_order_acct_card"] = getOutOrderAcctCard()
    // 全域资金开户手续费
    // dto["out_order_acct_open_fees"] = getOutOrderAcctOpenFees()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getCollectionSettleConfigList() string {
    dto := make(map[string]interface{})
    // 归集留存金(元)
    // dto["out_resv_amt"] = "test"
    // 转入商户号
    // dto["in_huifu_id"] = "test"
    // 转入账户
    // dto["in_acct_id"] = "test"
    // 生效日期
    // dto["valid_date"] = "test"
    // 功能开关
    // dto["switch_state"] = ""
    // 转出账户
    // dto["out_acct_id"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getWxZlPayConfList() interface{} {
    dto := make(map[string]interface{})
    // 申请服务
    // dto["service_code"] = "test"
    // 功能服务appid
    // dto["sub_app_id"] = "test"
    // 功能开关
    // dto["switch_state"] = "test"
    // 功能费率(%)
    // dto["fee_rate"] = "test"

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getWxZlConf() string {
    dto := make(map[string]interface{})
    // 微信子商户号
    // dto["sub_mch_id"] = "test"
    // 配置集合
    // dto["wx_zl_pay_conf_list"] = getWxZlPayConfList()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

