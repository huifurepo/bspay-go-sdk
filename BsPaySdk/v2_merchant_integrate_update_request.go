/**
 * 商户统一变更接口(2022)
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantIntegrateUpdateRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付ID
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 渠道商汇付ID
    DealType string `json:"deal_type" structs:"deal_type"` // 业务处理类型

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantIntegrateUpdateRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantIntegrateUpdateRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantIntegrateUpdateRequest(reqParam)
}

func (bp *BsPay) V2MerchantIntegrateUpdateRequest(reqParam V2MerchantIntegrateUpdateRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_INTEGRATE_UPDATE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
