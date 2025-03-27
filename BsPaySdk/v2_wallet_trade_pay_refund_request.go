/**
 * 钱包支付退款
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2WalletTradePayRefundRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    UserHuifuId string `json:"user_huifu_id" structs:"user_huifu_id"` // 钱包用户ID
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 退款金额
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原交易请求日期

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2WalletTradePayRefundRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2WalletTradePayRefundRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2WalletTradePayRefundRequest(reqParam)
}

func (bp *BsPay) V2WalletTradePayRefundRequest(reqParam V2WalletTradePayRefundRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_WALLET_TRADE_PAY_REFUND
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
