/**
 * 企业用户基本信息开户
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2UserBasicdataEntRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    RegName string `json:"reg_name" structs:"reg_name"` // 企业用户名称
    LicenseCode string `json:"license_code" structs:"license_code"` // 营业执照编号
    LicenseValidityType string `json:"license_validity_type" structs:"license_validity_type"` // 证照有效期类型
    LicenseBeginDate string `json:"license_begin_date" structs:"license_begin_date"` // 证照有效期起始日期
    LicenseEndDate string `json:"license_end_date" structs:"license_end_date"` // 证照有效期结束日期日期格式：yyyyMMdd; 非长期有效时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20320905&lt;/font&gt;
    RegProvId string `json:"reg_prov_id" structs:"reg_prov_id"` // 注册地址(省)
    RegAreaId string `json:"reg_area_id" structs:"reg_area_id"` // 注册地址(市)
    RegDistrictId string `json:"reg_district_id" structs:"reg_district_id"` // 注册地址(区)
    RegDetail string `json:"reg_detail" structs:"reg_detail"` // 注册地址(详细信息)
    LegalName string `json:"legal_name" structs:"legal_name"` // 法人姓名
    LegalCertType string `json:"legal_cert_type" structs:"legal_cert_type"` // 法人证件类型
    LegalCertNo string `json:"legal_cert_no" structs:"legal_cert_no"` // 法人证件号码
    LegalCertValidityType string `json:"legal_cert_validity_type" structs:"legal_cert_validity_type"` // 法人证件有效期类型
    LegalCertBeginDate string `json:"legal_cert_begin_date" structs:"legal_cert_begin_date"` // 法人证件有效期开始日期
    LegalCertEndDate string `json:"legal_cert_end_date" structs:"legal_cert_end_date"` // 法人证件有效期截止日期日期格式：yyyyMMdd; 非长期有效时必填，长期有效为空；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20320905&lt;/font&gt;
    LegalCertNationality string `json:"legal_cert_nationality" structs:"legal_cert_nationality"` // 法人国籍法人的证件类型为外国人居留证时，必填，参见《[国籍编码](https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/opps/api/prod/download_file/area/%E5%9B%BD%E7%B1%8D.xlsx)》&lt;font color&#x3D;&quot;green&quot;&gt;示例值：CHN&lt;/font&gt;
    ContactName string `json:"contact_name" structs:"contact_name"` // 联系人姓名
    ContactMobile string `json:"contact_mobile" structs:"contact_mobile"` // 联系人手机号
    LoginName string `json:"login_name" structs:"login_name"` // 管理员账号如需短信通知则必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：Lg20220222013747&lt;/font&gt;

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2UserBasicdataEntRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2UserBasicdataEntRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2UserBasicdataEntRequest(reqParam)
}

func (bp *BsPay) V2UserBasicdataEntRequest(reqParam V2UserBasicdataEntRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_USER_BASICDATA_ENT
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
