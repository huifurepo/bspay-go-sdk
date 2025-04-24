/**
 * 聚合反扫
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V3TradePaymentMicropayRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 交易金额
    GoodsDesc string `json:"goods_desc" structs:"goods_desc"` // 商品描述
    AuthCode string `json:"auth_code" structs:"auth_code"` // 支付授权码
    RiskCheckData string `json:"risk_check_data" structs:"risk_check_data"` // 安全信息

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV3TradePaymentMicropayRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V3TradePaymentMicropayRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V3TradePaymentMicropayRequest(reqParam)
}

func (bp *BsPay) V3TradePaymentMicropayRequest(reqParam V3TradePaymentMicropayRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V3_TRADE_PAYMENT_MICROPAY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
