/**
 * 不明来账列表查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePaymentZxeUnknownincomeQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    TransStartDate string `json:"trans_start_date" structs:"trans_start_date"` // 交易开始日期
    TransEndDate string `json:"trans_end_date" structs:"trans_end_date"` // 交易结束日期

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePaymentZxeUnknownincomeQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePaymentZxeUnknownincomeQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePaymentZxeUnknownincomeQueryRequest(reqParam)
}

func (bp *BsPay) V2TradePaymentZxeUnknownincomeQueryRequest(reqParam V2TradePaymentZxeUnknownincomeQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYMENT_ZXE_UNKNOWNINCOME_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
