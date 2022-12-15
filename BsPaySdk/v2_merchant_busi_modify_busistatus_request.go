/**
 * 商户状态变更
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantBusiModifyBusistatusRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id
    Status string `json:"status" structs:"status"` // 状态
    UpdStatusReason string `json:"upd_status_reason" structs:"upd_status_reason"` // 状态变更原因

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantBusiModifyBusistatusRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantBusiModifyBusistatusRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantBusiModifyBusistatusRequest(reqParam)
}

func (bp *BsPay) V2MerchantBusiModifyBusistatusRequest(reqParam V2MerchantBusiModifyBusistatusRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_BUSI_MODIFY_BUSISTATUS
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
