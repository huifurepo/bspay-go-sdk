/**
 * 账单退款接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V3BillpayOrderPaymentRefundRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    BillNo string `json:"bill_no" structs:"bill_no"` // 账单编号
    RefAmt string `json:"ref_amt" structs:"ref_amt"` // 退款金额
    BankInfoData string `json:"bank_info_data" structs:"bank_info_data"` // 大额转账支付账户信息数据jsonObject格式；银行大额转账支付交易的退款申请,付款方账户类型为对公时必填

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV3BillpayOrderPaymentRefundRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V3BillpayOrderPaymentRefundRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V3BillpayOrderPaymentRefundRequest(reqParam)
}

func (bp *BsPay) V3BillpayOrderPaymentRefundRequest(reqParam V3BillpayOrderPaymentRefundRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V3_BILLPAY_ORDER_PAYMENT_REFUND
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
