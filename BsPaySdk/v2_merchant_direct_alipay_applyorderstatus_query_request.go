/**
 * 支付宝直连-查询申请状态
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantDirectAlipayApplyorderstatusQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户汇付Id
    AppId string `json:"app_id" structs:"app_id"` // 开发者的应用ID

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantDirectAlipayApplyorderstatusQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantDirectAlipayApplyorderstatusQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantDirectAlipayApplyorderstatusQueryRequest(reqParam)
}

func (bp *BsPay) V2MerchantDirectAlipayApplyorderstatusQueryRequest(reqParam V2MerchantDirectAlipayApplyorderstatusQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_DIRECT_ALIPAY_APPLYORDERSTATUS_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
