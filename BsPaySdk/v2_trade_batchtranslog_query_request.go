/**
 * 批量出金交易查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeBatchtranslogQueryRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    BeginDate string `json:"begin_date" structs:"begin_date"` // 开始日期
    EndDate string `json:"end_date" structs:"end_date"` // 结束日期

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeBatchtranslogQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeBatchtranslogQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeBatchtranslogQueryRequest(reqParam)
}

func (bp *BsPay) V2TradeBatchtranslogQueryRequest(reqParam V2TradeBatchtranslogQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_BATCHTRANSLOG_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
