/**
 * 电子账户付款
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePaymentZxeAcctpyerRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    OutHuifuId string `json:"out_huifu_id" structs:"out_huifu_id"` // 出款方商户号
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 订单金额
    ThirdPayData string `json:"third_pay_data" structs:"third_pay_data"` // 三方支付数据

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePaymentZxeAcctpyerRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePaymentZxeAcctpyerRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePaymentZxeAcctpyerRequest(reqParam)
}

func (bp *BsPay) V2TradePaymentZxeAcctpyerRequest(reqParam V2TradePaymentZxeAcctpyerRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYMENT_ZXE_ACCTPYER
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
