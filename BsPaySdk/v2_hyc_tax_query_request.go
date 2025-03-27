/**
 * 完税凭证查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2HycTaxQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户汇付id
    StartDate string `json:"start_date" structs:"start_date"` // 开始时间
    EndDate string `json:"end_date" structs:"end_date"` // 结束时间

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2HycTaxQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2HycTaxQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2HycTaxQueryRequest(reqParam)
}

func (bp *BsPay) V2HycTaxQueryRequest(reqParam V2HycTaxQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_HYC_TAX_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
