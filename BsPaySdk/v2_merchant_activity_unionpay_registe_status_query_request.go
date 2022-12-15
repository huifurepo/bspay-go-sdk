/**
 * 银联活动商户入驻状态查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantActivityUnionpayRegisteStatusQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id
    SerialNo string `json:"serial_no" structs:"serial_no"` // 工单号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantActivityUnionpayRegisteStatusQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantActivityUnionpayRegisteStatusQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantActivityUnionpayRegisteStatusQueryRequest(reqParam)
}

func (bp *BsPay) V2MerchantActivityUnionpayRegisteStatusQueryRequest(reqParam V2MerchantActivityUnionpayRegisteStatusQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_ACTIVITY_UNIONPAY_REGISTE_STATUS_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
