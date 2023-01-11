/**
 * 批量交易状态查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeTransstatQueryRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    PageNo string `json:"page_no" structs:"page_no"` // 页码
    PageSize string `json:"page_size" structs:"page_size"` // 页大小
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeTransstatQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeTransstatQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeTransstatQueryRequest(reqParam)
}

func (bp *BsPay) V2TradeTransstatQueryRequest(reqParam V2TradeTransstatQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_TRANSSTAT_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
