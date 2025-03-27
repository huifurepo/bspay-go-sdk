/**
 * 修改子账户配置
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantSettleModifyRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户汇付Id
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 上级汇付Id
    AcctId string `json:"acct_id" structs:"acct_id"` // 子账户号
    SettleConfig string `json:"settle_config" structs:"settle_config"` // 结算规则配置
    CardInfo string `json:"card_info" structs:"card_info"` // 结算卡信息配置新账户绑定的结算银行账户。jsonObject格式。若结算规则中上送token_no，则card_info不填。

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantSettleModifyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantSettleModifyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantSettleModifyRequest(reqParam)
}

func (bp *BsPay) V2MerchantSettleModifyRequest(reqParam V2MerchantSettleModifyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_SETTLE_MODIFY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
