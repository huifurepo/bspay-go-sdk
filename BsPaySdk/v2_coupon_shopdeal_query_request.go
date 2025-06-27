/**
 * 美团非餐饮获取团购信息
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2CouponShopdealQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    BindId string `json:"bind_id" structs:"bind_id"` // 门店绑定流水号
    Offset string `json:"offset" structs:"offset"` // 页码
    Limit string `json:"limit" structs:"limit"` // 页大小
    Source string `json:"source" structs:"source"` // 售卖平台

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2CouponShopdealQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2CouponShopdealQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2CouponShopdealQueryRequest(reqParam)
}

func (bp *BsPay) V2CouponShopdealQueryRequest(reqParam V2CouponShopdealQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_COUPON_SHOPDEAL_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
