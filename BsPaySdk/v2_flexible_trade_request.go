/**
 * 灵工支付
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2FlexibleTradeRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    OutHuifuId string `json:"out_huifu_id" structs:"out_huifu_id"` // 出款方商户号
    OutAcctId string `json:"out_acct_id" structs:"out_acct_id"` // 出款方账户号
    StageOperationType string `json:"stage_operation_type" structs:"stage_operation_type"` // 交易阶段操作类型
    PhaseHfSeqId string `json:"phase_hf_seq_id" structs:"phase_hf_seq_id"` // 前段交易流水号** 当交易阶段操作类型为02时，该字段必填。填写的是交易阶段操作类型为01时交易已完成的交易全局流水号。 &lt;font color&#x3D;&quot;green&quot;&gt;示例值：20250620112533115566896&lt;/font&gt;
    OrdAmt string `json:"ord_amt" structs:"ord_amt"` // 支付金额

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2FlexibleTradeRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2FlexibleTradeRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2FlexibleTradeRequest(reqParam)
}

func (bp *BsPay) V2FlexibleTradeRequest(reqParam V2FlexibleTradeRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_FLEXIBLE_TRADE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
