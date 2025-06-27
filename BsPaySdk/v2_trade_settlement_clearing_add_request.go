/**
 * 电子账户资金清分
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeSettlementClearingAddRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    FileId string `json:"file_id" structs:"file_id"` // 清分文件ID
    TransDate string `json:"trans_date" structs:"trans_date"` // 交易日期

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeSettlementClearingAddRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeSettlementClearingAddRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeSettlementClearingAddRequest(reqParam)
}

func (bp *BsPay) V2TradeSettlementClearingAddRequest(reqParam V2TradeSettlementClearingAddRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_SETTLEMENT_CLEARING_ADD
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
