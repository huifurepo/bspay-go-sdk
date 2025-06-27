/**
 * 美团卡券核销
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2CouponMeituanConsumeRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    BindId string `json:"bind_id" structs:"bind_id"` // 门店绑定流水号
    ReceiptCodeInfos string `json:"receipt_code_infos" structs:"receipt_code_infos"` // 核销券
    AppShopAccount string `json:"app_shop_account" structs:"app_shop_account"` // 登录账号
    AppShopAccountName string `json:"app_shop_account_name" structs:"app_shop_account_name"` // 登录用户名

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2CouponMeituanConsumeRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2CouponMeituanConsumeRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2CouponMeituanConsumeRequest(reqParam)
}

func (bp *BsPay) V2CouponMeituanConsumeRequest(reqParam V2CouponMeituanConsumeRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_COUPON_MEITUAN_CONSUME
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
