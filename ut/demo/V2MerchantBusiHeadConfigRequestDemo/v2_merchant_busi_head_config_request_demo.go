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
    // extendInfoMap["ali_conf_list"] = get913c1602480a4d3f9272Ec0e767d7aeb()
    // 微信配置对象
    // extendInfoMap["wx_conf_list"] = get7c4de91918c54844A91660cc790a6279()
    // 银联二维码配置对象
    // extendInfoMap["union_conf_list"] = get5bb3480e73e74c89A5b8E8d55fcd727f()
    // 银联卡配置对象
    // extendInfoMap["bank_card_config"] = get43610f0c07234fd6Aaa572407a3afc14()
    // 分账配置对象
    // extendInfoMap["split_config"] = get6966f9e50a5440a58786Ff2442383ca1()
    // 微信直连配置对象
    // extendInfoMap["wx_zl_conf_list"] = getB9b15afd4a8845c68a93E1842aff41af()
    // 支付宝直连配置对象
    // extendInfoMap["ali_zl_conf"] = getA690bc363cdd4dd98ff27380d9a5e3ad()
    // 线上配置对象
    // extendInfoMap["online_fee_conf_list"] = getF65bded5Bed64adfB2b93e7b63e655ef()
    // 余额支付配置对象
    // extendInfoMap["balance_pay_config"] = getE2ea5610E72a4eb7BcbfB22e421f2280()
    // 补贴支付配置对象
    // extendInfoMap["combine_pay_config"] = get954e91de94c1460c84a09c566d830256()
    // 银行大额转账配置对象
    // extendInfoMap["bank_big_amt_pay_config"] = getF455cd9778ea4b56994dF548ccc8f072()
    // 全域资金管理配置对象（华通银行）
    // extendInfoMap["out_order_funds_config"] = getD7cd57bdAdb745cfA5dd146451d79cfa()
    // 全域资金管理配置(XW银行)
    // extendInfoMap["out_order_funds_new_net_config"] = getBf284cc60b5b4672887749ad39dc8524()
    // 结算配置对象
    // extendInfoMap["settle_config_list"] = getE6d4b7d816ee4f47887fDdd12d36aa23()
    // 取现配置对象
    // extendInfoMap["cash_config_list"] = get5d6c5e557bb04748B1254a84df2d23ef()
    // 外扣配置对象
    // extendInfoMap["out_fee_config"] = get4ad35a9b5ff64f65A4e1A84915733f4d()
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
    // 大额支付配置
    // extendInfoMap["large_amt_pay_config_list"] = get91761c415cb347e59369Ada12376ef40()
    // 全域资金管理配置(苏商)
    // extendInfoMap["out_order_funds_su_shang_config"] = getD8d27591E0424fc9B36115f70c06480b()
    // 托管支付开关
    // extendInfoMap["half_pay_host_flag"] = ""
    // 全域资金费用配置对象
    // extendInfoMap["out_order_funds_fee_list"] = getE42011c14cae453788fb872edae6a744()
    // 本地生活生活配置对象
    // extendInfoMap["lla_withhold_config"] = getEd96185593a44f3480a2F84da188b3ee()
    return extendInfoMap
}

func get913c1602480a4d3f9272Ec0e767d7aeb() string {
    dto := make(map[string]interface{})
    // 支付场景
    // dto["pay_scene"] = "test"
    // 手续费
    // dto["fee_rate"] = "test"
    // 允许开通该业务
    // dto["open_flag"] = "test"
    // 最低收取手续费（元）
    // dto["fee_min_amt"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get7c4de91918c54844A91660cc790a6279() string {
    dto := make(map[string]interface{})
    // 支付场景
    // dto["pay_scene"] = "test"
    // 手续费
    // dto["fee_rate"] = "test"
    // 允许开通该场景业务
    // dto["open_flag"] = "test"
    // 最低收取手续费（元）
    // dto["fee_min_amt"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get5bb3480e73e74c89A5b8E8d55fcd727f() string {
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

func get43610f0c07234fd6Aaa572407a3afc14() string {
    dto := make(map[string]interface{})
    // 借记卡手续费（%）
    // dto["debit_fee_rate"] = "test"
    // 贷记卡手续费（%）
    // dto["credit_fee_rate"] = "test"
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

func get6966f9e50a5440a58786Ff2442383ca1() string {
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

func getB9b15afd4a8845c68a93E1842aff41af() string {
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

func getA690bc363cdd4dd98ff27380d9a5e3ad() string {
    dto := make(map[string]interface{})
    // 手续费（%）
    // dto["fee_rate"] = "test"
    // 允许开通支付宝直连业务
    // dto["open_flag"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getF65bded5Bed64adfB2b93e7b63e655ef() string {
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
    // 手续费最小值（元）
    // dto["fee_min_amt"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getE2ea5610E72a4eb7BcbfB22e421f2280() string {
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

func get954e91de94c1460c84a09c566d830256() string {
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

func getF455cd9778ea4b56994dF548ccc8f072() string {
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

func getD7cd57bdAdb745cfA5dd146451d79cfa() string {
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

func getBf284cc60b5b4672887749ad39dc8524() string {
    dto := make(map[string]interface{})
    // 手续费（%）
    // dto["fee_rate"] = "test"
    // 保底手续费(元)
    // dto["fee_min_amt"] = "test"
    // 对公固定手续费(元)
    // dto["public_fee_fix_amt"] = "test"
    // 对私固定手续费(元)
    // dto["private_fee_fix_amt"] = "test"
    // 允许开通全域资金业务(XW)
    // dto["open_flag"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getE6d4b7d816ee4f47887fDdd12d36aa23() string {
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

func get5d6c5e557bb04748B1254a84df2d23ef() string {
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

func get4ad35a9b5ff64f65A4e1A84915733f4d() string {
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

func get91761c415cb347e59369Ada12376ef40() string {
    dto := make(map[string]interface{})
    // 手续费（%）
    // dto["fee_rate"] = "test"
    // 固定手续费(元)
    // dto["fee_fix_amt"] = "test"
    // 允许开通大额转账业务
    // dto["open_flag"] = "test"
    // 大额支付业务模式
    // dto["business_model"] = "test"
    // 允许用户入账
    // dto["allow_user_deposit_flag"] = ""
    // 银行卡绑定支付权限
    // dto["mer_same_card_recharge_flag"] = ""
    // 备付金固定账号模式自动退款
    // dto["provisions_auto_refund_flag"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getD8d27591E0424fc9B36115f70c06480b() string {
    dto := make(map[string]interface{})
    // 手续费（%）
    // dto["fee_rate"] = "test"
    // 保底手续费(元)
    // dto["fee_min_amt"] = "test"
    // 对公固定手续费(元)
    // dto["public_fee_fix_amt"] = "test"
    // 对私固定手续费(元)
    // dto["private_fee_fix_amt"] = "test"
    // 允许开通全域资金业务(苏商)
    // dto["open_flag"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getE42011c14cae453788fb872edae6a744() string {
    dto := make(map[string]interface{})
    // 业务类型
    // dto["bus_type"] = "test"
    // 手续费（百分比/%）
    // dto["fee_rate"] = "test"
    // 手续费（固定/元）
    // dto["fee_fix_amt"] = "test"
    // 是否允许开通
    // dto["open_flag"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getEd96185593a44f3480a2F84da188b3ee() interface{} {
    dto := make(map[string]interface{})
    // 本地生活开关
    // dto["llaWithholdFlag"] = "test"
    // 佣金收取手续费率
    // dto["fee_rate"] = "test"

    dtoList := [1]interface{}{dto}
    return dtoList
}

