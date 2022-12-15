/**
 * 注销终端
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TerminaldeviceDeviceinfoCancelRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id
    DeviceId string `json:"device_id" structs:"device_id"` // 终端号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TerminaldeviceDeviceinfoCancelRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TerminaldeviceDeviceinfoCancelRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TerminaldeviceDeviceinfoCancelRequest(reqParam)
}

func (bp *BsPay) V2TerminaldeviceDeviceinfoCancelRequest(reqParam V2TerminaldeviceDeviceinfoCancelRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TERMINALDEVICE_DEVICEINFO_CANCEL
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
