/**
 * 全渠道订单分账接收方查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2OcoOrderDetailListRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    BusiSource string `json:"busi_source" structs:"busi_source"` // 分账数据源
    OcoOrderId string `json:"oco_order_id" structs:"oco_order_id"` // 业务订单号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2OcoOrderDetailListRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2OcoOrderDetailListRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2OcoOrderDetailListRequest(reqParam)
}

func (bp *BsPay) V2OcoOrderDetailListRequest(reqParam V2OcoOrderDetailListRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_OCO_ORDER_DETAIL_LIST
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
