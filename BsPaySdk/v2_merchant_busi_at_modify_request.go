/**
 * 微信支付宝入驻信息修改
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantBusiAtModifyRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    AtRegList string `json:"at_reg_list" structs:"at_reg_list"` // AT信息修改列表

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantBusiAtModifyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantBusiAtModifyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantBusiAtModifyRequest(reqParam)
}

func (bp *BsPay) V2MerchantBusiAtModifyRequest(reqParam V2MerchantBusiAtModifyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_BUSI_AT_MODIFY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
