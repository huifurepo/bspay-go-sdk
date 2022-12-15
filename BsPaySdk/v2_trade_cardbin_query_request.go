/**
 * 卡bin信息查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeCardbinQueryRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    BankCardNoCrypt string `json:"bank_card_no_crypt" structs:"bank_card_no_crypt"` // 银行卡号密文

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeCardbinQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeCardbinQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeCardbinQueryRequest(reqParam)
}

func (bp *BsPay) V2TradeCardbinQueryRequest(reqParam V2TradeCardbinQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_CARDBIN_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
