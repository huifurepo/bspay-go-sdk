/**
 * 支付分扣款
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePayscorePayPayscorepayRequest struct {
    OutTradeNo string `json:"out_trade_no" structs:"out_trade_no"` // 微信扣款单号
    GoodsDesc string `json:"goods_desc" structs:"goods_desc"` // 商品描述
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    RiskCheckData string `json:"risk_check_data" structs:"risk_check_data"` // 安全信息

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePayscorePayPayscorepayRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePayscorePayPayscorepayRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePayscorePayPayscorepayRequest(reqParam)
}

func (bp *BsPay) V2TradePayscorePayPayscorepayRequest(reqParam V2TradePayscorePayPayscorepayRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYSCORE_PAY_PAYSCOREPAY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
