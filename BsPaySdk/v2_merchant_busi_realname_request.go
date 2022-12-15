/**
 * 微信实名认证
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantBusiRealnameRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付ID
    Name string `json:"name" structs:"name"` // 联系人姓名
    Mobile string `json:"mobile" structs:"mobile"` // 联系人手机号
    IdCardNumber string `json:"id_card_number" structs:"id_card_number"` // 联系人身份证号码
    ContactType string `json:"contact_type" structs:"contact_type"` // 联系人类型

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantBusiRealnameRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantBusiRealnameRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantBusiRealnameRequest(reqParam)
}

func (bp *BsPay) V2MerchantBusiRealnameRequest(reqParam V2MerchantBusiRealnameRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_BUSI_REALNAME
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
