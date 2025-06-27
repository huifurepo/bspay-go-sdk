/**
 * 美团卡券撤销
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2CouponMeituanCancelRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    BindId string `json:"bind_id" structs:"bind_id"` // 门店绑定流水号
    AppShopAccount string `json:"app_shop_account" structs:"app_shop_account"` // 登录账号
    AppShopAccountName string `json:"app_shop_account_name" structs:"app_shop_account_name"` // 登录用户名
    ReceiptCode string `json:"receipt_code" structs:"receipt_code"` // 券码

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2CouponMeituanCancelRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2CouponMeituanCancelRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2CouponMeituanCancelRequest(reqParam)
}

func (bp *BsPay) V2CouponMeituanCancelRequest(reqParam V2CouponMeituanCancelRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_COUPON_MEITUAN_CANCEL
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
