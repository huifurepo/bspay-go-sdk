/**
 * 用户列表查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2UserListQueryRequest struct {
    LegalCertNo string `json:"legal_cert_no" structs:"legal_cert_no"` // 法人证件号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2UserListQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2UserListQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2UserListQueryRequest(reqParam)
}

func (bp *BsPay) V2UserListQueryRequest(reqParam V2UserListQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_USER_LIST_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
