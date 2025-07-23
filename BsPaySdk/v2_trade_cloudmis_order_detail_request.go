/**
 * 云MIS订单详情查询接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeCloudmisOrderDetailRequest struct {
    ReqId string `json:"req_id" structs:"req_id"` // 请求流水号
    OrgHuifuId string `json:"org_huifu_id" structs:"org_huifu_id"` // 原MIS请求商户号
    OrgDeviceId string `json:"org_device_id" structs:"org_device_id"` // 原MIS请求终端号
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原MIS请求日期

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeCloudmisOrderDetailRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeCloudmisOrderDetailRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeCloudmisOrderDetailRequest(reqParam)
}

func (bp *BsPay) V2TradeCloudmisOrderDetailRequest(reqParam V2TradeCloudmisOrderDetailRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_CLOUDMIS_ORDER_DETAIL
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
