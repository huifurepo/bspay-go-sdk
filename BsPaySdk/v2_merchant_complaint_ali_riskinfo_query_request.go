/**
 * 支付宝投诉查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantComplaintAliRiskinfoQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求汇付流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求汇付时间
    BeginDate string `json:"begin_date" structs:"begin_date"` // 开始日期
    EndDate string `json:"end_date" structs:"end_date"` // 结束日期

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantComplaintAliRiskinfoQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantComplaintAliRiskinfoQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantComplaintAliRiskinfoQueryRequest(reqParam)
}

func (bp *BsPay) V2MerchantComplaintAliRiskinfoQueryRequest(reqParam V2MerchantComplaintAliRiskinfoQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_COMPLAINT_ALI_RISKINFO_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
