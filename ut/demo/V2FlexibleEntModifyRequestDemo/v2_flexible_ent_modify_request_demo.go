/**
 * 灵工企业商户业务修改 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2FlexibleEntModifyRequestDemo

import (
	"encoding/json"
	"fmt"

	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2FlexibleEntModifyRequestDemo() {
	// 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

	// 2.组装请求参数
	dgReq := BsPaySdk.V2FlexibleEntModifyRequest{
		// 请求流水号
		ReqSeqId: tool.GetReqSeqId(),
		// 请求日期
		ReqDate: tool.GetCurrentDate(),
		// 商户汇付id
		HuifuId: "6666000108903754",
		// 渠道商汇付ID
		UpperHuifuId: "6666000108329682",
		// 商户基本信息jsonObject格式；其中的contact_info和legal_info联系人和法人信息可能在卡信息修改时需要
		BasicInfo: getAf9330e0B7264377B857891dc3eb74b3(),
		// 签约人
		SignUserInfo: getDe548fa487934ee186afEa5616d3ccec(),
	}
	// 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

	fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
	fmt.Println("请求数据:", string(respStr))

	// 3. 发起API调用
	resp, err := dgSDK.V2FlexibleEntModifyRequest(dgReq)
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
	// 卡信息配置实体
	// extendInfoMap["card_info"] = getA2ffaf5c578d4b6bB755C3b8e7d5b8ad()
	// 取现配置列表
	extendInfoMap["cash_config"] = getA24f3b43221241a093199927ac24f2c2()
	// 大额支付配置
	extendInfoMap["large_amt_pay_config"] = get82e2c731C2f54b44B1d092c62832ed01()
	// 是否开通网银充值
	extendInfoMap["online_recharge_flag"] = "Y"
	// 线上费率配置
	extendInfoMap["online_fee_conf_list"] = get8ff6cf9e57f54a5a9483718f03c3f835()
	// 线上手续费承担方配置
	extendInfoMap["online_pay_fee_conf_list"] = get1af27596Cf084258Bc586a3b7cb7c5b7()
	// 灵工支付配置
	// extendInfoMap["flexible_pay_config"] = getE3a042a10b974c549f6c1d254f199c71()
	// 扩展资料包
	// extendInfoMap["extended_material_list"] = getFc39cf6e632c458b923dFf4a8a8ca660()
	// 电子协议异步通知地址
	extendInfoMap["agreement_async_return_url"] = ""
	// 异步请求地址
	extendInfoMap["async_return_url"] = ""
	// 业务开通结果异步消息接收地址
	extendInfoMap["busi_async_return_url"] = ""
	return extendInfoMap
}

func get653797d254e64a5087bcB4b7943d72c5() interface{} {
	dto := make(map[string]interface{})
	// 营业执照类型
	dto["license_type"] = "CERTIFICATE_TYPE_0001"
	// 营业执照编号
	// dto["license_code"] = "test"
	// 营业执照有效期开始日期
	dto["license_begin_date"] = "20210101"
	// 营业执照有效期截止日期
	dto["license_end_date"] = "20410101"
	// 营业执照有效期类型
	dto["license_validity_type"] = "0"
	// 注册省
	dto["reg_prov_id"] = "340000"
	// 注册市
	dto["reg_area_id"] = "340100"
	// 注册区
	dto["reg_district_id"] = "340102"
	// 注册详细地址
	dto["reg_detail"] = "瑶海区1号"
	// 证照图片
	dto["license_pic"] = "67cce967-bd22-32b4-a250-58b82e78154a"

	return dto
}

func get8c2a41ad32e54400Af7b162f388c2ad2() interface{} {
	dto := make(map[string]interface{})
	// 店铺门头照
	dto["store_header_pic"] = "ececec87-e772-331a-b4fc-fae92732d992"
	// 经营简称
	dto["short_name"] = "简称888"

	return dto
}

func get26d8f85d50ab4f5aAcb66d645323ce74() interface{} {
	dto := make(map[string]interface{})
	// 法人姓名
	dto["legal_name"] = "岑晓"
	// 法人证件类型
	dto["legal_cert_type"] = "00"
	// 法人证件号码
	dto["legal_cert_no"] = "513701199108105217"
	// 法人证件有效期类型
	dto["legal_cert_validity_type"] = ""
	// 法人证件有效期开始日期
	dto["legal_cert_begin_date"] = ""
	// 法人证件有效期截止日期
	dto["legal_cert_end_date"] = ""
	// 法人手机号
	dto["legal_mobile_no"] = ""
	// 法人证件地址
	dto["legal_addr"] = ""
	// 法人身份证国徽面
	dto["legal_cert_back_pic"] = ""
	// 法人身份证人像面
	dto["legal_cert_front_pic"] = ""

	return dto
}

func get95a38d50674543749f622f58fe3e01d3() interface{} {
	dto := make(map[string]interface{})
	// 联系人姓名
	dto["contact_name"] = "岑晓"
	// 联系人手机号
	dto["contact_mobile_no"] = "18777716992"
	// 联系人电子邮箱
	dto["contact_email"] = "lianxi@huifu.com"

	return dto
}

func getAf9330e0B7264377B857891dc3eb74b3() string {
	dto := make(map[string]interface{})
	// 营业执照信息
	dto["license_info"] = get653797d254e64a5087bcB4b7943d72c5()
	// 经营信息
	dto["company_info"] = get8c2a41ad32e54400Af7b162f388c2ad2()
	// 法人信息
	dto["legal_info"] = get26d8f85d50ab4f5aAcb66d645323ce74()
	// 联系人信息
	dto["contact_info"] = get95a38d50674543749f622f58fe3e01d3()

	dtoByte, _ := json.Marshal(dto)
	return string(dtoByte)
}

func getDe548fa487934ee186afEa5616d3ccec() string {
	dto := make(map[string]interface{})
	// 签约人类型
	dto["type"] = "CONTACT"
	// 姓名
	dto["name"] = "施忠晶"
	// 手机号
	dto["mobile_no"] = "13812345231"
	// 身份证
	dto["cert_no"] = "513701201104022297"

	dtoByte, _ := json.Marshal(dto)
	return string(dtoByte)
}

func getA2ffaf5c578d4b6bB755C3b8e7d5b8ad() interface{} {
	dto := make(map[string]interface{})
	// 银行账户名
	// dto["card_name"] = "test"
	// 银行账号
	// dto["card_no"] = "test"
	// 银行所在省
	// dto["prov_id"] = "test"
	// 银行所在市
	// dto["area_id"] = "test"
	// 银行编码
	// dto["bank_code"] = "test"
	// 联行号
	// dto["branch_code"] = "test"

	return dto
}

func getA24f3b43221241a093199927ac24f2c2() interface{} {
	dto := make(map[string]interface{})
	// 是否开通取现
	dto["switch_state"] = ""
	// 业务类型
	dto["cash_type"] = "D1"
	// 提现手续费（固定/元）开通提现业务时fix_amt与fee_rate至少填写一项；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.00&lt;/font&gt;
	dto["fix_amt"] = "2"
	// D1工作日取现手续费固定金额单位元，需保留小数点后两位。不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.00&lt;/font&gt;cash_type&#x3D;T1时，不生效 ；cash_type取现类型为D1时，遇工作日按此费率结算，若未配置则默认按照节假日手续费计算
	dto["weekday_fix_amt"] = "2"
	// D1工作日取现手续费率单位%，需保留小数点后两位。取值范围[0.00，100.00]，不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.05&lt;/font&gt;cash_type&#x3D;T1时，不生效 ；cash_type取现类型为D1时，遇工作日按此费率结算 ，若未配置则默认按照节假日手续费计算
	dto["weekday_fee_rate"] = "3"
	// 取现手续费率（%）
	dto["fee_rate"] = "10.00"
	// 是否交易手续费外扣
	dto["out_fee_flag"] = "1"
	// 手续费承担方
	dto["out_fee_huifu_id"] = "6666000108329682"
	// 手续费外扣的账户类型
	dto["out_fee_acct_type"] = "01"
	// 是否优先到账
	dto["is_priority_receipt"] = "Y"

	dtoList := [1]interface{}{dto}
	return dtoList
}

func getCaa423dd4dbf41b0981a73e8906bd5d5() interface{} {
	dto := make(map[string]interface{})
	// 费率（%）开通大额业务时必须填写一种收费方式；大于0,保留2位小数；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.00&lt;/font&gt;
	dto["fee_rate"] = "4"
	// 交易手续费（固定/元）开通大额业务时必须填写一种收费方式；大于0,保留2位小数；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：10.00&lt;/font&gt;
	dto["fee_fix_amt"] = "5"
	// 功能开关
	dto["switch_state"] = "1"
	// 大额调账标识申请类型
	dto["biz_type"] = "01"
	// 是否允许绑卡支付
	dto["mer_same_card_recharge_flag"] = "Y"
	// 是否允许用户入账
	dto["allow_user_deposit_flag"] = "N"
	// 备付金固定账号模式自动退款
	dto["provisions_auto_refund_flag"] = "Y"

	dtoList := [1]interface{}{dto}
	return dtoList
}

func get82e2c731C2f54b44B1d092c62832ed01() string {
	dto := make(map[string]interface{})
	// 大额支付配置列表
	dto["large_amt_pay_config_info_list"] = getCaa423dd4dbf41b0981a73e8906bd5d5()
	// 交易手续费外扣huifuId交易手续费外扣时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000108854952&lt;/font&gt;
	dto["out_fee_huifu_id"] = "6666000108329682"
	// 交易手续费外扣账户号交易手续费外扣时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：F00598602&lt;/font&gt;
	dto["out_fee_acct_id"] = "C02714529"
	// 交易手续费外扣标记
	dto["out_fee_flag"] = "1"

	dtoByte, _ := json.Marshal(dto)
	return string(dtoByte)
}

func get8ff6cf9e57f54a5a9483718f03c3f835() string {
	dto := make(map[string]interface{})
	// 银行编码
	dto["bank_id"] = "01050000"
	// 功能开关状态
	dto["stat_flag"] = "1"
	// 借贷标志
	dto["dc_flag"] = "D"
	// 银行名称
	dto["bank_name"] = "中国建设银行股份有限公司上海市中支行"
	// 银行中文简称
	dto["bank_short_chn"] = "上海市中支行"
	// 手续费（固定/元）
	dto["fix_amt"] = "6"
	// 费率（%）
	dto["fee_rate"] = "6"

	dtoList := [1]interface{}{dto}
	dtoByte, _ := json.Marshal(dtoList)
	return string(dtoByte)
}

func get1af27596Cf084258Bc586a3b7cb7c5b7() string {
	dto := make(map[string]interface{})
	// 业务类型
	dto["pay_type"] = "ONLINE_ENT_TOP_UP"
	// 手续费外扣时的账户类型
	dto["out_fee_acct_type"] = "01"
	// 手续费外扣汇付ID
	dto["out_fee_huifuid"] = "6666000108329682"
	// 是否交易手续费外扣
	dto["out_fee_flag"] = "1"

	dtoList := [1]interface{}{dto}
	dtoByte, _ := json.Marshal(dtoList)
	return string(dtoByte)
}

func getE3a042a10b974c549f6c1d254f199c71() string {
	dto := make(map[string]interface{})
	// 是否交易手续费外扣
	// dto["out_fee_flag"] = "test"
	// 功能开关
	// dto["switch_state"] = ""
	// 手续费(%)
	// dto["fee_rate"] = ""
	// 手续费（固定/元）
	// dto["fix_amt"] = ""
	// 外扣规则
	// dto["out_charge_mode"] = ""
	// 手续费外扣时的账户ID
	// dto["out_fee_acct_id"] = ""
	// 手续费外扣汇付ID
	// dto["out_fee_huifu_id"] = ""

	dtoByte, _ := json.Marshal(dto)
	return string(dtoByte)
}

func getFc39cf6e632c458b923dFf4a8a8ca660() string {
	dto := make(map[string]interface{})
	// 文件id
	// dto["file_id"] = "test"
	// 文件类型
	// dto["file_type"] = "test"

	dtoList := [1]interface{}{dto}
	dtoByte, _ := json.Marshal(dtoList)
	return string(dtoByte)
}
