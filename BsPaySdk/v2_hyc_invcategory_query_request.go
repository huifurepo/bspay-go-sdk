/**
 * 开票类目查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2HycInvcategoryQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    MinorAgentId string `json:"minor_agent_id" structs:"minor_agent_id"` // 落地公司机构号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2HycInvcategoryQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2HycInvcategoryQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2HycInvcategoryQueryRequest(reqParam)
}

func (bp *BsPay) V2HycInvcategoryQueryRequest(reqParam V2HycInvcategoryQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_HYC_INVCATEGORY_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
