/**
 * 花呗分期贴息查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2PcreditOrderQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id
    SolutionId string `json:"solution_id" structs:"solution_id"` // 贴息方案id
    StartTime string `json:"start_time" structs:"start_time"` // 活动开始时间
    EndTime string `json:"end_time" structs:"end_time"` // 活动结束时间

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2PcreditOrderQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2PcreditOrderQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2PcreditOrderQueryRequest(reqParam)
}

func (bp *BsPay) V2PcreditOrderQueryRequest(reqParam V2PcreditOrderQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_PCREDIT_ORDER_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
