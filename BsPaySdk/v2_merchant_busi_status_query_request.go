/**
 * 统一进件页面版查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantBusiStatusQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    StoreId string `json:"store_id" structs:"store_id"` // 门店号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantBusiStatusQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantBusiStatusQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantBusiStatusQueryRequest(reqParam)
}

func (bp *BsPay) V2MerchantBusiStatusQueryRequest(reqParam V2MerchantBusiStatusQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_BUSI_STATUS_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
