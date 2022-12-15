/**
 * 余额支付
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeAcctpaymentPayRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    OutHuifuId string `json:"out_huifu_id" structs:"out_huifu_id"` // 出款方商户号
    OrdAmt string `json:"ord_amt" structs:"ord_amt"` // 支付金额
    AcctSplitBunch string `json:"acct_split_bunch" structs:"acct_split_bunch"` // 分账对象
    RiskCheckData string `json:"risk_check_data" structs:"risk_check_data"` // 安全信息

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeAcctpaymentPayRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeAcctpaymentPayRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeAcctpaymentPayRequest(reqParam)
}

func (bp *BsPay) V2TradeAcctpaymentPayRequest(reqParam V2TradeAcctpaymentPayRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ACCTPAYMENT_PAY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
