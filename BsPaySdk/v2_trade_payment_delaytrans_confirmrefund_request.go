/**
 * 交易确认退款
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePaymentDelaytransConfirmrefundRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原交易请求日期
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 原交易请求流水号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePaymentDelaytransConfirmrefundRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePaymentDelaytransConfirmrefundRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePaymentDelaytransConfirmrefundRequest(reqParam)
}

func (bp *BsPay) V2TradePaymentDelaytransConfirmrefundRequest(reqParam V2TradePaymentDelaytransConfirmrefundRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYMENT_DELAYTRANS_CONFIRMREFUND
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
