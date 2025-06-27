/**
 * 全渠道订单分账查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2OcoOrderListRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    BusiSource string `json:"busi_source" structs:"busi_source"` // 分账数据源
    TransDate string `json:"trans_date" structs:"trans_date"` // 交易时间
    PageNum string `json:"page_num" structs:"page_num"` // 页码
    PageSize string `json:"page_size" structs:"page_size"` // 每页大小

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2OcoOrderListRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2OcoOrderListRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2OcoOrderListRequest(reqParam)
}

func (bp *BsPay) V2OcoOrderListRequest(reqParam V2OcoOrderListRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_OCO_ORDER_LIST
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
