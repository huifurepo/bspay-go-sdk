/**
 * 开票员查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2InvoiceClerkListQueryRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2InvoiceClerkListQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2InvoiceClerkListQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2InvoiceClerkListQueryRequest(reqParam)
}

func (bp *BsPay) V2InvoiceClerkListQueryRequest(reqParam V2InvoiceClerkListQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_INVOICE_CLERK_LIST_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
