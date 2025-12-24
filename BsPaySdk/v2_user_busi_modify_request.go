/**
 * 用户业务入驻修改
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2UserBusiModifyRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付ID
    LjhData string `json:"ljh_data" structs:"ljh_data"` // 乐接活配置当合作平台为乐接活，必填
    SignUserInfo string `json:"sign_user_info" structs:"sign_user_info"` // 签约人信息当电子回单配置开关为开通时必填
    HxyData string `json:"hxy_data" structs:"hxy_data"` // 汇薪云配置当合作平台为汇薪云时，该参数必填

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2UserBusiModifyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2UserBusiModifyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2UserBusiModifyRequest(reqParam)
}

func (bp *BsPay) V2UserBusiModifyRequest(reqParam V2UserBusiModifyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_USER_BUSI_MODIFY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
