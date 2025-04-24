/**
 * 银联统一在线收银台解约接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeOnlinepaymentUnioncancelRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原请求日期
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 原请求流水号
    NotifyUrl string `json:"notify_url" structs:"notify_url"` // 异步通知地址

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeOnlinepaymentUnioncancelRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeOnlinepaymentUnioncancelRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeOnlinepaymentUnioncancelRequest(reqParam)
}

func (bp *BsPay) V2TradeOnlinepaymentUnioncancelRequest(reqParam V2TradeOnlinepaymentUnioncancelRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ONLINEPAYMENT_UNIONCANCEL
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
