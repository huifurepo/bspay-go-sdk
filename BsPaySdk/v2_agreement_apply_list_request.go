/**
 * 协议申请列表查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2AgreementApplyListRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    AgreementType string `json:"agreement_type" structs:"agreement_type"` // 协议类型
    PageNum string `json:"page_num" structs:"page_num"` // 当前页
    PageSize string `json:"page_size" structs:"page_size"` // 分页大小

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2AgreementApplyListRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2AgreementApplyListRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2AgreementApplyListRequest(reqParam)
}

func (bp *BsPay) V2AgreementApplyListRequest(reqParam V2AgreementApplyListRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_AGREEMENT_APPLY_LIST
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
