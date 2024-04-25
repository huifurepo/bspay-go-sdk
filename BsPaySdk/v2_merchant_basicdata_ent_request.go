/**
 * 企业商户基本信息入驻(2022)
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantBasicdataEntRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 直属渠道号
    RegName string `json:"reg_name" structs:"reg_name"` // 商户名称
    ShortName string `json:"short_name" structs:"short_name"` // 商户简称
    EntType string `json:"ent_type" structs:"ent_type"` // 公司类型
    LicenseCode string `json:"license_code" structs:"license_code"` // 营业执照编号
    LicenseValidityType string `json:"license_validity_type" structs:"license_validity_type"` // 营业执照有效期类型
    LicenseBeginDate string `json:"license_begin_date" structs:"license_begin_date"` // 营业执照有效期开始日期
    LicenseEndDate string `json:"license_end_date" structs:"license_end_date"` // 营业执照有效期截止日期日期格式：yyyyMMdd，以北京时间为准。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20220125&lt;/font&gt;&lt;br/&gt;  当license_validity_type&#x3D;0时必填  ;当license_validity_type&#x3D;1时为空；当使用总部资质时不填
    RegProvId string `json:"reg_prov_id" structs:"reg_prov_id"` // 注册省
    RegAreaId string `json:"reg_area_id" structs:"reg_area_id"` // 注册市
    RegDistrictId string `json:"reg_district_id" structs:"reg_district_id"` // 注册区
    RegDetail string `json:"reg_detail" structs:"reg_detail"` // 注册详细地址
    LegalName string `json:"legal_name" structs:"legal_name"` // 法人姓名
    LegalCertType string `json:"legal_cert_type" structs:"legal_cert_type"` // 法人证件类型
    LegalCertNo string `json:"legal_cert_no" structs:"legal_cert_no"` // 法人证件号码
    LegalCertValidityType string `json:"legal_cert_validity_type" structs:"legal_cert_validity_type"` // 法人证件有效期类型
    LegalCertBeginDate string `json:"legal_cert_begin_date" structs:"legal_cert_begin_date"` // 法人证件有效期开始日期
    LegalCertEndDate string `json:"legal_cert_end_date" structs:"legal_cert_end_date"` // 法人证件有效期截止日期日期格式：yyyyMMdd，以北京时间为准。  &lt;font color&#x3D;&quot;green&quot;&gt;示例值：20220125&lt;/font&gt;&lt;br/&gt;当legal_cert_validity_type&#x3D;0时必填 ；当legal_cert_validity_type&#x3D;1时为空 ；当使用总部资质时不填
    ProvId string `json:"prov_id" structs:"prov_id"` // 经营省
    AreaId string `json:"area_id" structs:"area_id"` // 经营市
    DistrictId string `json:"district_id" structs:"district_id"` // 经营区
    DetailAddr string `json:"detail_addr" structs:"detail_addr"` // 经营详细地址
    ContactName string `json:"contact_name" structs:"contact_name"` // 联系人姓名
    ContactMobileNo string `json:"contact_mobile_no" structs:"contact_mobile_no"` // 联系人手机号
    ContactEmail string `json:"contact_email" structs:"contact_email"` // 联系人电子邮箱
    ServicePhone string `json:"service_phone" structs:"service_phone"` // 客服电话
    BusiType string `json:"busi_type" structs:"busi_type"` // 经营类型
    ReceiptName string `json:"receipt_name" structs:"receipt_name"` // 小票名称
    Mcc string `json:"mcc" structs:"mcc"` // 所属行业
    CardInfo string `json:"card_info" structs:"card_info"` // 结算卡信息配置
    OpenLicenceNo string `json:"open_licence_no" structs:"open_licence_no"` // 基本存款账户编号或核准号基本存款账户信息请填写基本存款账户编号；开户许可证请填写核准号 ；&lt;br/&gt;当注册地址或经营地址为如下地区时必填：浙江,海南,重庆,河南,江苏,宁波市,黑龙江,吉林,湖南,贵州,陕西,湖北 &lt;br/&gt;当使用总部资质时不填 ；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：J2900123456789&lt;/font&gt;
    HeadHuifuId string `json:"head_huifu_id" structs:"head_huifu_id"` // 总部汇付Id如果headOfficeFlag&#x3D;0，useHeadInfoFlag&#x3D;Y,且head_huifu_id不为空则基本信息部分复用总部的基本信息。&lt;br/&gt;如果head_office_flag&#x3D;0，则字段必填,如果head_office_flag&#x3D;1，总部汇付Id不可传&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000123123123&lt;/font&gt;

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantBasicdataEntRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantBasicdataEntRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantBasicdataEntRequest(reqParam)
}

func (bp *BsPay) V2MerchantBasicdataEntRequest(reqParam V2MerchantBasicdataEntRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_BASICDATA_ENT
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
