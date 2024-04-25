/**
 * 钱包支付下单
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2WalletTradePayBalanceRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    UserHuifuId string `json:"user_huifu_id" structs:"user_huifu_id"` // 钱包用户ID
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 订单金额
    FrontUrl string `json:"front_url" structs:"front_url"` // 跳转地址

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2WalletTradePayBalanceRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2WalletTradePayBalanceRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2WalletTradePayBalanceRequest(reqParam)
}

func (bp *BsPay) V2WalletTradePayBalanceRequest(reqParam V2WalletTradePayBalanceRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_WALLET_TRADE_PAY_BALANCE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
