/**
 * 汇付入账查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeOnlinepaymentTransferRemittanceorderRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    OrgReqStartDate string `json:"org_req_start_date" structs:"org_req_start_date"` // 原请求开始日期
    OrgReqEndDate string `json:"org_req_end_date" structs:"org_req_end_date"` // 原请求结束日期

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeOnlinepaymentTransferRemittanceorderRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeOnlinepaymentTransferRemittanceorderRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeOnlinepaymentTransferRemittanceorderRequest(reqParam)
}

func (bp *BsPay) V2TradeOnlinepaymentTransferRemittanceorderRequest(reqParam V2TradeOnlinepaymentTransferRemittanceorderRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ONLINEPAYMENT_TRANSFER_REMITTANCEORDER
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
