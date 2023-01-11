/**
 * 微信直连-微信关注配置
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantDirectWechatSubscribeConfigRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付ID
    AppId string `json:"app_id" structs:"app_id"` // 开发者的应用ID
    MchId string `json:"mch_id" structs:"mch_id"` // 商户号
    SubMchid string `json:"sub_mchid" structs:"sub_mchid"` // 特约商户号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantDirectWechatSubscribeConfigRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantDirectWechatSubscribeConfigRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantDirectWechatSubscribeConfigRequest(reqParam)
}

func (bp *BsPay) V2MerchantDirectWechatSubscribeConfigRequest(reqParam V2MerchantDirectWechatSubscribeConfigRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_DIRECT_WECHAT_SUBSCRIBE_CONFIG
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
