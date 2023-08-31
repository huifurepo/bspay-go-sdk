/**
 * 商户统一进件接口(2022) - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantIntegrateRegRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantIntegrateRegRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantIntegrateRegRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 渠道商汇付id
        UpperHuifuId:"6666000105215341",
        // 公司类型
        EntType:"3",
        // 商户名称
        RegName:"天马微电子股份有限公司",
        // 经营类型
        BusiType:"1",
        // 经营详细地址
        DetailAddr:"深圳市宝安区新安街道海旺社区N26区海秀路2021号荣超滨海大厦A座2111",
        // 经营省
        ProvId:"310000",
        // 经营市
        AreaId:"310100",
        // 经营区
        DistrictId:"310104",
        // 联系人信息
        ContactInfo:getContactInfo(),
        // 卡信息配置实体
        CardInfo:getCardInfo(),
        // 取现配置列表jsonArray格式 ；
        CashConfig:getCashConfig(),
        // 结算配置jsonObject格式；
        SettleConfig:getSettleConfig(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantIntegrateRegRequest(dgReq)
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
    extendInfoMap["short_name"] = "姜雨"
    // 小票名称
    extendInfoMap["receipt_name"] = "汇付天下"
    // 商户英文名称
    extendInfoMap["mer_en_name"] = "huifu"
    // 所属行业
    extendInfoMap["mcc"] = "7273"
    // 营业执照类型
    extendInfoMap["license_type"] = "NATIONAL_LEGAL_MERGE"
    // 营业执照编号
    extendInfoMap["license_code"] = "914403001921834459"
    // 营业执照有效期类型
    extendInfoMap["license_validity_type"] = "0"
    // 营业执照有效期开始日期
    extendInfoMap["license_begin_date"] = "19831108"
    // 营业执照有效期截止日期
    extendInfoMap["license_end_date"] = "20380831"
    // 注册详细地址
    extendInfoMap["reg_detail"] = "深圳市宝安区新安街道海旺社区N26区海秀路2021号荣超滨海大厦A座2111"
    // 注册省
    extendInfoMap["reg_prov_id"] = "310000"
    // 注册市
    extendInfoMap["reg_area_id"] = "310100"
    // 注册区
    extendInfoMap["reg_district_id"] = "310104"
    // 客服电话
    extendInfoMap["service_phone"] = "15556622000"
    // 商户主页URL
    extendInfoMap["mer_url"] = "http://www.baidu.com"
    // 商户ICP备案编号
    extendInfoMap["mer_icp"] = "苏ICP备15042526号"
    // 开户许可证核准号
    extendInfoMap["open_licence_no"] = "123456789"
    // 法人信息
    extendInfoMap["legal_info"] = getLegalInfo()
    // 签约人
    // extendInfoMap["sign_user_info"] = getSignUserInfo()
    // 管理员账号
    extendInfoMap["login_name"] = "LG02022072707540497330158089012"
    // 是否通知商户标识
    extendInfoMap["sms_send_flag"] = "1"
    // 协议信息实体
    extendInfoMap["agreement_info"] = getAgreementInfo()
    // 业务开关配置
    extendInfoMap["biz_conf"] = getBizConf()
    // 微信配置对象
    extendInfoMap["wx_conf_list"] = getWxConfList()
    // 实名认证信息
    extendInfoMap["wx_realname_info"] = getWxRealnameInfo()
    // 支付宝配置对象
    extendInfoMap["ali_conf_list"] = getAliConfList()
    // 银联二维码配置
    extendInfoMap["union_conf_list"] = getUnionConfList()
    // 银联小微入驻信息实体
    extendInfoMap["union_micro_info"] = getUnionMicroInfo()
    // 余额支付配置实体
    extendInfoMap["balance_pay_config"] = getBalancePayConfig()
    // 银行卡业务配置实体
    extendInfoMap["bank_card_conf"] = getBankCardConf()
    // 花呗分期费率配置实体
    extendInfoMap["hb_fq_fee_config"] = getHbFqFeeConfig()
    // 补贴支付
    extendInfoMap["combine_pay_config"] = getCombinePayConfig()
    // 商户业务类型
    // extendInfoMap["mer_bus_type"] = ""
    // 线上费率配置
    // extendInfoMap["online_fee_conf_list"] = getOnlineFeeConfList()
    // 线上手续费承担方配置
    // extendInfoMap["online_pay_fee_conf_list"] = getOnlinePayFeeConfList()
    // 文件列表
    extendInfoMap["file_info"] = getFileInfo()
    // 异步消息接收地址
    extendInfoMap["async_return_url"] = "rmq://C_SSPM_YMFZ_AUDIT"
    // 业务开通结果异步消息接收地址
    extendInfoMap["busi_async_return_url"] = "virgo://http://192.168.85.157:30031/sspm/testVirgo"
    // 交易异步应答地址
    extendInfoMap["recon_resp_addr"] = "archer://C_SSPM_NSPOSM_BUSIRESULT"
    return extendInfoMap
}

func getLegalInfo() string {
    dto := make(map[string]interface{})
    // 法人姓名
    dto["legal_name"] = "张三"
    // 法人证件类型
    dto["legal_cert_type"] = "00"
    // 法人证件号码
    dto["legal_cert_no"] = "210411198701140000"
    // 法人证件有效期开始日期
    dto["legal_cert_begin_date"] = "20180201"
    // 法人证件有效期类型
    dto["legal_cert_validity_type"] = "0"
    // 法人证件有效期截止日期
    dto["legal_cert_end_date"] = "20380201"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getContactInfo() string {
    dto := make(map[string]interface{})
    // 联系人姓名
    dto["contact_name"] = "张三"
    // 联系人手机号
    dto["contact_mobile_no"] = "15657470000"
    // 联系人电子邮箱
    dto["contact_email"] = "jeff.peng@huifu.com"
    // 联系人身份证号
    // dto["contact_cert_no"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getSignUserInfo() string {
    dto := make(map[string]interface{})
    // 签约人类型
    // dto["type"] = "test"
    // 姓名
    // dto["sign_name"] = ""
    // 手机号
    // dto["sign_mobile_no"] = ""
    // 身份证
    // dto["sign_cert_no"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getCardInfo() string {
    dto := make(map[string]interface{})
    // 结算类型
    dto["card_type"] = "2"
    // 银行所在省参考银行省份编码；参考[地区码](https://paas.huifu.com/partners/api/#/csfl/api_csfl_dqbm)，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：310000&lt;/font&gt;；如修改省市要级联修改&lt;br/&gt;当card_type&#x3D;0时非必填， 当card_type&#x3D;1或2时必填
    dto["prov_id"] = "310000"
    // 银行所在市参考省市区编码；参考[地区码](https://paas.huifu.com/partners/api/#/csfl/api_csfl_dqbm)，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：310100 &lt;/font&gt;；如修改省市要级联修改&lt;br/&gt;当card_type&#x3D;0时非必填， 当card_type&#x3D;1或2时必填
    dto["area_id"] = "310100"
    // 结算账户名
    dto["card_name"] = "张华"
    // 结算账号
    dto["card_no"] = "621485121290000"
    // 银行编码
    dto["bank_code"] = "01050000"
    // 联行号
    dto["branch_code"] = "105290075067"
    // 支行名称
    dto["branch_name"] = "中国建设银行股份有限公司上海五角场支行"
    // 持卡人证件有效期类型
    dto["cert_validity_type"] = "1"
    // 持卡人证件有效期（起始）
    dto["cert_begin_date"] = "20210201"
    // 持卡人证件有效期（截止）
    dto["cert_end_date"] = ""
    // 持卡人证件号码
    dto["cert_no"] = "110101199305182000"
    // 持卡人证件类型
    dto["cert_type"] = "00"
    // 结算人手机号
    dto["mp"] = "1865590000"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getCashConfig() string {
    dto := make(map[string]interface{})
    // 取现手续费率（%）fix_amt与fee_rate至少填写一项，单位%，需保留小数点后两位，取值范围[0.00,100.00]，不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.05&lt;/font&gt;&lt;br/&gt;注：如果fix_amt与fee_rate都填写了则手续费&#x3D;fix_amt+支付金额*fee_rate
    dto["fee_rate"] = "2"
    // 业务类型
    dto["cash_type"] = "D1"
    // 提现手续费（固定/元）
    dto["fix_amt"] = "0.04"
    // 是否交易手续费外扣
    dto["out_fee_flag"] = "1"
    // 手续费承担方
    // dto["out_fee_huifu_id"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getSettleConfig() string {
    dto := make(map[string]interface{})
    // 结算周期
    dto["settle_cycle"] = "T1"
    // 节假日结算手续费率%
    dto["fixed_ratio"] = "3"
    // 起结金额
    dto["min_amt"] = "0.40"
    // 结算手续费外扣时的账户类型
    dto["out_settle_acct_type"] = "01"
    // 结算手续费外扣时的汇付ID
    dto["out_settle_huifuid"] = "6666000105215340"
    // 手续费外扣标记
    dto["out_settle_flag"] = "1"
    // 留存金额
    dto["remained_amt"] = "100"
    // 结算摘要
    dto["settle_abstract"] = "结算测试"
    // 结算方式
    dto["settle_pattern"] = "P0"
    // 结算批次号
    dto["settle_batch_no"] = "300"
    // 是否优先到账
    dto["is_priority_receipt"] = "N"
    // 自定义结算处理时间
    dto["settle_time"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getAgreementInfo() string {
    dto := make(map[string]interface{})
    // 协议类型
    dto["agreement_type"] = "1"
    // 协议号
    dto["agreement_no"] = "202010200100000203"
    // 协议模板号
    dto["agreement_model"] = "202010200100000203"
    // 协议模板名称
    dto["agreement_name"] = "电子协议签约模板"
    // 签约日期
    dto["sign_date"] = "20200325"
    // 协议开始日期
    dto["agree_begin_date"] = "20200325"
    // 协议结束日期
    dto["agree_end_date"] = "20400325"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getBizConf() string {
    dto := make(map[string]interface{})
    // 延迟入账开关
    dto["delay_flag"] = "Y"
    // 商户开通强制延迟标记
    dto["forced_delay_flag"] = "Y"
    // 是否开通网银
    dto["online_flag"] = "Y"
    // 是否开通快捷
    dto["quick_flag"] = "Y"
    // 是否开通代扣
    dto["withhold_flag"] = "Y"
    // 是否开通微信预授权
    dto["wechatpay_pre_auth_flag"] = "Y"
    // 是否开通支付宝预授权
    dto["alipay_pre_auth_flag"] = "Y"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getWxConfList() string {
    dto := make(map[string]interface{})
    // 支付场景
    dto["pay_scene"] = "1"
    // 手续费（%）
    dto["fee_rate"] = "0.38"
    // 费率规则id
    dto["fee_rule_id"] = "758"
    // 商户经营类目
    dto["mcc"] = "111"
    // 子渠道号
    dto["pay_channel_id"] = "JP00001"
    // 公众号支付Appid
    dto["wx_woa_app_id"] = "wx_woa_app_id"
    // 微信公众号授权目录
    dto["wx_woa_path"] = "wx_woa_path "
    // 微信小程序APPID
    dto["wx_applet_app_id"] = "wx_applet_app_id"
    // 微信公众号APPID对应的秘钥
    dto["wx_woa_secret"] = "wx_woa_secret"
    // 微信小程序APPID对应的秘钥
    dto["wx_applet_secret"] = "wx_applet_secret"
    // 申请服务 
    dto["service_codes"] = "['JSAPI','PAP']"
    // 交易手续费外扣时的账户类型
    dto["out_fee_acct_type"] = "01"
    // 交易手续费外扣汇付ID
    dto["out_fee_huifuid"] = "6666000108840829"
    // 是否交易手续费外扣
    dto["out_fee_flag"] = "1"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getUboInfo() string {
    dto := make(map[string]interface{})
    // 证件类型
    dto["ubo_id_doc_type"] = "00"
    // 证件正面照片
    dto["ubo_id_doc_copy"] = "c7faf0e6-39e9-3c35-9075-2312ad6f4ea4"
    // 证件姓名
    dto["ubo_id_doc_name"] = "王五"
    // 证件号码
    dto["ubo_id_doc_number"] = "110101199003072631"
    // 证件居住地址
    dto["ubo_id_doc_address"] = "上海市徐汇区宜山路789号"
    // 证件有效期开始时间
    dto["ubo_period_begin"] = "19900307"
    // 证件有效期结束时间
    dto["ubo_period_end"] = "20330909"
    // 证件反面照片
    dto["ubo_id_doc_copy_back"] = "c7faf0e6-39e9-3c35-9075-2312ad6f4ea4"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getQualificationInfoList() string {
    dto := make(map[string]interface{})
    // 行业类目id
    // dto["mcc_code"] = "test"
    // 行业经营许可证资质照片
    // dto["image_list"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getWxRealnameInfo() string {
    dto := make(map[string]interface{})
    // 支付场景
    dto["pay_scene"] = "1"
    // 联系人姓名
    dto["name"] = "詹振"
    // 联系人手机号
    dto["mobile"] = "15657470001"
    // 联系人证件号码
    dto["contact_id_card_number"] = "210411198701140000"
    // 实名认证类型
    // dto["realname_info_type"] = ""
    // 子渠道号
    dto["pay_channel_id"] = "JP00001"
    // 联系人类型
    dto["contact_type"] = "SUPER"
    // 联系人证件类型
    dto["contact_id_doc_type"] = "00"
    // 联系人证件有效期开始时间
    dto["contact_period_begin_date"] = "20210101"
    // 联系人证件有效期结束时间
    dto["contact_period_end_date"] = "20410101"
    // 机构证书类型
    dto["cert_type"] = "CERTIFICATE_TYPE_2389"
    // 机构证书编号
    dto["cert_number"] = "1234567892"
    // 经营者/法人是否为受益人
    dto["owner"] = "Y"
    // 法人证件居住地址
    dto["legal_identification_address"] = "上海祁连山路1号"
    // 小微经营类型
    dto["micro_biz_type"] = ""
    // 门店名称
    dto["store_name"] = ""
    // 门店省市编码
    dto["store_address_code"] = ""
    // 门店地址
    dto["store_address"] = ""
    // 特殊行业Id
    dto["category_id"] = ""
    // 是否金融机构
    dto["finance_institution_flag"] = "N"
    // 金融机构类型
    dto["finance_type"] = ""
    // 受益人信息
    dto["ubo_info_list"] = getUboInfo()
    // 经营许可证
    // dto["qualification_info_list"] = getQualificationInfoList()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getAliConfList() string {
    dto := make(map[string]interface{})
    // 支付场景
    dto["pay_scene"] = "1"
    // 手续费（%）
    dto["fee_rate"] = "0.38"
    // 商户经营类目
    dto["mcc"] = "2015063000020189"
    // 子渠道号
    dto["pay_channel_id"] = "JQF00001"
    // 拟申请的间联商户等级
    dto["indirect_level"] = "INDIRECT_LEVEL_M3"
    // 交易手续费外扣时的账户类型
    dto["out_fee_acct_type"] = "01"
    // 交易手续费外扣汇付ID
    dto["out_fee_huifuid"] = "6666000105215340"
    // 是否交易手续费外扣
    dto["out_fee_flag"] = "1"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getUnionConfList() string {
    dto := make(map[string]interface{})
    // 借记卡手续费1000以上（%）
    dto["debit_fee_rate_up"] = "0.55"
    // 银联二维码业务贷记卡手续费1000以上（%）
    dto["credit_fee_rate_up"] = "0.56"
    // 借记卡手续费1000以下（%）
    dto["debit_fee_rate_down"] = "0.38"
    // 银联二维码业务贷记卡手续费1000以下（%）
    dto["credit_fee_rate_down"] = "0.38"
    // 银联业务手续费类型
    dto["charge_cate_code"] = "03"
    // 借记卡封顶1000以上（元）
    dto["debit_fee_limit_up"] = "20"
    // 借记卡封顶1000以下（元）
    dto["debit_fee_limit_down"] = "10"
    // 商户经营类目
    dto["mcc"] = "5411"
    // 交易手续费外扣时的账户类型
    dto["out_fee_acct_type"] = "01"
    // 交易手续费外扣汇付ID
    dto["out_fee_huifuid"] = "6666000105215340"
    // 是否交易手续费外扣
    dto["out_fee_flag"] = "1"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getUnionMicroInfo() string {
    dto := make(map[string]interface{})
    // 银联商户类别
    dto["mchnt_type"] = "0101"
    // 商户经度
    dto["mer_lng"] = "121.472644"
    // 商户纬度
    dto["mer_lat"] = "31.231706"
    // 店铺名称
    dto["shop_name"] = "测试"
    // 商户经营类目
    dto["mcc"] = "5992"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getBalancePayConfig() string {
    dto := make(map[string]interface{})
    // 支付手续费(%)
    dto["fee_rate"] = "3"
    // 支付固定手续费(元)
    dto["fee_fix_amt"] = "0.80"
    // 交易手续费外扣时的账户类型
    dto["out_fee_acct_type"] = "01"
    // 交易手续费外扣汇付ID
    dto["out_fee_huifuid"] = "6666000105215340"
    // 是否交易手续费外扣
    dto["out_fee_flag"] = "1"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getBankCardConf() string {
    dto := make(map[string]interface{})
    // 借记卡手续费（%）
    dto["debit_fee_rate"] = "0.38"
    // 贷记卡手续费（%）
    dto["credit_fee_rate"] = "0.39"
    // 商户经营类目
    dto["mcc"] = "5411"
    // 银行业务手续费类型
    dto["charge_cate_code"] = "02"
    // 借记卡封顶值（元）
    dto["debit_fee_limit"] = "0.56"
    // 云闪付借记卡手续费1000以上（%）
    dto["cloud_debit_fee_rate_up"] = "0.56"
    // 云闪付借记卡封顶1000以上(元)
    dto["cloud_debit_fee_limit_up"] = "12"
    // 云闪付贷记卡手续费1000以上（%）
    dto["cloud_credit_fee_rate_up"] = "0.59"
    // 云闪付借记卡手续费1000以下（%）
    dto["cloud_debit_fee_rate_down"] = "0.37"
    // 云闪付借记卡封顶1000以下(元)
    dto["cloud_debit_fee_limit_down"] = "5"
    // 云闪付贷记卡手续费1000以下（%）
    dto["cloud_credit_fee_rate_down"] = "0.36"
    // 是否开通小额双免
    dto["is_open_small_flag"] = "0"
    // 小额双免单笔限额(元)
    dto["small_free_amt"] = "1000"
    // 小额双免手续费（%）
    dto["small_fee_amt"] = "0.33"
    // 交易手续费外扣时的账户类型
    dto["out_fee_acct_type"] = "01"
    // 交易手续费外扣汇付ID
    dto["out_fee_huifuid"] = "6666000105215340"
    // 是否交易手续费外扣
    dto["out_fee_flag"] = "1"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getHbFqFeeConfig() string {
    dto := make(map[string]interface{})
    // 花呗收单分期3期（%）分期费率不为空时，收单费率必填，大于0，保留2位小数，不小于渠道商成本；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.0&lt;/font&gt;代表费率为1.00%
    dto["acq_three_period"] = "1.30"
    // 花呗收单分期6期（%）分期费率不为空时，收单费率必填，大于0，保留2位小数，不小于渠道商成本；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.0&lt;/font&gt;代表费率为1.00%
    dto["acq_six_period"] = "4.60"
    // 花呗收单分期12期（%）分期费率不为空时，收单费率必填，大于0，保留2位小数，不小于渠道商成本；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.0&lt;/font&gt;代表费率为1.00%
    dto["acq_twelve_period"] = "9.12"
    // 花呗分期3期（%）
    dto["three_period"] = "1.80"
    // 花呗分期6期（%）
    dto["six_period"] = "4.60"
    // 花呗分期12期（%）
    dto["twelve_period"] = "9.12"
    // 商户经营类目
    dto["ali_mcc"] = "2015063000020189"
    // 支付场景
    dto["pay_scene"] = "1"
    // 交易手续费外扣时的账户类型
    dto["out_fee_acct_type"] = "01"
    // 交易手续费外扣汇付ID
    dto["out_fee_huifuid"] = "6666000105215340"
    // 是否交易手续费外扣
    dto["out_fee_flag"] = "1"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getCombinePayConfig() string {
    dto := make(map[string]interface{})
    // 支付手续费(%)
    dto["fee_rate"] = "5"
    // 支付固定手续费(元)
    dto["fee_fix_amt"] = "0.56"
    // 交易手续费外扣时的账户类型
    dto["out_fee_acct_type"] = "01"
    // 交易手续费外扣汇付ID
    dto["out_fee_huifuid"] = "6666000105215340"
    // 是否交易手续费外扣
    dto["out_fee_flag"] = "1"

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

func getFileInfo() string {
    dto := make(map[string]interface{})
    // D1结算协议图片
    dto["settle_agree_pic"] = "测试2022062910491030461"
    // 授权委托书
    dto["auth_enturst_pic"] = "0215232e-b595-368e-8a68-8c15b04f875b"
    // 商务协议
    dto["ba_pic"] = "d1451277-85c6-3177-ac3d-a8be47b9ae9d"
    // 公司照片一
    dto["store_header_pic"] = "d1451277-85c6-3177-ac3d-a8be47b9ae9d"
    // 公司照片二
    dto["store_indoor_pic"] = "009cd33c-01be-31f0-8e8c-615460949b96"
    // 公司照片三
    dto["store_cashier_desk_pic"] = "2020022204231214021970311"
    // 法人身份证反面
    dto["legal_cert_back_pic"] = "2020022204231214021970321"
    // 法人身份证正面
    dto["legal_cert_front_pic"] = "2022071804231214021970321"
    // 营业执照图片
    dto["license_pic"] = "36ac0355-a54d-3178-b4b5-9aecd07367e6"
    // 组织机构代码证
    dto["org_code_pic"] = "5bd7fea7-c8e4-31fc-a672-755adbcd4a4c"
    // 开户许可证
    dto["reg_acct_pic"] = "d1d50615-0ff4-3488-b415-0ac21a556c4a"
    // 结算卡反面
    dto["settle_card_back_pic"] = "9aa840b9-517e-3828-9e7e-7098eff89f24"
    // 结算卡正面
    dto["settle_card_front_pic"] = "f90c3466-bbb8-3309-b0d5-961abe6567b1"
    // 结算人身份证反面
    dto["settle_cert_back_pic"] = "7d4fd561-b70f-3b2e-b958-85f5328d226f"
    // 结算人身份证正面
    dto["settle_cert_front_pic"] = "a95a035a-d7e3-39fa-a869-dcc7d25a3349"
    // 税务登记证
    dto["tax_reg_pic"] = "d13832f9-2244-3a3b-ba09-936b100a8ce9"
    // 实名登记证书照片
    dto["cert_pic"] = "90dab223-0b14-3ec8-8c65-f388f1c6b1f1"
    // 个人商户身份证件正面照片
    dto["identification_front_pic"] = "b6fc4738-fe0a-3940-98e8-0987fee50b42"
    // 个人商户身份证件反面照片
    dto["identification_back_pic"] = "30492625-20c5-3796-822f-a10e63ba76e5"
    // 单位证明函照片
    dto["company_prove_pic"] = "36ac0355-a54d-3178-b4b5-9aecd07367e6"
    // 金融机构许可证图片1
    dto["finance_license_pic1"] = "ff647802-0ba1-36c0-952e-e94623cf0e7c"
    // 金融机构许可证图片2
    dto["finance_license_pic2"] = "42cecea7-1aef-33fb-bf04-c2bc621b0302"
    // 金融机构许可证图片3
    dto["finance_license_pic3"] = "48157e9b-44cc-33e1-8169-a0fe8c1c0848"
    // 金融机构许可证图片4
    dto["finance_license_pic4"] = "ca1cbd42-b14e-326b-9aef-288d45cf8b42"
    // 金融机构许可证图片5
    dto["finance_license_pic5"] = "2e74d95f-fd16-3766-ab39-c407c5b1c004"
    // 联系人身份证正面照
    dto["contact_id_front_pic"] = "b6fc4738-fe0a-3940-98e8-0987fee50b42"
    // 联系人身份证反面照
    dto["contact_id_back_pic"] = "71da066c-5d15-3658-a86d-4e85ee67808a"
    // 联系人护照人像面
    dto["contact_passport_img_pic"] = "8958a61c-970c-3ad8-a091-80238ef80a8b"
    // 联系人证件照正面
    dto["contact_cert_front_pic"] = "75ef9587-2faf-3b2c-820b-9ea447e754e3"
    // 联系人证件照反面
    dto["contact_cert_back_pic"] = "d42c010b-9316-369f-80ed-4ce4bda13602"
    // 微信业务办理授权函
    dto["contact_wx_busi_auth_pic"] = "49ac7d9b-851c-31b4-8b21-2983ea97b98c"
    // 行业经营许可证资质照片一
    dto["industry_busi_license_pic1"] = "1931c359-e42f-3e5f-875e-e22fc695aefd"
    // 行业经营许可证资质照片二
    dto["industry_busi_license_pic2"] = "0ddea6a0-6991-39ac-a68d-155d5d00d840"
    // 行业经营许可证资质照片三
    dto["industry_busi_license_pic3"] = "b5d77b0f-391f-3447-9843-386fc4096649"
    // 行业经营许可证资质照片四
    dto["industry_busi_license_pic4"] = "2af4514d-3d9c-3545-bc45-2424e80ab7e4"
    // 行业经营许可证资质照片五
    dto["industry_busi_license_pic5"] = "c3421d61-df60-3b2d-bcf1-e3709da867f2"
    // 行业经营许可证资质照片六
    dto["industry_busi_license_pic6"] = "b56c5cb1-4724-3574-ae38-7e8d5510b607"
    // 法人护照人像面
    dto["legal_passport_img_pic"] = "893dd8c7-c0a6-3cbd-a6c2-a52baf40398c"
    // 法人港澳台通行证正面
    dto["legal_hk_aom_front_pic"] = "8cb60559-e51c-344e-bcbf-96f3011acbd4"
    // 法人其他证件照片
    dto["legal_other_cert_pic"] = "562511a9-aa29-3e9e-9647-a97430ea9766"
    // 持卡人身份证人像面
    dto["cert_front_pic"] = "01c81d4f-0efd-3771-a8b2-660fe37f3859"
    // 持卡人身份证国徽面
    dto["cert_back_pic"] = "9f459914-4021-3c54-a5d9-31f1a64dc392"
    // 持卡人护照人像面
    dto["cert_passport_img_pic"] = "a501f0c6-a9ee-30d0-aedb-cec882da6d21"
    // 持卡人港澳台通行证正面
    dto["cert_hk_aom_front_pic"] = "cdcae795-6a9d-32f8-8033-d7bad4008974"
    // 持卡人其它证件照片
    dto["cert_other_pic"] = "398bbd13-40c9-37ce-8265-f6c1ecd317fa"
    // 签约人身份证照片-人像面
    // dto["sign_identity_front_file_id"] = ""
    // 签约人身份证照片-国徽面
    // dto["sign_identity_back_file_id"] = ""
    // 签约人法人授权书
    // dto["sign_auth_file_id"] = ""
    // 支付宝授权函照片
    // dto["contact_ali_busi_auth_pic"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

