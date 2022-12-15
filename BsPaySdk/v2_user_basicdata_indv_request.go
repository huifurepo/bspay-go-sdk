/**
 * 个人用户基本信息开户
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2UserBasicdataIndvRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    Name string `json:"name" structs:"name"` // 个人姓名
    CertType string `json:"cert_type" structs:"cert_type"` // 个人证件类型
    CertNo string `json:"cert_no" structs:"cert_no"` // 个人证件号码
    CertValidityType string `json:"cert_validity_type" structs:"cert_validity_type"` // 个人证件有效期类型
    CertBeginDate string `json:"cert_begin_date" structs:"cert_begin_date"` // 个人证件有效期开始日期
    MobileNo string `json:"mobile_no" structs:"mobile_no"` // 手机号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2UserBasicdataIndvRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2UserBasicdataIndvRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2UserBasicdataIndvRequest(reqParam)
}

func (bp *BsPay) V2UserBasicdataIndvRequest(reqParam V2UserBasicdataIndvRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_USER_BASICDATA_INDV
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
