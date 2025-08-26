/**
 * 抖音卡券撤销
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2CouponDouyinCancelRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    BindId string `json:"bind_id" structs:"bind_id"` // 门店绑定流水号
    EncryptedCode string `json:"encrypted_code" structs:"encrypted_code"` // 抖音券码
    VerifyId string `json:"verify_id" structs:"verify_id"` // 核销标识

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2CouponDouyinCancelRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2CouponDouyinCancelRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2CouponDouyinCancelRequest(reqParam)
}

func (bp *BsPay) V2CouponDouyinCancelRequest(reqParam V2CouponDouyinCancelRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_COUPON_DOUYIN_CANCEL
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
