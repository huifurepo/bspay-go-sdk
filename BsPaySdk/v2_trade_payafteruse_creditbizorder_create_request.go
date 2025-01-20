/**
 * 服务单创建
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePayafteruseCreditbizorderCreateRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 订单总金额
    BuyerId string `json:"buyer_id" structs:"buyer_id"` // 支付宝用户ID
    Title string `json:"title" structs:"title"` // 订单标题
    MerchantBizType string `json:"merchant_biz_type" structs:"merchant_biz_type"` // 订单类型
    Path string `json:"path" structs:"path"` // 订单详情地址
    ZmServiceId string `json:"zm_service_id" structs:"zm_service_id"` // 芝麻信用服务ID
    ItemInfos string `json:"item_infos" structs:"item_infos"` // 商品详细信息

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePayafteruseCreditbizorderCreateRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePayafteruseCreditbizorderCreateRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePayafteruseCreditbizorderCreateRequest(reqParam)
}

func (bp *BsPay) V2TradePayafteruseCreditbizorderCreateRequest(reqParam V2TradePayafteruseCreditbizorderCreateRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYAFTERUSE_CREDITBIZORDER_CREATE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
