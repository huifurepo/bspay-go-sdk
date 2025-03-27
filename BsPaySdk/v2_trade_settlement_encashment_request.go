/**
 * 取现
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeSettlementEncashmentRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    CashAmt string `json:"cash_amt" structs:"cash_amt"` // 取现金额
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 取现方ID号
    IntoAcctDateType string `json:"into_acct_date_type" structs:"into_acct_date_type"` // 到账日期类型
    TokenNo string `json:"token_no" structs:"token_no"` // 取现卡序列号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeSettlementEncashmentRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeSettlementEncashmentRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeSettlementEncashmentRequest(reqParam)
}

func (bp *BsPay) V2TradeSettlementEncashmentRequest(reqParam V2TradeSettlementEncashmentRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_SETTLEMENT_ENCASHMENT
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
