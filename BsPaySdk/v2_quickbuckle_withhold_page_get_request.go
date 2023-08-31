/**
 * 代扣绑卡页面版
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2QuickbuckleWithholdPageGetRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付Id
    OrderId string `json:"order_id" structs:"order_id"` // 订单号
    OrderDate string `json:"order_date" structs:"order_date"` // 订单日期

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2QuickbuckleWithholdPageGetRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2QuickbuckleWithholdPageGetRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2QuickbuckleWithholdPageGetRequest(reqParam)
}

func (bp *BsPay) V2QuickbuckleWithholdPageGetRequest(reqParam V2QuickbuckleWithholdPageGetRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_QUICKBUCKLE_WITHHOLD_PAGE_GET
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
