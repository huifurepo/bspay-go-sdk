/**
 * 支付宝直连-申请当面付代签约
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantDirectAlipayFacetofacesignApplyRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 上级主体ID
    DirectCategory string `json:"direct_category" structs:"direct_category"` // 支付宝经营类目
    AppId string `json:"app_id" structs:"app_id"` // 开发者的应用ID
    ContactName string `json:"contact_name" structs:"contact_name"` // 联系人姓名
    ContactMobileNo string `json:"contact_mobile_no" structs:"contact_mobile_no"` // 联系人手机号
    ContactEmail string `json:"contact_email" structs:"contact_email"` // 联系人电子邮箱
    Account string `json:"account" structs:"account"` // 商户账号
    Rate string `json:"rate" structs:"rate"` // 服务费率（%）0.38~3之间，精确到0.01。当签约且授权sign_and_auth&#x3D;Y时，必填。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.38&lt;/font&gt;
    FileList string `json:"file_list" structs:"file_list"` // 文件列表

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantDirectAlipayFacetofacesignApplyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantDirectAlipayFacetofacesignApplyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantDirectAlipayFacetofacesignApplyRequest(reqParam)
}

func (bp *BsPay) V2MerchantDirectAlipayFacetofacesignApplyRequest(reqParam V2MerchantDirectAlipayFacetofacesignApplyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_DIRECT_ALIPAY_FACETOFACESIGN_APPLY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
