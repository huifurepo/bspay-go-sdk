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
    // 银联配置对象
    extendInfoMap["union_conf_list"] = getUnionConfList()
    // 银行卡支付配置信息
    extendInfoMap["bank_card_conf"] = getBankCardConf()
    // 支付宝配置对象
    extendInfoMap["ali_conf_list"] = getAliConfList()
    // 开通支付宝预授权
    extendInfoMap["alipay_pre_auth_flag"] = "Y"
    // 微信配置对象
    extendInfoMap["wx_conf_list"] = getWxConfList()
    // 开通微信预授权
    // extendInfoMap["wechatpay_pre_auth_flag"] = ""
    // 营销补贴
    extendInfoMap["combine_pay_config"] = getCombinePayConfig()
    // 余额支付配置对象
    extendInfoMap["balance_pay_config"] = getBalancePayConfig()
    // 异步消息接收地址
    extendInfoMap["async_return_url"] = "[http://www.baidu55.com](http://www.baidu55.com/)"
    // 业务开通结果异步消息接收地址
    extendInfoMap["busi_async_return_url"] = ""
    // 交易异步应答地址
    extendInfoMap["recon_resp_addr"] = "[http://192.168.85.157:30031/sspm/testVirgo](http://192.168.85.157:30031/sspm/testVirgo)"
    // 线上费率配置
    // extendInfoMap["online_fee_conf_list"] = getOnlineFeeConfList()
    // 商户业务类型
    // extendInfoMap["mer_bus_type"] = ""
    // 线上手续费承担方配置
    // extendInfoMap["online_pay_fee_conf_list"] = getOnlinePayFeeConfList()
    // 银行大额转账对象
    // extendInfoMap["bank_big_amt_pay_config"] = getBankBigAmtPayConfig()
    return extendInfoMap
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

func getWxConfList() string {
    dto := make(map[string]interface{})
    // 手续费（%）
    dto["fee_rate"] = "2.15"
    // 支付场景
    dto["pay_scene"] = "10"
    // 商户经营类目
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

