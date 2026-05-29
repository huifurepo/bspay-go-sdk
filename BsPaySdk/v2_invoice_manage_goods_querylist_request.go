/**
 * 开票商品查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2InvoiceManageGoodsQuerylistRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2InvoiceManageGoodsQuerylistRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2InvoiceManageGoodsQuerylistRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2InvoiceManageGoodsQuerylistRequest(reqParam)
}

func (bp *BsPay) V2InvoiceManageGoodsQuerylistRequest(reqParam V2InvoiceManageGoodsQuerylistRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_INVOICE_MANAGE_GOODS_QUERYLIST
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
