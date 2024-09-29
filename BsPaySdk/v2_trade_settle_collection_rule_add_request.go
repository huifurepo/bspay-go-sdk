/**
 * 新增归集配置
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeSettleCollectionRuleAddRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    InHuifuId string `json:"in_huifu_id" structs:"in_huifu_id"` // 转入方商户号
    OutHuifuId string `json:"out_huifu_id" structs:"out_huifu_id"` // 转出方商户号
    SignUserMobileNo string `json:"sign_user_mobile_no" structs:"sign_user_mobile_no"` // 签约人手机号协议类型为电子协议时必填，必须为法人手机号。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：13911111111&lt;/font&gt;
    FileId string `json:"file_id" structs:"file_id"` // 协议文件Id协议类型为纸质协议时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e529&lt;/font&gt;

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeSettleCollectionRuleAddRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeSettleCollectionRuleAddRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeSettleCollectionRuleAddRequest(reqParam)
}

func (bp *BsPay) V2TradeSettleCollectionRuleAddRequest(reqParam V2TradeSettleCollectionRuleAddRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_SETTLE_COLLECTION_RULE_ADD
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
