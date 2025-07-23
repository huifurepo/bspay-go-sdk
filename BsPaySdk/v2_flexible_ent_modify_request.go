/**
 * 灵工企业商户业务修改
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2FlexibleEntModifyRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户汇付id
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 渠道商汇付ID
    BasicInfo string `json:"basic_info" structs:"basic_info"` // 商户基本信息jsonObject格式；其中的contact_info和legal_info联系人和法人信息可能在卡信息修改时需要
    SignUserInfo string `json:"sign_user_info" structs:"sign_user_info"` // 签约人

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2FlexibleEntModifyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2FlexibleEntModifyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2FlexibleEntModifyRequest(reqParam)
}

func (bp *BsPay) V2FlexibleEntModifyRequest(reqParam V2FlexibleEntModifyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_FLEXIBLE_ENT_MODIFY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
