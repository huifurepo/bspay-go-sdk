/**
 * 微信特约商户进件 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantDirectWechatSignRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantDirectWechatSignRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantDirectWechatSignRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户汇付Id
        HuifuId:"6666000003134509",
        // 渠道商汇付Id
        UpperHuifuId:"6666000003134508",
        // 开发者的应用ID
        AppId:"20200506780602902",
        // 商户号
        MchId:"748851385",
        // 经营者/法人是否为受益人
        Owner:"Y",
        // 超级管理员信息
        ContactInfo:getContactInfo(),
        // 经营场景类型
        SalesScenesType:"SALES_SCENES_STORE,SALES_SCENES_MP,SALES_SCENES_MINI_PROGRAM,SALES_SCENES_WEB,SALES_SCENES_APP,SALES_SCENES_WEWORK",
        // 经营场景
        SalesInfo:getSalesInfo(),
        // 结算信息
        SettlementInfo:getSettlementInfo(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantDirectWechatSignRequest(dgReq)
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
    // 服务商微信公众号APPID对应密钥
    extendInfoMap["wx_woa_secret"] = "1234567890"
    // 公司类型
    extendInfoMap["ent_type"] = "2"
    // 登记证书
    extendInfoMap["certificate_info"] = getCertificateInfo()
    // 最终受益人信息
    extendInfoMap["ubo_info"] = getUboInfo()
    // 银行账户信息
    extendInfoMap["bank_account_info"] = getBankAccountInfo()
    // 补充说明
    extendInfoMap["business_addition_msg"] = "补充说明"
    // 文件列表
    extendInfoMap["file_list"] = getFileList()
    return extendInfoMap
}

func getSettlementInfo() string {
    dto := make(map[string]interface{})
    // 入驻结算规则ID
    dto["settlement_id"] = "716"
    // 所属行业
    dto["qualification_type"] = "餐饮"
    // 文件列表
    dto["file_list"] = getFileList()
    // 优惠费率活动ID
    dto["activities_id"] = "20191030111cff5b5e"
    // 优惠费率活动值
    dto["activities_rate"] = "0.60"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getBankAccountInfo() string {
    dto := make(map[string]interface{})
    // 账户类型
    dto["bank_account_type"] = "BANK_ACCOUNT_TYPE_CORPORATE"
    // 开户名称
    dto["account_name"] = "门店名称"
    // 开户银行
    dto["account_bank"] = "中国农业银行"
    // 开户银行省市编码
    dto["bank_address_code"] = "01030000"
    // 开户银行联行号1、17家直连银行无需填写，如为其他银行，则开户银行全称（含支行）和开户银行联行号二选一。&lt;br/&gt;2、详细参见[开户银行全称（含支行）对照表](https://pay.weixin.qq.com/wiki/doc/apiv3/terms_definition/chapter1_1_3.shtml#part-6)。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：102100020270&lt;/font&gt;
    dto["bank_branch_id"] = "102110001296"
    // 开户银行全称（含支行)1、17家直连银行无需填写，如为其他银行，则开户银行全称（含支行）和 开户银行联行号二选一。&lt;br/&gt;2、需填写银行全称，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：中国工商银行北京海运仓支行&lt;/font&gt;，详细参见[开户银行全称（含支行）对照表](https://pay.weixin.qq.com/wiki/doc/apiv3/terms_definition/chapter1_1_3.shtml#part-6)。
    dto["bank_name"] = "102110001296"
    // 银行账号
    dto["account_number"] = "102110001296"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getFileList() string {
    dto := make(map[string]interface{})
    // 文件类型
    dto["file_type"] = "F85"
    // 文件jfileID
    dto["file_id"] = "42204258-967e-373c-88d2-1afa4c7bb8ef"
    // 文件名称
    dto["file_name"] = "微信直连额外补充材料一"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getCertificateInfo() string {
    dto := make(map[string]interface{})
    // 登记证书类型
    dto["cert_type"] = "所有场景类型"
    // 证书号
    dto["cert_no"] = "证书号"
    // 证书商户名称
    dto["cert_mer_name"] = "证书商户名称"
    // 注册地址
    dto["reg_detail"] = "注册地址"
    // 法人姓名
    dto["legal_name"] = "法人姓名"
    // 证书有效期类型
    dto["cert_validity_type"] = "1"
    // 证书有效期开始日期
    dto["cert_begin_date"] = "20200420"
    // 文件列表
    // dto["file_list"] = getFileList()
    // 证书有效期截止日期
    dto["cert_end_date"] = "20400420"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getSalesInfo() string {
    dto := make(map[string]interface{})
    // 服务商公众号APPID公众号场景必传(与mp_sub_appid二选一) 。可填写当前服务商商户号已绑定的公众号APPID。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wx5934540532 &lt;/font&gt;
    dto["mp_appid"] = "服务商公众号APPID"
    // 商家公众号APPID公众号场景必传(与mp_appid二选一)。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wx5934540532 &lt;/font&gt; &lt;br/&gt;1、可填写与商家主体一致且已认证的公众号APPID，需是 已认证的服务号、政府或媒体类型的订阅号。&lt;br/&gt;2、审核通过后，系统将发起特约商家商户号与该AppID的 绑定（即配置为sub_appid），服务商随后可在发起支付时 选择传入该appid，以完成支付，并获取sub_openid用于数 据统计，营销等业务场景 。
    dto["mp_sub_appid"] = "商家公众号APPID"
    // 服务商小程序APPID小程序场景必传(与mmini_program_sub_appid二选一)。 可填写当前服务商商户号已绑定的小程序APPID。&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wx852a790f100000fe&lt;/font&gt;
    dto["mini_program_appid"] = "服务商小程序APPID"
    // 商家小程序APPID小程序场景必传(与mini_program_appid二选一)，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wx852a790f100000fe&lt;/font&gt; &lt;br/&gt;1、可填写与商家主体一致且已认证的小程序APPID，需是 已认证的小程序。&lt;br/&gt; 2、审核通过后，系统将发起特约商家商户号与该AppID的 绑定（即配置为sub_appid），服务商随后可在发起支付时 选择传入该appid，以完成支付，并获取sub_openid用于数 据统计，营销等业务场景。
    dto["mini_program_sub_appid"] = "商家小程序APPID"
    // 服务商应用APPIDAPP场景必传(与 app_sub_appid 二选一)。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wx852a790f100000fe&lt;/font&gt;&lt;br/&gt; 可填写当前服务商商户号已绑定的应用APPID。
    dto["app_appid"] = "服务商应用APPID"
    // 商家应用APPIDAPP场景必传(与app_appid二选一);&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wx852a790f100000fe&lt;/font&gt; &lt;br/&gt;1、可填写与商家主体一致且已认证的应用APPID，需是已 认证的APP。 &lt;br/&gt;2、审核通过后，系统将发起特约商家商户号与该AppID的 绑定（即配置为sub_appid），服务商随后可在发起支付时 选择传入该appid，以完成支付，并获取sub_openid用于数 据统计，营销等业务场景。
    dto["app_sub_appid"] = "商家应用APPID"
    // 文件列表
    dto["file_list"] = getFileList()
    // 门店名称
    dto["biz_store_name"] = "门店名称"
    // 门店省市编码
    dto["biz_address_code"] = "门店省市编码"
    // 门店地址
    dto["biz_store_address"] = "门店地址"
    // 线下场所对应的商家APPID
    dto["biz_sub_appid"] = "线下场所对应的商家APPID"
    // 互联网网站域名
    dto["domain"] = "互联网网站域名"
    // 互联网网站对应的商家APPID
    dto["web_appid"] = "互联网网站对应的商家APPID"
    // 商家企业微信CorpID
    dto["sub_corp_id"] = "商家企业微信CorpID"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getContactInfo() string {
    dto := make(map[string]interface{})
    // 超级管理员姓名
    dto["contact_name"] = "超级管理员姓名7586"
    // 超级管理员证件号码&lt;font color&#x3D;&quot;green&quot;&gt;示例值：320926198512025834&lt;/font&gt;&lt;br/&gt;1、“超级管理员证件号码”与“超级管理员微信openid”，二选一必填。&lt;br/&gt;2、超级管理员签约时，校验微信号绑定的银行卡实名信息，是否与该证件号码一致。&lt;br/&gt;3、可传身份证、来往内地通行证、来往大陆通行证、护照等证件号码。
    dto["contact_cert_no"] = "530102198206242725"
    // 超级管理员微信openid&lt;font color&#x3D;&quot;green&quot;&gt;示例值：oGhiSxIAPtEnPfe9Xo000000A&lt;/font&gt;&lt;br/&gt;1、“超级管理员身份证件号码”与“超级管理员微信openid”，二选一必填。&lt;br/&gt;2、超级管理员签约时，校验微信号是否与该微信openid一致。
    dto["openid"] = "openid"
    // 超级管理员手机号
    dto["contact_mobile_no"] = "13778851762"
    // 超级管理员电子邮箱
    dto["contact_email"] = "89007865@qq.com"
    // 超级管理员证件类型
    dto["cert_type"] = "00"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getUboInfo() string {
    dto := make(map[string]interface{})
    // 证件类型
    dto["cert_type"] = "00"
    // 证件号码
    dto["cert_no"] = "530102198206242725"
    // 姓名
    dto["name"] = "姓名"
    // 证件有效类型
    dto["cert_validity_type"] = "1"
    // 证件有效期开始日期
    dto["cert_begin_date"] = "20200420"
    // 文件列表
    dto["file_list"] = getFileList()
    // 证件有效期截止日期
    dto["cert_end_date"] = "20400420"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

