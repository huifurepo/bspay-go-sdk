/**
 * 扫码交易关单查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePaymentScanpayClosequeryRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原交易请求日期

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePaymentScanpayClosequeryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePaymentScanpayClosequeryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePaymentScanpayClosequeryRequest(reqParam)
}

func (bp *BsPay) V2TradePaymentScanpayClosequeryRequest(reqParam V2TradePaymentScanpayClosequeryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYMENT_SCANPAY_CLOSEQUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
