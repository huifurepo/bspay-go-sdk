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
        // *协议信息实体*jsonObject字符串；[签约协议材料说明](https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/opps/api/prod/download_file/kyc/KYC-%E5%95%86%E6%88%B7%E5%90%88%E5%90%8C%E7%AD%BE%E7%BA%A6%E8%A7%84%E5%88%99.xlsx)&lt;br/&gt;若未签署过协议的情况下，调用该接口时必填；&lt;br/&gt;若已签署过纸质或电子协议下，调用该接口时必填，且只可以选择纸质或电子协议；&lt;br/&gt;若已签署过挂网协议下，调用该接口时，选填。
        AgreementInfo:getFa55923fC2d6483fB40034b57533c679(),
        // *线上业务类型编码*开通快捷、网银、余额支付Pro版、银行大额转账、分账比例&gt;30%、代扣必填；参见[线上业务类型编码及补充材料说明](https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/opps/api/prod/download_file/kyc/KYC-%E7%BA%BF%E4%B8%8A%E4%B8%9A%E5%8A%A1%E7%B1%BB%E5%9E%8B%E7%BC%96%E7%A0%81%E5%8F%8A%E8%A1%A5%E5%85%85%E6%9D%90%E6%96%99%E8%AF%B4%E6%98%8E.xlsx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：H7999AL&lt;/font&gt;
        // OnlineBusiType:"test",
        // 签约人jsonObject字符串；agreement_info中选择电子签约时必填；个人商户填本人信息。
        // SignUserInfo:getA7b56f0b78b74c61AbecC4ed3ea15bf6(),
        // 代扣场景withhold_flag开通需填写 1:权益投流(订阅),2:非权益投流(订阅),3:保险缴纳,4:道路通行费(ETC),5:公共事业(水电煤)缴费,6:物业缴费
        // WithholdPayScene:"test",
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
    // 微信支付宝商户简称
    extendInfoMap["short_name"] = ""
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
    // extendInfoMap["online_fee_conf_list"] = get4ac5fff288a34ac3Bdc5989b00d1edf1()
    // 支付宝配置对象
    extendInfoMap["ali_conf_list"] = get535af30dE2f14ad69841E9fcb7a95716()
    // 微信配置对象
    extendInfoMap["wx_conf_list"] = get36bf1c2f0bf24f8a8fd39675ed1b021f()
    // 银联二维码配置对象
    extendInfoMap["union_conf_list"] = get9ea086279d6a4d0d9330E6e40b55bddf()
    // 银行卡支付配置信息
    extendInfoMap["bank_card_conf"] = get790f276d862e48d8B01fE2e690e0e250()
    // *余额支付配置对象*
    extendInfoMap["balance_pay_config"] = getFacd221974d440bf8a9eB1984f5943b1()
    // 补贴支付
    extendInfoMap["combine_pay_config"] = get6725a4189ab7447a8ebfE033dc3cc458()
    // 线上手续费承担方配置
    // extendInfoMap["online_pay_fee_conf_list"] = get5379d0a89c6e4c89A02926c6b05b5b98()
    // 全域资金管理配置(华通银行)
    // extendInfoMap["out_order_funds_config"] = getCababba719dd457083e99d18c38532bb()
    // 异步消息接收地址
    extendInfoMap["async_return_url"] = "http://callback.example.com/sspm/testVirgo"
    // 业务开通结果异步消息接收地址
    extendInfoMap["busi_async_return_url"] = ""
    // 交易异步应答地址
    extendInfoMap["recon_resp_addr"] = "http://callback.example.com/sspm/testVirgo"
    // *运营媒介*
    // extendInfoMap["online_media_info_list"] = getB25246ad3b9447198d0a4fac0ab10450()
    // *补充文件信息*
    // extendInfoMap["extended_material_list"] = get61966f16B1eb4ca193e7F3bbdb034e6b()
    // 商户开通强制延迟标记
    // extendInfoMap["forced_delay_flag"] = ""
    // 微信直连配置对象
    // extendInfoMap["wx_zl_conf"] = getBc3c15758a314d418122291dda43bf2d()
    // 支付宝直连配置对象
    // extendInfoMap["ali_zl_conf"] = getFff133434e6a4cbd9e49C33a11c2da87()
    // 开户费用值(元)
    // extendInfoMap["enter_fee"] = ""
    // 开户费用类型
    // extendInfoMap["enter_fee_flag"] = ""
    // 是否开通在线退款
    // extendInfoMap["online_refund"] = ""
    // 是否支持平台退款
    // extendInfoMap["platform_refund"] = ""
    // 是否支持撤销
    // extendInfoMap["support_revoke"] = ""
    // 分账配置信息
    // extendInfoMap["split_conf_info"] = get2e79e0f8Dfa641aaAbe00cfc52d792d9()
    // 银联线上收银台
    // extendInfoMap["uni_app_payment_config"] = getB0f5629fD96f42779e96415a5e034b93()
    // 资金归集开通标记
    // extendInfoMap["fund_collection_flag"] = ""
    // 代发配置
    // extendInfoMap["surrogate_config_list"] = getF8203121Db1444f9A724632a461bb794()
    // 大额支付配置
    // extendInfoMap["large_amt_pay_config"] = get7994c6b214bb43269d9e471aad9c9619()
    // 统一收银台开关
    // extendInfoMap["half_pay_host_flag"] = ""
    // 代发复核配置
    // extendInfoMap["agent_recheck_config"] = get73b318df8691470f9dedD51fe0080064()
    // 商户开通网银充值开关
    // extendInfoMap["online_recharge_flag"] = ""
    // 是否开通垫资退款
    // extendInfoMap["refund_mend_open_flag"] = ""
    // 捷行付配置
    // extendInfoMap["jxf_conf_info"] = get95e6f12344e7482cBc98Cea89e932cb4()
    // 合单支付配置信息
    // extendInfoMap["merge_pay_config_info"] = get7865978bDaa1425bA4de340294477601()
    // 是否开通账单支付
    // extendInfoMap["bill_payment_flag"] = ""
    // 交易应答返回推送系统号
    // extendInfoMap["recon_resp_sys_id"] = ""
    // 抖音配置对象对象
    // extendInfoMap["dy_conf_list"] = getF0c1086a526544e39b7125fb4b0b66d9()
    return extendInfoMap
}

func getFa55923fC2d6483fB40034b57533c679() string {
    dto := make(map[string]interface{})
    // *协议类型*
    dto["agreement_type"] = "0"
    // *挂网协议地址*挂网协议必填；**必须按示例值填写**，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/opps/api/prod/download_file/PaymentServiceAgreement.htm&lt;/font&gt;；
    // dto["agreement_url"] = "test"
    // 纸质协议开始日期
    dto["agree_begin_date"] = "20200325"
    // 纸质协议结束日期
    dto["agree_end_date"] = "20400325"
    // 电子协议签约短信发送标识
    // dto["message_send_type"] = ""
    // 电子协议异步通知地址
    // dto["agreement_async_return_url"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get4ac5fff288a34ac3Bdc5989b00d1edf1() string {
    dto := make(map[string]interface{})
    // 业务类型
    // dto["fee_type"] = ""
    // 银行编码
    // dto["bank_id"] = ""
    // 借贷标志
    // dto["dc_flag"] = ""
    // 手续费（固定/元）
    // dto["fix_amt"] = ""
    // 费率（%）
    // dto["fee_rate"] = ""
    // 银行名称
    // dto["bank_name"] = ""
    // 银行中文简称
    // dto["bank_short_chn"] = ""
    // 费率状态
    // dto["stat_flag"] = ""
    // 手续费最小值（元）
    // dto["fee_min_amt"] = ""
    // 代扣手续费计费模式
    // dto["fee_formula_type"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get60d6fcb103b140bcA1adE1696875a531() string {
    dto := make(map[string]interface{})
    // 联系人身份证号码
    // dto["id_card_number"] = "test"
    // 联系人姓名
    // dto["name"] = ""
    // 联系人手机号
    // dto["mobile"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get6fd865f540934e538994F6a3fb1e054e() string {
    dto := make(map[string]interface{})
    // 证件持有人类型
    // dto["legal_type"] = ""
    // 证件类型
    // dto["card_type"] = ""
    // 姓名
    // dto["person_name"] = ""
    // 证件号码
    // dto["card_no"] = ""
    // 证件生效时间
    // dto["effect_time"] = ""
    // 证件过期时间
    // dto["expire_time"] = ""
    // 证件正面照
    // dto["card_front_img"] = ""
    // 证件反面照
    // dto["card_back_img"] = ""
    // 授权函照片
    // dto["auth_letter_img"] = ""
    // 是否为受益人
    // dto["is_benefit_person"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getDb4bebebA9964e7487ca24611484d99b() string {
    dto := make(map[string]interface{})
    // 是否金融机构
    // dto["finance_institution_flag"] = ""
    // 金融机构类型
    // dto["finance_type"] = ""
    // 证书类型
    // dto["cert_type"] = ""
    // 小微经营类型
    // dto["micro_biz_type"] = ""
    // 特殊行业id
    // dto["special_category_id"] = ""
    // 联系人信息对象
    // dto["contact_person_info"] = get60d6fcb103b140bcA1adE1696875a531()
    // 法人身份信息
    // dto["legal_person_info"] = get6fd865f540934e538994F6a3fb1e054e()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get535af30dE2f14ad69841E9fcb7a95716() string {
    dto := make(map[string]interface{})
    // *功能开关*
    // dto["switch_state"] = "test"
    // 手续费（%）
    dto["fee_rate"] = "3.15"
    // 最低收取手续费（元）
    // dto["fee_min_amt"] = ""
    // 商户经营类目
    dto["mcc"] = "2016062900190337"
    // 支付场景
    dto["pay_scene"] = "1"
    // 子渠道号
    dto["pay_channel_id"] = "10000001"
    // 是否需要实名认证
    // dto["is_check_real_name"] = ""
    // 实名认证信息
    // dto["al_real_name_info"] = getDb4bebebA9964e7487ca24611484d99b()

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getDb6f633eF7be4b7d8a37485550d8e27d() string {
    dto := make(map[string]interface{})
    // 联系人姓名联系人类型contact_type&#x3D;SUPER时必填。示例值：张三
    // dto["name"] = "test"
    // 联系人手机号联系人类型contact_type&#x3D;SUPER时必填。示例值：13917364538
    // dto["mobile"] = "test"
    // 联系人证件类型联系人类型contact_type&#x3D;SUPER时必填。&lt;br/&gt;枚举值参见《自然人证件类型》说明，示例值：00&lt;br/&gt;个体户/企业/事业单位/社会组织：可选择任一证件类型，政府机关仅支持身份证类型。
    // dto["contact_id_doc_type"] = "test"
    // 联系人证件号码联系人类型contact_type&#x3D;SUPER时必填。示例值：320936198512035017
    // dto["id_card_number"] = "test"
    // 联系人证件有效期开始时间联系人类型contact_type&#x3D;SUPER时必填&lt;br/&gt;格式：yyyy-MM-dd；示例值：2019-06-06
    // dto["contact_period_begin"] = "test"
    // 联系人证件有效期结束时间联系人类型contact_type&#x3D;SUPER时必填&lt;br/&gt;格式：yyyy-MM-dd；示例值：2029-06-06&lt;br/&gt;结束时间大于开始时间;若证件有效期为长期，请填写：长期
    // dto["contact_period_end"] = "test"
    // 是否金融机构
    // dto["finance_institution_flag"] = ""
    // 金融机构类型
    // dto["finance_type"] = ""
    // 证书类型
    // dto["cert_type"] = ""
    // 小微经营类型
    // dto["micro_biz_type"] = ""
    // 特殊行业id
    // dto["special_category_id"] = ""
    // 联系人类型
    // dto["contact_type"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get36bf1c2f0bf24f8a8fd39675ed1b021f() string {
    dto := make(map[string]interface{})
    // 开关状态
    // dto["switch_state"] = "test"
    // 手续费（%）
    dto["fee_rate"] = "2.15"
    // 最低收取手续费（元）
    // dto["fee_min_amt"] = ""
    // 支付场景
    dto["pay_scene"] = "10"
    // 费率规则ID
    dto["fee_rule_id"] = "765"
    // 子渠道号
    dto["pay_channel_id"] = "JP00001"
    // 申请服务
    dto["service_codes"] = ""
    // 是否需要实名认证
    // dto["is_check_real_name"] = ""
    // 实名认证信息
    // dto["wx_real_name_info"] = getDb6f633eF7be4b7d8a37485550d8e27d()

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get9ea086279d6a4d0d9330E6e40b55bddf() string {
    dto := make(map[string]interface{})
    // 银行业务手续费类型
    dto["charge_cate_code"] = ""
    // 功能开关
    // dto["switch_state"] = ""
    // 借记卡手续费（%）
    // dto["debit_fee_rate"] = ""
    // 贷记卡手续费1000以上（%）
    dto["credit_fee_rate_up"] = "6.566"
    // 贷记卡手续费1000及以下（%）
    dto["credit_fee_rate_down"] = "1"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get790f276d862e48d8B01fE2e690e0e250() string {
    dto := make(map[string]interface{})
    // 借记卡手续费（%）
    dto["debit_fee_rate"] = "3"
    // 贷记卡手续费（%）
    dto["credit_fee_rate"] = "6"
    // 银行业务手续费类型
    dto["charge_cate_code"] = "02"
    // 借记卡封顶值
    dto["debit_fee_limit"] = "5"
    // 是否开通小额双免
    dto["is_open_small_flag"] = "0"
    // 小额双免单笔限额(元)
    dto["small_free_amt"] = "500"
    // 小额双免手续费（%）
    dto["small_fee_amt"] = "1"
    // 功能开关
    // dto["switch_state"] = ""
    // 云闪付开通状态
    // dto["is_open_cloud_flag"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getA53c8a3a9b0649cbB3a85f572616c48e() string {
    dto := make(map[string]interface{})
    // *业务模式说明*请详细说明商户的业务模式，包括不限于（业务参与方、各方的权利及义务关系等）。最大500个汉字，必填&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：商户向企业采购产品，并将其展示在其电商平台上，用户注册后购买其产品，产品主要有家电、服饰、鞋帽等，用户购买产品的资金作为商户所得，商户根据订单情况进行物流配送。&lt;/font&gt;
    // dto["busi_instruction"] = "test"
    // *资金流向说明*请详细说明商户的资金流向，包括但不限于（交易所涉及的各方账户、资金流入流出过程等）。最大500个汉字，必填&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：用户选购产品后根据订单支付给商户，商户根据结算周期，结算到对公账户。&lt;/font&gt;
    // dto["capital_instruction"] = "test"
    // *功能开通用途说明*请详述所申请功能的原因、场景、涉及对象类型。最大500个汉字，必填&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：企业工资代发&lt;/font&gt;
    // dto["function_instruction"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getFacd221974d440bf8a9eB1984f5943b1() string {
    dto := make(map[string]interface{})
    // *业务模式*
    // dto["balance_model"] = "test"
    // *业务情况说明*
    // dto["description_info"] = getA53c8a3a9b0649cbB3a85f572616c48e()
    // 支付手续费(%)
    dto["fee_rate"] = "2"
    // 支付固定手续费(元)
    dto["fee_fix_amt"] = "1"
    // 功能开关
    dto["switch_state"] = "1"
    // 交易手续费外扣时的账户类型
    // dto["out_fee_acct_type"] = ""
    // 交易手续费外扣汇付ID
    // dto["out_fee_huifuid"] = ""
    // 是否交易手续费外扣
    // dto["out_fee_flag"] = ""
    // 扣费模式
    // dto["charge_mode"] = ""
    // 余额支付模式
    // dto["pay_mode"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get6725a4189ab7447a8ebfE033dc3cc458() string {
    dto := make(map[string]interface{})
    // 功能开关
    dto["switch_state"] = "0"
    // 是否交易手续费外扣
    // dto["out_fee_flag"] = "test"
    // 交易手续费外扣汇付ID
    // dto["out_fee_huifuid"] = "test"
    // 支付手续费(%)
    dto["fee_rate"] = "10"
    // 支付固定手续费(元)
    dto["fee_fix_amt"] = "5"
    // 交易手续费外扣时的账户类型
    // dto["out_fee_acct_type"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get5379d0a89c6e4c89A02926c6b05b5b98() string {
    dto := make(map[string]interface{})
    // 业务类型
    // dto["pay_type"] = ""
    // 交易手续费外扣时的账户类型
    // dto["out_fee_acct_type"] = ""
    // 交易手续费外扣汇付ID
    // dto["out_fee_huifuid"] = ""
    // 是否交易手续费外扣
    // dto["out_fee_flag"] = ""
    // 功能开关
    // dto["switch_state"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getD32d25ce68b146c095750e13d507dfe1() string {
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

func getE3698fa9963041338df62247ec2e7de3() string {
    dto := make(map[string]interface{})
    // 开户手续费(元)
    // dto["fee_fix_amt"] = "test"
    // 开户手续费外扣时的账户类型01-基本户，02-现金户，05-充值户，09-营销户，不填默认01；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：01&lt;/font&gt;&lt;br/&gt;注：fee_fix_amt：开户手续费大于0时必填
    // dto["out_fee_acct_type"] = "test"
    // 开户手续费外扣汇付ID开通手续费外扣业务时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000109812123&lt;/font&gt;&lt;br/&gt;注：fee_fix_amt：开户手续费大于0时必填
    // dto["out_fee_huifuid"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getCababba719dd457083e99d18c38532bb() string {
    dto := make(map[string]interface{})
    // 功能开关
    // dto["switch_state"] = "test"
    // 自动入账开关
    // dto["out_order_auto_acct_flag"] = "test"
    // 全渠道资金管理补充材料涉及文件类型：F504-全渠道资金管理补充材料；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e529&lt;/font&gt;
    // dto["other_payment_institutions_pic"] = "test"
    // 支付手续费(%)
    // dto["fee_rate"] = ""
    // 手续费最小值(元)
    // dto["fee_min_amt"] = ""
    // 交易手续费外扣时的账户类型
    // dto["out_fee_acct_type"] = ""
    // 交易手续费外扣标记
    // dto["out_fee_flag"] = ""
    // 交易手续费外扣汇付ID
    // dto["out_fee_huifuid"] = ""
    // 全域资金开户使用的银行卡信息
    // dto["out_order_acct_card"] = getD32d25ce68b146c095750e13d507dfe1()
    // 全域资金开户手续费
    // dto["out_order_acct_open_fees"] = getE3698fa9963041338df62247ec2e7de3()
    // 全域支付业务模式
    // dto["business_model"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getB25246ad3b9447198d0a4fac0ab10450() string {
    dto := make(map[string]interface{})
    // *运营媒介类型*
    // dto["media_type"] = "test"
    // *媒介名称*PC网站域名／APP名称／小程序名称／公众号名称；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：汇付服务&lt;/font&gt;&lt;br/&gt;运营媒介类型为 ：S1/S2/S3/S4时，必填；
    // dto["media_name"] = "test"
    // ICP备案/许可证号运营媒介类型为 ：S1时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：沪ICP备06046402号-28&lt;/font&gt;
    // dto["mer_icp"] = "test"
    // *其他有效信息*其他有效信息或链接地址/APP下载地址；类型为S5或S2必填&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：http://download.huifu.com&lt;/font&gt;
    // dto["other_info"] = "test"
    // *媒介主体与商户主体是否一致*Y/N；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：Y&lt;/font&gt;&lt;br/&gt;运营媒介为S1、S2、S3、S4且企业商户开通快捷或网银，或大额转账，或余额支付或分账业务（20%（不含）-100%），或为个人商户开通分账业务（10%（不含）-100%），必填&lt;br/&gt;若不一致，则需提供ICP备案/APP/微信公众号/小程序主体与商户的使用授权或开发证明材料；
    // dto["media_mer_common_flag"] = "test"
    // *授权或开发证明材料*运营媒介为S1、S2、S3、S4且媒介主体与商户主体不一致时，且企业商户开通快捷或网银，或大额转账，或余额支付或分账业务（20%（不含）-100%），或为个人商户开通分账业务（10%（不含）-100%），必填&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
    // dto["authorize_materials"] = "test"
    // 微信APP补充材料运营媒介为S2且开通微信下app支付时选填，具体见[图片上传接口](https://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc)
    // dto["other_material"] = "test"
    // appId运营媒介为S2且开通微信下app支付时必填
    // dto["app_id"] = "test"
    // appId认证主体名称运营媒介为S2且开通微信下app支付时必填
    // dto["app_name"] = "test"
    // 补充说明运营媒介为S2且开通微信下app支付时选填
    // dto["supplement"] = "test"
    // *测试账号*
    // dto["test_account"] = ""
    // *测试密码*
    // dto["test_secret"] = ""
    // *运营媒介-首页*
    // dto["media_front_page"] = ""
    // *运营媒介-服务/商品明细页面*
    // dto["media_service_page"] = ""
    // *运营媒介-下单场景页面*
    // dto["media_order_page"] = ""
    // *运营媒介-支付页面*
    // dto["media_pay_page"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get61966f16B1eb4ca193e7F3bbdb034e6b() string {
    dto := make(map[string]interface{})
    // *文件id*
    // dto["file_id"] = "test"
    // *文件类型*
    // dto["file_type"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getA7b56f0b78b74c61AbecC4ed3ea15bf6() string {
    dto := make(map[string]interface{})
    // 签约人类型
    // dto["type"] = "test"
    // 姓名签约人类型&#x3D;其他，必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：张三&lt;/font&gt;
    // dto["name"] = "test"
    // 手机号签约人类型&#x3D;法人/其他 ，必填；注意：**签约人会做姓名+身份证+手机号验证，请正确填写**；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：13917463536&lt;/font&gt;
    // dto["mobile_no"] = "test"
    // 身份证签约人类型&#x3D;联系人/其他，必填 ；注意：**签约人会做姓名+身份证+手机号验证，请正确填写**；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：320946195712025082&lt;/font&gt;
    // dto["cert_no"] = "test"
    // 邮箱
    // dto["email"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get4536af781cda4cf182fdF43699b63451() string {
    dto := make(map[string]interface{})
    // 文件类型
    // dto["file_type"] = "test"
    // 文件jfileId
    // dto["file_id"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getD26a78b679344baaAf21Bbbf3384cac4() string {
    dto := make(map[string]interface{})
    // 联系人类型
    // dto["contact_type"] = "test"
    // 联系人姓名联系人类型为经办人时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：张三&lt;/font&gt;
    // dto["contact_name"] = "test"
    // 联系人证件类型联系人类型为经办人时必填；00：身份证01:护照11：港澳台同胞通行证12：外国人居留证13：港澳居民证14：台湾居民证&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：00&lt;/font&gt;；
    // dto["cert_type"] = "test"
    // 联系人证件号码联系人类型为经办人时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：32090297512026402&lt;/font&gt;
    // dto["contact_cert_no"] = "test"
    // 联系人手机号
    // dto["contact_mobile_no"] = "test"
    // 联系人电子邮箱
    // dto["contact_email"] = "test"
    // 联系人资料联系人类型为经办人时必填F28-联系人身份证国徽面   F29-联系人身份证人像面F227-微信业务办理授权函&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：&lt;/font&gt;
    // dto["contact_file_list"] = get4536af781cda4cf182fdF43699b63451()
    // 证件有效期类型
    // dto["contact_cert_validity_type"] = "test"
    // 证件有效期开始日期
    // dto["contact_cert_begin_date"] = "test"
    // 证件有效期截止日期
    // dto["contact_cert_end_date"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getFf71b3fa4dbd42dcB2a10ab5478cb4f1() string {
    dto := make(map[string]interface{})
    // 文件类型
    // dto["file_type"] = "test"
    // 文件jfileId
    // dto["file_id"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getF787aec3F8124b4b8a3e211f9250fabb() string {
    dto := make(map[string]interface{})
    // 登记证书类型
    // dto["cert_type"] = "test"
    // 证书号
    // dto["cert_no"] = "test"
    // 证书商户名称
    // dto["cert_mer_name"] = "test"
    // 注册地址
    // dto["reg_detail"] = "test"
    // 法人姓名
    // dto["legal_name"] = "test"
    // 证书有效期类型
    // dto["cert_validity_type"] = "test"
    // 证书有效期开始日期
    // dto["cert_begin_date"] = "test"
    // 文件列表
    // dto["cert_file_list"] = getFf71b3fa4dbd42dcB2a10ab5478cb4f1()
    // 证书有效期截止日期
    // dto["cert_end_date"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get473b5001B1b248eeA5a02988059f4f6d() string {
    dto := make(map[string]interface{})
    // 文件类型
    // dto["file_type"] = "test"
    // 文件jfileId
    // dto["file_id"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getAd8fa475C574421d9368A1ba9ec02aad() string {
    dto := make(map[string]interface{})
    // 证件类型
    // dto["ubo_cert_type"] = "test"
    // 证件号码
    // dto["ubo_cert_no"] = "test"
    // 姓名
    // dto["ubo_name"] = "test"
    // 受益人证件居住地址
    // dto["ubo_cert_doc_address"] = "test"
    // 证件有效类型
    // dto["ubo_cert_validity_type"] = "test"
    // 证件有效期开始日期
    // dto["ubo_cert_begin_date"] = "test"
    // 文件列表
    // dto["ubo_file_list"] = get473b5001B1b248eeA5a02988059f4f6d()
    // 证件有效期截止日期
    // dto["ubo_cert_end_date"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get1155723d35564e5b971885f7aa10af3f() string {
    dto := make(map[string]interface{})
    // 经营者/法人是否为受益人
    // dto["ubo_type"] = "test"
    // 受益人信息列表jsonArray格式,当ubo_type为Y时可不填
    // dto["ubo_info_list"] = getAd8fa475C574421d9368A1ba9ec02aad()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getD38027409bc845a7B5edE086643479b2() string {
    dto := make(map[string]interface{})
    // 文件类型
    // dto["file_type"] = "test"
    // 文件jfileId
    // dto["file_id"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get896e37dbAc084215B709B98e07d96276() string {
    dto := make(map[string]interface{})
    // 文件类型
    // dto["file_type"] = "test"
    // 文件jfileId
    // dto["file_id"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get916fe3547b8747ecAceb6ffed4051636() string {
    dto := make(map[string]interface{})
    // 经营场景类型
    // dto["sales_scenes_type"] = "test"
    // 功能费率仅支持渠道商传入该字段。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.06&lt;/font&gt;平台商户为子商户开通微信直连支付时，不支持该字段，取平台商户费率上送微信。
    // dto["fee_rate"] = "test"
    // 线下场所对应的商家公众号APPID开通线下门店场景时，填入。都填入时，取公众号的上送微信。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wx51aa91a575359ff5&lt;/font&gt;
    // dto["biz_sub_jsapi_app_id"] = "test"
    // 线下场所对应的商家小程序APPID开通线下门店场景时，填入。都填入时，取公众号的上送微信。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wxea9c30a90fs8d3fe&lt;/font&gt;
    // dto["biz_sub_mini_app_id"] = "test"
    // 服务商公众号 ID开通公众号场景时，直连服务商和商户的公众号 APP ID，二选一填入。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wx51aa91a575359ff5&lt;/font&gt;
    // dto["jsapi_app_id"] = "test"
    // 商家公众号APPID开通公众号场景时，直连服务商和商户的公众号 APP ID，二选一填入。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wxea9c30a90fs8d3fe&lt;/font&gt;
    // dto["jsapi_sub_app_id"] = "test"
    // 服务商小程序APPID开通小程序场景时，直连服务商和商户的小程序 APP ID，二选一填入。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wx51aa91a575359ff5&lt;/font&gt;
    // dto["mini_app_id"] = "test"
    // 商家小程序APPID开通小程序场景时，直连服务商和商户的小程序 APP ID，二选一填入。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wx51aa91a575359ff5&lt;/font&gt;
    // dto["mini_sub_app_id"] = "test"
    // 服务商应用APPID开通 APP 场景时，直连服务商和商户的 APP ID，二选一填入。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wx51aa91a575359ff5&lt;/font&gt;
    // dto["app_app_id"] = "test"
    // 商家应用APPID开通 APP 场景时，直连服务商和商户的 APP ID，二选一填入。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wx51aa91a575359ff5&lt;/font&gt;
    // dto["app_sub_app_id"] = "test"
    // 互联网网站域名开通互联网场景时填入；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：huifu.com&lt;/font&gt;
    // dto["web_domain"] = "test"
    // 互联网网站对应的商家APPID开通互联网场景时填入；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wx51aa91a575359ff5&lt;/font&gt;
    // dto["web_app_id"] = "test"
    // 商家企业微信CorpID开通企业微信场景时填入；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：xxxxx&lt;/font&gt;
    // dto["sub_corp_id"] = "test"
    // 文件列表
    // dto["sales_scenes_file_list"] = get896e37dbAc084215B709B98e07d96276()
    // 文件列表
    // dto["fee_rate_file_list"] = getD38027409bc845a7B5edE086643479b2()
    // 门店名称
    // dto["biz_store_name"] = ""
    // 门店省市编码
    // dto["biz_address_code"] = ""
    // 门店地址
    // dto["biz_store_address"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get5e98d3c918ac479eAa1eF571fe37fe9e() string {
    dto := make(map[string]interface{})
    // 账户类型
    // dto["card_type"] = "test"
    // 开户名称
    // dto["card_name"] = "test"
    // 开户银行
    // dto["bank_code"] = "test"
    // 开户银行省编码
    // dto["prov_id"] = "test"
    // 开户银行市编码
    // dto["area_id"] = "test"
    // 开户银行联行号开户银行联行号与开户银行全称（含支行)二选一；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：102290026507&lt;/font&gt;
    // dto["branch_code"] = "test"
    // 开户银行全称（含支行)开户银行联行号与开户银行全称（含支行)二选一；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：中国工商银行股份有限公司上海市中山北路支行&lt;/font&gt;
    // dto["branch_name"] = "test"
    // 银行账号
    // dto["card_no"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getDc06c574Fc0c466c9b389d3d465f39d2() string {
    dto := make(map[string]interface{})
    // 文件类型
    // dto["file_type"] = "test"
    // 文件jfileId
    // dto["file_id"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getBe178278D11a4c44A764B5225abdf5f9() string {
    dto := make(map[string]interface{})
    // 申请服务
    // dto["service_code"] = "test"
    // 功能服务appid
    // dto["sub_app_id"] = "test"
    // 功能开关
    // dto["switch_state"] = "test"
    // 功能费率(%)
    // dto["fee_rate"] = "test"
    // 操作类型ADD-新增， UPDATE-修改， 默认新增；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：ADD&lt;/font&gt;
    // dto["operate_type"] = "test"
    // 联系人信息jsonObject字符串,新增时必填
    // dto["contact_info"] = getD26a78b679344baaAf21Bbbf3384cac4()
    // 特殊主体登记证书jsonObject字符串，商户营业执照类型为政府机关/事业单位/其他组织时，传入相应信息。新增时需填入
    // dto["certificate_info"] = getF787aec3F8124b4b8a3e211f9250fabb()
    // 最终受益人信息jsonObject字符串，商户类型为企业时，微信侧必填。（如果基本信息里有的话，可以不传取 huifu_id 对应的信息）。新增时填入
    // dto["ubo_info"] = get1155723d35564e5b971885f7aa10af3f()
    // 经营场景jsonObject字符串，新增时填入
    // dto["sales_info"] = get916fe3547b8747ecAceb6ffed4051636()
    // 银行账户信息jsonObject字符串，该字段不填时，取商户在汇付系统录入的结算账号信息。新增或修改时填入，修改时必填
    // dto["wx_card_info"] = get5e98d3c918ac479eAa1eF571fe37fe9e()
    // 补充说明信息
    // dto["business_addition_msg"] = ""
    // 补充说明文件列表
    // dto["addition_file_list"] = getDc06c574Fc0c466c9b389d3d465f39d2()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getBc3c15758a314d418122291dda43bf2d() string {
    dto := make(map[string]interface{})
    // 微信子商户号
    // dto["sub_mch_id"] = "test"
    // 配置集合
    // dto["wx_zl_pay_conf_list"] = getBe178278D11a4c44A764B5225abdf5f9()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getB4f4917b30434068B9c4Ddfbc19df363() string {
    dto := make(map[string]interface{})
    // 文件类型
    // dto["file_type"] = "test"
    // 文件jfileId
    // dto["file_id"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getFff133434e6a4cbd9e49C33a11c2da87() string {
    dto := make(map[string]interface{})
    // 申请类型
    // dto["apply_type"] = "test"
    // 商户支付宝账号
    // dto["account"] = "test"
    // 服务费率仅支持渠道商。平台商户调用不支持该字段服务费率（%），0.38~3之间，精确到0.01。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.06&lt;/font&gt;
    // dto["fee_rate"] = "test"
    // 文件列表
    // dto["file_list"] = getB4f4917b30434068B9c4Ddfbc19df363()
    // 联系人姓名
    // dto["contact_name"] = ""
    // 联系人手机号
    // dto["contact_mobile_no"] = ""
    // 联系人电子邮箱
    // dto["contact_email"] = ""
    // 订单授权凭证
    // dto["order_ticket"] = ""
    // 营业执照编号
    // dto["license_code"] = ""
    // 营业执照有效期类型
    // dto["license_validity_type"] = ""
    // 营业执照有效期开始日期
    // dto["license_begin_date"] = ""
    // 营业执照有效期截止日期
    // dto["license_end_date"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get997e98a3112147bfA72341fdf8a69322() interface{} {
    dto := make(map[string]interface{})
    // 分账比例
    // dto["fee_rate"] = "test"
    // 汇付Id
    // dto["huifu_id"] = "test"

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getC333fc13843a4f748041541d071720a3() interface{} {
    dto := make(map[string]interface{})
    // *业务模式说明*
    // dto["busi_instruction"] = "test"
    // *资金流向说明*
    // dto["capital_instruction"] = "test"
    // *功能开通用途说明*
    // dto["function_instruction"] = "test"

    return dto;
}

func get2e79e0f8Dfa641aaAbe00cfc52d792d9() string {
    dto := make(map[string]interface{})
    // 分账规则来源
    // dto["rule_origin"] = "test"
    // 分账开关
    // dto["div_flag"] = "test"
    // 最大分账比例%
    // dto["apply_ratio"] = "test"
    // 分账模式
    // dto["scene"] = "test"
    // 分账明细
    // dto["acct_split_bunch_list"] = get997e98a3112147bfA72341fdf8a69322()
    // 生效类型
    // dto["start_type"] = ""
    // 交易手续费外扣开关
    // dto["out_fee_flag"] = ""
    // 交易手续费外扣时的账户类型
    // dto["out_fee_acct_type"] = ""
    // 交易手续费外扣汇付ID
    // dto["out_fee_huifuid"] = ""
    // 手续费%
    // dto["split_fee_rate"] = ""
    // 固定手续费
    // dto["per_amt"] = ""
    // 业务情况说明
    // dto["split_ext_info"] = getC333fc13843a4f748041541d071720a3()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getB0f5629fD96f42779e96415a5e034b93() string {
    dto := make(map[string]interface{})
    // 借记手续费（%）借记卡费率与贷记卡费率不能同时为空;保留2位小数，最大值100.00，最小值0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.02&lt;/font&gt;
    // dto["debit_fee_rate"] = "test"
    // 贷记手续费（%）借记卡费率与贷记卡费率不能同时为空;保留2位小数，最大值100.00，最小值0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.06&lt;/font&gt;
    // dto["credit_fee_rate"] = "test"
    // 状态开关
    // dto["switch_state"] = ""
    // 是否交易手续费外扣
    // dto["out_fee_flag"] = ""
    // 交易手续费外扣汇付ID
    // dto["out_fee_huifuid"] = ""
    // 交易手续费外扣时的账户类型
    // dto["out_fee_acct_type"] = ""
    // 云闪付免密支付开通标识
    // dto["cloud_quick_pass_secret_free_flag"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getF8203121Db1444f9A724632a461bb794() string {
    dto := make(map[string]interface{})
    // 代发业务类型
    // dto["surrogate_type"] = "test"
    // 手续费（固定/元）手续费（固定/元），保留小数点后两位；fee_formula_type为05，06时必填 &lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.00&lt;/font&gt;
    // dto["fix_amt"] = "test"
    // 手续费（百分比/%）手续费（百分比/%），保留小数点后两位；取值范围[0.00,100.00]；fee_formula_type为01，02，03，06时必填 &lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.05&lt;/font&gt;
    // dto["fee_rate"] = "test"
    // 代发手续费计费模式01: 百分比手续费，无封顶值或保底值&lt;br/&gt;02: 百分比手续费，有保底值&lt;br/&gt;03: 百分比手续费，有封顶值&lt;br/&gt;05: 固定手续费 &lt;br/&gt;06: 固定手续费+百分比手续费&lt;br/&gt;为空默认06 &lt;font color&#x3D;&quot;green&quot;&gt;示例值：06&lt;/font&gt;
    // dto["fee_formula_type"] = "test"
    // 手续费封顶值（固定/元）手续费封顶值（固定/元），保留小数点后两位；fee_formula_type为03时必填 &lt;font color&#x3D;&quot;green&quot;&gt;示例值：100.00&lt;/font&gt;
    // dto["fee_max_amt"] = "test"
    // 手续费保底值（固定/元）手续费保底值（固定/元），保留小数点后两位；fee_formula_type为02时必填 &lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.00&lt;/font&gt;
    // dto["fee_min_amt"] = "test"
    // 开通状态
    // dto["switch_state"] = ""
    // 是否交手续费外扣标记
    // dto["out_fee_flag"] = ""
    // 交易手续费外扣时账户类型
    // dto["out_fee_acct_type"] = ""
    // 交易手续费外扣汇付ID
    // dto["out_fee_huifu_id"] = ""
    // 是否允许对私代发
    // dto["surrogate_private_flag"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get9f8c46555bc444deB8729eb31b805596() string {
    dto := make(map[string]interface{})
    // 大额支付业务模式
    // dto["business_model"] = "test"
    // 费率（%）开通大额业务时必须填写一种收费方式；大于0,保留2位小数；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.00&lt;/font&gt;
    // dto["fee_rate"] = "test"
    // 交易手续费（固定/元）开通大额业务时必须填写一种收费方式；大于0,保留2位小数；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：10.00&lt;/font&gt;
    // dto["fee_fix_amt"] = "test"
    // 功能开关
    // dto["switch_state"] = ""
    // 大额调账标识申请类型
    // dto["biz_type"] = ""
    // 是否允许绑卡支付
    // dto["mer_same_card_recharge_flag"] = ""
    // 是否允许用户入账
    // dto["allow_user_deposit_flag"] = ""
    // 备付金固定账号模式自动退款
    // dto["provisions_auto_refund_flag"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get7994c6b214bb43269d9e471aad9c9619() string {
    dto := make(map[string]interface{})
    // 大额支付配置列表
    // dto["large_amt_pay_config_info_list"] = get9f8c46555bc444deB8729eb31b805596()
    // 交易手续费外扣huifuId交易手续费外扣时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000108854952&lt;/font&gt;
    // dto["out_fee_huifu_id"] = "test"
    // 交易手续费外扣账户号交易手续费外扣时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：F00598602&lt;/font&gt;
    // dto["out_fee_acct_id"] = "test"
    // 交易手续费外扣标记
    // dto["out_fee_flag"] = ""
    // 商户付款方卡类型
    // dto["mer_payer_card_type"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get73b318df8691470f9dedD51fe0080064() string {
    dto := make(map[string]interface{})
    // 代发复核开关
    // dto["agent_recheck_flag"] = ""
    // 复核授权商户号
    // dto["agent_recheck_huifu_id"] = ""
    // 复核类型
    // dto["agent_recheck_type"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get5e0fb8ec802a4dc4B228171b958b061b() string {
    dto := make(map[string]interface{})
    // 操作人姓名银行名称非平安银行时，需填写；示例：张三
    // dto["user_name"] = "test"
    // 操作人手机号银行名称非平安银行时，需填写；示例：13333333333
    // dto["operator_phone"] = "test"
    // 付款方账号
    // dto["acct_no"] = "test"
    // 付款方账户名称
    // dto["acct_name"] = "test"
    // 开户行联行号
    // dto["bank_branch_code"] = "test"
    // 开户行行名
    // dto["bank_branch_name"] = "test"
    // 签约申请生效日期银行名称非平安银行时，需填写；示例：20251219
    // dto["sign_start_date"] = "test"
    // 签约申请失效日期银行名称非平安银行时，需填写；示例：20261219
    // dto["sign_end_date"] = "test"
    // 机构编码
    // dto["org_code"] = "test"
    // 付款方账户类型
    // dto["acct_type"] = ""
    // 附言
    // dto["remark"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get30fd27c156684f318ea385ab700ddab2() interface{} {
    dto := make(map[string]interface{})
    // 捷行付手续费外扣汇付ID捷行付设置手续费时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000108854952&lt;/font&gt;
    // dto["out_fee_huifuid"] = "test"
    // 捷行付手续费扣款账户类型捷行付手续费外扣时必填；01-基本户(默认)，05-充值户；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：01&lt;/font&gt;
    // dto["out_fee_acct_type"] = "test"
    // 虚拟卡开卡手续费（元）
    // dto["open_card_fix_fee"] = ""
    // 虚拟卡额度发放手续费（%）
    // dto["quota_issue_fee_rate"] = ""
    // 捷行付手续费内外扣标记
    // dto["out_fee_flag"] = ""

    return dto;
}

func get95e6f12344e7482cBc98Cea89e932cb4() string {
    dto := make(map[string]interface{})
    // 账户类型
    // dto["acct_type"] = "test"
    // 开通状态
    // dto["open_flag"] = "test"
    // 银行名称
    // dto["bank_name"] = "test"
    // 银行账户信息当账户类型为银行账户时必填
    // dto["jxf_bank_info"] = get5e0fb8ec802a4dc4B228171b958b061b()
    // 付款确认请求URL付款确认标识为Y时，该字段必填
    // dto["pay_confirm_req_url"] = "test"
    // 手续费信息
    // dto["jxf_fee_info"] = get30fd27c156684f318ea385ab700ddab2()
    // 付款确认标识
    // dto["pay_confirm_flag"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get01d5830d69ee428e8fc8Cdfebe320e95() string {
    dto := make(map[string]interface{})
    // 合单支付模式
    // dto["merge_pay_mode"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get7865978bDaa1425bA4de340294477601() string {
    dto := make(map[string]interface{})
    // 合单支付开关
    // dto["merge_pay_flag"] = "test"
    // 合单支付配置合单支付配置 （当合单支付开关为开时必填）
    // dto["merge_pay_config"] = get01d5830d69ee428e8fc8Cdfebe320e95()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getF0c1086a526544e39b7125fb4b0b66d9() string {
    dto := make(map[string]interface{})
    // 手续费（%）
    // dto["fee_rate"] = ""
    // 最低收取手续费（元）
    // dto["fee_min_amt"] = ""
    // 是否交易手续费外扣
    // dto["out_fee_flag"] = ""
    // 交易手续费外扣汇付ID
    // dto["out_fee_huifuid"] = ""
    // 交易手续费外扣时的账户类型
    // dto["out_fee_acct_type"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

