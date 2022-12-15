/**
 * 支付宝实名申请提交
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantBusiAliRealnameApplyRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付ID
    ContactPersonInfo string `json:"contact_person_info" structs:"contact_person_info"` // 联系人信息
    AuthIdentityInfo string `json:"auth_identity_info" structs:"auth_identity_info"` // 主体信息

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantBusiAliRealnameApplyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantBusiAliRealnameApplyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantBusiAliRealnameApplyRequest(reqParam)
}

func (bp *BsPay) V2MerchantBusiAliRealnameApplyRequest(reqParam V2MerchantBusiAliRealnameApplyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_BUSI_ALI_REALNAME_APPLY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
