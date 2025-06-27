/**
 * 抖音券状态批量查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2CouponDouyinCertificateQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    BindId string `json:"bind_id" structs:"bind_id"` // 门店绑定流水号
    EncryptedCode string `json:"encrypted_code" structs:"encrypted_code"` // 验券准备接口返回的加密券码encrypted_code和order_id二选一必传，encrypted_code和order_id不能同时传入
    OrderId string `json:"order_id" structs:"order_id"` // 订单id验券准备等接口获得，encrypted_code和order_id二选一必传，encrypted_code和order_id不能同时传入

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2CouponDouyinCertificateQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2CouponDouyinCertificateQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2CouponDouyinCertificateQueryRequest(reqParam)
}

func (bp *BsPay) V2CouponDouyinCertificateQueryRequest(reqParam V2CouponDouyinCertificateQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_COUPON_DOUYIN_CERTIFICATE_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
