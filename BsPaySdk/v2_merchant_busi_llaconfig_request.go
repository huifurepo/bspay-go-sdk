/**
 * 代运营代扣业务配置
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantBusiLlaconfigRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户汇付id
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 所属渠道商
    LlaAgencyConfig string `json:"lla_agency_config" structs:"lla_agency_config"` // 代运营配置json字符串，业务角色为AGENCY:代运营时必填
    LlaMerchantConfig string `json:"lla_merchant_config" structs:"lla_merchant_config"` // 商家配置json字符串，业务角色为MERCHANT:商家时必填
    PaperAgreementFile string `json:"paper_agreement_file" structs:"paper_agreement_file"` // 纸质协议文件协议类型为纸质且业务角色不为空时必填，文件类型F633；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
    SignUserInfo string `json:"sign_user_info" structs:"sign_user_info"` // 签约人信息json字符串，协议类型为电子时必填

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantBusiLlaconfigRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantBusiLlaconfigRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantBusiLlaconfigRequest(reqParam)
}

func (bp *BsPay) V2MerchantBusiLlaconfigRequest(reqParam V2MerchantBusiLlaconfigRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_BUSI_LLACONFIG
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
