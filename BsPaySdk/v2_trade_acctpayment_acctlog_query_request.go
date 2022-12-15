/**
 * 账务流水查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeAcctpaymentAcctlogQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 渠道/代理/商户/用户编号
    AcctDate string `json:"acct_date" structs:"acct_date"` // 账务日期

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeAcctpaymentAcctlogQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeAcctpaymentAcctlogQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeAcctpaymentAcctlogQueryRequest(reqParam)
}

func (bp *BsPay) V2TradeAcctpaymentAcctlogQueryRequest(reqParam V2TradeAcctpaymentAcctlogQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ACCTPAYMENT_ACCTLOG_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
