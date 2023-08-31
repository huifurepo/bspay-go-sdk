/**
 * 创建支付分订单
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePayscoreServiceorderCreateRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 商户申请单号
    ServiceIntroduction string `json:"service_introduction" structs:"service_introduction"` // 服务信息
    RiskFund string `json:"risk_fund" structs:"risk_fund"` // 服务风险金
    TimeRange string `json:"time_range" structs:"time_range"` // 服务时间
    NotifyUrl string `json:"notify_url" structs:"notify_url"` // 商户回调地址

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePayscoreServiceorderCreateRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePayscoreServiceorderCreateRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePayscoreServiceorderCreateRequest(reqParam)
}

func (bp *BsPay) V2TradePayscoreServiceorderCreateRequest(reqParam V2TradePayscoreServiceorderCreateRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYSCORE_SERVICEORDER_CREATE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
