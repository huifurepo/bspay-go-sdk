/**
 * 银联统一在线收银台签解约查询接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeOnlinepaymentUnionsignqueryRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原请求日期
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 原请求流水号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeOnlinepaymentUnionsignqueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeOnlinepaymentUnionsignqueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeOnlinepaymentUnionsignqueryRequest(reqParam)
}

func (bp *BsPay) V2TradeOnlinepaymentUnionsignqueryRequest(reqParam V2TradeOnlinepaymentUnionsignqueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ONLINEPAYMENT_UNIONSIGNQUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
