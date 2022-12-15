/**
 * 网银支持银行列表查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeOnlinepaymentBankpayBanklistRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    GateType string `json:"gate_type" structs:"gate_type"` // 网关支付类型
    OrderType string `json:"order_type" structs:"order_type"` // 订单类型

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeOnlinepaymentBankpayBanklistRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeOnlinepaymentBankpayBanklistRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeOnlinepaymentBankpayBanklistRequest(reqParam)
}

func (bp *BsPay) V2TradeOnlinepaymentBankpayBanklistRequest(reqParam V2TradeOnlinepaymentBankpayBanklistRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ONLINEPAYMENT_BANKPAY_BANKLIST
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
