/**
 * 支付分扣款
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePayscorePayPayscorepayRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    DeductReqSeqId string `json:"deduct_req_seq_id" structs:"deduct_req_seq_id"` // 扣款登记创建请求流水号deduct_req_seq_id与deduct_hf_seq_id二选一；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：2022012614120615001&lt;/font&gt;
    DeductHfSeqId string `json:"deduct_hf_seq_id" structs:"deduct_hf_seq_id"` // 扣款登记返回的汇付全局流水号deduct_req_seq_id与deduct_hf_seq_id二选一；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：00470topo1A211015160805P090ac132fef00000&lt;/font&gt;
    OutTradeNo string `json:"out_trade_no" structs:"out_trade_no"` // 微信扣款单号
    GoodsDesc string `json:"goods_desc" structs:"goods_desc"` // 商品描述
    RiskCheckData string `json:"risk_check_data" structs:"risk_check_data"` // 安全信息

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePayscorePayPayscorepayRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePayscorePayPayscorepayRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePayscorePayPayscorepayRequest(reqParam)
}

func (bp *BsPay) V2TradePayscorePayPayscorepayRequest(reqParam V2TradePayscorePayPayscorepayRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYSCORE_PAY_PAYSCOREPAY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
