/**
 * 商户分账配置
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantSplitConfigRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户汇付Id
    RuleOrigin string `json:"rule_origin" structs:"rule_origin"` // 分账规则来源
    RepealFlag string `json:"repeal_flag" structs:"repeal_flag"` // 分账是否支持撤销交易
    RefundFlag string `json:"refund_flag" structs:"refund_flag"` // 分账是否支持退货交易
    DivFlag string `json:"div_flag" structs:"div_flag"` // 分账开关
    ApplyRatio string `json:"apply_ratio" structs:"apply_ratio"` // 最大分账比例
    StartType string `json:"start_type" structs:"start_type"` // 生效类型

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantSplitConfigRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantSplitConfigRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantSplitConfigRequest(reqParam)
}

func (bp *BsPay) V2MerchantSplitConfigRequest(reqParam V2MerchantSplitConfigRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_SPLIT_CONFIG
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
