/**
 * 快捷支付
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V3TradeOnlinepaymentQuickpayPayRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    SmsCode string `json:"sms_code" structs:"sms_code"` // 短信验证码

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV3TradeOnlinepaymentQuickpayPayRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V3TradeOnlinepaymentQuickpayPayRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V3TradeOnlinepaymentQuickpayPayRequest(reqParam)
}

func (bp *BsPay) V3TradeOnlinepaymentQuickpayPayRequest(reqParam V3TradeOnlinepaymentQuickpayPayRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V3_TRADE_ONLINEPAYMENT_QUICKPAY_PAY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
