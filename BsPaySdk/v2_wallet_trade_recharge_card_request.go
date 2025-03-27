/**
 * 钱包绑卡充值下单
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2WalletTradeRechargeCardRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    UserHuifuId string `json:"user_huifu_id" structs:"user_huifu_id"` // 钱包用户ID
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 订单金额
    WxRechareInfo string `json:"wx_rechare_info" structs:"wx_rechare_info"` // 微信充值信息微信充值必填
    AlipayRechargeInfo string `json:"alipay_recharge_info" structs:"alipay_recharge_info"` // 支付宝充值信息支付宝充值必填

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2WalletTradeRechargeCardRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2WalletTradeRechargeCardRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2WalletTradeRechargeCardRequest(reqParam)
}

func (bp *BsPay) V2WalletTradeRechargeCardRequest(reqParam V2WalletTradeRechargeCardRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_WALLET_TRADE_RECHARGE_CARD
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
