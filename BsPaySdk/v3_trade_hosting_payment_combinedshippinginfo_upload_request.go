/**
 * 发货信息合单录入
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V3TradeHostingPaymentCombinedshippinginfoUploadRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    OrderKey string `json:"order_key" structs:"order_key"` // 订单信息
    UploadTime string `json:"upload_time" structs:"upload_time"` // 上传时间
    Payer string `json:"payer" structs:"payer"` // 支付者

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV3TradeHostingPaymentCombinedshippinginfoUploadRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V3TradeHostingPaymentCombinedshippinginfoUploadRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V3TradeHostingPaymentCombinedshippinginfoUploadRequest(reqParam)
}

func (bp *BsPay) V3TradeHostingPaymentCombinedshippinginfoUploadRequest(reqParam V3TradeHostingPaymentCombinedshippinginfoUploadRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V3_TRADE_HOSTING_PAYMENT_COMBINEDSHIPPINGINFO_UPLOAD
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
