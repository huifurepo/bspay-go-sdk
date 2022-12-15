/**
 * 微信支付宝预授权完成撤销
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePaymentPreauthpaycancelRefundRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 客户号
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原预授权完成交易请求日期
    OrdAmt string `json:"ord_amt" structs:"ord_amt"` // 完成撤销金额
    RiskCheckInfo string `json:"risk_check_info" structs:"risk_check_info"` // 风控信息

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePaymentPreauthpaycancelRefundRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePaymentPreauthpaycancelRefundRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePaymentPreauthpaycancelRefundRequest(reqParam)
}

func (bp *BsPay) V2TradePaymentPreauthpaycancelRefundRequest(reqParam V2TradePaymentPreauthpaycancelRefundRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYMENT_PREAUTHPAYCANCEL_REFUND
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
