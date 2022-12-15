/**
 * 直付通商户入驻
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantDirectZftRegRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付ID
    AppId string `json:"app_id" structs:"app_id"` // 开发者的应用ID
    Name string `json:"name" structs:"name"` // 进件的二级商户名称
    MerchantType string `json:"merchant_type" structs:"merchant_type"` // 商家类型
    Mcc string `json:"mcc" structs:"mcc"` // 商户经营类目
    CertType string `json:"cert_type" structs:"cert_type"` // 商户证件类型
    CertNo string `json:"cert_no" structs:"cert_no"` // 商户证件编号
    CertName string `json:"cert_name" structs:"cert_name"` // 证件名称目前只有个体工商户商户类型要求填入本字段，填写值为个体工商户营业执照上的名称。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：xxxx小卖铺&lt;/font&gt;
    LegalName string `json:"legal_name" structs:"legal_name"` // 法人名称仅个人商户非必填，其他必填。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：张三&lt;/font&gt;
    LegalCertNo string `json:"legal_cert_no" structs:"legal_cert_no"` // 法人证件号码仅个人商户非必填，其他必填。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：3209261975120284333&lt;/font&gt;
    ServicePhone string `json:"service_phone" structs:"service_phone"` // 客服电话
    ProvId string `json:"prov_id" structs:"prov_id"` // 经营省
    AreaId string `json:"area_id" structs:"area_id"` // 经营市
    DistrictId string `json:"district_id" structs:"district_id"` // 经营区
    DetailAddr string `json:"detail_addr" structs:"detail_addr"` // 经营详细地址
    ContactName string `json:"contact_name" structs:"contact_name"` // 联系人姓名
    ContactTag string `json:"contact_tag" structs:"contact_tag"` // 商户联系人业务标识
    ContactType string `json:"contact_type" structs:"contact_type"` // 联系人类型
    ContactMobileNo string `json:"contact_mobile_no" structs:"contact_mobile_no"` // 联系人手机号
    ZftCardInfoList string `json:"zft_card_info_list" structs:"zft_card_info_list"` // 商户结算卡信息jsonArray格式。本业务当前只允许传入一张结算卡。与支付宝账号字段二选一必填
    AlipayLogonId string `json:"alipay_logon_id" structs:"alipay_logon_id"` // 商户支付宝账号商户支付宝账号，用作结算账号。与银行卡对象字段二选一必填。&lt;br/&gt;本字段要求支付宝账号的名称与商户名称mch_name同名，且是实名认证过的支付宝账户。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：test@huifu.com&lt;/font&gt;
    IndustryQualificationType string `json:"industry_qualification_type" structs:"industry_qualification_type"` // 商户行业资质类型当商户是特殊行业时必填，具体取值[参见表格](https://mif-pub.alipayobjects.com/QualificationType.xlsx)。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：310&lt;/font&gt;
    Service string `json:"service" structs:"service"` // 商户使用服务
    SignTimeWithIsv string `json:"sign_time_with_isv" structs:"sign_time_with_isv"` // 商户与服务商的签约时间
    BindingAlipayLogonId string `json:"binding_alipay_logon_id" structs:"binding_alipay_logon_id"` // 商户支付宝账户用于协议确认。目前商业场景（除医疗、中小学教育等）下必填。本字段要求上送的支付宝账号的名称与商户名称name同名，且是实名认证支付宝账户。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：test@huifu.com&lt;/font&gt;
    DefaultSettleType string `json:"default_settle_type" structs:"default_settle_type"` // 默认结算类型
    FileList string `json:"file_list" structs:"file_list"` // 文件列表

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantDirectZftRegRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantDirectZftRegRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantDirectZftRegRequest(reqParam)
}

func (bp *BsPay) V2MerchantDirectZftRegRequest(reqParam V2MerchantDirectZftRegRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_DIRECT_ZFT_REG
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
