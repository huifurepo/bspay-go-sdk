/**
 * 个人商户基本信息入驻
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantBasicdataIndvRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 上级主体ID
    RegName string `json:"reg_name" structs:"reg_name"` // 商户名
    ProvId string `json:"prov_id" structs:"prov_id"` // 经营省
    AreaId string `json:"area_id" structs:"area_id"` // 经营市
    DistrictId string `json:"district_id" structs:"district_id"` // 经营区
    DetailAddr string `json:"detail_addr" structs:"detail_addr"` // 经营详细地址
    ContactName string `json:"contact_name" structs:"contact_name"` // 联系人姓名
    ContactMobileNo string `json:"contact_mobile_no" structs:"contact_mobile_no"` // 联系人手机号
    ContactEmail string `json:"contact_email" structs:"contact_email"` // 联系人电子邮箱
    CardInfo string `json:"card_info" structs:"card_info"` // 结算卡信息配置

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantBasicdataIndvRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantBasicdataIndvRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantBasicdataIndvRequest(reqParam)
}

func (bp *BsPay) V2MerchantBasicdataIndvRequest(reqParam V2MerchantBasicdataIndvRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_BASICDATA_INDV
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
