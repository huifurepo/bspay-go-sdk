/**
 * 支付宝申诉查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantComplaintQueryStatusRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求汇付流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求汇付时间
    RiskBizId string `json:"risk_biz_id" structs:"risk_biz_id"` // 支付宝推送流水号
    BankMerCode string `json:"bank_mer_code" structs:"bank_mer_code"` // 申诉的商户

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantComplaintQueryStatusRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantComplaintQueryStatusRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantComplaintQueryStatusRequest(reqParam)
}

func (bp *BsPay) V2MerchantComplaintQueryStatusRequest(reqParam V2MerchantComplaintQueryStatusRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_COMPLAINT_QUERY_STATUS
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
