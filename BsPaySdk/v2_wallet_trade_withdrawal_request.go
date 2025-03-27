/**
 * 钱包提现下单
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2WalletTradeWithdrawalRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    UserHuifuId string `json:"user_huifu_id" structs:"user_huifu_id"` // 钱包用户ID
    TokenNo string `json:"token_no" structs:"token_no"` // 银行卡序列号
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 提现金额
    FrontUrl string `json:"front_url" structs:"front_url"` // 跳转地址
    NotifyUrl string `json:"notify_url" structs:"notify_url"` // 异步通知地址
    IntoAcctDateType string `json:"into_acct_date_type" structs:"into_acct_date_type"` // 到账日期类型

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2WalletTradeWithdrawalRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2WalletTradeWithdrawalRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2WalletTradeWithdrawalRequest(reqParam)
}

func (bp *BsPay) V2WalletTradeWithdrawalRequest(reqParam V2WalletTradeWithdrawalRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_WALLET_TRADE_WITHDRAWAL
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
