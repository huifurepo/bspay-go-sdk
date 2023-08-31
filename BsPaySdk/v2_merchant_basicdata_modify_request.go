/**
 * 商户基本信息修改(2022)
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantBasicdataModifyRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 上级主体ID
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantBasicdataModifyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantBasicdataModifyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantBasicdataModifyRequest(reqParam)
}

func (bp *BsPay) V2MerchantBasicdataModifyRequest(reqParam V2MerchantBasicdataModifyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_BASICDATA_MODIFY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
