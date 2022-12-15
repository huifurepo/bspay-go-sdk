/**
 * 结算记录查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantBasicdataSettlementQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id
    BeginDate string `json:"begin_date" structs:"begin_date"` // 结算开始日期
    EndDate string `json:"end_date" structs:"end_date"` // 结算结束日期
    PageSize string `json:"page_size" structs:"page_size"` // 分页条数

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantBasicdataSettlementQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantBasicdataSettlementQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantBasicdataSettlementQueryRequest(reqParam)
}

func (bp *BsPay) V2MerchantBasicdataSettlementQueryRequest(reqParam V2MerchantBasicdataSettlementQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_BASICDATA_SETTLEMENT_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
