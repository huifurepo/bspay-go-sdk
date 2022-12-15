/**
 * 账户余额信息查询接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeAcctpaymentBalanceQueryRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeAcctpaymentBalanceQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeAcctpaymentBalanceQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeAcctpaymentBalanceQueryRequest(reqParam)
}

func (bp *BsPay) V2TradeAcctpaymentBalanceQueryRequest(reqParam V2TradeAcctpaymentBalanceQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ACCTPAYMENT_BALANCE_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
