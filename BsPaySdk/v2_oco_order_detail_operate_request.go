/**
 * 全渠道订单分账明细操作
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2OcoOrderDetailOperateRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    BusiSource string `json:"busi_source" structs:"busi_source"` // 分账数据源
    OcoOrderId string `json:"oco_order_id" structs:"oco_order_id"` // 业务订单号
    OperateType string `json:"operate_type" structs:"operate_type"` // 操作类型
    PayType string `json:"pay_type" structs:"pay_type"` // 支付方式枚举：BALANCE-余额支付 EFP-全域资金付款；备注：当operate_type&#x3D;SPLIT 立即分账时，pay_type必填，且若为退款，按原交易类型原路返回；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：BALANCE&lt;/font&gt;

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2OcoOrderDetailOperateRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2OcoOrderDetailOperateRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2OcoOrderDetailOperateRequest(reqParam)
}

func (bp *BsPay) V2OcoOrderDetailOperateRequest(reqParam V2OcoOrderDetailOperateRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_OCO_ORDER_DETAIL_OPERATE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
