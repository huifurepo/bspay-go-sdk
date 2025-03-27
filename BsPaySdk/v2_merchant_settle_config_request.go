/**
 * 子账户开通配置
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantSettleConfigRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户汇付Id
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 上级汇付Id
    AcctType string `json:"acct_type" structs:"acct_type"` // 子账户类型
    AcctName string `json:"acct_name" structs:"acct_name"` // 账户名称
    CardInfo string `json:"card_info" structs:"card_info"` // 结算卡信息配置新账户绑定的结算银行账户。jsonObject格式。若结算规则中上送token_no，则card_info不填。

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantSettleConfigRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantSettleConfigRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantSettleConfigRequest(reqParam)
}

func (bp *BsPay) V2MerchantSettleConfigRequest(reqParam V2MerchantSettleConfigRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_SETTLE_CONFIG
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
