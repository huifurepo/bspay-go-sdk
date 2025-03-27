/**
 * 个人签约发起
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2HycPersonsignCreateRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 用户汇付id
    MinorAgentId string `json:"minor_agent_id" structs:"minor_agent_id"` // 落地公司机构号
    LjhData string `json:"ljh_data" structs:"ljh_data"` // 乐接活请求参数jsonObject格式 合作平台为乐接活时必传

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2HycPersonsignCreateRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2HycPersonsignCreateRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2HycPersonsignCreateRequest(reqParam)
}

func (bp *BsPay) V2HycPersonsignCreateRequest(reqParam V2HycPersonsignCreateRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_HYC_PERSONSIGN_CREATE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
