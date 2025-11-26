/**
 * 商户基本信息修改 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantBasicdataModifyRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantBasicdataModifyRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantBasicdataModifyRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 直属渠道号
        UpperHuifuId:"6666000021000000",
        // 汇付客户Id
        HuifuId:"6666000107932702",
        // 签约人jsonObject格式；agreement_info中选择电子签约时必填；个人商户填本人信息。
        // SignUserInfo:get7b55b7d7391d43908eed85e6b73af3fd(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantBasicdataModifyRequest(dgReq)
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
    // 商户名称
    // extendInfoMap["reg_name"] = ""
    // 商户简称
    extendInfoMap["short_name"] = "尼博网络"
    // 商户英文名称
    // extendInfoMap["mer_en_name"] = ""
    // 公司类型
    // extendInfoMap["ent_type"] = ""
    // 所属行业
    extendInfoMap["mcc"] = "5411"
    // 营业执照类型
    // extendInfoMap["license_type"] = ""
    // 营业执照编号
    // extendInfoMap["license_code"] = ""
    // 营业执照有效期类型
    extendInfoMap["license_validity_type"] = "0"
    // 营业执照有效期开始日期
    extendInfoMap["license_begin_date"] = "20200814"
    // 营业执照有效期截止日期
    extendInfoMap["license_end_date"] = "20400813"
    // 注册省
    extendInfoMap["reg_prov_id"] = "310000"
    // 注册市
    extendInfoMap["reg_area_id"] = "310100"
    // 注册区
    extendInfoMap["reg_district_id"] = "310120"
    // 注册详细地址
    extendInfoMap["reg_detail"] = "台州市宜山路700解放2路715"
    // 注册资本
    // extendInfoMap["reg_capital"] = ""
    // 经营范围
    // extendInfoMap["business_scope"] = ""
    // 成立时间
    // extendInfoMap["found_date"] = ""
    // 法人姓名
    extendInfoMap["legal_name"] = "沈荣"
    // 法人证件类型
    extendInfoMap["legal_cert_type"] = "00"
    // 法人证件号码
    extendInfoMap["legal_cert_no"] = "320923199111206319"
    // 法人证件有效期类型
    extendInfoMap["legal_cert_validity_type"] = "0"
    // 法人证件有效期开始日期
    extendInfoMap["legal_cert_begin_date"] = "20200814"
    // 法人证件有效期截止日期
    extendInfoMap["legal_cert_end_date"] = "20400813"
    // 法人身份证地址
    // extendInfoMap["legal_addr"] = ""
    // 法人手机号
    // extendInfoMap["legal_mobile_no"] = ""
    // *负责人职业*
    // extendInfoMap["occupation"] = ""
    // 经营省
    extendInfoMap["prov_id"] = "310000"
    // 经营市
    extendInfoMap["area_id"] = "310100"
    // 经营区
    extendInfoMap["district_id"] = "310105"
    // 经营详细地址
    extendInfoMap["detail_addr"] = "徐州市徐汇区宜山路7497号"
    // 联系人姓名
    extendInfoMap["contact_name"] = "我是联系人"
    // 联系人电子邮箱
    extendInfoMap["contact_email"] = "mei.zhang@huifu.com"
    // 客服电话
    extendInfoMap["service_phone"] = "15556622368"
    // 小票名称
    extendInfoMap["receipt_name"] = "小票上的名称"
    // *结算卡信息配置*
    extendInfoMap["card_info"] = get7b629edd5720480eA2d1527ccceead1a()
    // 结算协议图片文件
    // extendInfoMap["settle_agree_pic"] = ""
    // 基本存款账户编号或核准号
    // extendInfoMap["open_licence_no"] = ""
    // 取现信息配置
    extendInfoMap["cash_config"] = getB4c12b9e620844179e02A63a50bae972()
    // 结算规则配置
    extendInfoMap["settle_config"] = getA2194d6e98dd4dce9d79C15b8a52a54c()
    // 商户主页URL
    // extendInfoMap["mer_url"] = ""
    // 商户ICP备案编号
    // extendInfoMap["mer_icp"] = ""
    // 受益人列表
    // extendInfoMap["beneficiary_info"] = get92d7aa547c6c4e40Bfba92b33a3e6691()
    // 协议信息
    // extendInfoMap["agreement_info"] = getAd408404B0764f60896838dca28e34d6()
    // 营业执照图片
    // extendInfoMap["license_pic"] = ""
    // 授权委托书
    // extendInfoMap["auth_entrust_pic"] = ""
    // 场景类型
    // extendInfoMap["scene_type"] = ""
    // 店铺门头照
    // extendInfoMap["store_header_pic"] = ""
    // 店铺内景/工作区域照
    // extendInfoMap["store_indoor_pic"] = ""
    // 店铺收银台/公司前台照
    // extendInfoMap["store_cashier_desk_pic"] = ""
    // 扩展资料包
    // extendInfoMap["extended_material_list"] = get0e652d5c278b4b6289416af49f5827aa()
    // 异步通知地址
    extendInfoMap["async_return_url"] = "archer://C_SSPM_NSPOSM_BUSIRESULT"
    // 斗拱e账户功能配置
    // extendInfoMap["elec_acct_config"] = get0c477a60A33445bf8a96464788abe8e9()
    // 股东信息
    // extendInfoMap["share_holder_info_list"] = get48c5d1f8330f4e8b880aBefad36b0aaf()
    // 外部商户号
    // extendInfoMap["ext_mer_id"] = ""
    // 备注
    // extendInfoMap["remark"] = ""
    return extendInfoMap
}

func get7b629edd5720480eA2d1527ccceead1a() string {
    dto := make(map[string]interface{})
    // 结算账户类型
    dto["card_type"] = "1"
    // 结算账户名
    dto["card_name"] = "万方肖"
    // 结算账号
    dto["card_no"] = "1001265009300682579"
    // 银行所在市
    dto["area_id"] = "410100"
    // 支行联行号当card_type&#x3D;0时必填。参考：[银行支行编码](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_yhzhbm) 当card_type&#x3D;0时必填， 当card_type&#x3D;1或2时非必填 &lt;font color&#x3D;&quot;green&quot;&gt;示例值：102290026507&lt;/font&gt;
    dto["branch_code"] = "102491002812"
    // 持卡人证件类型参见《[自然人证件类型](https://paas.huifu.com/open/doc/api/#/api_ggcsbm?id&#x3D;%e8%87%aa%e7%84%b6%e4%ba%ba%e8%af%81%e4%bb%b6%e7%b1%bb%e5%9e%8b)》。&lt;br/&gt; 当card_type&#x3D;0时为空， 当card_type&#x3D;1或2时必填； &lt;font color&#x3D;&quot;green&quot;&gt;示例值：00&lt;/font&gt;&lt;br/&gt;持卡人证件类型04、11需要补充 F101【结算人】港澳台居民来往内地通行证&lt;br/&gt;持卡人证件类型13需要补充 F513【结算人】外国人居留证&lt;br/&gt;持卡人证件类型14、15需要补充 F514【结算人】港澳台居住证&lt;br/&gt;其它持卡人证件类型补充F102【结算人】其它证件材料&lt;br/&gt;补充材料填写在extended_material_list 扩展资料包中
    dto["cert_type"] = "00"
    // 持卡人证件有效期类型0：非长期有效, 1：长期有效, &lt;font color&#x3D;&quot;green&quot;&gt;示例值：0&lt;/font&gt;&lt;br/&gt;当card_type&#x3D;0时为空； 当card_type&#x3D;1或2时必填；
    dto["cert_validity_type"] = "0"
    // 持卡人证件有效期开始日期日期格式：yyyyMMdd，以北京时间为准； 当card_type&#x3D;0时为空， 当card_type&#x3D;1或2时必填， &lt;font color&#x3D;&quot;green&quot;&gt;示例值：20220125 &lt;/font&gt;
    dto["cert_begin_date"] = "20140905"
    // 持卡人证件有效期截止日期日期格式：yyyyMMdd，以北京时间为准。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20220125&lt;/font&gt;&lt;br/&gt;  当cert_validity_type&#x3D;0时必填；当cert_validity_type&#x3D;1时为空
    dto["cert_end_date"] = "20240905"
    // 开户许可证开户许可证图片文件，请填写图片file_id，可通过 [商户图片上传](https://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc) 接口获取，文件类型F08；&lt;br/&gt;企业商户需要，结算账号为对公账户必填；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e529&lt;/font&gt;
    dto["reg_acct_pic"] = "d1451277-85c6-3177-ac3d-a8be47b9ae9d"
    // 法人身份证人像面请填写图片file_id，可通过 [商户图片上传](https://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc) 接口获取，文件类型F02；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e529 &lt;/font&gt;
    dto["legal_cert_front_pic"] = ""
    // 法人身份证国徽面请填写图片file_id，可通过 [商户图片上传](https://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc) 接口获取，文件类型F03；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e529&lt;/font&gt;
    dto["legal_cert_back_pic"] = ""
    // 银行卡卡号面银行卡卡号面图片文件对私结算必填，请填写图片file_id，可通过 [商户图片上传](https://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc) 接口获取，文件类型F13；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e529&lt;/font&gt;
    dto["settle_card_front_pic"] = "d1451277-85c6-3177-ac3d-a8be47b9ae9d"
    // 结算人身份证人像面对私结算必填，请填写图片file_id，可通过 [商户图片上传](https://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc) 接口获取，文件类型F55；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e529&lt;/font&gt;
    dto["settle_cert_front_pic"] = "d1451277-85c6-3177-ac3d-a8be47b9ae9d"
    // 结算人身份证国徽面对私结算必填，请填写图片file_id，可通过 [商户图片上传](https://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc) 接口获取，文件类型F56；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e529  &lt;/font&gt;
    dto["settle_cert_back_pic"] = "d1451277-85c6-3177-ac3d-a8be47b9ae9d"
    // 短信验证码结算卡、结算人变更需要短信验证码；当【变更结算人核验方式 &#x3D; 0】时必填。 请先调用商户短信发送接口；示例值：851249
    // dto["verify_code"] = "test"
    // 核验人类型当【变更结算人核验方式 &#x3D; 1】时必填。企业用户选择0:法人，1:原持卡人；&lt;font  color&#x3D;&quot;green&quot;&gt;示例值：0&lt;/font&gt;
    // dto["authentication_person_type"] = "test"
    // 核验人手机号当【变更结算人核验方式 &#x3D; 1】时必填。请根据核验人类型传入核验人实名手机号。&lt;font  color&#x3D;&quot;green&quot;&gt;示例值：13911111111&lt;/font&gt;
    // dto["authentication_person_mobile"] = "test"
    // 银行所在省
    dto["prov_id"] = "310000"
    // 持卡人证件号码
    dto["cert_no"] = "320923199111206319"
    // 结算人手机号
    dto["mp"] = "18221987178"
    // 变更结算人核验方式
    // dto["authentication_type"] = ""
    // 适用子账户
    // dto["acct_ids"] = ""
    // 默认结算卡标志
    // dto["is_settle_default"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getB4c12b9e620844179e02A63a50bae972() string {
    dto := make(map[string]interface{})
    // 状态
    dto["switch_state"] = "1"
    // 业务类型
    dto["cash_type"] = "D0"
    // 提现手续费（固定/元）fix_amt与fee_rate至少填写一项， 需保留小数点后两位，不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.00&lt;/font&gt;注：当cash_type&#x3D;D1时为节假日取现手续费
    dto["fix_amt"] = "4.00"
    // 提现手续费率（%）fix_amt与fee_rate至少填写一项，需保留小数点后两位，取值范围[0.00,100.00]，不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.05&lt;/font&gt;注：1、如果fix_amt与fee_rate都填写了则手续费&#x3D;fix_amt+支付金额\*fee_rate2、当cash_type&#x3D;D1时为节假日取现手续费
    dto["fee_rate"] = "5.50"
    // D1工作日取现手续费固定金额单位元，需保留小数点后两位。不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.00&lt;/font&gt;cash_type&#x3D;T1时，不生效 ；cash_type取现类型为D1时，遇工作日按此费率结算，若未配置则默认按照节假日手续费计算
    // dto["weekday_fix_amt"] = "test"
    // D1工作日取现手续费率单位%，需保留小数点后两位。取值范围[0.00，100.00]，不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.05&lt;/font&gt;cash_type&#x3D;T1时，不生效 ；cash_type取现类型为D1时，遇工作日按此费率结算 ，若未配置则默认按照节假日手续费计算
    // dto["weekday_fee_rate"] = "test"
    // 是否交易手续费外扣
    // dto["out_fee_flag"] = ""
    // 手续费承担方
    // dto["out_fee_huifu_id"] = ""
    // 手续费外扣的账户类型
    // dto["out_fee_acct_type"] = ""
    // 是否优先到账
    // dto["is_priority_receipt"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getA2194d6e98dd4dce9d79C15b8a52a54c() string {
    dto := make(map[string]interface{})
    // 结算开关
    dto["settle_status"] = "1"
    // 结算周期
    dto["settle_cycle"] = "T1"
    // 结算手续费外扣商户号结算手续费外扣商户号，填写承担手续费的汇付商户号；当out_settle_flag&#x3D;1时必填，否则非必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000123123123&lt;/font&gt;
    // dto["out_settle_huifuid"] = "test"
    // 节假日结算手续费率settle_cycle为D1、TS时必填。单位%，需保留小数点后两位。取值范围[0.00，100.00]，不收费请填写0.00；settle_cycle&#x3D;T1时，不生效 ；settle_cycle为D1时，遇节假日按此费率结算 ；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.05&lt;/font&gt;
    dto["fixed_ratio"] = "0"
    // 节假日结算手续费固定金额settle_cycle为D1、TS时必填。单位元，需保留小数点后两位。不收费请填写0.00；settle_cycle结算周期为D1时，遇节假日按此费率结算 ；&lt;br/&gt; &lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.00&lt;/font&gt;
    dto["constant_amt"] = "0"
    // 起结金额
    // dto["min_amt"] = ""
    // 留存金额
    // dto["remained_amt"] = ""
    // 结算摘要
    // dto["settle_abstract"] = ""
    // 手续费外扣标记
    dto["out_settle_flag"] = "2"
    // 结算手续费外扣时的账户类型
    dto["out_settle_acct_type"] = "09"
    // 结算批次号
    dto["settle_batch_no"] = "100"
    // 是否优先到账
    dto["is_priority_receipt"] = "N"
    // 自定义结算处理时间
    // dto["settle_time"] = ""
    // 工作日结算手续费率
    // dto["workday_fixed_ratio"] = ""
    // 工作日结算手续费固定金额
    // dto["workday_constant_amt"] = ""
    // 结算方式
    dto["settle_pattern"] = "P0"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get92d7aa547c6c4e40Bfba92b33a3e6691() string {
    dto := make(map[string]interface{})
    // 受益人名称
    // dto["bo_name"] = "test"
    // 受益人证件类型
    // dto["bo_type"] = "test"
    // 受益人证件号
    // dto["bo_no"] = "test"
    // 受益人证件有效期开始时间
    // dto["bo_date_start"] = "test"
    // 受益人证件有效期结束时间
    // dto["bo_deadline"] = "test"
    // 受益人地址
    // dto["bo_address"] = "test"
    // 受益人手机号开通全域资金管理业务时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：13911111111&lt;/font&gt;
    // dto["bo_mobile_no"] = "test"
    // 最终受益人受益方式A01：直接或间接控股25%（含）以上 &lt;br/&gt;A02：通过人事、财务等其他方式对公司进行控制 &lt;br/&gt;A03：高级管理人员 &lt;br/&gt;A04：法人或公司负责人 &lt;br/&gt;A05：其他&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：A01&lt;/font&gt; &lt;br/&gt;开通全域资金管理业务时必填
    // dto["final_beneficiary_mode"] = "test"
    // 受益唯一ID
    // dto["bo_id"] = ""
    // 受益人有效标识
    // dto["bo_stat"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get7b55b7d7391d43908eed85e6b73af3fd() string {
    dto := make(map[string]interface{})
    // 签约人类型
    // dto["type"] = "test"
    // 姓名
    // dto["name"] = ""
    // 手机号
    // dto["mobile_no"] = ""
    // 身份证
    // dto["cert_no"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getAd408404B0764f60896838dca28e34d6() string {
    dto := make(map[string]interface{})
    // 协议类型
    // dto["agreement_type"] = "test"
    // 挂网协议地址挂网协议必填；**必须按示例值填写**，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/opps/api/prod/download_file/PaymentServiceAgreement.htm&lt;/font&gt;；
    // dto["agreement_url"] = "test"
    // 纸质协议开始日期
    // dto["agree_begin_date"] = ""
    // 纸质协议结束日期
    // dto["agree_end_date"] = ""
    // 电子协议异步通知地址
    // dto["agreement_async_return_url"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get0e652d5c278b4b6289416af49f5827aa() string {
    dto := make(map[string]interface{})
    // 文件id
    // dto["file_id"] = "test"
    // 文件类型
    // dto["file_type"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get284f0185740d4dd3Bf1409a33a5e2f61() interface{} {
    dto := make(map[string]interface{})
    // 银行编码
    // dto["bank_code"] = "test"
    // 支行联行号
    // dto["branch_code"] = "test"
    // 支行名称
    // dto["branch_name"] = "test"
    // 结算账户名
    // dto["card_name"] = "test"
    // 银行卡号
    // dto["card_no"] = "test"
    // 卡类型
    // dto["card_type"] = "test"
    // 银行绑定手机号个人必填
    // dto["mp"] = "test"
    // 用户授权协议版本号该字段在绑定个人账户时必填，取值商户自定义。与个人用户签约的电子协议版本号，通过该版本号能够确定协议的具体内容
    // dto["auth_version"] = "test"
    // 用户授权协议号该字段在绑定个人账户时必填，取值商户自定义。与个人用户签约的授权交易流水号，通过该流水号应能确定电子协议版本号、签约人、签约时间
    // dto["auth_no"] = "test"
    // 银行所在省
    // dto["prov_id"] = ""
    // 银行所在市
    // dto["area_id"] = ""
    // 默认卡标识
    // dto["default_cash_flag"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get0c477a60A33445bf8a96464788abe8e9() string {
    dto := make(map[string]interface{})
    // 电子账户开关
    // dto["switch_state"] = "test"
    // 账户类型
    // dto["acct_type"] = "test"
    // 电子账户提现手续费承担方下级商户必填；1:总部 2:其他；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1&lt;/font&gt;
    // dto["cash_fee_party"] = "test"
    // 场景类型
    // dto["scene"] = "test"
    // 角色类型
    // dto["role_type"] = "test"
    // 银行卡信息
    // dto["elec_card_list"] = get284f0185740d4dd3Bf1409a33a5e2f61()
    // 中信签约短信流水号
    // dto["elec_acct_sign_seq_id"] = ""
    // 签约成功标志
    // dto["sign_success_flag"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get48c5d1f8330f4e8b880aBefad36b0aaf() string {
    dto := make(map[string]interface{})
    // 股东姓名
    // dto["name"] = "test"
    // 股东证件类型
    // dto["cert_type"] = "test"
    // 股东证件号码
    // dto["cert_no"] = "test"
    // 股东证件有效期类型
    // dto["cert_validity_type"] = "test"
    // 股东证件有效期起始日
    // dto["cert_begin_date"] = "test"
    // 股东编号
    // dto["share_holder_id"] = ""
    // 股东证件有效期到期日
    // dto["cert_end_date"] = ""
    // 股东状态
    // dto["state"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

