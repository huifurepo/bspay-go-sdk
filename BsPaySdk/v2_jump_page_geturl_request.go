/**
 * 获取控台页面跳转链接
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2JumpPageGeturlRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    ExternalUserId string `json:"external_user_id" structs:"external_user_id"` // 外部系统用户标识
    JumpFunctionType string `json:"jump_function_type" structs:"jump_function_type"` // 功能菜单

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2JumpPageGeturlRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2JumpPageGeturlRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2JumpPageGeturlRequest(reqParam)
}

func (bp *BsPay) V2JumpPageGeturlRequest(reqParam V2JumpPageGeturlRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_JUMP_PAGE_GETURL
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
