/**
 * 上架下架分期活动接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2PcreditStatueModifyRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    SolutionId string `json:"solution_id" structs:"solution_id"` // 贴息方案实例id
    Status string `json:"status" structs:"status"` // 状态控制

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2PcreditStatueModifyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2PcreditStatueModifyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2PcreditStatueModifyRequest(reqParam)
}

func (bp *BsPay) V2PcreditStatueModifyRequest(reqParam V2PcreditStatueModifyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_PCREDIT_STATUE_MODIFY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
