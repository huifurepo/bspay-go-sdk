/**
 * 用户业务入驻
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2UserBusiOpenRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付ID
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 渠道商/商户汇付Id
    LjhData string `json:"ljh_data" structs:"ljh_data"` // 乐接活配置当合作平台为乐接活，必填

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2UserBusiOpenRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2UserBusiOpenRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2UserBusiOpenRequest(reqParam)
}

func (bp *BsPay) V2UserBusiOpenRequest(reqParam V2UserBusiOpenRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_USER_BUSI_OPEN
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
