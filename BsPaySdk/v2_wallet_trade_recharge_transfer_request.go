/**
 * 用户补贴
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2WalletTradeRechargeTransferRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 出款方商户号
    UserHuifuId string `json:"user_huifu_id" structs:"user_huifu_id"` // 收款方用户号
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 转账金额

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2WalletTradeRechargeTransferRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2WalletTradeRechargeTransferRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2WalletTradeRechargeTransferRequest(reqParam)
}

func (bp *BsPay) V2WalletTradeRechargeTransferRequest(reqParam V2WalletTradeRechargeTransferRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_WALLET_TRADE_RECHARGE_TRANSFER
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
