/**
 * 电子账户资金清分结果查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeSettlementClearingQueryRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    FileId string `json:"file_id" structs:"file_id"` // 清分文件ID

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeSettlementClearingQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeSettlementClearingQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeSettlementClearingQueryRequest(reqParam)
}

func (bp *BsPay) V2TradeSettlementClearingQueryRequest(reqParam V2TradeSettlementClearingQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_SETTLEMENT_CLEARING_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
