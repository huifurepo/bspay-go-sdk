/**
 * 灵工企业商户进件接口 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2FlexibleEntRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2FlexibleEntRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2FlexibleEntRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 渠道商号
        UpperHuifuId:"6666000108406684",
        // 商户角色
        // MerRole:"test",
        // 落地公司类型当选择落地公司时，必填;SELF-自有，COOPERATE-合作
        // LocalCompanyType:"test",
        // 商户名称
        RegName:"圆务铁白事",
        // 商户简称
        ShortName:"沈桂英",
        // 证照图片
        LicensePic:"f6fc539b-73ff-3645-8539-ad318971cc48",
        // 证照编号
        LicenseCode:"91440101MA20250618",
        // 证照有效期类型
        LicenseValidityType:"1",
        // 证照有效期开始日期
        LicenseBeginDate:"20240314",
        // 证照有效期截止日期格式：yyyyMMdd。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20220125&lt;/font&gt;&lt;br/&gt;  当license_validity_type&#x3D;0时必填；当license_validity_type&#x3D;1时为空；
        LicenseEndDate:"",
        // 成立时间
        FoundDate:"20240314",
        // 注册区
        RegDistrictId:"310104",
        // 注册详细地址
        RegDetail:"方立全气气目队得形",
        // 法人姓名
        LegalName:"天天优先",
        // 法人证件类型
        LegalCertType:"00",
        // 法人证件号码
        LegalCertNo:"110101199003072551",
        // 法人证件有效期类型
        LegalCertValidityType:"1",
        // 法人证件有效期开始日期
        LegalCertBeginDate:"19710718",
        // 法人证件有效期截止日期日期格式：yyyyMMdd， &lt;font color&#x3D;&quot;green&quot;&gt;示例值：20220125&lt;/font&gt;&lt;br/&gt;当legal_cert_validity_type&#x3D;0时必填；&lt;br/&gt;当legal_cert_validity_type&#x3D;1时为空；&lt;br/&gt;
        LegalCertEndDate:"",
        // 法人证件地址
        LegalAddr:"信约候再研情其常",
        // 法人身份证国徽面
        LegalCertBackPic:"f6fc539b-73ff-3645-8539-ad318971cc48",
        // 法人身份证人像面
        LegalCertFrontPic:"f6fc539b-73ff-3645-8539-ad318971cc48",
        // 店铺门头照
        StoreHeaderPic:"f6fc539b-73ff-3645-8539-ad318971cc48",
        // 联系人手机号
        ContactMobileNo:"13074561767",
        // 联系人电子邮箱
        ContactEmail:"c.vwpjkqx@urgr.be",
        // 管理员账号
        LoginName:"req2025061853130071",
        // 银行卡信息配置
        CardInfo:getCaa0c059Bc8544639eabA1cbfb39c354(),
        // 签约人
        // SignUserInfo:get22486c6a11674ab78c2898a036901bf9(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2FlexibleEntRequest(dgReq)
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
    // 商户通知标识
    extendInfoMap["sms_send_flag"] = "1"
    // 证照类型
    extendInfoMap["license_type"] = "NATIONAL_LEGAL_MERGE"
    // 注册省
    extendInfoMap["reg_prov_id"] = "310000"
    // 注册市
    extendInfoMap["reg_area_id"] = "310100"
    // 法人手机号
    extendInfoMap["legal_mobile_no"] = "13074561767"
    // 联系人姓名
    extendInfoMap["contact_name"] = "文超"
    // 取现业务配置
    // extendInfoMap["cash_config"] = get990b98a297e44a189e1aF069844dc589()
    // 大额支付配置
    // extendInfoMap["large_amt_pay_config"] = getB721f3523cb14df09177415a293eb494()
    // 是否开通网银充值
    // extendInfoMap["online_recharge_flag"] = ""
    // 线上费率配置
    // extendInfoMap["online_fee_conf_list"] = getFb310e7292804509Ba9e019648efe90c()
    // 线上手续费承担方配置
    // extendInfoMap["online_pay_fee_conf_list"] = get20b809b62f9f4239A59fE5b8fcd2d44d()
    // 灵工支付配置
    // extendInfoMap["flexible_pay_config"] = getAa3c46b582694c21A5a82f4e29d46802()
    // 电子协议异步通知地址
    // extendInfoMap["agreement_async_return_url"] = ""
    // 异步请求地址
    // extendInfoMap["async_return_url"] = ""
    // 业务开通结果异步消息接收地址
    // extendInfoMap["busi_async_return_url"] = ""
    // 扩展资料包
    // extendInfoMap["extended_material_list"] = getE932a627C70c4eacAe676058420b597b()
    return extendInfoMap
}

func getCaa0c059Bc8544639eabA1cbfb39c354() string {
    dto := make(map[string]interface{})
    // 银行账户名
    dto["card_name"] = "圆务铁白事"
    // 银行账号
    dto["card_no"] = "6222021287358404424"
    // 银行所在市
    dto["area_id"] = "310100"
    // 银行编码
    dto["bank_code"] = "01050000"
    // 联行号
    dto["branch_code"] = "105290071050"
    // 银行所在省
    dto["prov_id"] = "310000"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get990b98a297e44a189e1aF069844dc589() string {
    dto := make(map[string]interface{})
    // 取现手续费（固定/元）fix_amt与fee_rate至少填写一项， 需保留小数点后两位，不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.00&lt;/font&gt;注：当cash_type&#x3D;D1时为节假日取现手续费；当cash_type&#x3D;T1时为工作日取现手续费
    // dto["fix_amt"] = "test"
    // 取现手续费率（%）fix_amt与fee_rate至少填写一项，需保留小数点后两位，取值范围[0.00,100.00]，不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.05&lt;/font&gt;注：1、如果fix_amt与fee_rate都填写了则手续费&#x3D;fix_amt+支付金额\*fee_rate2、当cash_type&#x3D;D1时为节假日取现手续费；当cash_type&#x3D;T1时为工作日取现手续费
    // dto["fee_rate"] = "test"
    // D1工作日取现手续费固定金额单位元，需保留小数点后两位。不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.00&lt;/font&gt;cash_type&#x3D;T1时，不生效 ；cash_type取现类型为D1时，遇工作日按此费率结算，若未配置则默认按照节假日手续费计算
    // dto["weekday_fix_amt"] = "test"
    // D1工作日取现手续费率单位%，需保留小数点后两位。取值范围[0.00，100.00]，不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.05&lt;/font&gt;cash_type&#x3D;T1时，不生效 ；cash_type取现类型为D1时，遇工作日按此费率结算 ，若未配置则默认按照节假日手续费计算
    // dto["weekday_fee_rate"] = "test"
    // 手续费承担方手续费外扣时必需指定手续费承担方ID ；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000123123123&lt;/font&gt;
    // dto["out_fee_huifu_id"] = "test"
    // 取现类型
    // dto["cash_type"] = ""
    // 是否取现手续费外扣
    // dto["out_fee_flag"] = ""
    // 手续费外扣的账户类型
    // dto["out_fee_acct_type"] = ""
    // 是否优先到账
    // dto["is_priority_receipt"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getE76fe6ae64054e81AcacBde2a5371a92() interface{} {
    dto := make(map[string]interface{})
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
    return dtoList
}

func getB721f3523cb14df09177415a293eb494() string {
    dto := make(map[string]interface{})
    // 大额支付配置列表
    // dto["large_amt_pay_config_info_list"] = getE76fe6ae64054e81AcacBde2a5371a92()
    // 交易手续费外扣huifuId交易手续费外扣时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000108854952&lt;/font&gt;
    // dto["out_fee_huifu_id"] = "test"
    // 交易手续费外扣账户号交易手续费外扣时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：F00598602&lt;/font&gt;
    // dto["out_fee_acct_id"] = "test"
    // 交易手续费外扣标记
    // dto["out_fee_flag"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getFb310e7292804509Ba9e019648efe90c() string {
    dto := make(map[string]interface{})
    // 银行编码
    // dto["bank_id"] = "test"
    // 功能开关状态
    // dto["stat_flag"] = "test"
    // 借贷标志
    // dto["dc_flag"] = "test"
    // 银行名称
    // dto["bank_name"] = ""
    // 银行中文简称
    // dto["bank_short_chn"] = ""
    // 手续费（固定/元）
    // dto["fix_amt"] = ""
    // 费率（%）
    // dto["fee_rate"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get20b809b62f9f4239A59fE5b8fcd2d44d() string {
    dto := make(map[string]interface{})
    // 业务类型
    // dto["pay_type"] = ""
    // 手续费外扣时的账户类型
    // dto["out_fee_acct_type"] = ""
    // 手续费外扣汇付ID
    // dto["out_fee_huifuid"] = ""
    // 是否交易手续费外扣
    // dto["out_fee_flag"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getAa3c46b582694c21A5a82f4e29d46802() string {
    dto := make(map[string]interface{})
    // 是否交易手续费外扣
    // dto["out_fee_flag"] = "test"
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

func get22486c6a11674ab78c2898a036901bf9() string {
    dto := make(map[string]interface{})
    // 签约人类型
    // dto["type"] = "test"
    // 姓名签约人类型&#x3D;联系人（经办人），必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：张三&lt;/font&gt;
    // dto["name"] = "test"
    // 身份证签约人类型&#x3D;联系人（经办人），必填 ；注意：**签约人会做姓名+身份证+手机号验证，请正确填写**；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：320946195712025082&lt;/font&gt;
    // dto["cert_no"] = "test"
    // 手机号签约人类型&#x3D;法人 ，必填；注意：**签约人会做姓名+身份证+手机号验证，请正确填写**；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：13917463536&lt;/font&gt;
    // dto["mobile_no"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getE932a627C70c4eacAe676058420b597b() string {
    dto := make(map[string]interface{})
    // 文件id
    // dto["file_id"] = "test"
    // 文件类型
    // dto["file_type"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

