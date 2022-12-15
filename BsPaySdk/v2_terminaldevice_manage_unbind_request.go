/**
 * 终端解绑
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TerminaldeviceManageUnbindRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付ID
    DeviceId string `json:"device_id" structs:"device_id"` // 终端号
    Reason string `json:"reason" structs:"reason"` // 解绑原因

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TerminaldeviceManageUnbindRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TerminaldeviceManageUnbindRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TerminaldeviceManageUnbindRequest(reqParam)
}

func (bp *BsPay) V2TerminaldeviceManageUnbindRequest(reqParam V2TerminaldeviceManageUnbindRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TERMINALDEVICE_MANAGE_UNBIND
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
