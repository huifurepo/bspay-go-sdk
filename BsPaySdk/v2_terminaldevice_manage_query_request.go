/**
 * 服务商终端列表查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TerminaldeviceManageQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TerminaldeviceManageQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TerminaldeviceManageQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TerminaldeviceManageQueryRequest(reqParam)
}

func (bp *BsPay) V2TerminaldeviceManageQueryRequest(reqParam V2TerminaldeviceManageQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TERMINALDEVICE_MANAGE_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
