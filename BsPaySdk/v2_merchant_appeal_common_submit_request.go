/**
 * 提交申诉
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantAppealCommonSubmitRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    BusinessPattern string `json:"business_pattern" structs:"business_pattern"` // 商户经营模式
    AssistId string `json:"assist_id" structs:"assist_id"` // 协查单号
    AppealId string `json:"appeal_id" structs:"appeal_id"` // 申诉单号
    MerType string `json:"mer_type" structs:"mer_type"` // 商户类型
    AppealPersonName string `json:"appeal_person_name" structs:"appeal_person_name"` // 申诉人姓名
    AppealPersonCertNo string `json:"appeal_person_cert_no" structs:"appeal_person_cert_no"` // 申诉人身份证号
    AppealPersonPhoneNo string `json:"appeal_person_phone_no" structs:"appeal_person_phone_no"` // 申诉人联系电话
    LegalName string `json:"legal_name" structs:"legal_name"` // 法人姓名
    LegalCertNo string `json:"legal_cert_no" structs:"legal_cert_no"` // 法人身份证号
    LegalPhoneNo string `json:"legal_phone_no" structs:"legal_phone_no"` // 法人联系电话
    MainBusiness string `json:"main_business" structs:"main_business"` // 商户主营业务
    AppealDesc string `json:"appeal_desc" structs:"appeal_desc"` // 申诉理由

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantAppealCommonSubmitRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantAppealCommonSubmitRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantAppealCommonSubmitRequest(reqParam)
}

func (bp *BsPay) V2MerchantAppealCommonSubmitRequest(reqParam V2MerchantAppealCommonSubmitRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_APPEAL_COMMON_SUBMIT
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
