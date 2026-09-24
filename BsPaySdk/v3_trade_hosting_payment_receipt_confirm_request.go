/**
 * 确认收货提醒
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V3TradeHostingPaymentReceiptConfirmRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    ReceivedTime string `json:"received_time" structs:"received_time"` // 快递签收时间

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV3TradeHostingPaymentReceiptConfirmRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V3TradeHostingPaymentReceiptConfirmRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V3TradeHostingPaymentReceiptConfirmRequest(reqParam)
}

func (bp *BsPay) V3TradeHostingPaymentReceiptConfirmRequest(reqParam V3TradeHostingPaymentReceiptConfirmRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V3_TRADE_HOSTING_PAYMENT_RECEIPT_CONFIRM
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
