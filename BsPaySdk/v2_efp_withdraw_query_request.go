/**
 * 全渠道资金付款提现记录查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2EfpWithdrawQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户汇付id
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 付款或提现的请求流水号
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 付款或提现的请求日期

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2EfpWithdrawQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2EfpWithdrawQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2EfpWithdrawQueryRequest(reqParam)
}

func (bp *BsPay) V2EfpWithdrawQueryRequest(reqParam V2EfpWithdrawQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_EFP_WITHDRAW_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
