/**
 * 修改付款人信息
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2BillEntPayerUpdateRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    PayerId string `json:"payer_id" structs:"payer_id"` // 付款人
    PayerName string `json:"payer_name" structs:"payer_name"` // 付款人名称

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2BillEntPayerUpdateRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2BillEntPayerUpdateRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2BillEntPayerUpdateRequest(reqParam)
}

func (bp *BsPay) V2BillEntPayerUpdateRequest(reqParam V2BillEntPayerUpdateRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_BILL_ENT_PAYER_UPDATE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
