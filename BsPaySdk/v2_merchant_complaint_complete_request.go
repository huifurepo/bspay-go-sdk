/**
 * 反馈处理完成
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantComplaintCompleteRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    ComplaintId string `json:"complaint_id" structs:"complaint_id"` // 微信投诉单号
    ComplaintedMchid string `json:"complainted_mchid" structs:"complainted_mchid"` // 被诉商户微信号
    MchId string `json:"mch_id" structs:"mch_id"` // 微信商户号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantComplaintCompleteRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantComplaintCompleteRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantComplaintCompleteRequest(reqParam)
}

func (bp *BsPay) V2MerchantComplaintCompleteRequest(reqParam V2MerchantComplaintCompleteRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_COMPLAINT_COMPLETE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
