/**
 * 防断链子商户号列表查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantAtpreventQuerysubmerlistRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户汇付Id

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantAtpreventQuerysubmerlistRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantAtpreventQuerysubmerlistRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantAtpreventQuerysubmerlistRequest(reqParam)
}

func (bp *BsPay) V2MerchantAtpreventQuerysubmerlistRequest(reqParam V2MerchantAtpreventQuerysubmerlistRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_ATPREVENT_QUERYSUBMERLIST
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
