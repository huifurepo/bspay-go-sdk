/**
 * 商户微信支付宝ID查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantBusiMercodeQueryRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付ID
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    PayWay string `json:"pay_way" structs:"pay_way"` // 入驻通道类型

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantBusiMercodeQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantBusiMercodeQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantBusiMercodeQueryRequest(reqParam)
}

func (bp *BsPay) V2MerchantBusiMercodeQueryRequest(reqParam V2MerchantBusiMercodeQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_BUSI_MERCODE_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
