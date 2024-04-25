/**
 * 自助扫码开票
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2InvoiceSelfscanopenRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    IvcType string `json:"ivc_type" structs:"ivc_type"` // 发票类型
    OpenType string `json:"open_type" structs:"open_type"` // 开票类型
    OrderAmt string `json:"order_amt" structs:"order_amt"` // 含税合计金额（元）
    GoodsInfos string `json:"goods_infos" structs:"goods_infos"` // 开票商品信息
    PayerInfo string `json:"payer_info" structs:"payer_info"` // 开票人信息

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2InvoiceSelfscanopenRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2InvoiceSelfscanopenRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2InvoiceSelfscanopenRequest(reqParam)
}

func (bp *BsPay) V2InvoiceSelfscanopenRequest(reqParam V2InvoiceSelfscanopenRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_INVOICE_SELFSCANOPEN
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
