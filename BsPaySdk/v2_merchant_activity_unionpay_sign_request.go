/**
 * 银联活动报名
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantActivityUnionpaySignRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id
    ActivityId string `json:"activity_id" structs:"activity_id"` // 活动编号
    MerNo string `json:"mer_no" structs:"mer_no"` // 银联活动商户号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantActivityUnionpaySignRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantActivityUnionpaySignRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantActivityUnionpaySignRequest(reqParam)
}

func (bp *BsPay) V2MerchantActivityUnionpaySignRequest(reqParam V2MerchantActivityUnionpaySignRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_ACTIVITY_UNIONPAY_SIGN
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
