/**
 * 申诉客户信息查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantAppealCustinfoQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    AssistId string `json:"assist_id" structs:"assist_id"` // 协查单号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantAppealCustinfoQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantAppealCustinfoQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantAppealCustinfoQueryRequest(reqParam)
}

func (bp *BsPay) V2MerchantAppealCustinfoQueryRequest(reqParam V2MerchantAppealCustinfoQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_APPEAL_CUSTINFO_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
