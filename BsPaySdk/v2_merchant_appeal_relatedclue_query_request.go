/**
 * 关联线索查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantAppealRelatedclueQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    PageSize string `json:"page_size" structs:"page_size"` // 分页条数
    AssistId string `json:"assist_id" structs:"assist_id"` // 协查单号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantAppealRelatedclueQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantAppealRelatedclueQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantAppealRelatedclueQueryRequest(reqParam)
}

func (bp *BsPay) V2MerchantAppealRelatedclueQueryRequest(reqParam V2MerchantAppealRelatedclueQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_APPEAL_RELATEDCLUE_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
