/**
 * 终端云MIS交易
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeCloudmisDeviceInformationMisRequest struct {
    ReqId string `json:"req_id" structs:"req_id"` // 请求流水号
    DeviceId string `json:"device_id" structs:"device_id"` // 终端设备号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    JsonData string `json:"json_data" structs:"json_data"` // 交易信息

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeCloudmisDeviceInformationMisRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeCloudmisDeviceInformationMisRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeCloudmisDeviceInformationMisRequest(reqParam)
}

func (bp *BsPay) V2TradeCloudmisDeviceInformationMisRequest(reqParam V2TradeCloudmisDeviceInformationMisRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_CLOUDMIS_DEVICE_INFORMATION_MIS
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
