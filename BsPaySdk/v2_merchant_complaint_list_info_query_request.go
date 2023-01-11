/**
 * 查询投诉单列表及详情
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantComplaintListInfoQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    BeginDate string `json:"begin_date" structs:"begin_date"` // 开始日期
    EndDate string `json:"end_date" structs:"end_date"` // 结束日期

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantComplaintListInfoQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantComplaintListInfoQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantComplaintListInfoQueryRequest(reqParam)
}

func (bp *BsPay) V2MerchantComplaintListInfoQueryRequest(reqParam V2MerchantComplaintListInfoQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_COMPLAINT_LIST_INFO_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
