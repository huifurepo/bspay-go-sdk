/**
 * 银联统一在线收银台
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeOnlinepaymentUnionpayRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 订单金额
    OrderDesc string `json:"order_desc" structs:"order_desc"` // 商品描述
    RiskCheckData string `json:"risk_check_data" structs:"risk_check_data"` // 安全信息
    ThirdPayData string `json:"third_pay_data" structs:"third_pay_data"` // 三方支付数据jsonObject&lt;br/&gt;pay_scene&#x3D;U_JSAPI或pay_scene&#x3D;U_MINIAPP时，必填

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeOnlinepaymentUnionpayRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeOnlinepaymentUnionpayRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeOnlinepaymentUnionpayRequest(reqParam)
}

func (bp *BsPay) V2TradeOnlinepaymentUnionpayRequest(reqParam V2TradeOnlinepaymentUnionpayRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ONLINEPAYMENT_UNIONPAY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
