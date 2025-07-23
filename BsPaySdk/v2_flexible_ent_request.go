/**
 * 灵工企业商户进件接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2FlexibleEntRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 渠道商号
    MerRole string `json:"mer_role" structs:"mer_role"` // 商户角色
    LocalCompanyType string `json:"local_company_type" structs:"local_company_type"` // 落地公司类型当选择落地公司时，必填;SELF-自有，COOPERATE-合作
    RegName string `json:"reg_name" structs:"reg_name"` // 商户名称
    ShortName string `json:"short_name" structs:"short_name"` // 商户简称
    LicensePic string `json:"license_pic" structs:"license_pic"` // 证照图片
    LicenseCode string `json:"license_code" structs:"license_code"` // 证照编号
    LicenseValidityType string `json:"license_validity_type" structs:"license_validity_type"` // 证照有效期类型
    LicenseBeginDate string `json:"license_begin_date" structs:"license_begin_date"` // 证照有效期开始日期
    LicenseEndDate string `json:"license_end_date" structs:"license_end_date"` // 证照有效期截止日期格式：yyyyMMdd。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20220125&lt;/font&gt;&lt;br/&gt;  当license_validity_type&#x3D;0时必填；当license_validity_type&#x3D;1时为空；
    FoundDate string `json:"found_date" structs:"found_date"` // 成立时间
    RegDistrictId string `json:"reg_district_id" structs:"reg_district_id"` // 注册区
    RegDetail string `json:"reg_detail" structs:"reg_detail"` // 注册详细地址
    LegalName string `json:"legal_name" structs:"legal_name"` // 法人姓名
    LegalCertType string `json:"legal_cert_type" structs:"legal_cert_type"` // 法人证件类型
    LegalCertNo string `json:"legal_cert_no" structs:"legal_cert_no"` // 法人证件号码
    LegalCertValidityType string `json:"legal_cert_validity_type" structs:"legal_cert_validity_type"` // 法人证件有效期类型
    LegalCertBeginDate string `json:"legal_cert_begin_date" structs:"legal_cert_begin_date"` // 法人证件有效期开始日期
    LegalCertEndDate string `json:"legal_cert_end_date" structs:"legal_cert_end_date"` // 法人证件有效期截止日期日期格式：yyyyMMdd， &lt;font color&#x3D;&quot;green&quot;&gt;示例值：20220125&lt;/font&gt;&lt;br/&gt;当legal_cert_validity_type&#x3D;0时必填；&lt;br/&gt;当legal_cert_validity_type&#x3D;1时为空；&lt;br/&gt;
    LegalAddr string `json:"legal_addr" structs:"legal_addr"` // 法人证件地址
    LegalCertBackPic string `json:"legal_cert_back_pic" structs:"legal_cert_back_pic"` // 法人身份证国徽面
    LegalCertFrontPic string `json:"legal_cert_front_pic" structs:"legal_cert_front_pic"` // 法人身份证人像面
    StoreHeaderPic string `json:"store_header_pic" structs:"store_header_pic"` // 店铺门头照
    ContactMobileNo string `json:"contact_mobile_no" structs:"contact_mobile_no"` // 联系人手机号
    ContactEmail string `json:"contact_email" structs:"contact_email"` // 联系人电子邮箱
    LoginName string `json:"login_name" structs:"login_name"` // 管理员账号
    CardInfo string `json:"card_info" structs:"card_info"` // 银行卡信息配置
    SignUserInfo string `json:"sign_user_info" structs:"sign_user_info"` // 签约人

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2FlexibleEntRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2FlexibleEntRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2FlexibleEntRequest(reqParam)
}

func (bp *BsPay) V2FlexibleEntRequest(reqParam V2FlexibleEntRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_FLEXIBLE_ENT
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
