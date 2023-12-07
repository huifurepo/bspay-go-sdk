/**
 * 银行大额转账差错申请查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeOnlinepaymentTransferBankmistakeApplyqueryRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原请求日期
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 原请求流水号
    OrderType string `json:"order_type" structs:"order_type"` // 订单类型

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeOnlinepaymentTransferBankmistakeApplyqueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeOnlinepaymentTransferBankmistakeApplyqueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeOnlinepaymentTransferBankmistakeApplyqueryRequest(reqParam)
}

func (bp *BsPay) V2TradeOnlinepaymentTransferBankmistakeApplyqueryRequest(reqParam V2TradeOnlinepaymentTransferBankmistakeApplyqueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ONLINEPAYMENT_TRANSFER_BANKMISTAKE_APPLYQUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
