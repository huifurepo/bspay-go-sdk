/**
 * 删除付款人
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2BillEntPayerDeleteRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    PayerId string `json:"payer_id" structs:"payer_id"` // 付款人

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2BillEntPayerDeleteRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2BillEntPayerDeleteRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2BillEntPayerDeleteRequest(reqParam)
}

func (bp *BsPay) V2BillEntPayerDeleteRequest(reqParam V2BillEntPayerDeleteRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_BILL_ENT_PAYER_DELETE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
