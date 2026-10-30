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
        // 直属渠道号该商户进件完成后归属的渠道商huifu_id；；&lt;br/&gt;sys_id主体为渠道商时填写； &lt;br/&gt;sys_id主体为总部商户时选填；&lt;br/&gt; &lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000123123123&lt;/font&gt;
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
    // extendInfoMap["ali_conf_list"] = get3b39d92eC1c74d6dA4fbC29f0fc5de36()
    // 微信配置对象
    // extendInfoMap["wx_conf_list"] = getBe2c2ad8B8ff43bf98223776eafe76f6()
    // 银联二维码配置对象
    // extendInfoMap["union_conf_list"] = getA0588272Dd9543ef86e7B472eb82c55a()
    // 银联卡配置对象
    // extendInfoMap["bank_card_config"] = get13430b5bDfab4634992748b527981508()
    // 分账配置对象
    // extendInfoMap["split_config"] = getBc4c9bcbF9b942889486Cdbc5c8c9f1c()
    // 微信直连配置对象
    // extendInfoMap["wx_zl_conf_list"] = get7a42976cE3da4575B6feDd8932994c14()
    // 支付宝直连配置对象
    // extendInfoMap["ali_zl_conf"] = get622dcd3a7b7b45c38de2919594c16f41()
    // 线上配置对象
    // extendInfoMap["online_fee_conf_list"] = get0f786eef05944bf68fb5C92f5906a96a()
    // 余额支付配置对象
    // extendInfoMap["balance_pay_config"] = get0e160e96Cb17429686a4365e5d968f57()
    // 补贴支付配置对象
    // extendInfoMap["combine_pay_config"] = get0532062b23d34359A699Ef47a9c75acd()
    // 银行大额转账配置对象
    // extendInfoMap["bank_big_amt_pay_config"] = get5c94bd76B25040878755136ea12a78de()
    // 全域资金管理配置对象（华通银行）
    // extendInfoMap["out_order_funds_config"] = getC5eb39eeCa584d29B09e60103d064fb7()
    // 全域资金管理配置(XW银行)
    // extendInfoMap["out_order_funds_new_net_config"] = getAa4bcb7a48f0433eAda10206f4814f47()
    // 结算配置对象
    // extendInfoMap["settle_config_list"] = get0c4853b267924652B52e76b3d0fea998()
    // 取现配置对象
    // extendInfoMap["cash_config_list"] = get6a265ef9E88e4c008176975ce5c72a35()
    // 外扣配置对象
    // extendInfoMap["out_fee_config"] = get3a37f3557ff64fe6A11e8f0db37e709e()
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
    // extendInfoMap["large_amt_pay_config_list"] = get0e88fc4d2ab44c46A447F445fc6e6694()
    // 全域资金管理配置(苏商)
    // extendInfoMap["out_order_funds_su_shang_config"] = getD181c95e0e114a4eB294659df51793cc()
    // 统一收银台开关
    // extendInfoMap["half_pay_host_flag"] = ""
    // 全域资金费用配置对象
    // extendInfoMap["out_order_funds_fee_list"] = getE7e28aef95074e0c8886206a279a8cbb()
    // 本地生活生活配置对象
    // extendInfoMap["lla_withhold_config"] = get8f6fea3524f041fcB92e1b3200c1414b()
    // 是否开通账单支付
    // extendInfoMap["bill_payment_flag"] = ""
    // 抖音配置对象
    // extendInfoMap["dy_conf_list"] = get53df1aab578943b7B08cC1524584ea1d()
    return extendInfoMap
}

func get3b39d92eC1c74d6dA4fbC29f0fc5de36() string {
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

func getBe2c2ad8B8ff43bf98223776eafe76f6() string {
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

func getA0588272Dd9543ef86e7B472eb82c55a() string {
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

func get13430b5bDfab4634992748b527981508() string {
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

func getBc4c9bcbF9b942889486Cdbc5c8c9f1c() string {
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

func get7a42976cE3da4575B6feDd8932994c14() string {
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

func get622dcd3a7b7b45c38de2919594c16f41() string {
    dto := make(map[string]interface{})
    // 手续费（%）
    // dto["fee_rate"] = "test"
    // 允许开通支付宝直连业务
    // dto["open_flag"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get0f786eef05944bf68fb5C92f5906a96a() string {
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
    // 代扣手续费计费模式
    // dto["fee_formula_type"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get0e160e96Cb17429686a4365e5d968f57() string {
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

func get0532062b23d34359A699Ef47a9c75acd() string {
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

func get5c94bd76B25040878755136ea12a78de() string {
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

func getC5eb39eeCa584d29B09e60103d064fb7() string {
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

func getAa4bcb7a48f0433eAda10206f4814f47() string {
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

func get0c4853b267924652B52e76b3d0fea998() string {
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

func get6a265ef9E88e4c008176975ce5c72a35() string {
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

func get3a37f3557ff64fe6A11e8f0db37e709e() string {
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

func get0e88fc4d2ab44c46A447F445fc6e6694() string {
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

func getD181c95e0e114a4eB294659df51793cc() string {
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

func getE7e28aef95074e0c8886206a279a8cbb() string {
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

func get8f6fea3524f041fcB92e1b3200c1414b() interface{} {
    dto := make(map[string]interface{})
    // 本地生活开关
    // dto["llaWithholdFlag"] = "test"
    // 佣金收取手续费率
    // dto["fee_rate"] = "test"

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get53df1aab578943b7B08cC1524584ea1d() string {
    dto := make(map[string]interface{})
    // 抖音商户行业分类
    // dto["mer_industry_category"] = "test"
    // 手续费模式
    // dto["fee_model"] = "test"
    // 手续费（%）
    // dto["fee_rate"] = "test"
    // 最低收取手续费（元）
    // dto["fee_min_amt"] = ""
    // 是否允许开通
    // dto["open_flag"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

