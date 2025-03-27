/**
 * 交易确认查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V3TradePaymentDelaytransConfirmqueryRequest struct {
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原请求日期
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 原请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV3TradePaymentDelaytransConfirmqueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V3TradePaymentDelaytransConfirmqueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V3TradePaymentDelaytransConfirmqueryRequest(reqParam)
}

func (bp *BsPay) V3TradePaymentDelaytransConfirmqueryRequest(reqParam V3TradePaymentDelaytransConfirmqueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V3_TRADE_PAYMENT_DELAYTRANS_CONFIRMQUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
