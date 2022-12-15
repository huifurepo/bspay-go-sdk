/**
 * 银行卡分期支持银行查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeBankinstallmentinfoQueryRequest struct {
    PageNum string `json:"page_num" structs:"page_num"` // 页码
    PageSize string `json:"page_size" structs:"page_size"` // 每页条数
    ProductId string `json:"product_id" structs:"product_id"` // 产品号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeBankinstallmentinfoQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeBankinstallmentinfoQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeBankinstallmentinfoQueryRequest(reqParam)
}

func (bp *BsPay) V2TradeBankinstallmentinfoQueryRequest(reqParam V2TradeBankinstallmentinfoQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_BANKINSTALLMENTINFO_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc, false)
}
