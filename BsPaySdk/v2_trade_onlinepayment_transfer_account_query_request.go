/**
 * 银行大额资金流水查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeOnlinepaymentTransferAccountQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 原请求流水号
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原请求日期
    TransEndDate string `json:"trans_end_date" structs:"trans_end_date"` // 打款结束日期
    TransStartDate string `json:"trans_start_date" structs:"trans_start_date"` // 交易开始日期

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeOnlinepaymentTransferAccountQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeOnlinepaymentTransferAccountQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeOnlinepaymentTransferAccountQueryRequest(reqParam)
}

func (bp *BsPay) V2TradeOnlinepaymentTransferAccountQueryRequest(reqParam V2TradeOnlinepaymentTransferAccountQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ONLINEPAYMENT_TRANSFER_ACCOUNT_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
