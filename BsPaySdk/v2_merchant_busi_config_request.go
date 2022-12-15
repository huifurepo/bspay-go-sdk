/**
 * 微信商户配置
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantBusiConfigRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id
    FeeType string `json:"fee_type" structs:"fee_type"` // 业务开通类型
    WxWoaAppId string `json:"wx_woa_app_id" structs:"wx_woa_app_id"` // 公众号支付Appid条件必填，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wx3767c5bd01df5061&lt;/font&gt; ；wx_woa_app_id 和 wx_applet_app_id两者不能同时为空
    WxAppletAppId string `json:"wx_applet_app_id" structs:"wx_applet_app_id"` // 微信小程序APPID条件必填，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wx8523175fea790f10&lt;/font&gt; ；wx_woa_app_id，wx_applet_app_id两者不能同时为空

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantBusiConfigRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantBusiConfigRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantBusiConfigRequest(reqParam)
}

func (bp *BsPay) V2MerchantBusiConfigRequest(reqParam V2MerchantBusiConfigRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_BUSI_CONFIG
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
