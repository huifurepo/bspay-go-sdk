/**
 * 微信特约商户进件
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantDirectWechatSignRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户汇付Id
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 渠道商汇付Id
    AppId string `json:"app_id" structs:"app_id"` // 开发者的应用ID
    MchId string `json:"mch_id" structs:"mch_id"` // 商户号
    Owner string `json:"owner" structs:"owner"` // 经营者/法人是否为受益人
    ContactInfo string `json:"contact_info" structs:"contact_info"` // 超级管理员信息
    SalesScenesType string `json:"sales_scenes_type" structs:"sales_scenes_type"` // 经营场景类型
    SalesInfo string `json:"sales_info" structs:"sales_info"` // 经营场景
    SettlementInfo string `json:"settlement_info" structs:"settlement_info"` // 结算信息

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantDirectWechatSignRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantDirectWechatSignRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantDirectWechatSignRequest(reqParam)
}

func (bp *BsPay) V2MerchantDirectWechatSignRequest(reqParam V2MerchantDirectWechatSignRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_DIRECT_WECHAT_SIGN
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
