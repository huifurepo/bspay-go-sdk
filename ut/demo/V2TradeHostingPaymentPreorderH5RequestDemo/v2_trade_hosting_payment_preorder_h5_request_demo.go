/**
 * H5、PC预下单接口 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeHostingPaymentPreorderH5RequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeHostingPaymentPreorderH5RequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeHostingPaymentPreorderH5Request{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000111546360",
        // 交易金额
        TransAmt:"0.10",
        // 商品描述
        GoodsDesc:"支付托管消费",
        // 预下单类型
        PreOrderType:"1",
        // 半支付托管扩展参数集合
        HostingData:getHostingData(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeHostingPaymentPreorderH5Request(dgReq)
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
    // 收银台样式
    // extendInfoMap["style_id"] = ""
    // 是否延迟交易
    extendInfoMap["delay_acct_flag"] = "N"
    // 分账对象
    extendInfoMap["acct_split_bunch"] = getAcctSplitBunch()
    // 交易失效时间
    // extendInfoMap["time_expire"] = ""
    // 业务信息
    extendInfoMap["biz_info"] = getBizInfo()
    // 交易异步通知地址
    extendInfoMap["notify_url"] = "https://callback.service.com/xx"
    // 使用类型
    // extendInfoMap["usage_type"] = ""
    // 交易类型
    // extendInfoMap["trans_type"] = ""
    return extendInfoMap
}

func getAcctInfosRucan() interface{} {
    dto := make(map[string]interface{})
    // 分账金额
    dto["div_amt"] = "0.08"
    // 分账接收方ID
    dto["huifu_id"] = "6666000111546360"

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getAcctSplitBunch() string {
    dto := make(map[string]interface{})
    // 分账明细
    dto["acct_infos"] = getAcctInfosRucan()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getHostingData() string {
    dto := make(map[string]interface{})
    // 项目标题
    dto["project_title"] = "收银台标题"
    // 半支付托管项目号
    dto["project_id"] = "PROJECTID2022032912492559"
    // 请求类型P:PC页面版，默认：P；M:H5页面版；指定交易类型时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：M&lt;/font&gt;
    // dto["request_type"] = "test"
    // 商户私有信息
    dto["private_info"] = "商户私有信息test"
    // 回调地址
    dto["callback_url"] = "https://paas.huifu.com"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getPayerCheckAli() interface{} {
    dto := make(map[string]interface{})
    // 是否提供校验身份信息
    dto["need_check_info"] = "T"
    // 允许的最小买家年龄
    dto["min_age"] = "12"
    // 是否强制校验付款人身份信息
    dto["fix_buyer"] = "F"

    return dto;
}

func getPayerCheckWx() interface{} {
    dto := make(map[string]interface{})
    // 指定支付者
    dto["limit_payer"] = "ADULT"
    // 微信实名验证
    dto["real_name_flag"] = "Y"

    return dto;
}

func getPersonPayer() interface{} {
    dto := make(map[string]interface{})
    // 姓名
    dto["name"] = "张三"
    // 证件类型
    dto["cert_type"] = "IDENTITY_CARD"
    // 证件号
    dto["cert_no"] = "Mc5pjf+b/Keyi/t/wnHJtJYPHd1xXntq6tau0j8SjLzJx+q2xL2mOmKRDAYHu4uY1JSoPbWBhq9b7gT7Kxb1CYnkj7vmSlTYl8tVKfOPFyauOE66ew9cmkhmUzjzVTM1quoR63pP8+ESvZZrRPFE4YY9PXO9It9JINo8bjX22fQEFZKmXaEcqnSDcl2LUuJguvQ0LejI6zbxCJhfSHbz7HhHTIZTUchkWpKoy8YlfG27FumjXHU3rIjbrgmc+8pXbyndTNlui1+lTu6deibGKq/CpShA8z5FkHsn6/1O9ZEjLcnPnSLUwCnu75UlVVk66g+hR1OGdRrFMfYQnK7Lzw=="
    // 手机号
    dto["mobile"] = "15012345678"

    return dto;
}

func getBizInfo() string {
    dto := make(map[string]interface{})
    // 付款人验证（支付宝）
    dto["payer_check_ali"] = getPayerCheckAli()
    // 付款人验证（微信）
    dto["payer_check_wx"] = getPayerCheckWx()
    // 个人付款人信息
    dto["person_payer"] = getPersonPayer()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

