/**
 * 商户统一进件接口(2022)
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantIntegrateRegRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 渠道商汇付id
    EntType string `json:"ent_type" structs:"ent_type"` // 公司类型
    RegName string `json:"reg_name" structs:"reg_name"` // 商户名称
    BusiType string `json:"busi_type" structs:"busi_type"` // 经营类型
    DetailAddr string `json:"detail_addr" structs:"detail_addr"` // 经营详细地址
    ProvId string `json:"prov_id" structs:"prov_id"` // 经营省
    AreaId string `json:"area_id" structs:"area_id"` // 经营市
    DistrictId string `json:"district_id" structs:"district_id"` // 经营区
    ContactInfo string `json:"contact_info" structs:"contact_info"` // 联系人信息
    CardInfo string `json:"card_info" structs:"card_info"` // 卡信息配置实体
    CashConfig string `json:"cash_config" structs:"cash_config"` // 取现配置列表jsonArray格式 ；
    SettleConfig string `json:"settle_config" structs:"settle_config"` // 结算配置jsonObject格式；

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantIntegrateRegRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantIntegrateRegRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantIntegrateRegRequest(reqParam)
}

func (bp *BsPay) V2MerchantIntegrateRegRequest(reqParam V2MerchantIntegrateRegRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_INTEGRATE_REG
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
