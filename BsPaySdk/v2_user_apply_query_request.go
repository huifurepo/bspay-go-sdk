/**
 * 用户申请单状态查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2UserApplyQueryRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ApplyNo string `json:"apply_no" structs:"apply_no"` // 申请单号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2UserApplyQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2UserApplyQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2UserApplyQueryRequest(reqParam)
}

func (bp *BsPay) V2UserApplyQueryRequest(reqParam V2UserApplyQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_USER_APPLY_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
