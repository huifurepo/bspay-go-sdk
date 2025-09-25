/**
 * 代运营佣金代扣查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2LlaWithholdQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原请求日期
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 原请求流水号org_hf_seq_id与org_req_seq_id二选一必填。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：2021091708126665001&lt;/font&gt;
    OrgHfSeqId string `json:"org_hf_seq_id" structs:"org_hf_seq_id"` // 原全局流水号org_hf_seq_id与org_req_seq_id二选一必填。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：00470topo1A221019132207P068ac1362af00000&lt;/font&gt;
    AgencyHuifuId string `json:"agency_huifu_id" structs:"agency_huifu_id"` // 代运营汇付id

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2LlaWithholdQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2LlaWithholdQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2LlaWithholdQueryRequest(reqParam)
}

func (bp *BsPay) V2LlaWithholdQueryRequest(reqParam V2LlaWithholdQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_LLA_WITHHOLD_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
