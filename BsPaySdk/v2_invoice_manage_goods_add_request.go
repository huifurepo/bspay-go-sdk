/**
 * 开票商品新增
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2InvoiceManageGoodsAddRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    GoodsName string `json:"goods_name" structs:"goods_name"` // 商品名称
    TaxCode string `json:"tax_code" structs:"tax_code"` // 税收分类编码
    TaxRate string `json:"tax_rate" structs:"tax_rate"` // 税率
    IsDefault string `json:"is_default" structs:"is_default"` // 是否默认

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2InvoiceManageGoodsAddRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2InvoiceManageGoodsAddRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2InvoiceManageGoodsAddRequest(reqParam)
}

func (bp *BsPay) V2InvoiceManageGoodsAddRequest(reqParam V2InvoiceManageGoodsAddRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_INVOICE_MANAGE_GOODS_ADD
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
