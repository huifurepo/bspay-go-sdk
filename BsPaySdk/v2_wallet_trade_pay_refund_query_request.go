/**
 * 钱包支付退款查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2WalletTradePayRefundQueryRequest struct {
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原退款交易请求日期
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 原退款交易请求流水号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2WalletTradePayRefundQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2WalletTradePayRefundQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2WalletTradePayRefundQueryRequest(reqParam)
}

func (bp *BsPay) V2WalletTradePayRefundQueryRequest(reqParam V2WalletTradePayRefundQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_WALLET_TRADE_PAY_REFUND_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
