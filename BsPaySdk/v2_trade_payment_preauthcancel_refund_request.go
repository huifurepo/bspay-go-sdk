/**
 * 微信支付宝预授权撤销
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePaymentPreauthcancelRefundRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 客户号
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原交易请求日期
    OrdAmt string `json:"ord_amt" structs:"ord_amt"` // 撤销金额
    RiskCheckInfo string `json:"risk_check_info" structs:"risk_check_info"` // 风控信息

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePaymentPreauthcancelRefundRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePaymentPreauthcancelRefundRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePaymentPreauthcancelRefundRequest(reqParam)
}

func (bp *BsPay) V2TradePaymentPreauthcancelRefundRequest(reqParam V2TradePaymentPreauthcancelRefundRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYMENT_PREAUTHCANCEL_REFUND
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
