/**
 * 绑定终端信息查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TerminaldeviceDeviceinfoQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id
    PageSize string `json:"page_size" structs:"page_size"` // 分页大小
    PageNum string `json:"page_num" structs:"page_num"` // 当前页码

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TerminaldeviceDeviceinfoQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TerminaldeviceDeviceinfoQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TerminaldeviceDeviceinfoQueryRequest(reqParam)
}

func (bp *BsPay) V2TerminaldeviceDeviceinfoQueryRequest(reqParam V2TerminaldeviceDeviceinfoQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TERMINALDEVICE_DEVICEINFO_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
